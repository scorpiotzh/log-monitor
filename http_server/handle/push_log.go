package handle

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log-monitor/elastic"
	"log-monitor/utils"
	"net/http"
	"time"
)

type ReqPushLog struct {
	Index   string        `json:"index"`   //索引
	Method  string        `json:"method"`  //调用方法
	Ip      string        `json:"ip"`      //调用的IP地址
	Latency time.Duration `json:"latency"` //接口耗费时间
	ErrMsg  string        `json:"err_msg"` //错误消息
	ErrNo   int           `json:"err_no"`  //错误码
}

func (l *LogHttpHandle) PushLog(ctx *gin.Context) {
	var (
		req ReqPushLog

		funcName = "PushLog"
	)
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusOK, utils.ApiRespErr(http.StatusInternalServerError, fmt.Sprintf("ShouldBindJSON err:%s", err.Error())))
		return
	}
	log.Info("LogHttpHandle:", funcName, utils.Json(&req))
	la := elastic.LogApi{
		Method:   req.Method,
		Ip:       req.Ip,
		Latency:  req.Latency,
		CallTime: time.Now(),
		LogDate:  time.Now().Format("2006-01-02"),
		ErrMsg:   req.ErrMsg,
		ErrNo:    req.ErrNo,
	}
	if err := l.ela.PushIndex(req.Index, &la); err != nil {
		log.Error("PushIndex err:", err.Error(), utils.Json(&req))
		ctx.JSON(http.StatusOK, utils.ApiRespErr(http.StatusInternalServerError, fmt.Sprintf("PushIndex err:%s", err.Error())))
		return
	}
	ctx.JSON(http.StatusOK, utils.ApiRespOK(nil))
}
