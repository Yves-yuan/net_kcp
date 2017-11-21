/*
	日期: 2017-07-08
	作者: Yuan
	功能: CRecvPacket包装一个已经反序列化好的数据包, CRecvPacket实例有反序列化线程实例化
	      发送到逻辑处理队列.
*/
package packet

import (
	"server/core/net/session"
	"server/route"

	"github.com/golang/protobuf/proto"
)

type CRecvPacket struct {
	m_session *session.CSession
	m_handler *route.CHandler
	m_payload proto.Message
}

// 实例化CRecvPacket
func NewRecv(s *session.CSession, h *route.CHandler) *CRecvPacket {
	return &CRecvPacket{
		m_session: s,
		m_handler: h,
	}
}

// 发返Session
func (p *CRecvPacket) Session() *session.CSession {
	return p.m_session
}

// 返回处理器
func (p *CRecvPacket) Handler() *route.CHandler {
	return p.m_handler
}

// 返回反序列化好的消息
func (p *CRecvPacket) Payload() proto.Message {
	return p.m_payload
}

func (p *CRecvPacket) SetData(data []byte) {
	p.m_payload = p.m_handler.Payload()

	// 反序列化
	if err := proto.Unmarshal(data, p.m_payload); err != nil {
		logger.Errorf("CExternalKcpDecoder decode Unmarshal go wrong, err=%s", err.Error())
		return
	}
}
