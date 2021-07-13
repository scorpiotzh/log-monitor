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
	MethodDesc          string        `json:"method_desc"`
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
接口｜总次数｜成功率｜平均耗时
`
	indexStr := `
<font color="info">%s</font>
`
	methodStr := "> %s｜%d｜%s｜%s\n"
	for k, api := range apiMap {
		msg += fmt.Sprintf(indexStr, k)
		for _, m := range api {
			successRate := ""
			if m.SuccessRate < 0.9 {
				successRate = fmt.Sprintf(`<font color="warning">%.f%%</font>`, m.SuccessRate*100)
			} else {
				successRate = fmt.Sprintf(`%.f%%`, m.SuccessRate*100)
			}
			averageResponseTime := ""
			if m.AverageResponseTime.Seconds() > 1 {
				averageResponseTime = fmt.Sprintf(`<font color="warning">%.3f s</font>`, m.AverageResponseTime.Seconds())
			} else {
				averageResponseTime = fmt.Sprintf(`%.3f ms`, float64(m.AverageResponseTime.Microseconds()/1000))
			}
			msg += fmt.Sprintf(methodStr, m.Method, m.Total, successRate, averageResponseTime)
		}
	}
	return SendNotifyWx(NotifyWxTypeMarkdown, msg, key)
}
