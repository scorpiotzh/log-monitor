package utils

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
	"github.com/scorpiotzh/mylog"
	"time"
)

var (
	log = mylog.NewLogger("lark_notify", mylog.LevelDebug)
)

const (
	LarkNotifyUrl = "https://open.larksuite.com/open-apis/bot/v2/hook/%s"
)

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

func SendLarkTextNotify(key, title, text string) {
	if key == "" || text == "" {
		return
	}
	var data MsgData
	data.Email = ""
	data.MsgType = "post"
	data.Content.Post.ZhCn.Title = title
	data.Content.Post.ZhCn.Content = [][]MsgContent{
		{
			MsgContent{
				Tag:      "text",
				UnEscape: false,
				Text:     text,
			},
		},
	}
	url := fmt.Sprintf(LarkNotifyUrl, key)
	_, body, errs := gorequest.New().Post(url).Timeout(time.Second * 10).SendStruct(&data).End()
	if len(errs) > 0 {
		log.Error("sendLarkTextNotify req err:", errs)
	} else {
		log.Info("sendLarkTextNotify req:", body)
	}
}

func GetSendLarkNotifyApiInfoStr(rate, duration time.Duration, apiMap map[string][]ApiInfo) string {
	msg := fmt.Sprintf(`接口监控（频率 - 时长：%d分钟 - %d分钟）
接口	｜总次数	｜成功率	｜平均耗时
`, rate, duration)
	indexStr := `[ %s ]
`
	methodStr := `- %-6s	｜%-6d	｜%-6s	｜%s
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
