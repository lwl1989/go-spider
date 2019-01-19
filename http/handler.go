package http

import (
	"github.com/lwl1989/go-spider/config"
	"net/http"
	"sync"
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


}