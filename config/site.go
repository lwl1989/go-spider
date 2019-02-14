package config

type HttpConfig struct {
	Port int `json:"port"`
	Proxy string `json:"proxy"`
	HttpProxy string `json:"http_proxy"`

	CallBack string `json:"call_back"`
}