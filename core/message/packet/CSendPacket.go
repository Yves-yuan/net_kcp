/*
	日期: 2017-07-08
	作者: Long Heng
	功能: CSendPacke包装一个发送到客户端的消息, 消息还未被序列化, CSendPacket实例会
	      发送到序列化线程, 序列化线程序列化后通过网络发送出去
*/
package packet

import (
	"server/message"

	"server/core/log"

	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
	"server/core/net/session"
)

var (
	ErrUnsupportedPayload = errors.New("发送数据包不支持的Payload类型")

	logger = log.New("packet")
)

type CSendPacket struct {
	m_id   int            // 消息ID
	m_type int            // 消息类型
	m_conn session.Sender // 底层网络实体, 可以是TCP连接或者KCP连接
	m_data interface{}    // 消息可以是proto.Message或[]byte, 如果是[]byte表示广播消息, 已经提前序列化好
}

// 返回一个新实例
func NewSend(conn session.Sender, id message.MSGID, data interface{}, typ int) *CSendPacket {
	return &CSendPacket{
		m_conn: conn,
		m_id:   int(id),
		m_data: data,
		m_type: typ,
	}
}

// 返回消息ID
func (p *CSendPacket) ID() int {
	return p.m_id
}

// 返回消息类型
func (p *CSendPacket) Type() int {
	return p.m_type
}

// 返回底层网络实体
func (p *CSendPacket) Conn() session.Sender {
	return p.m_conn
}

// 返回序列化后的数据
func (p *CSendPacket) Payload() ([]byte, error) {
	switch v := p.m_data.(type) {
	case []byte:
		return v, nil
	case proto.Message:
		return proto.Marshal(v)
	default:
		return nil, ErrUnsupportedPayload
	}
}
