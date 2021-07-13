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
	go func() {
		if err := l.doApiCheckAll(); err != nil {
			log.Error("doApiCheckAll err:", err)
		}
	}()
	tickerApi := time.NewTicker(time.Minute * config.Cfg.TimerServer.ApiNotifyTicker)
	tickerAll := time.NewTicker(time.Minute * config.Cfg.TimerServer.ApiNotifyAllTicker)
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
			case <-tickerAll.C:
				log.Info("RunApiCheck doApiCheckAll start ...")
				if err := l.doApiCheckAll(); err != nil {
					log.Error("doApiCheckAll err:", err)
				}
				log.Info("RunApiCheck doApiCheckAll end ...")
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
			successRate := float64(1)
			if total > 0 {
				avgTime = time.Duration(res.Aggregations.TotalTime.Value) / time.Duration(total)
				successRate = float64(okCount) / float64(total)
			}
			log.Warnf("doApiCheck: API[%s],方法[%s],总数[%d],成功[%d],失败[%d],平均时间[%.3g s]", index, m.Method, total, okCount, errCount, avgTime.Seconds())
			if total < config.Cfg.TimerServer.ApiNotifyMinCallNum && total < m.Num {
				continue
			}
			if successRate > 0.9 {
				continue
			} else { //	成功率低于90%告警
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
	}
	return utils.SendNotifyWxApiInfo(config.Cfg.TimerServer.ApiNotifyWxKey, config.Cfg.TimerServer.ApiNotifyTicker, config.Cfg.TimerServer.ApiNotifyCheckTime, apiMap)
}

func (l *LogTimer) doApiCheckAll() error {
	apiMap := make(map[string][]utils.ApiInfo)
	for index, methods := range config.Cfg.TimerServer.CheckIndexList {
		for _, m := range methods {
			res, err := l.ela.SearchLogApiInfo(index, m.Method, -time.Minute*config.Cfg.TimerServer.ApiNotifyCheckAllTime)
			if err != nil {
				return fmt.Errorf("SearchLogApiInfo err:%s [%s]", err.Error(), m.Method)
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
	return utils.SendNotifyWxApiInfo(config.Cfg.TimerServer.ApiNotifyWxKey, config.Cfg.TimerServer.ApiNotifyAllTicker, config.Cfg.TimerServer.ApiNotifyCheckAllTime, apiMap)
}
