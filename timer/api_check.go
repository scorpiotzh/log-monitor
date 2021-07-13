package timer

import (
	"context"
	"fmt"
	"log-monitor/config"
	"log-monitor/utils"
	"sync"
	"time"
)

func (l *LogTimer) RunApiCheck(ctx context.Context, wg *sync.WaitGroup) {
	tickerApi := time.NewTicker(time.Minute * config.Cfg.TimerServer.ApiNotifyTicker)
	wg.Add(1)
	go func() {
		for {
			select {
			case <-tickerApi.C:
				log.Info("RunApiCheck doApiCheck start ...")
				if err := l.doApiCheck(); err != nil {
					log.Error("doApiCheck err:", err)
				}
				log.Info("RunApiCheck doApiCheck end ...")
			case <-ctx.Done():
				wg.Done()
				return
			}
		}
	}()
}

func (l *LogTimer) doApiCheck() error {
	apiMap := make(map[string][]utils.ApiInfo)
	for index, methods := range config.Cfg.TimerServer.CheckIndexList {
		for _, m := range methods {
			res, err := l.ela.SearchLogApiInfo(index, m.Method, -time.Minute*config.Cfg.TimerServer.ApiNotifyCheckTime)
			if err != nil {
				return fmt.Errorf("SearchLogApiInfo err:%s [%s]", err.Error(), m.Method)
			}
			total := res.Aggregations.TotalCount.Value
			okCount := res.Aggregations.SuccessCount.DocCount
			errCount := res.Aggregations.ErrorCount.DocCount
			avgTime := time.Duration(0)
			if total > 0 {
				avgTime = time.Duration(res.Aggregations.TotalTime.Value) / time.Duration(total)
			}
			successRate := float64(1)
			log.Warnf("doApiCheck: API[%s],方法[%s],总数[%d],成功[%d],失败[%d],平均时间[%d ms]", index, m.Method, total, okCount, errCount, avgTime.Microseconds())
			if total < config.Cfg.TimerServer.ApiNotifyMinCallNum && total < m.Num {
				continue
			} else if total > 0 {
				successRate = float64(okCount) / float64(total)
			}
			if successRate > 0.9 {
				continue
			} else { //	成功率低于90%告警
				apiMap[index] = append(apiMap[index], utils.ApiInfo{
					Method:              m.Method,
					Total:               total,
					OkCount:             okCount,
					FailCount:           errCount,
					SuccessRate:         successRate,
					AverageResponseTime: avgTime,
				})
			}
		}
	}
	return utils.SendNotifyWxApiInfo(config.Cfg.TimerServer.ApiNotifyWxKey, apiMap)
}