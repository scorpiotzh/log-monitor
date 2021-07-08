package timer

import (
	"github.com/olivere/elastic/v7"
	"log-monitor/logger"
	"time"
)

func (l *LogTimer) DeleteByLogDate(logIndex string) {
	if err := l.task.AddFunc("0 0 8 1/1 * ?", func() {
		logDate := time.Now().AddDate(0, 0, -3).Format("2006-01-02")
		q := elastic.NewTermQuery("log_date", logDate)
		if err := l.ela.DeleteByQuery(logIndex, q); err != nil {
			logger.Error("DeleteByQuery err:", err.Error(), logDate)
		}
	}); err != nil {
		logger.Error("AddFunc err:", err.Error())
	}
	l.task.Start()
}
