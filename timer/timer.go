package timer

import (
	"github.com/robfig/cron"
	"log-monitor/elastic"
)

type LogTimer struct {
	task *cron.Cron
	ela  *elastic.Elastic
}

func Initialize(ela *elastic.Elastic) *LogTimer {
	return &LogTimer{task: cron.New(), ela: ela}
}
