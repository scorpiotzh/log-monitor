package handle

import (
	"github.com/gin-gonic/gin"
	"log-monitor/logger"
	"time"
)

func (l *LogHttpHandle) Version(ctx *gin.Context) {
	logger.Info("Version:", time.Now().String())
}
