package spider

import (
	"errors"
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/queue"
	"regexp"
	"strings"
	"sync"
)

var MapOnce sync.Once
var Collys CollyMaps

type CollyMaps map[string]*CollySpider
type CollySpider struct {
	c *colly.Collector
	Rule *Rule
	hash string
}


func (m CollyMaps) addColly(spider *CollySpider) {
	hash,err := spider.GetHash()
	if err == nil {
		m[hash] = spider
	}


}

func (spider *CollySpider) GetHash() (string,error) {
	if spider.Rule == nil {
		return "",errors.New("未配置规则")
	}
	//need change a new rule, test only
	if spider.hash == "" {
		spider.hash = spider.Rule.Index
	}

	return spider.hash, nil
}

func (spider *CollySpider) Run() (error) {
	spider.c = colly.NewCollector(
		colly.URLFilters(
			regexp.MustCompile(spider.Rule.PageReg),
		),
	)
	c := spider.c

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("visiting", r.URL)
	})
	c.OnHTML("meta")
	c.OnHTML(spider.Rule.IndexDom, func(e *colly.HTMLElement) {
		q, _ := queue.New(
			2, // Number of consumer threads
			&queue.InMemoryQueueStorage{MaxSize: spider.Rule.MaxPage}, // Use default queue storage,
		)
		s2 := e.ChildAttrs(spider.Rule.IndexListDom,spider.Rule.IndexListAttr)
		scheme := e.Request.URL.Scheme
		for _,v := range s2 {
			if strings.HasPrefix(v, "//") {
				q.AddURL(scheme+":"+v)
			}
		}
		q.Run(c)
		//fmt.Println(s2)
		//s,e1 := e.DOM.Html()
		//fmt.Println(s,e1)
	})
	c.OnHTML(spider.Rule.PageDom, func(e *colly.HTMLElement) {
		fmt.Println(e.Name)
		//q, _ := queue.New(
		//	2, // Number of consumer threads
		//	&queue.InMemoryQueueStorage{MaxSize: 10000}, // Use default queue storage,
		//)

		s,e1 := e.DOM.Html()
		fmt.Println(s,e1)
	})
	c.Visit(spider.Rule.Index)

	return nil
}


func getOneNewColly() *colly.Collector {
	return  colly.NewCollector()
}

func getCollyMaps() CollyMaps {
	MapOnce.Do(func() {
		Collys = make(map[string]*CollySpider)
	})
	return Collys
}