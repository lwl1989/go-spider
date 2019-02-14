package http

import (
	"github.com/lwl1989/go-spider/spider"
	"sync"
	"github.com/lwl1989/timing"
	"time"
	"fmt"
	"reflect"
)

type Task struct {
	Job  	*spider.CollySpider `json:"job"`
	Uuid    string `json:"uuid"`
	RunTime int64 `json:"run_time"`
	Spacing int64 `json:"spacing_time"`
	EndTime int64 `json:"end_time"`
	Number  int	  `json:"number"`
}

type TaskTimes   struct {
	RunTime int64  `json:"run_time"`
	Spacing int64  `json:"spacing_time"`
	EndTime int64  `json:"end_time,omitempty"`
	Number  int  `json:"number,omitempty"`
}
var taskJob *timing.TaskScheduler
var initTaskOnce sync.Once

func InitTask() {
	initTaskOnce.Do(func(){
		taskJob = timing.NewScheduler()
		taskJob.Start()
	})
}

func GetAllTask() []interface{} {
	tasks := taskJob.Export()
	list := make([]interface{}, 0)
	if len(tasks) > 1 {
		for _,t := range tasks{
			t1 := reflect.TypeOf(t.Job)
			fmt.Println(t1)
			switch ty := t.Job.(type) {
			case *spider.CollySpider:
			default:
				fmt.Println(ty)
			}
			msg,ok := t.Job.(*spider.CollySpider)
			if !ok {
				fmt.Println("err change")
			}else {
				list = append(list, &spider.CollySpider{
					Rule: msg.Rule,
					Times: msg.Times,
				})
			}
		}
	}
	return list
}

func AddToTask(spider *spider.CollySpider) string {
	job := taskJob
	
	runTime := spider.Times.RunTime
	if runTime == 0 && spider.Times.Spacing > 0 {
		runTime = time.Now().UnixNano()+ spider.Times.Spacing * int64(time.Second)
	} else if runTime < time.Now().Unix() {
		runTime = (time.Now().Unix() + 1) * int64(time.Second)
	}
	uuid := job.AddTask(&timing.Task{
		Job:     spider,
		RunTime: runTime,
		Spacing: spider.Times.Spacing,
		EndTime: spider.Times.EndTime,
		Number: spider.Times.Number,
	})
	return uuid
}