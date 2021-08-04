package handle

import (
	"fmt"
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
			res, err := l.ela.SearchLogApiInfo(index, m.Method, -time.Minute*config.Cfg.TimerServer.ApiNotifyCheckAllTime)
			if err != nil {
				ctx.JSON(http.StatusOK, utils.ApiRespErr(500, fmt.Sprintf("SearchLogApiInfo err:%s [%s]", err.Error(), m.Method)))
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
			} else {
				continue
			}
			log.Warnf("doApiCheck: API[%s],方法[%s],总数[%d],成功[%d],失败[%d],平均时间[%.3g s]", index, m.Method, total, okCount, errCount, avgTime.Seconds())
			apiMap[index] = append(apiMap[index], utils.ApiInfo{
				Method:              m.Method,
				MethodDesc:          m.Desc,
				Total:               total,
				OkCount:             okCount,
				FailCount:           errCount,
				SuccessRate:         successRate,
				AverageResponseTime: avgTime,
			})
		}
	}
	//_ = utils.SendNotifyWxApiInfo(config.Cfg.TimerServer.ApiNotifyWxKey, 1, 1, apiMap)
	ctx.String(http.StatusOK, getNotifyStr(config.Cfg.TimerServer.ApiNotifyAllTicker, config.Cfg.TimerServer.ApiNotifyCheckAllTime, apiMap))
}

func getNotifyStr(rate, duration time.Duration, apiMap map[string][]utils.ApiInfo) string {
	msg := fmt.Sprintf(`接口告警 (频率 - 时长：%d分钟 - %d分钟)
接口	｜总次数	｜成功率	｜平均耗时
`, rate, duration)
	indexStr := `[ %s ]
`
	methodStr := `- %10s	｜%5d	｜%4s	｜%s
`
	for k, api := range apiMap {
		msg += fmt.Sprintf(indexStr, k)
		for _, m := range api {
			successRate := fmt.Sprintf("%.f%%", m.SuccessRate*100)
			averageResponseTime := ""
			if m.AverageResponseTime.Seconds() > 1 {
				averageResponseTime = fmt.Sprintf("%.3g s", m.AverageResponseTime.Seconds())
			} else {
				averageResponseTime = fmt.Sprintf("%.3g ms", float64(m.AverageResponseTime.Microseconds()/1000))
			}
			msg += fmt.Sprintf(methodStr, m.MethodDesc, m.Total, successRate, averageResponseTime)
		}
	}
	return msg
}

func (l *LogHttpHandle) DelLog(ctx *gin.Context) {
	_ = l.ela.Purge("das2-index")
	ctx.JSON(http.StatusOK, utils.ApiRespOK(nil))
}
