package spider

import (
	"errors"
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/queue"
	"regexp"
	"strings"
	"sync"
	"github.com/lwl1989/go-spider/config"
	"github.com/gocolly/colly/proxy"
)

var MapOnce sync.Once
var Collys CollyMaps
var Cf *config.Config

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

	if spider.Rule.PageReg != "" {
		opt := colly.URLFilters(
			regexp.MustCompile(spider.Rule.PageReg),
		)
		spider.c = colly.NewCollector(
			opt,
		)
	}else{
		spider.c = colly.NewCollector()
	}

	c := spider.c
	rp, err := proxy.RoundRobinProxySwitcher("socks5://127.0.0.1:1086")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	c.SetProxyFunc(rp)
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("抓取到页面", r.URL)
	})
	//c.OnHTML("meta")
	c.OnHTML(spider.Rule.IndexDom, func(e *colly.HTMLElement) {
		s2 := e.ChildAttrs(spider.Rule.IndexListDom, spider.Rule.IndexListAttr)
		length := len(s2)
		//index html got, new a queue with list
		//spider.Rule.MaxPage
		if length > 0 {
			q, _ := queue.New(
				1, // Number of consumer threads
				&queue.InMemoryQueueStorage{MaxSize: length}, // Use default queue storage,
			)
			for _,v := range s2 {
				if strings.HasPrefix(v, "//") {
					scheme := e.Request.URL.Scheme
					q.AddURL(scheme+":"+v)
				}else if strings.HasPrefix(v, "/") {
					scheme := e.Request.URL.Scheme
					q.AddURL(scheme+"://"+e.Request.URL.Host + v)
				}else {
					q.AddURL(v)
				}
			}
			// run queue
			q.Run(c)
		}


	})
	c.OnHTML(spider.Rule.PageDom.GetArticle(),func(e *colly.HTMLElement) {
		var contents = make([]string, 0)
		if spider.Rule.PageDom.GetArticle() == spider.Rule.PageDom.GetContent() {
			//return all
			s,e1 := e.DOM.Html()
			if e1 != nil {
				fmt.Println(e1)
			}else {
				fmt.Println(s, e1)
			}
			contents = append(contents, s)
		}else{
			//get children string
		}
	})
	//for _,v := range spider.Rule.PageDom {
	//	c.OnHTML(v, func(e *colly.HTMLElement) {
	//		s,e1 := e.DOM.Html()
	//		if e1 != nil {
	//			fmt.Println(e1)
	//		}else {
	//			fmt.Println(s, e1)
	//		}
	//	})
	//}

	c.Visit(spider.Rule.Index)

	return nil
}


func getOneNewColly() *colly.Collector {
	return colly.NewCollector()
}

func getCollyMaps() CollyMaps {
	MapOnce.Do(func() {
		Collys = make(map[string]*CollySpider)
	})
	return Collys
}