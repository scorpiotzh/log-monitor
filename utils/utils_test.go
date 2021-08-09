package utils

import (
	"testing"
	"time"
)

func TestSendNotifyWxApiInfo(t *testing.T) {
	key := ""
	apiMap := map[string][]ApiInfo{
		"das-rpc-index": {ApiInfo{
			Method:              "register",
			Total:               100,
			OkCount:             90,
			FailCount:           10,
			SuccessRate:         0.9,
			AverageResponseTime: time.Second * 2,
		}},
		"das-pay-index": {ApiInfo{
			Method:              "pay",
			Total:               100,
			OkCount:             95,
			FailCount:           5,
			SuccessRate:         0.95,
			AverageResponseTime: time.Second * 3,
		}},
	}
	_ = SendNotifyWxApiInfo(key, 1, 1, apiMap)
}

func TestLarkNotifyBot(t *testing.T) {
	SendLarkTextNotify("", "", "aaaa")
}
