package http_server

import (
	"github.com/eager7/elog"
	"github.com/gin-gonic/gin"
	"log-monitor/elastic"
	"log-monitor/http_server/handle"
)

var log = elog.NewLogger("log_http_server", elog.NoticeLevel)

type LogHttpServer struct {
	inAddress string
	internal  *gin.Engine
	h         *handle.LogHttpHandle
}

func Initialize(ela *elastic.Elastic, inAddress string) *LogHttpServer {
	return &LogHttpServer{
		inAddress: inAddress,
		internal:  gin.New(),
		h:         handle.Initialize(ela),
	}
}

func (d *LogHttpServer) Run() {
	d.initRouter()
	go func() {
		if err := d.internal.Run(d.inAddress); err != nil {
			log.Error("http_server run internal err:", err)
		}
	}()
}
