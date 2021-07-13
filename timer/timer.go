package timer

import (
	"github.com/eager7/elog"
	"github.com/robfig/cron"
	"log-monitor/elastic"
)

var log = elog.NewLogger("log_timer", elog.NoticeLevel)

type LogTimer struct {
	task *cron.Cron
	ela  *elastic.Elastic
}

func Initialize(ela *elastic.Elastic) *LogTimer {
	return &LogTimer{task: cron.New(), ela: ela}
}
