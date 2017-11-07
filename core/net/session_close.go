/*
  作    者: Long
  版    本: v1
  完成日期: 2017-07-14
  功能描述: 部分模块需要关注session关闭, 关闭session时, 执行注册的回调函数
*/
package net

import (
	"sync"

	"server/core/net/session"
)

type SessionClosedHandler func(session *session.CSession)

// 管理器单例, 管理所有的回调函数
var manager = struct {
	sync.RWMutex
	handlers []SessionClosedHandler
}{}

// TODO(warning): 由于golang中不能直接比较函数的相等性, 所以这里可能重复添加
// TODO(notice):
// handlers通常来说在程序启动时, 已经在各个模块的启动部分注册, 运行过程中不会注册
// 如果以后出现在运行过程注册, 需要处理并发
func OnSessionClosed(handler SessionClosedHandler) {
	manager.Lock()
	defer manager.Unlock()

	manager.handlers = append(manager.handlers, handler)
}
