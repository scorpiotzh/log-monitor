package handle

import (
	"github.com/gin-gonic/gin"
	"log-monitor/config"
	"log-monitor/utils"
	"net/http"
	"time"
)

func (l *LogHttpHandle) Version(ctx *gin.Context) {
	log.Info("Version:", time.Now().String())
	ctx.JSON(http.StatusOK, utils.ApiRespOK(nil))
}

func (l *LogHttpHandle) SearchLogApiInfo(ctx *gin.Context) {
	apiMap := make(map[string][]utils.ApiInfo)
	for index, methods := range config.Cfg.TimerServer.CheckIndexList {
		for _, m := range methods {
			res, err := l.ela.SearchLogApiInfo(index, m.Method, -time.Hour*24)
			if err != nil {
				log.Error("SearchLogApiInfo err:", err.Error(), m.Method)
				ctx.JSON(http.StatusOK, utils.ApiRespErr(http.StatusInternalServerError, err.Error()))
				return
			}
			total := res.Aggregations.TotalCount.Value
			okCount := res.Aggregations.SuccessCount.DocCount
			errCount := res.Aggregations.ErrorCount.DocCount
			avgTime := time.Duration(0)
			successRate := float64(1)
			if total > 0 {
				avgTime = time.Duration(res.Aggregations.TotalTime.Value) / time.Duration(total)
				successRate = float64(okCount) / float64(total)
			}
			log.Warnf("doApiCheck: API[%s],方法[%s],总数[%d],成功[%d],失败[%d],平均时间[%d ms]", index, m.Method, total, okCount, errCount, avgTime.Microseconds())
			apiMap[index] = append(apiMap[index], utils.ApiInfo{
				Method:              m.Desc,
				Total:               total,
				OkCount:             okCount,
				FailCount:           errCount,
				SuccessRate:         successRate,
				AverageResponseTime: avgTime,
			})
		}
	}
	_ = utils.SendNotifyWxApiInfo(config.Cfg.TimerServer.ApiNotifyWxKey, apiMap)
	ctx.JSON(http.StatusOK, utils.ApiRespOK(nil))
}
