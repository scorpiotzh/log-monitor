package handle

import (
	"github.com/scorpiotzh/mylog"
	"log-monitor/elastic"
)

var log = mylog.NewLogger("http_handle", mylog.LevelDebug)

type LogHttpHandle struct {
	ela *elastic.Elastic
}

func Initialize(ela *elastic.Elastic) *LogHttpHandle {
	return &LogHttpHandle{ela: ela}
}
