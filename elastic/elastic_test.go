package elastic

import (
	"context"
	"fmt"
	"testing"
	"time"
)

const (
	Url      = "http://localhost:9200/"
	Username = ""
	Password = ""
	Index    = "test"
)

func TestPushIndex(t *testing.T) {
	ela, err := Initialize(context.TODO(), Url, Username, Password)
	if err != nil {
		t.Fatal(err)
	}
	la := LogApi{
		Method:  "test",
		Ip:      "127.0.0.1",
		Latency: 100,
		LctTime: time.Now(),
		ErrMsg:  "internal err",
		ErrNo:   500,
	}
	err = ela.PushIndex(Index, la)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDocList(t *testing.T) {
	ela, err := Initialize(context.TODO(), Url, Username, Password)
	if err != nil {
		t.Fatal(err)
	}
	res, err := ela.DocList(Index, 1, 10)
	if err != nil {
		t.Fatal(err)
	}
	list := SearchResultToLogApi(res)
	fmt.Println(list)
}
