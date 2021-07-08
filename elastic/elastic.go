package elastic

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"strconv"
	"strings"
	"time"
)

type Elastic struct {
	ctx    context.Context
	client *elastic.Client
}

func Initialize(ctx context.Context, url, username, password string) (*Elastic, error) {
	client, err := elastic.NewClient(elastic.SetURL(url), elastic.SetSniff(false), elastic.SetBasicAuth(username, password))
	if err != nil {
		return nil, fmt.Errorf("es init client err:%s", err.Error())
	}
	resp, err := client.NodesInfo().Do(ctx)
	if err == nil {
		fmt.Println("cluster name:", resp.ClusterName)
		for k, node := range resp.Nodes {
			fmt.Println("id:", k, "name:", node.Name, "host:", node.Host, "version:", node.Version)
		}
	}
	return &Elastic{ctx: ctx, client: client}, nil
}

// 插入
func (e *Elastic) PushIndex(index string, body interface{}) error {
	id := strconv.Itoa(int(time.Now().UnixNano())) //相当于主键
	_, err := e.client.Index().Index(index).Id(id).BodyJson(body).Do(e.ctx)
	if err != nil {
		return fmt.Errorf("es push err:%s", err.Error())
	}
	return nil
}

// 删除
func (e *Elastic) Purge(index string) error {
	_, err := e.client.DeleteIndex(index).Do(e.ctx)
	if err != nil {
		return fmt.Errorf("es delete index err:%s", err.Error())
	}
	return nil
}

// 分页读取
func (e *Elastic) DocList(index string, page, size int) (*elastic.SearchResult, error) {
	res, err := e.client.Search(index).From((page - 1) * size).Size(size).Do(e.ctx)
	if err != nil {
		if strings.Contains(err.Error(), "type=index_not_found_exception") {
			fmt.Println("empty es")
			return nil, nil
		}
		return nil, fmt.Errorf("search err:%s", err.Error())
	}
	return res, nil
}

// 遍历
func (e *Elastic) Traverse(index string) (*elastic.SearchResult, error) {
	res, err := e.client.Search(index).Do(e.ctx)
	if err != nil {
		if strings.Contains(err.Error(), "type=index_not_found_exception") {
			fmt.Println("empty es")
			return nil, nil
		}
		return nil, fmt.Errorf("search err:%s", err.Error())
	}
	return res, nil
}
