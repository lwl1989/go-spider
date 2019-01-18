package spider

import (
	"github.com/lwl1989/logger"
	"sync"
)

var spiderLog *logger.TTLog
var logOnce sync.Once

func GetLog() *logger.TTLog  {
	logOnce.Do(func() {
		spiderLog = logger.GetFileLogger(Cf.LogConfig)
	})
	return spiderLog
}