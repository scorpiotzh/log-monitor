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
	Url      = "http://127.0.0.1:9200/"
	Username = "elastic"
	Password = "654321"
	Index    = "test-index"
)

func TestPushIndex(t *testing.T) {
	ela, err := Initialize(context.TODO(), Url, Username, Password)
	if err != nil {
		t.Fatal(err)
	}
	la := LogApi{
		Method:  "test",
		Ip:      "127.0.0.1",
		Latency: time.Second * 1,
		CalTime: time.Now(),
		LogDate: time.Now().Format("2006-01-02"),
		ErrMsg:  "",
		ErrNo:   0,
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

func TestSearch(t *testing.T) {
	ela, err := Initialize(context.TODO(), Url, Username, Password)
	if err != nil {
		t.Fatal(err)
	}
	res, err := ela.SearchLogApiInfo(Index, "test1", -time.Hour*24)
	if err != nil {
		t.Fatal(err)
	} else {
		fmt.Println(utils.Json(&res))
	}
	if res2, err := ela.SearchLogApiErrCount(Index, "test1", -time.Hour, res.Aggregations.ErrorCount.DocCount); err != nil {
		t.Fatal(err)
	} else {
		fmt.Println(utils.Json(&res2))
	}
}
