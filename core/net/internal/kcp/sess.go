package kcp

import (
	"container/heap"
	"crypto/rand"
	"encoding/binary"
	"hash/crc32"
	"io"
	"log"
	"net"
	"sync"
	"sync/atomic"
	"time"

	"server/core/message/packet"
	"server/core/net/session"
	"server/message"

	"github.com/pkg/errors"
	"golang.org/x/net/ipv4"
)

type errTimeout struct {
	error
}

func (errTimeout) Timeout() bool   { return true }
func (errTimeout) Temporary() bool { return true }
func (errTimeout) Error() string   { return "i/o timeout" }

const (
	// 16-bytes magic number for each packet
	nonceSize = 16

	// 4-bytes packet checksum
	crcSize = 4

	// overall crypto header size
	cryptHeaderSize = nonceSize + crcSize

	// maximum packet size
	mtuLimit = 1500

	// FEC keeps rxFECMulti* (dataShard+parityShard) ordered packets in memory
	rxFECMulti = 3

	// accept backlog
	acceptBacklog = 128

	// prerouting(to session) queue
	qlen = 128

	// packet queue size, 64k
	packetBacklog = 1 << 16

	// 等待删除队列大小(极端情况就是所有玩家同时断线)
	deleteBacklog = 1 << 12
)

const (
	errBrokenPipe       = "broken pipe"
	errInvalidOperation = "invalid operation"
)

var (
	// global packet buffer
	// shared among sending/receiving/FEC
	xmitBuf sync.Pool

	// monotonic session id
	sid uint32
)

func init() {
	xmitBuf.New = func() interface{} {
		return make([]byte, mtuLimit)
	}
}

type IDecoder interface {
	Decode(session *session.CSession, input []byte) (out []*packet.CRecvPacket)
}

type IEncoder interface {
	Encode(id, typ int, payload []byte) []byte
}

type (
	// UDPSession defines a KCP session implemented by UDP
	UDPSession struct {
		conn     net.PacketConn    // the underlying packet connection
		remote   net.Addr          // remote peer address
		block    BlockCrypt        // block encryption
		decoder  IDecoder          // 解码接口
		encoder  IEncoder          // 编码接口
		cSession *session.CSession // 应用称session
		kcp      *KCP              // KCP ARQ protocol
		l        *Listener         // point to the Listener if it's accepted by Listener

		// FEC
		fecDecoder *FECDecoder
		fecEncoder *FECEncoder

		ts         int64  // 接收到消息时的时间戳
		updaterIdx int    // session最小堆索引
		dup        int    // duplicate udp packets
		headerSize int    // the overall header size added before KCP frame
		sid        uint32 // session id(monotonic)

		ackNoDelay bool // send ack immediately for each incoming packet
		writeDelay bool // delay kcp.flush() for Write() for bulk transfer
		isClosed   bool // flag the session has Closed

		rd time.Time // read deadline
		wd time.Time // write deadline
		mu sync.Mutex

		// notifications
		die     chan struct{} // notify session has Closed
		chWrite chan []byte   // 写入数据的管道
		ext     []byte        // extended output buffer(with header)
	}

	setReadBuffer interface {
		SetReadBuffer(bytes int) error
	}

	setWriteBuffer interface {
		SetWriteBuffer(bytes int) error
	}
)

// newUDPSession create a new udp session for client or server
func newUDPSession(conv uint32, dataShards, parityShards int, l *Listener, conn net.PacketConn, remote net.Addr, block BlockCrypt) *UDPSession {
	sess := new(UDPSession)
	sess.sid = atomic.AddUint32(&sid, 1)
	sess.die = make(chan struct{})
	sess.remote = remote
	sess.conn = conn
	sess.l = l
	sess.block = block

	sess.ts = time.Now().Unix()

	sess.chWrite = make(chan []byte, 512)

	// FEC initialization
	sess.fecDecoder = newFECDecoder(rxFECMulti*(dataShards+parityShards), dataShards, parityShards)
	if sess.block != nil {
		sess.fecEncoder = newFECEncoder(dataShards, parityShards, cryptHeaderSize)
	} else {
		sess.fecEncoder = newFECEncoder(dataShards, parityShards, 0)
	}

	// calculate header size
	if sess.block != nil {
		sess.headerSize += cryptHeaderSize
	}
	if sess.fecEncoder != nil {
		sess.headerSize += fecHeaderSizePlus2
	}

	// only allocate extended packet buffer
	// when the extra header is required
	if sess.headerSize > 0 {
		sess.ext = make([]byte, mtuLimit)
	}

	sess.kcp = NewKCP(conv, func(buf []byte, size int) {
		if size >= IKCP_OVERHEAD {
			sess.output(buf[:size])
		}
	})
	sess.kcp.SetMtu(IKCP_MTU_DEF - sess.headerSize)

	// add current session to the global updater,
	// which periodically calls sess.update()
	updater.addSession(sess)

	if sess.l == nil { // it's a client connection
		panic("this code is for server")
		atomic.AddUint64(&DefaultSnmp.ActiveOpens, 1)
	} else {
		atomic.AddUint64(&DefaultSnmp.PassiveOpens, 1)
	}
	currestab := atomic.AddUint64(&DefaultSnmp.CurrEstab, 1)
	maxconn := atomic.LoadUint64(&DefaultSnmp.MaxConn)
	if currestab > maxconn {
		atomic.CompareAndSwapUint64(&DefaultSnmp.MaxConn, maxconn, currestab)
	}

	// 实例化应用层session(用户层session与传输层session不同)
	sess.cSession = session.New(sess)

	// 实例化codec
	sess.encoder = l.encoderFactory()
	sess.decoder = l.decoderFactory()

	go sess.write()

	return sess
}

func (s *UDPSession) Encoder() IEncoder {
	return s.encoder
}

func (s *UDPSession) Decoder() IDecoder {
	return s.decoder
}

func (s *UDPSession) Session() *session.CSession {
	return s.cSession
}

// 发送消息到序列化线程, TODO(warning): 此接口只应该由应用层session调用
func (s *UDPSession) Send(msgid message.MSGID, data interface{}) error {
	if s.isClosed {
		return errors.New(errBrokenPipe)
	}

	s.l.chSerializer <- packet.NewSend(s, msgid, data, 0)
	return nil
}

func (s *UDPSession) write() {
	for data := range s.chWrite {
		if _, err := s.conn.WriteTo(data, s.remote); err != nil {
			log.Println(err.Error())
		}
	}
}

func (l *Listener) Start() {
	// 处理UDP消息
	go l.monitor()

	// KCP控制线程
	go l.controller()
}

func (l *Listener) controller() {
	var (
		timer     <-chan time.Time
		h         = &updater
		chTimeout = make(chan *UDPSession, 128)

		// capacity of channel
		chTimeoutLen    = cap(chTimeout)
		chReadPacketLen = cap(l.chReadPacket)
	)

	for {
		select {
		case p, ok := <-l.chInput: // 对读取的UPD包, 将读取的数据Input进相应的kcp对象
			if !ok {
				break
			}
			now := time.Now().Unix()
		READ:
			// input
			s := l.dealInput(p)
			if s == nil || s.isClosed {
				continue
			}

			if s.ackNoDelay {
				s.kcp.flush(true)
			}

			// receive all packet
			n := s.kcp.PeekSize()
			for n > 0 && len(l.chReadPacket) < chReadPacketLen {
				buf := make([]byte, n)
				if n := s.kcp.Recv(buf); n < 0 {
					continue
				}
				s.ts = now
				l.chReadPacket <- &SessionPacket{s, buf}
				n = s.kcp.PeekSize()
			}

			for len(l.chInput) > 0 {
				p, ok = <-l.chInput
				if !ok {
					break
				}

				goto READ
			}

		case p, ok := <-l.chSendPacket: // 等待发送的数据包, Send进kcp对象
			if !ok {
				continue
			}

		SEND:
			p.From.kcp.Send(p.Data)

			if !p.From.writeDelay {
				p.From.kcp.flush(false)
			}

			for len(l.chSendPacket) > 0 {
				p, ok = <-l.chSendPacket
				if !ok {
					break
				}

				goto SEND
			}

		case <-timer: // 定时刷新
			hlen := h.Len()
			now := time.Now()
			for i := 0; i < hlen; i++ {
				entry := heap.Pop(h).(entry)
				if time.Now().Unix()-entry.s.ts > l.heartbeatInternal && len(chTimeout) < chTimeoutLen {
					chTimeout <- entry.s
				}
				if now.After(entry.ts) {
					entry.ts = now.Add(entry.s.update())
					heap.Push(h, entry)
				} else {
					heap.Push(h, entry)
					break
				}
			}
			if h.Len() > 0 {
				timer = time.After(h.entries[0].ts.Sub(now))
			}

		case <-h.chWakeUp: // 唤醒刷新
			if h.Len() > 0 {
				now := time.Now()
				timer = time.After(h.entries[0].ts.Sub(now))
			}

		case s, ok := <-chTimeout: // 处理超时udp session
			if !ok {
				continue
			}
			s.Close()

		case s, ok := <-l.chDelete: // 删除session
			if !ok {
				continue
			}
			delete(l.sessions, s)

		case <-l.die: // listener关闭
			log.Println("kcp controller quit")
			return
		}
	}
}

// Close closes the connection.
// TODO(warning): session的close可能在以下线程调用, 需要放入队列, 在逻辑线程中处理回调
// 1. kcp controller中超时
// 2. decoder中由于玩家包验证失败调用
// 3. 应用层玩家主动关闭close
func (s *UDPSession) Close() error {
	updater.removeSession(s)
	s.mu.Lock()
	defer s.mu.Unlock()

	// 验证session状态
	if s.isClosed {
		return errors.New(errBrokenPipe)
	}
	s.isClosed = true

	// 关闭相关channel
	close(s.die)
	close(s.chWrite)

	// 移除从listener的sessions中移除session
	atomic.AddUint64(&DefaultSnmp.CurrEstab, ^uint64(0))
	if s.l == nil { // client socket close
		return s.conn.Close()
	} else {
		s.l.chDelete <- s.RemoteAddr().String()
	}

	// 放入队列, 逻辑线程取出后, 调用所有的回调函数,
	// 放出逻辑线程调用是防止并发
	s.l.chClosedSession <- s.cSession

	return nil
}

// LocalAddr returns the local network address. The Addr returned is shared by all invocations of LocalAddr, so do not modify it.
func (s *UDPSession) LocalAddr() net.Addr { return s.conn.LocalAddr() }

// RemoteAddr returns the remote network address. The Addr returned is shared by all invocations of RemoteAddr, so do not modify it.
func (s *UDPSession) RemoteAddr() net.Addr { return s.remote }

// SetDeadline sets the deadline associated with the listener. A zero time value disables the deadline.
func (s *UDPSession) SetDeadline(t time.Time) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.rd = t
	s.wd = t
	return nil
}

// SetReadDeadline implements the Conn SetReadDeadline method.
func (s *UDPSession) SetReadDeadline(t time.Time) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.rd = t
	return nil
}

// SetWriteDeadline implements the Conn SetWriteDeadline method.
func (s *UDPSession) SetWriteDeadline(t time.Time) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.wd = t
	return nil
}

// SetWriteDelay delays write for bulk transfer until the next update interval
func (s *UDPSession) SetWriteDelay(delay bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.writeDelay = delay
}

// SetWindowSize set maximum window size
func (s *UDPSession) SetWindowSize(sndwnd, rcvwnd int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.kcp.WndSize(sndwnd, rcvwnd)
}

// SetMtu sets the maximum transmission unit(not including UDP header)
func (s *UDPSession) SetMtu(mtu int) bool {
	if mtu > mtuLimit {
		return false
	}

	s.mu.Lock()
	defer s.mu.Unlock()
	s.kcp.SetMtu(mtu - s.headerSize)
	return true
}

// SetStreamMode toggles the stream mode on/off
func (s *UDPSession) SetStreamMode(enable bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if enable {
		s.kcp.stream = 1
	} else {
		s.kcp.stream = 0
	}
}

func (s *UDPSession) SetFastMode() {
	s.kcp.NoDelay(1, 10, 2, 1)
}

// SetACKNoDelay changes ack flush option, set true to flush ack immediately,
func (s *UDPSession) SetACKNoDelay(nodelay bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.ackNoDelay = nodelay
}

// SetDUP duplicates udp packets for kcp output, for testing purpose only
func (s *UDPSession) SetDUP(dup int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.dup = dup
}

// SetNoDelay calls nodelay() of kcp
func (s *UDPSession) SetNoDelay(nodelay, interval, resend, nc int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.kcp.NoDelay(nodelay, interval, resend, nc)
}

// SetDSCP sets the 6bit DSCP field of IP header, no effect if it's accepted from Listener
func (s *UDPSession) SetDSCP(dscp int) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.l == nil {
		if nc, ok := s.conn.(*ConnectedUDPConn); ok {
			return ipv4.NewConn(nc.Conn).SetTOS(dscp << 2)
		} else if nc, ok := s.conn.(net.Conn); ok {
			return ipv4.NewConn(nc).SetTOS(dscp << 2)
		}
	}
	return errors.New(errInvalidOperation)
}

// SetReadBuffer sets the socket read buffer, no effect if it's accepted from Listener
func (s *UDPSession) SetReadBuffer(bytes int) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.l == nil {
		if nc, ok := s.conn.(setReadBuffer); ok {
			return nc.SetReadBuffer(bytes)
		}
	}
	return errors.New(errInvalidOperation)
}

// SetWriteBuffer sets the socket write buffer, no effect if it's accepted from Listener
func (s *UDPSession) SetWriteBuffer(bytes int) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.l == nil {
		if nc, ok := s.conn.(setWriteBuffer); ok {
			return nc.SetWriteBuffer(bytes)
		}
	}
	return errors.New(errInvalidOperation)
}

// output pipeline entry
// steps for output data processing:
// 0. Header extends
// 1. FEC
// 2. CRC32
// 3. Encryption
// 4. WriteTo kernel
func (s *UDPSession) output(buf []byte) {
	var ecc [][]byte

	// extend buf's header space
	ext := buf
	if s.headerSize > 0 {
		ext = s.ext[:s.headerSize+len(buf)]
		copy(ext[s.headerSize:], buf)
	}

	// FEC stage
	if s.fecEncoder != nil {
		ecc = s.fecEncoder.Encode(ext)
	}

	// encryption stage
	if s.block != nil {
		io.ReadFull(rand.Reader, ext[:nonceSize])
		checksum := crc32.ChecksumIEEE(ext[cryptHeaderSize:])
		binary.LittleEndian.PutUint32(ext[nonceSize:], checksum)
		s.block.Encrypt(ext, ext)

		if ecc != nil {
			for k := range ecc {
				io.ReadFull(rand.Reader, ecc[k][:nonceSize])
				checksum := crc32.ChecksumIEEE(ecc[k][cryptHeaderSize:])
				binary.LittleEndian.PutUint32(ecc[k][nonceSize:], checksum)
				s.block.Encrypt(ecc[k], ecc[k])
			}
		}
	}

	// WriteTo kernel
	nbytes := 0
	npkts := 0
	// if mrand.Intn(100) < 50 {
	for i := 0; i < s.dup+1; i++ {
		//if n, err := s.conn.WriteTo(ext, s.remote); err == nil {
		//	nbytes += n
		//	npkts++
		//}
		//if len(ext) == IKCP_OVERHEAD && ext[4] == IKCP_CMD_PUSH {
		//	panic(fmt.Sprintf("send segment does not contain body: % x", ext))
		//}
		if s.isClosed {
			return
		}
		data := make([]byte, len(ext))
		copy(data, ext)
		s.chWrite <- data
	}
	// }

	if ecc != nil {
		for k := range ecc {
			if n, err := s.conn.WriteTo(ecc[k], s.remote); err == nil {
				nbytes += n
				npkts++
			}
		}
	}
	atomic.AddUint64(&DefaultSnmp.OutPkts, uint64(npkts))
	atomic.AddUint64(&DefaultSnmp.OutBytes, uint64(nbytes))
}

// kcp update, returns interval for next calling
func (s *UDPSession) update() (interval time.Duration) {
	s.kcp.flush(false)
	return time.Duration(s.kcp.interval) * time.Millisecond
}

// GetConv gets conversation id of a session
func (s *UDPSession) GetConv() uint32 {
	return s.kcp.conv
}

func (s *UDPSession) kcpInput(data []byte) {
	var kcpInErrors, fecErrs, fecRecovered, fecParityShards uint64
	if s.fecDecoder != nil {
		f := s.fecDecoder.decodeBytes(data)
		s.mu.Lock()
		if f.flag == typeData {
			if ret := s.kcp.Input(data[fecHeaderSizePlus2:], true, s.ackNoDelay); ret != 0 {
				kcpInErrors++
			}
		}

		if f.flag == typeData || f.flag == typeFEC {
			if f.flag == typeFEC {
				fecParityShards++
			}

			if recovers := s.fecDecoder.Decode(f); recovers != nil {
				for _, r := range recovers {
					if len(r) >= 2 { // must be larger than 2bytes
						sz := binary.LittleEndian.Uint16(r)
						if int(sz) <= len(r) && sz >= 2 {
							if ret := s.kcp.Input(r[2:sz], false, s.ackNoDelay); ret == 0 {
								fecRecovered++
							} else {
								kcpInErrors++
							}
						} else {
							fecErrs++
						}
					} else {
						fecErrs++
					}
				}
			}
		}

		s.mu.Unlock()
	} else {
		s.mu.Lock()
		if ret := s.kcp.Input(data, true, s.ackNoDelay); ret != 0 {
			kcpInErrors++
		}
		s.mu.Unlock()
	}
	atomic.AddUint64(&DefaultSnmp.InPkts, 1)
	atomic.AddUint64(&DefaultSnmp.InBytes, uint64(len(data)))
	if fecParityShards > 0 {
		atomic.AddUint64(&DefaultSnmp.FECParityShards, fecParityShards)
	}
	if kcpInErrors > 0 {
		atomic.AddUint64(&DefaultSnmp.KCPInErrors, kcpInErrors)
	}
	if fecErrs > 0 {
		atomic.AddUint64(&DefaultSnmp.FECErrs, fecErrs)
	}
	if fecRecovered > 0 {
		atomic.AddUint64(&DefaultSnmp.FECRecovered, fecRecovered)
	}
}

type (
	// Listener defines a server listening for connections
	Listener struct {
		block        BlockCrypt     // block encryption
		dataShards   int            // FEC data shard
		parityShards int            // FEC parity shard
		fecDecoder   *FECDecoder    // FEC mock initialization
		conn         net.PacketConn // the underlying packet connection

		sessions map[string]*UDPSession // all sessions accepted by this Listener
		chInput  chan *udpPacket        // UDP端口读取到的UDP数据包
		chDelete chan string            // 等待删除的session远程地址

		chReadPacket    chan<- *SessionPacket      // 从KCP读取到的字节数组
		chSendPacket    <-chan *SessionPacket      // 等待写入到KCP的字节数组(已经有codec处理)
		chSerializer    chan<- *packet.CSendPacket // 等待序列化的原始消息
		chClosedSession chan<- *session.CSession   // 已经关闭的session, 等待逻辑线程调用注册在net.OnSessionClosed的回调函数

		flushSessions map[uint32]*UDPSession

		chAccepts  chan *UDPSession // Listen() backlog
		headerSize int              // the overall header size added before KCP frame
		die        chan struct{}    // notify the listener has closed
		rd         atomic.Value     // read deadline for Accept()
		wd         atomic.Value
		mu         sync.Mutex
		smu        sync.Mutex
		fast       bool

		encoderFactory func() IEncoder // 获取Encoder的工厂函数
		decoderFactory func() IDecoder // 获取Decoder的工厂函数

		heartbeatInternal int64 // 心跳超时时间(秒)
	}

	// incoming packet
	udpPacket struct {
		from net.Addr
		data []byte
	}

	// 来自KCP Recv的字节数组
	SessionPacket struct {
		From *UDPSession
		Data []byte
	}
)

func (l *Listener) SetHeartbeatInternal(second int64) {
	l.heartbeatInternal = second
}

func (l *Listener) SetEncoderFactory(factory func() IEncoder) {
	l.encoderFactory = factory
}

func (l *Listener) SetDecoderFactory(factory func() IDecoder) {
	l.decoderFactory = factory
}

func (l *Listener) SetReadPacketChan(ch chan<- *SessionPacket) {
	l.chReadPacket = ch
}

func (l *Listener) SetSendPacketChan(ch <-chan *SessionPacket) {
	l.chSendPacket = ch
}

func (l *Listener) SetSerializerChan(ch chan<- *packet.CSendPacket) {
	l.chSerializer = ch
}

func (l *Listener) SetClosedSessionChan(ch chan<- *session.CSession) {
	l.chClosedSession = ch
}

func (l *Listener) SetFastMode() {
	l.fast = true
}

func (l *Listener) monitor() {
	go l.receiver()
	for {
		select {
		case <-l.die:
			return
		}
	}
}

func (l *Listener) dealInput(p *udpPacket) (outs *UDPSession) {
	raw := p.data
	data := p.data
	from := p.from
	dataValid := false
	if l.block != nil {
		l.block.Decrypt(data, data)
		data = data[nonceSize:]
		checksum := crc32.ChecksumIEEE(data[crcSize:])
		if checksum == binary.LittleEndian.Uint32(data) {
			data = data[crcSize:]
			dataValid = true
		} else {
			atomic.AddUint64(&DefaultSnmp.InCsumErrors, 1)
		}
	} else if l.block == nil {
		dataValid = true
	}

	if dataValid {
		addr := from.String()
		s, ok := l.sessions[addr]
		if !ok { // new session
			if len(l.chAccepts) < cap(l.chAccepts) { // do not let new session overwhelm accept queue
				var conv uint32
				convValid := false
				if l.fecDecoder != nil {
					isfec := binary.LittleEndian.Uint16(data[4:])
					if isfec == typeData {
						conv = binary.LittleEndian.Uint32(data[fecHeaderSizePlus2:])
						convValid = true
					}
				} else {
					conv = binary.LittleEndian.Uint32(data)
					convValid = true
				}

				if convValid {
					s := newUDPSession(conv, l.dataShards, l.parityShards, l, l.conn, from, l.block)
					if l.fast {
						s.SetFastMode()
					}
					s.kcpInput(data)
					l.flushSessions[s.sid] = s
					l.sessions[addr] = s

					outs = s
				}
			}
		} else {
			s.kcpInput(data)
			l.flushSessions[s.sid] = s
			outs = s
		}
	}

	xmitBuf.Put(raw)
	return
}

func (l *Listener) receiver() {
	for {
		data := xmitBuf.Get().([]byte)[:mtuLimit]
		if n, from, err := l.conn.ReadFrom(data); err == nil && n >= l.headerSize+IKCP_OVERHEAD {
			l.chInput <- &udpPacket{from, data[:n]}
		} else if err != nil {
			return
		} else {
			atomic.AddUint64(&DefaultSnmp.InErrs, 1)
		}
	}
}

// SetReadBuffer sets the socket read buffer for the Listener
func (l *Listener) SetReadBuffer(bytes int) error {
	if nc, ok := l.conn.(setReadBuffer); ok {
		return nc.SetReadBuffer(bytes)
	}
	return errors.New(errInvalidOperation)
}

// SetWriteBuffer sets the socket write buffer for the Listener
func (l *Listener) SetWriteBuffer(bytes int) error {
	if nc, ok := l.conn.(setWriteBuffer); ok {
		return nc.SetWriteBuffer(bytes)
	}
	return errors.New(errInvalidOperation)
}

// SetDSCP sets the 6bit DSCP field of IP header
func (l *Listener) SetDSCP(dscp int) error {
	if nc, ok := l.conn.(net.Conn); ok {
		return ipv4.NewConn(nc).SetTOS(dscp << 2)
	}
	return errors.New(errInvalidOperation)
}

// AcceptKCP accepts a KCP connection
func (l *Listener) AcceptKCP() (*UDPSession, error) {
	var timeout <-chan time.Time
	if tdeadline, ok := l.rd.Load().(time.Time); ok && !tdeadline.IsZero() {
		timeout = time.After(tdeadline.Sub(time.Now()))
	}

	select {
	case <-timeout:
		return nil, &errTimeout{}
	case c := <-l.chAccepts:
		return c, nil
	case <-l.die:
		return nil, errors.New(errBrokenPipe)
	}
}

// SetDeadline sets the deadline associated with the listener. A zero time value disables the deadline.
func (l *Listener) SetDeadline(t time.Time) error {
	l.SetReadDeadline(t)
	l.SetWriteDeadline(t)
	return nil
}

// SetReadDeadline implements the Conn SetReadDeadline method.
func (l *Listener) SetReadDeadline(t time.Time) error {
	l.rd.Store(t)
	return nil
}

// SetWriteDeadline implements the Conn SetWriteDeadline method.
func (l *Listener) SetWriteDeadline(t time.Time) error {
	l.wd.Store(t)
	return nil
}

// Close stops listening on the UDP address. Already Accepted connections are not closed.
func (l *Listener) Close() error {
	close(l.die)
	return l.conn.Close()
}

// Addr returns the listener's network address, The Addr returned is shared by all invocations of Addr, so do not modify it.
func (l *Listener) Addr() net.Addr {
	return l.conn.LocalAddr()
}

// ListenWithOptions listens for incoming KCP packets addressed to the local address laddr on the network "udp" with packet encryption,
// dataShards, parityShards defines Reed-Solomon Erasure Coding parameters
func ListenWithOptions(laddr string, block BlockCrypt, dataShards, parityShards int) (*Listener, error) {
	udpaddr, err := net.ResolveUDPAddr("udp", laddr)
	if err != nil {
		return nil, errors.Wrap(err, "net.ResolveUDPAddr")
	}
	conn, err := net.ListenUDP("udp", udpaddr)
	if err != nil {
		return nil, errors.Wrap(err, "net.ListenUDP")
	}

	return ServeConn(block, dataShards, parityShards, conn)
}

// ServeConn serves KCP protocol for a single packet connection.
func ServeConn(block BlockCrypt, dataShards, parityShards int, conn net.PacketConn) (*Listener, error) {
	l := new(Listener)
	l.conn = conn
	l.sessions = make(map[string]*UDPSession)
	l.chInput = make(chan *udpPacket, packetBacklog)
	l.chDelete = make(chan string, deleteBacklog)
	l.flushSessions = make(map[uint32]*UDPSession)
	l.chAccepts = make(chan *UDPSession, acceptBacklog)
	l.die = make(chan struct{})
	l.dataShards = dataShards
	l.parityShards = parityShards
	l.block = block
	l.fecDecoder = newFECDecoder(rxFECMulti*(dataShards+parityShards), dataShards, parityShards)

	// calculate header size
	if l.block != nil {
		l.headerSize += cryptHeaderSize
	}
	if l.fecDecoder != nil {
		l.headerSize += fecHeaderSizePlus2
	}

	return l, nil
}

func currentMs() uint32 {
	return uint32(time.Now().UnixNano() / int64(time.Millisecond))
}

// ConnectedUDPConn is a wrapper for net.UDPConn which converts WriteTo syscalls
// to Write syscalls that are 4 times faster on some OS'es. This should only be
// used for connections that were produced by a net.Dial* call.
type ConnectedUDPConn struct {
	*net.UDPConn
	Conn net.Conn // underlying connection if any
}

// WriteTo redirects all writes to the Write syscall, which is 4 times faster.
func (c *ConnectedUDPConn) WriteTo(b []byte, addr net.Addr) (int, error) {
	return c.Write(b)
}
