package handle

import (
	"github.com/gin-gonic/gin"
	"time"
)

func (l *LogHttpHandle) Version(ctx *gin.Context) {
	log.Info("Version:", time.Now().String())
}
