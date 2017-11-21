/*************************************************
  Copyright (C), 2006-2010, 8760 Tech. Co., Ltd.
  作    者: Yuan
  版    本: v1
  完成日期: 2017-07-08
  功能描述: CTransportDecoder从网络层接收到的二进制进行粘包处理, 然后反序列化
            1. Decode 解码
*************************************************/
package codec

import (
	"server/core/message/packet"
	"server/core/net/session"
	"server/message"
	"server/route"

	"bytes"
	"encoding/binary"
)

type CTransportDecoder struct {
	*bytes.Buffer
	m_size      int    // 当前包的大小
	m_lastTime  uint64 // 上一个包的时间
	m_lastOrder uint32 // 上一个包的序号
}

// 实例化
func NewTransportDecoder() *CTransportDecoder {
	return &CTransportDecoder{
		Buffer: &bytes.Buffer{},
		m_size: -1,
	}
}

// 通过输入解码数据包
// sess:  用户session
// input: 数据
func (d *CTransportDecoder) Decode(sess *session.CSession, input []byte) (out []*packet.CRecvPacket) {
	// 写入数据
	d.Write(input)

	// 检查长度
	if d.Len() < HEADER_SIZE {
		return
	}

	// 上次没有读取长度
	if d.m_size < 0 {
		d.m_size = int(binary.BigEndian.Uint32(d.Next(recvSizeLength)))
	}

	// 读取所有的包
	for d.m_size <= d.Len() {
		buffer := d.Next(d.m_size)

		// 时间校验
		time := binary.BigEndian.Uint64(buffer[:recvTimeSize])
		if d.m_lastTime > time {
			logger.Errorf("Session时间不吻合, 上次时间=%d, 本次时间=%d", d.m_lastTime, time)
			sess.Close()
			return
		}

		// 序号验证
		order := binary.BigEndian.Uint32(buffer[recvOrderOffset : recvOrderOffset+recvOrderLength])
		//order = order ^ (0xFE98 << 8)
		//order = order ^ uint32(d.m_size)
		if d.m_lastOrder != 0 && order != d.m_lastOrder+1 {
			logger.Errorf("Order检查失败, 上次序号=%d, 本次序号=%d", d.m_lastOrder, order)
			sess.Close()
			return
		}
		d.m_lastOrder = order

		// 读取消息ID
		msgid := binary.BigEndian.Uint16(buffer[recvProtoOffset : recvProtoOffset+recvProtoSize])
		// 获取handler
		handler, err := route.FetchHandler(message.MSGID(msgid))
		if err != nil {
			logger.Errorf("获取Handler失败, msgid=%s, err=%s", msgid, err.Error())
			return
		}

		pkt := packet.NewRecv(sess, handler)
		pkt.SetData(buffer[recvPayloadOffset:d.m_size])

		// 添加到返回列表
		out = append(out, pkt)

		// 读取连续的消息
		if d.Len() < HEADER_SIZE {
			d.m_size = -1
			break
		} else {
			d.m_size = int(binary.BigEndian.Uint32(d.Next(recvSizeLength)))
		}
	}

	return
}
