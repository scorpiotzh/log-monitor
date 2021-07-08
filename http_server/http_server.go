package http_server

import (
	"github.com/gin-gonic/gin"
	"log-monitor/http_server/handle"
	"log-monitor/logger"
)

type LogHttpServer struct {
	inAddress string
	internal  *gin.Engine
	h         *handle.LogHttpHandle
}

func Initialize(inAddress string) *LogHttpServer {
	return &LogHttpServer{
		inAddress: inAddress,
		internal:  gin.New(),
		h:         handle.Initialize(),
	}
}

func (d *LogHttpServer) Run() {
	d.initRouter()
	go func() {
		if err := d.internal.Run(d.inAddress); err != nil {
			logger.Error("http_server run internal err:", err)
		}
	}()
}
