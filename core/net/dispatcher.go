package net

import (
	"time"

	"server/core/message/packet"
	. "server/core/net/session"
	"server/core/timer"
	"server/message"
)

// 逻辑线程, 分发用户逻辑
func dispatch(chDeserialize <-chan *packet.CRecvPacket, chClosedSession <-chan *CSession) {
	ticker := time.NewTicker(timer.Precision())
	defer func() {
		ticker.Stop()
	}()

	for {
		select {
		case p, ok := <-chDeserialize: // 已经反序列化好的包, 执行逻辑
			if !ok {
				return
			}
			handler := p.Handler()
			session := p.Session()
			payload := p.Payload()

			// 屏蔽心跳消息
			if handler.MsgID() != message.MSGID_ReqHeartbeatE && handler.MsgID() != message.MSGID_ReqKCPCheckDelayE {

				if debug {
					logger.Debugf("===> 消息ID=%s, 消息内容=%s", handler.MsgID().String(), payload.String())
				}
			}

			// 收集handler处理时间
			start := time.Now()
			if err := handler.Handle(session, payload); err != nil {
				logger.Errorf("处理消息错误, 消息ID=%s, 错误信息=%s", handler.MsgID().String(), err.Error())
			}
			handlerSpendStats.WithLabelValues(handler.MsgID().String()).Set(float64(time.Since(start).Nanoseconds()))

		case s, ok := <-chClosedSession: // 已经关闭的session, 执行回调
			if !ok {
				return
			}
			// TODO: handlers通常来说在程序启动时, 已经在各个模块的启动部分注册, 运行过程中不会注册
			// 如果以后出现在运行过程注册, 需要处理并发
			handlers := manager.handlers

			// 调用所有的session close回调
			for i := range handlers {
				h := handlers[i]
				h(s)
			}

		case <-ticker.C:
			timer.Cron()
		}
	}
}
