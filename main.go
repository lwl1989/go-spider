package main

import (
	"github.com/lwl1989/go-spider/config"
	"github.com/lwl1989/go-spider/http"
	"github.com/lwl1989/go-spider/spider"
)


func main()  {
	initSpider()
}

func initSpider()  {
	//init spider config
	spider.Cf = config.GetConfig("/etc/go-spider.json")
	//init spider log
	spider.GetLog()
	//init http task
	//the task can loop task and  once task
 	http.InitTask()
	//init handler
	http.GetHandler(spider.Cf)
}