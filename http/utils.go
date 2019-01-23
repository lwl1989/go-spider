package http

import "encoding/json"

type httpError struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

func InterfaceToJson(any interface{}) []byte {
	var b []byte
	var err error
	switch any.(type) {
		case error:{
			b,_ = json.Marshal(GetHttpErrorWithError(any.(error), "500"))
		}
	default:
		b,err = json.Marshal(any)
		if err != nil {
			b,_ = json.Marshal(GetHttpErrorWithError(err, "500"))
			return b
		}
	}
	return b
}

func GetHttpErrorWithError(err error, code string) *httpError {
	return &httpError{
		Code:code,
		Msg:err.Error(),
	}
}
