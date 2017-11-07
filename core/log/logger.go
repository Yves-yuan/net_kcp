package log

import (
	"strings"

	"server/core/log/internal/sirupsen/logrus"
)

type Logger struct {
	*logrus.Entry
}

// 日志等级
const (
	PanicLevel = "panic"
	FatalLevel = "fatal"
	ErrorLevel = "error"
	InfoLevel  = "info"
	DebugLevel = "debug"
	WarnLevel  = "warn"
)

var (
	textFormatter = &logrus.TextFormatter{TimestampFormat: "2006-01-02T15:04:05.000Z07:00"}
)

// 设置日志输出级别
func SetLevel(level string) {
	level = strings.TrimSpace(level)
	if lvl, err := logrus.ParseLevel(level); err != nil {
		logrus.Error(err.Error())
	} else {
		logrus.SetLevel(lvl)
	}
}

// 设置日志输出格式, 可选text/json
func SetFormat(format string) {
	format = strings.TrimSpace(format)
	if format != "text" && format != "json" {
		logrus.Errorf("不可识别的日志输出格式, 只支持json/text, 收到: %s", format)
		return
	}

	// 默认就是text输出格式
	if format == "json" {
		logrus.SetFormatter(&logrus.JSONFormatter{})
	}
}

func (l *Logger) Enable(enable bool) {
	if !enable {
		l.Logger.Level = logrus.ErrorLevel
	} else {
		l.Level = logrus.GetLevel()
	}
}

func (l *Logger) SetLevel(level string) {
	level = strings.TrimSpace(level)
	if lvl, err := logrus.ParseLevel(level); err != nil {
		logrus.Error(err.Error())
	} else {
		l.Logger.Level = lvl
	}
}
