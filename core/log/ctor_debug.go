// +build debug

package log

import "server/core/log/internal/sirupsen/logrus"

func New(module string) *Logger {
	l := &Logger{
		logrus.New().WithField("module", module),
	}
	l.Logger.Hooks.Add(NewHook())
	l.Logger.Formatter = textFormatter
	return l
}
