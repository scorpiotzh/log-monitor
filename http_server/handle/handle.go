package handle

import (
	"github.com/eager7/elog"
	"log-monitor/elastic"
)

var log = elog.NewLogger("log_http_handle", elog.NoticeLevel)

type LogHttpHandle struct {
	ela *elastic.Elastic
}

func Initialize(ela *elastic.Elastic) *LogHttpHandle {
	return &LogHttpHandle{ela: ela}
}
