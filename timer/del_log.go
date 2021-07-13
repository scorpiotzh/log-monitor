package timer

import (
	"github.com/olivere/elastic/v7"
	"time"
)

func (l *LogTimer) RunDeleteLogByLogDate(indexList []string) {
	if err := l.task.AddFunc("0 0 8 1/1 * ?", func() {
		logDate := time.Now().AddDate(0, 0, -3).Format("2006-01-02")
		q := elastic.NewTermQuery("log_date", logDate)
		log.Info("DeleteByLogDate:", logDate, indexList)
		for _, v := range indexList {
			if err := l.ela.DeleteByQuery(v, q); err != nil {
				log.Error("DeleteByQuery err:", err.Error(), logDate)
			}
		}
	}); err != nil {
		log.Error("AddFunc err:", err.Error())
	}
	l.task.Start()
}
