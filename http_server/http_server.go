package http_server

import (
	"github.com/gin-gonic/gin"
	"github.com/scorpiotzh/mylog"
	"log-monitor/elastic"
	"log-monitor/http_server/handle"
)

var log = mylog.NewLogger("http_server", mylog.LevelDebug)

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
