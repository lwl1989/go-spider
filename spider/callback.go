package spider

import (
	"io"
	"net/http"
	"errors"
	"sync"
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

func (call *HttpCall) Do()  {
	//
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