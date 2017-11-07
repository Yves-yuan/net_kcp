package timer

import "server/core/log"

var logger = log.New("timer")

func init() {
	logger.SetLevel(log.InfoLevel)
}
