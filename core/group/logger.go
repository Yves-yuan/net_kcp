package group

import "server/core/log"

var logger = log.New("group")

func init() {
	logger.SetLevel(log.DebugLevel)
}
