package config

import (
	"fmt"
	"github.com/go-yaml/yaml"
	"io/ioutil"
	"os"
	"time"
)

type LogConfig struct {
	HttpServer struct {
		InAddress string `json:"in_address" yaml:"in_address"`
	} `json:"http_server" yaml:"http_server"`
	ElasticServer struct {
		Url      string `json:"url" yaml:"url"`
		Username string `json:"username" yaml:"username"`
		Password string `json:"password" yaml:"password"`
	} `json:"elastic_server" yaml:"elastic_server"`
	TimerServer struct {
		ApiNotifyMinCallNum   int           `json:"api_notify_min_call_num" yaml:"api_notify_min_call_num"`
		ApiNotifyTicker       time.Duration `json:"api_notify_ticker" yaml:"api_notify_ticker"`
		ApiNotifyCheckTime    time.Duration `json:"api_notify_check_time" yaml:"api_notify_check_time"`
		ApiNotifyAllTicker    time.Duration `json:"api_notify_all_ticker" yaml:"api_notify_all_ticker"`
		ApiNotifyCheckAllTime time.Duration `json:"api_notify_check_all_time" yaml:"api_notify_check_all_time"`
		ApiNotifyWxKey        string        `json:"api_notify_wx_key" yaml:"api_notify_wx_key"`
		DeleteIndexList       []string      `json:"delete_index_list" yaml:"delete_index_list"`
		CheckIndexList        map[string][]struct {
			Method string `json:"method" yaml:"method"`
			Desc   string `json:"desc" yaml:"desc"`
			Num    int    `json:"num" yaml:"num"`
		} `json:"check_index_list" yaml:"check_index_list"`
	} `json:"timer_server" yaml:"timer_server"`
}

var Cfg LogConfig

func InitCfgFromFile(filepath string, receiver interface{}) error {
	file, err := os.OpenFile(filepath, os.O_RDONLY, 0666)
	if err != nil {
		return fmt.Errorf("OpenFile err:%s", err.Error())
	}
	bys, err := ioutil.ReadAll(file)
	if err != nil {
		return fmt.Errorf("ReadAll err:%s", err.Error())
	}
	if err = yaml.Unmarshal(bys, receiver); err != nil {
		return fmt.Errorf("Unmarshal err:%s", err.Error())
	}
	return nil
}
