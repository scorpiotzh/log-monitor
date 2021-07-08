package elastic

import (
	"github.com/olivere/elastic/v7"
	"reflect"
	"time"
)

type LogApi struct {
	Method  string        `json:"method"`  //调用方法
	Ip      string        `json:"ip"`      //调用的IP地址
	Latency time.Duration `json:"latency"` //接口耗费时间
	LctTime time.Time     `json:"lctTime"` //调用时间点
	ErrMsg  string        `json:"err_msg"` //错误消息
	ErrNo   int           `json:"err_no"`  //错误码
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
