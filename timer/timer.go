package timer

import (
	"github.com/robfig/cron"
	"github.com/scorpiotzh/mylog"
	"log-monitor/elastic"
)

var log = mylog.NewLogger("timer", mylog.LevelDebug)

type LogTimer struct {
	task *cron.Cron
	ela  *elastic.Elastic
}

func Initialize(ela *elastic.Elastic) *LogTimer {
	return &LogTimer{task: cron.New(), ela: ela}
}
