/*
  Copyright (C), 2006-2010, 8760 Tech. Co., Ltd.
  作    者: Yuan
  版    本: v1
  完成日期: 2017-07-08
  功能描述:玩家Session抽象
*/
package session

import (
	"server/message"

	"sync/atomic"

	"net"

	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
)

var (
	uniqueSessionID int64 = 0 // SessionID

	ErrUnsupportedMessageType = errors.New("不支持的消息类型, 消息类型只能为proto.Message或[]byte")
)

// 底层网络抽象
type Sender interface {
	RemoteAddr() net.Addr
	Send(id message.MSGID, data interface{}) error
	Close() error
}

// 玩家session
type CSession struct {
	id     int64
	sender Sender      // 底层网络
	entity interface{} // 用户实体
}

// 实例化一个session
func New(sender Sender) *CSession {
	id := atomic.AddInt64(&uniqueSessionID, 1)
	return &CSession{id: id, sender: sender}
}

// 获取session ID
func (s *CSession) ID() int64 {
	return s.id
}

// 发送消息给客户端
func (s *CSession) Send(id message.MSGID, data interface{}) error {
	switch v := data.(type) {
	case proto.Message:
		if id != message.MSGID_ResHeartbeatE && id != message.MSGID_ResKCPCheckDelayE {
			if debug {
				println("====>", id.String(), v.String())
			}
		}
	case []byte:
		// expect
	default:
		return ErrUnsupportedMessageType
	}
	return s.sender.Send(id, data)
}

// 应用层调用Close来关闭当前session
func (s *CSession) Close() error {
	return s.sender.Close()
}

// 获取应用层实体
func (s *CSession) Entity() interface{} {
	return s.entity
}

// 设置应用层实体
func (s *CSession) SetEntity(entity interface{}) {
	s.entity = entity
}

// 获取远程地址
func (s *CSession) RemoteAddr() net.Addr {
	return s.sender.RemoteAddr()
}
