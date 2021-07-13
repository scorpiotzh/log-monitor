package elastic

import (
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
	"log-monitor/utils"
	"reflect"
	"strconv"
	"time"
)

type LogApi struct {
	Method   string        `json:"method"`    //调用方法
	Ip       string        `json:"ip"`        //调用的IP地址
	Latency  time.Duration `json:"latency"`   //接口耗费时间
	CallTime time.Time     `json:"call_time"` //调用时间点
	LogDate  string        `json:"log_date"`  //日期
	ErrMsg   string        `json:"err_msg"`   //错误消息
	ErrNo    int           `json:"err_no"`    //错误码
}

func SearchResultToLogApi(sr *elastic.SearchResult) []LogApi {
	var logApi LogApi
	var laList []LogApi
	list := sr.Each(reflect.TypeOf(logApi))
	for _, v := range list {
		if la, ok := v.(LogApi); ok {
			laList = append(laList, la)
		}
	}
	return laList
}

type ResultSearchLogApiInfo struct {
	Aggregations struct {
		ErrorCount struct {
			DocCount int `json:"doc_count"`
		} `json:"error_count"`
		SuccessCount struct {
			DocCount int `json:"doc_count"`
		} `json:"success_count"`
		TotalCount struct {
			Value int `json:"value"`
		} `json:"total_count"`
		TotalTime struct {
			Value float64 `json:"value"`
		} `json:"total_time"`
	} `json:"aggregations"`
}

// 查询接口调用次数统计
func (e *Elastic) SearchLogApiInfo(index, method string, before time.Duration) (*ResultSearchLogApiInfo, error) {
	gte := strconv.FormatInt(time.Now().Add(before).UnixNano()/1e6, 10)
	q := elastic.NewBoolQuery().Filter(
		elastic.NewRangeQuery("call_time").Gte(gte),
		elastic.NewTermQuery("method", method),
	)

	sumLatency := elastic.NewSumAggregation().Field("latency")
	countErrno := elastic.NewValueCountAggregation().Field("err_no")
	okQuery := elastic.NewTermQuery("err_no", 0)
	errQuery := elastic.NewBoolQuery().MustNot(okQuery)
	okAgg := elastic.NewFilterAggregation().Filter(okQuery).SubAggregation("success_count", countErrno)
	errAgg := elastic.NewFilterAggregation().Filter(errQuery).SubAggregation("error_count", countErrno)

	res, err := e.client.Search().Index(index).Query(q).
		Aggregation("total_time", sumLatency).
		Aggregation("total_count", countErrno).
		Aggregation("success_count", okAgg).
		Aggregation("error_count", errAgg).Size(0).Do(e.ctx)
	if err != nil {
		return nil, fmt.Errorf("SearchLogApiInfo err:%s", err.Error())
	}
	var result ResultSearchLogApiInfo
	_ = json.Unmarshal([]byte(utils.Json(&res)), &result)
	return &result, nil
}

type ResultSearchLogApiErrCount struct {
	Hits struct {
		Hits []struct {
			Source struct {
				ErrNo  int    `json:"err_no"`
				ErrMsg string `json:"err_msg"`
			} `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}

// 错误码数量统计
func (e *Elastic) SearchLogApiErrCount(index, method string, before time.Duration, size int) (*ResultSearchLogApiErrCount, error) {
	gte := strconv.FormatInt(time.Now().Add(before).UnixNano()/1e6, 10)
	q := elastic.NewBoolQuery().Filter(
		elastic.NewRangeQuery("call_time").Gte(gte),
		elastic.NewTermQuery("method", method),
	).MustNot(elastic.NewTermQuery("err_no", 0))
	res, err := e.client.Search().Index(index).Query(q).Size(100).Do(e.ctx)
	if err != nil {
		return nil, fmt.Errorf("SearchLogApiErrCount err:%s", err.Error())
	}
	var result ResultSearchLogApiErrCount
	_ = json.Unmarshal([]byte(utils.Json(&res)), &result)
	return &result, nil
}
