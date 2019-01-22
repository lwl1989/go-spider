package http

import (
	"github.com/lwl1989/go-spider/spider"
	"sync"
)

type Task struct {
	Rule 	*spider.Rule `json:"rule"`
	Name    string `json:"name"`
	hash	string
	Uuid    string `json:"uuid"`
	RunTime int64 `json:"run_time"`
	Spacing int64 `json:"spacing_time"`
	EndTime int64 `json:"end_time"`
}
type Tasks map[string]*Task
var tasks Tasks
var initTaskOnce sync.Once

func InitTask() {
	initTaskOnce.Do(func(){
		tasks = make(Tasks)
	})
}
func GetAllTask() map[string]*Task {
	return tasks
}

func (tasks Tasks) addTask(task *Task) {
	hash := task.GetHash()
	if _, ok := tasks[hash]; ok {
		tasks[hash] = task //覆盖
	}else{
		tasks[hash] = task
	}
}


func (task *Task) GetHash() string {
	if task.hash == "" {
		task.hash = task.Name
	}

	return task.hash
}
