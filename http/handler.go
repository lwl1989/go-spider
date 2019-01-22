package http

import (
	"github.com/lwl1989/go-spider/config"
	"net/http"
	"sync"
	"net/url"
	"io"
	"io/ioutil"
)

type handler struct {
	conf  *config.Config
}

var httpHandler *handler
var handlerOnce sync.Once

//init http server and init request handler
func GetHandler(cf *config.Config) *handler {
	handlerOnce.Do(func() {
		httpHandler = &handler{
			conf:cf,
		}

		err := http.ListenAndServe(cf.GetServerPort(), httpHandler)
		if err != nil {
			panic(err)
		}
	})
	return httpHandler
}

//handler http request
//do any actions
func (handler *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//獲取所有正在執行的任務
	if r.URL.Path == "/tasks" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(InterfaceToJson(GetAllTask()))
		return
	}

	if r.Method == "POST" {
		_,body := initRequest(r)
		if url.URL.Path == "addRule" {
			task,err := newTask(body)
			if err != nil {
				errResponse(err, w)
				return
			}

		}


	}

}

func errResponse(err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(InterfaceToJson(err))
}
func initRequest(r *http.Request) (contentType string, body []byte) {
	contentType = r.Header.Get("Content-Type")
	var reader io.Reader = r.Body

	maxFormSize := int64(2 << 20) // 2 MB is a lot of text.
	reader = io.LimitReader(r.Body, maxFormSize+1)

	b, e := ioutil.ReadAll(reader)


	if e != nil {
		return contentType, nil
	}
	return contentType, b
}