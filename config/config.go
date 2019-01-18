package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
		panic("load config err " + path)
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