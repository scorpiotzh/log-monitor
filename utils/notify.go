package utils

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
	"net/http"
	"time"
)

const (
	NotifyWxTypeText     = "text"
	NotifyWxTypeMarkdown = "markdown"
	NotifyUrlWx          = "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=%s"
)

type NotifyDataWx struct {
	MsgType  string        `json:"msgtype"`
	Markdown NotifyContent `json:"markdown"`
	Text     NotifyContent `json:"text"`
}

type NotifyContent struct {
	Content string `json:"content"`
}

func SendNotifyWx(msgType, msg, key string) error {
	if key == "" {
		return nil
	}
	data := NotifyDataWx{
		MsgType:  msgType,
		Markdown: NotifyContent{},
		Text:     NotifyContent{},
	}
	content := NotifyContent{Content: msg}
	switch msgType {
	case NotifyWxTypeText:
		data.Text = content
	case NotifyWxTypeMarkdown:
		data.Markdown = content
	default:
		data.MsgType = NotifyWxTypeText
		data.Text = content
	}
	url := fmt.Sprintf(NotifyUrlWx, key)
	resp, _, errs := gorequest.New().Post(url).Timeout(time.Second * 10).SendStruct(&data).End()
	if len(errs) > 0 {
		return fmt.Errorf("errs:%v", errs)
	} else if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("http code:%d", resp.StatusCode)
	}
	return nil
}

type ApiInfo struct {
	Method              string        `json:"method"`
	Total               int           `json:"total"`
	OkCount             int           `json:"ok_count"`
	FailCount           int           `json:"fail_count"`
	SuccessRate         float64       `json:"success_rate"`
	AverageResponseTime time.Duration `json:"average_response_time"`
}

func SendNotifyWxApiInfo(key string, apiMap map[string][]ApiInfo) error {
	if len(apiMap) == 0 {
		return nil
	}
	msg := `<font color="warning">接口告警</font>
方法 | 次数 | 成功率 | 时间
`
	indexStr := `
<font color="info">%s</font>
`
	methodStr := "> %s|%d|%.f%%|%.3f ms\n"
	methodStr2 := "> %s|%d|%.f%%|%.3f s\n"
	for k, api := range apiMap {
		msg += fmt.Sprintf(indexStr, k)
		for _, m := range api {
			if m.AverageResponseTime.Seconds() > 1 {
				msg += fmt.Sprintf(methodStr2, m.Method, m.Total, m.SuccessRate*100, m.AverageResponseTime.Seconds())
			} else {
				msg += fmt.Sprintf(methodStr, m.Method, m.Total, m.SuccessRate*100, float64(m.AverageResponseTime.Microseconds())/1000)
			}
		}
	}
	return SendNotifyWx(NotifyWxTypeMarkdown, msg, key)
}
