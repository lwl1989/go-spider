package spider

import (
	"io"
	"net/http"
	"errors"
	"sync"
	"time"
	"encoding/json"
	"bytes"
	"log"
)

//this interface must use after call data to other way
//if get error , other may be empty
type ICallResponse interface{
	GetResponseBody() io.Reader
	String() string
	GetResponseStatus() string
}
type CallResponse func(response *ICallResponse)
type CallResponseError func(error error)
//call the other sys, and get a response
type ICall interface{
	Do()
	OnResponse(response CallResponse)
	OnError() (err CallResponseError)
}

type HttpCall struct{
	Url string `json:"url"`
	Method string `json:"method,omitempty"`
	Params interface{} `json:"params"`
	lock   *sync.RWMutex
	response *ICallResponse
}

type HttpCallResponse struct {
	Response *http.Response
	Error error
}

type RedisCall struct {
	Params interface{} `json:"params"`
}

type RedisResponse struct {
	Msg string `json:"msg"`
}

func (r *RedisCall) Do() {
	cf := Cf.Connections
	c := GetRedis(cf)
	buf := bytes.NewBufferString("")
	encoder := json.NewEncoder(buf)
	encoder.SetEscapeHTML(false)

	if err := encoder.Encode(&r.Params); err != nil {
		c.LPush("MofNews", err.Error())
	} else {
		c.LPush("MofNews", buf.String())
	}
	//values,err := json.Marshal(r.Params)
	//if err != nil {
	//
	//}else{
	//
	//}
}

func (call *HttpCall) Do()  {
	//
	c := http.DefaultClient
	c.Timeout = time.Duration(20 * time.Second)

	method := call.Method
	switch method {
		case "GET":
			c.Get(call.Url)
		case "POST":
			fallthrough
		default:
			data,err := json.Marshal(call.Params)
			if err == nil {
				_,err := c.Post(call.Url, "application/json", bytes.NewBuffer(data[:]))
				if err != nil {
					log.Panicln("error")
				}
			}else{
				log.Panicln("error")
			}
	}
}

func (call *HttpCall) OnResponse(response CallResponse)  {
	call.lock.Lock()
	defer call.lock.Unlock()
}

func (call *HttpCall) OnError(err CallResponseError)   {
	call.lock.Lock()
	defer call.lock.Unlock()
}

func (response *HttpCallResponse) GetResponseBody() io.Reader {
	return response.Response.Body
}

func (response *HttpCallResponse) String() string {
	if response.Response == nil || response.Response.Body == nil {
		response.Error = errors.New("no body received")
		return response.Error.Error()
	}
	defer response.Response.Body.Close()
	content := make([]byte,0)
	for ; ;  {
		b := make([]byte,1024)
		n,err := response.Response.Body.Read(b)
		if err != nil {
			if err == io.EOF {
				content = append(content, b[:n]...)
				break
			}else {
				return  err.Error()
			}
		}
		if n > 0 {
			content = append(content, b[:n]...)
		}
	}

	return string(content[:])
}

func (response *HttpCallResponse) GetStatus() string {
	return response.Response.Status
}

func (response *HttpCallResponse) GetError() error {
	return response.Error
}