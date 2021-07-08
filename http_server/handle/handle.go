package handle

import "log-monitor/elastic"

type LogHttpHandle struct {
	ela *elastic.Elastic
}

func Initialize(ela *elastic.Elastic) *LogHttpHandle {
	return &LogHttpHandle{ela: ela}
}
