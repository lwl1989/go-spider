package http

import (
	"github.com/lwl1989/go-spider/spider"
	"sync"
	"encoding/json"
	"github.com/lwl1989/timing"
	"time"
)

type Task struct {
	Rule 	*spider.Rule `json:"rule"`
	Name    string `json:"name"`
	hash	string
	Times   *TaskTimes  `json:"times"`
}
type TaskTimes   struct {
	Uuid    string `json:"uuid"`
	RunTime int64  `json:"run_time"`
	Spacing int64  `json:"spacing_time"`
	EndTime int64  `json:"end_time"`
}
type Tasks map[string]*Task
var tasks Tasks
var taskJob *timing.TaskScheduler
var initTaskOnce sync.Once

func InitTask() {
	initTaskOnce.Do(func(){
		tasks = make(Tasks)
		taskJob = timing.NewScheduler()
		taskJob.Start()
	})
}
func GetAllTask() map[string]*Task {
	return tasks
}
func newTask(content []byte) (*Task,error) {
	t := &Task{
		Rule:&spider.Rule{

		},
		Times:&TaskTimes{

		},
	}
	err := json.Unmarshal(content, t)
	return t,err
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
func (task *Task) Run() {
	sp := &spider.CollySpider{
		Rule: task.Rule,
	}
	sp.Run()
}

func AddToTask(j *Task) string {
	job := taskJob
	
	runTime := j.Times.RunTime
	if runTime == 0 && j.Times.Spacing > 0 {
		runTime = time.Now().UnixNano()+ j.Times.Spacing * int64(time.Second)
	} else if runTime < time.Now().Unix() {
		runTime = (time.Now().Unix() + 1) * int64(time.Second)
	}
	uuid := job.AddTask(&timing.Task{
		Job:     j,
		RunTime: runTime,
		Spacing: j.Times.Spacing,
		EndTime: j.Times.EndTime,
	})
	return uuid
}