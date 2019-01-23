package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
)
type Config struct {
	HttpConfig *HttpConfig	`json:"http_config"`
	Connections *Connections `json:"connections"`

	LogConfig *LogFileConfig `json:"log_config"`
}


func GetConfig(path string) *Config {
	fmt.Println("load:", path)
	data, err := ioutil.ReadFile(path)
	if err != nil {
		data, err = ioutil.ReadFile("config.json")
		if err != nil {
			panic("load config err " + path)
		}
	}
	b := []byte(data)
	config := getConfig()
	err = json.Unmarshal(b, config)
	if err != nil {
		panic("load config err " + path)
	}
	//if config.SaveSpace == 0 {
	//	config.SaveSpace = 600
	//}
	return config
}

func getConfig() *Config  {
	return &Config{
		Connections: &Connections{

		},
		LogConfig: &LogFileConfig{
			FilePath: "/tmp/spider_log",
		},
	}
}

func (conf *Config) GetServerListen() string {
	return "0.0.0.0:"+strconv.FormatInt(int64(conf.HttpConfig.Port), 10)
}