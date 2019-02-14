package config

type CallBack struct {
	Method string `json:"method"`
	Http struct{
		Url string `json:"url"`
		Method string `json:"method,omitempty"`
	}`json:"http"`
}
