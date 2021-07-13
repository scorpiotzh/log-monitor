package handle

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"log-monitor/config"
	"log-monitor/utils"
	"net/http"
	"strconv"
	"time"
)

func (l *LogHttpHandle) Version(ctx *gin.Context) {
	log.Info("Version:", time.Now().String())

	gte := strconv.FormatInt(time.Now().Add(-time.Hour).UnixNano()/1e6, 10)
	fmt.Println(time.Now().Add(-time.Hour), gte)
	q := elastic.NewBoolQuery().Filter(
		elastic.NewRangeQuery("call_time").Gte(gte),
	)
	res, err := l.ela.Client().Search().Index("das2-index").Query(q).Size(100).Do(context.TODO())
	if err != nil {
		log.Error(err)
	}
	log.Info("Version:", utils.Json(&res))

	ctx.JSON(http.StatusOK, utils.ApiRespOK(nil))
}

func (l *LogHttpHandle) SearchLogApiInfo(ctx *gin.Context) {
	res, _ := l.ela.Traverse("das2-index")
	//list := elastic.SearchResultToLogApi(res)
	log.Info("Traverse:", utils.Json(&res))

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

func (l *LogHttpHandle) DelLog(ctx *gin.Context) {
	_ = l.ela.Purge("das2-index")
	ctx.JSON(http.StatusOK, utils.ApiRespOK(nil))
}
