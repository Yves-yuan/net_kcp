/*
	日期: 2017-07-07
	作者: Long Heng
	功能: KCP Server启动器, 启动时需要根据客户端的编解码设置不同的Decoder/Encoder
*/
package net

import (
	"server/core/message/packet"
	"server/core/net/codec"
	"server/core/net/internal/kcp"

	"server/core/net/session"

	"github.com/pkg/errors"
)

const (
	// 数据包堆积大小
	packetBacklog = 64 * 1024

	// session关闭等待执行的回调队列大小
	closedSessionBacklog = 1 << 12
)

// 启动一个KCP服务器
func StartKCPServer(addr string, heartbeat int64) error {
	listener, err := kcp.ListenWithOptions(addr, nil, 0, 0)
	if err != nil {
		return errors.WithStack(err)
	}

	// 设置实例化Encoder的工厂函数
	listener.SetEncoderFactory(func() kcp.IEncoder {
		return codec.NewTransportEncoder()
	})

	// 设置实例化Decoder的工厂函数
	listener.SetDecoderFactory(func() kcp.IDecoder {
		return codec.NewTransportDecoder()
	})

	// 设置发送队列
	chRead := make(chan *kcp.SessionPacket, packetBacklog) // 接收队列
	chSend := make(chan *kcp.SessionPacket, packetBacklog) // 发送队列
	listener.SetReadPacketChan(chRead)
	listener.SetSendPacketChan(chSend)

	// 序列化和反序列化队列
	chDeserialize := make(chan *packet.CRecvPacket, packetBacklog) // 反序列化以后的消息包
	chSerialized := make(chan *packet.CSendPacket, packetBacklog)  // 序列化的原始消息队列
	listener.SetSerializerChan(chSerialized)

	// session关闭的回调函数队列, 需要在逻辑线程中处理session关闭回调
	chClosedSession := make(chan *session.CSession, closedSessionBacklog)
	listener.SetClosedSessionChan(chClosedSession)

	// 心跳时间
	// TODO: 临时设置
	listener.SetHeartbeatInternal(heartbeat)

	// 启动codec线程
	go codec.StartDecode(chRead, chDeserialize)
	go codec.StartEncode(chSerialized, chSend)

	// 启动逻辑分发线程(1. 处理已经反序列化的包, 2. session close的回调也应该在逻辑线程中调用)
	go dispatch(chDeserialize, chClosedSession)

	listener.Start()

	// 注册性能监测工具
	registerMetrics()

	return nil
}
