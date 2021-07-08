package config

import (
	"fmt"
	"github.com/go-yaml/yaml"
	"io/ioutil"
	"os"
)

type LogConfig struct {
	HttpServer struct {
		InAddress string `json:"in_address" yaml:"in_address"`
	} `json:"http_server" yaml:"http_server"`
	ElasticServer struct {
		Url      string `json:"url" yaml:"url"`
		Username string `json:"username" yaml:"username"`
		Password string `json:"password" yaml:"password"`
		LogIndex string `json:"log_index" yaml:"log_index"`
	} `json:"elastic_server" yaml:"elastic_server"`
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
