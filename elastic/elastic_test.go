package elastic

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"log-monitor/utils"
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
		Latency: 200,
		CalTime: time.Now(),
		LogDate: time.Now().Format("2006-01-02"),
		ErrMsg:  "internal err",
		ErrNo:   400,
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
	fmt.Println(utils.Json(list))
}

func TestDelete(t *testing.T) {
	ela, err := Initialize(context.TODO(), Url, Username, Password)
	if err != nil {
		t.Fatal(err)
	}
	q := elastic.NewTermQuery("method", "test")
	err = ela.DeleteByQuery(Index, q)
	if err != nil {
		t.Fatal(err)
	}
}
