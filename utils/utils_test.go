package utils

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
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
	//fmt.Println(fmt.Sprintf("%s	｜\n%s	｜","嗷嗷","啊啊啊啊"))
	url := "https://open.larksuite.com/open-apis/bot/v2/hook/b44a2fa0-a994-45ed-8964-d99d3aa77a68"
	type MsgContent struct {
		Tag      string `json:"tag"`
		UnEscape bool   `json:"un_escape"`
		Text     string `json:"text"`
	}
	type MsgData struct {
		Email   string `json:"email"`
		MsgType string `json:"msg_type"`
		Content struct {
			Post struct {
				ZhCn struct {
					Title   string         `json:"title"`
					Content [][]MsgContent `json:"content"`
				} `json:"zh_cn"`
			} `json:"post"`
		} `json:"content"`
	}
	var data MsgData
	data.Email = ""
	data.MsgType = "post"
	data.Content.Post.ZhCn.Title = "通知测试"
	data.Content.Post.ZhCn.Content = [][]MsgContent{
		{
			MsgContent{
				Tag:      "text",
				UnEscape: true,
				Text: `接口告警 (频率 - 时长：120分钟 - 240分钟)
接口	｜总次数	｜成功率	｜平均耗时
[ das-index ]
- 配置信息  	｜275  ｜100%｜0 ms
- 代币列表  	｜100  ｜100%｜0 ms
- 注册中   	｜66   ｜100%｜1 ms
- 我的奖励  	｜63   ｜100%｜9 ms
- 我的账号  	｜126  ｜100%｜6 ms
- 提现记录  	｜63   ｜100%｜8 ms
- 账户查询  	｜25   ｜100%｜3 ms
- 注册下单  	｜1    ｜100%｜16 ms
- 账户详情  	｜54   ｜100%｜6 ms
- 解析记录  	｜57   ｜100%｜6 ms
- 修改记录  	｜9    ｜56% ｜69 ms
- 签名订单  	｜4    ｜100%｜48 ms
- 查询交易  	｜56   ｜100%｜1 ms
- 可转出金额 	｜63   ｜100%｜6 ms`,
			},
		},
	}
	_, body, errs := gorequest.New().Post(url).Timeout(time.Second * 10).SendStruct(&data).End()
	if len(errs) > 0 {
		t.Fatal(errs)
	}
	fmt.Println(body)
}
