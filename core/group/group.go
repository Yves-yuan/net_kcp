package group

import (
	"errors"
	"sync"
	"sync/atomic"

	"server/core/net/session"
	"server/message"

	"github.com/golang/protobuf/proto"
)

const (
	statusWorking = 0
	statusClosed  = 1
)

var (
	ErrCloseClosedGroup   = errors.New("close closed group")
	ErrClosedGroup        = errors.New("group closed")
	ErrSessionDuplication = errors.New("session has existed in the current group")
)

// SessionFilter represents a filter which was used to filter session when Multicast,
// the session will receive the message while filter returns true.
type SessionFilter func(*session.CSession) bool

// CGroup represents a session group which used to manage a number of
// sessions, data send to the group will send to all session in it.
type CGroup struct {
	sync.RWMutex
	status   int32                       // channel current status
	name     string                      // channel name
	sessions map[int64]*session.CSession // session id map to session instance
}

// NewGroup returns a new group instance
func New(name string) *CGroup {
	return &CGroup{
		status:   statusWorking,
		name:     name,
		sessions: make(map[int64]*session.CSession),
	}
}

// Push message to partial client, which filter return true
func (c *CGroup) Multicast(msgid message.MSGID, v proto.Message, filter SessionFilter) error {
	if c.isClosed() {
		return ErrClosedGroup
	}

	if debug {
		println("++++>", msgid.String(), v.String())
	}
	// 序列化
	data, err := proto.Marshal(v)
	if err != nil {
		logger.Error(err)
		return err
	}

	c.RLock()
	defer c.RUnlock()

	for _, s := range c.sessions {
		if !filter(s) {
			continue
		}
		if err := s.Send(msgid, data); err != nil {
			logger.Error(err.Error())
		}
	}

	return nil
}

// Push message to all client
func (c *CGroup) Broadcast(msgid message.MSGID, v proto.Message) error {
	if c.isClosed() {
		return ErrClosedGroup
	}

	if debug {
		println("++++>", msgid.String(), v.String())
	}
	// 序列化
	data, err := proto.Marshal(v)
	if err != nil {
		logger.Error(err)
		return err
	}

	c.RLock()
	defer c.RUnlock()

	for _, s := range c.sessions {
		if err := s.Send(msgid, data); err != nil {
			logger.Warnf("CGroup.Broadcast: %s", err.Error())
		}
	}

	return nil
}

// Add add session to group
func (c *CGroup) Add(session *session.CSession) error {
	if c.isClosed() {
		return ErrClosedGroup
	}

	c.Lock()
	defer c.Unlock()

	id := session.ID()
	_, ok := c.sessions[session.ID()]
	if ok {
		return ErrSessionDuplication
	}

	c.sessions[id] = session
	return nil
}

// Leave remove specified UID related session from group
func (c *CGroup) Leave(session *session.CSession) error {
	if c.isClosed() {
		return ErrClosedGroup
	}

	c.Lock()
	defer c.Unlock()

	delete(c.sessions, session.ID())
	return nil
}

// LeaveAll clear all sessions in the group
func (c *CGroup) LeaveAll() error {
	if c.isClosed() {
		return ErrClosedGroup
	}

	c.sessions = make(map[int64]*session.CSession)
	return nil
}

// Count get current member amount in the group
func (c *CGroup) Count() int {
	c.RLock()
	defer c.RUnlock()

	return len(c.sessions)
}

func (c *CGroup) isClosed() bool {
	if atomic.LoadInt32(&c.status) == statusClosed {
		return true
	}
	return false
}

// Close destroy group, which will release all resource in the group
func (c *CGroup) Close() error {
	if c.isClosed() {
		return ErrCloseClosedGroup
	}

	c.LeaveAll()
	atomic.StoreInt32(&c.status, statusClosed)
	return nil
}
