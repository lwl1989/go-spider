package http

import (
	"errors"
	"github.com/lwl1989/go-spider/config"
	"io"
	"io/ioutil"
	"net/http"
	"sync"
	"github.com/lwl1989/go-spider/spider"
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

		err := http.ListenAndServe(cf.GetServerListen(), httpHandler)
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
		successResponse(GetAllTask(), w)
		return
	}

	if r.Method == "POST" {
		_,body := initRequest(r)
		if r.URL.Path == "/rule/add" {
			task,err := spider.NewSpider(body)
			if err != nil {
				errResponse(err, w)
				return
			}
			successResponse(AddToTask(task), w)
			return
		}
		if r.URL.Path == "/rule" {
			uuid := r.FormValue("uuid")
			arr := GetAllTask()
			for _,v := range arr{
				switch t:=v.(type) {
				case *Task:
					if uuid == t.Uuid {
						successResponse(v, w)
						return
					}
				default:
					successResponse(t, w)
				}

			}
			errResponse(errors.New("Not found this uuid : "+uuid), w)
			return
		}
		if r.URL.Path == "/rule/stop" {
			uuid := r.FormValue("uuid")
			taskJob.StopOnce(uuid)
			successResponse("stop success", w)
			return
		}

	}
	errResponse(errors.New("Not allowed this method! "), w)
}


func successResponse(data interface{}, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(InterfaceToJson(data))
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