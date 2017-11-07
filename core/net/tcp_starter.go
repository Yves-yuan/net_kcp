package net

import (
	"net"

	"server/core/log"
	"server/core/message/packet"
	"server/core/net/codec"
	"server/core/net/session"
	"server/message"

	"github.com/pkg/errors"
)

const bufferSize = 512 // TCP接收缓冲区大小

type agentStatus byte

const (
	agentStatusInit = iota
	agentStatusWork
	agentStatusClosed
)

var (
	logger = log.New("net")

	ErrAgentClosed = errors.New("closed agent") // 发送数据, 但是客户端agent已经关闭
)

func init() {
	logger.SetLevel(log.DebugLevel)
}

type agent struct {
	connection net.Conn                 // 底层网络连接
	status     agentStatus              // agent状态
	chDie      chan struct{}            // agent关闭信号
	chWrite    chan *packet.CSendPacket // 发送数据队列
	session    *session.CSession        // 用户session
	encoder    *codec.CTransportEncoder // Encoder
	decoder    *codec.CTransportDecoder // Decoder
}

// 返回agent实例
func newAgent(conn net.Conn) *agent {
	a := &agent{
		connection: conn,
		status:     agentStatusInit,
		chWrite:    make(chan *packet.CSendPacket, 64),
		chDie:      make(chan struct{}, 1),
	}

	// 初始化用户session并绑定到agent
	s := session.New(a)
	a.session = s

	// 编码与解码
	a.encoder = codec.NewTransportEncoder()
	a.decoder = codec.NewTransportDecoder()

	return a
}

// 写入数据到TCP底层缓冲区
func (a *agent) write() {
	// 写入协程退出时, 关闭写入channel
	defer close(a.chWrite)

	for {
		select {
		case pkt := <-a.chWrite: // 写入数据到底层网络连接
			payload, err := pkt.Payload()
			if err != nil {
				logger.Error(err)
				break
			}

			// 编码
			data := a.encoder.Encode(int(pkt.ID()), pkt.Type(), payload)

			// 如果底层网络断开时关闭底层网络
			if _, err := a.connection.Write(data); err != nil {
				logger.Error(err)
				a.connection.Close()
				return
			}

		case <-a.chDie: // agent关闭信号
			return
		}
	}
}

func (a *agent) RemoteAddr() net.Addr {
	return a.connection.RemoteAddr()
}

func (a *agent) Send(msgid message.MSGID, msg interface{}) error {
	if a.status == agentStatusClosed {
		return ErrAgentClosed
	}

	data := packet.NewSend(a, msgid, msg, 0)
	a.chWrite <- data

	return nil
}

func (a *agent) Close() error {
	if a.status == agentStatusClosed {
		return ErrAgentClosed
	}

	// 关闭网络连接, connection.Read会立即返回, 其他相关清理在Read返回后处理
	return a.connection.Close()
}

// 启动TCPServer
func StartTCPServer(addr string) error {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	// 队列
	chDeserialize := make(chan *packet.CRecvPacket, packetBacklog)        // 已序列化好, 等待处理的队列
	chClosedSession := make(chan *session.CSession, closedSessionBacklog) // 已经关闭session, 等待回调session close队列

	// 监听TCP端口是否有新连接
	go monitor(listener, chDeserialize, chClosedSession)

	// 开启逻辑线程, 逻辑分发
	go dispatch(chDeserialize, chClosedSession)

	// 注册性能监测工具
	registerMetrics()

	return nil
}

// 从TCP端口接收连接
func monitor(listener net.Listener, chDeserialize chan<- *packet.CRecvPacket, chClosedSession chan<- *session.CSession) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			logger.Error(err)
			continue
		}

		go handle(conn, chDeserialize, chClosedSession)
	}
}

// 处理TCP连接
func handle(conn net.Conn, chDeserialize chan<- *packet.CRecvPacket, chClosedSession chan<- *session.CSession) {
	logger.Infof("新连接建立: %s", conn.RemoteAddr().String())

	// 客户端的网络层agent
	a := newAgent(conn)

	// 开启写协程, 将数据写入底层网络缓冲区
	go a.write()

	// TCP连接断开时需要进行清理工作
	defer func() {
		logger.Infof("TCP连接断开, 清理连接相关数据")
		a.Close()
		a.status = agentStatusClosed
		close(a.chDie)
		chClosedSession <- a.session
	}()

	a.status = agentStatusWork

	buffer := make([]byte, bufferSize)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			logger.Errorf("TCP连接读取数据错误: %s", err.Error())
			return
		}

		// 将解出的所有包放入逻辑处理队列
		packets := a.decoder.Decode(a.session, buffer[:n])
		for i := range packets {
			chDeserialize <- packets[i]
		}
	}
}
