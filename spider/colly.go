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
	"encoding/json"
	"net/url"
	"time"
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
		return "", errors.New("未配置规则")
	}
	//need change a new rule, test only
	if spider.hash == "" {
		spider.hash = spider.Rule.Index
	}

	return spider.hash, nil
}


func (spider *CollySpider) runHtmlBody() (error) {
	//spider.c.OnHTML("#app", func(element *colly.HTMLElement) {
	//	fmt.Println(element.Text)
	//})
	spider.c.OnHTML(spider.Rule.PageDom.GetArticle(),func(e *colly.HTMLElement) {
		fmt.Println("get")
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

	return nil
}

func (spider *CollySpider) runJsonList() (error) {
	a,_ := GetJson(spider.Rule.Index)
	m := make(map[string]interface{})
	err := json.Unmarshal(a, &m)
	if err != nil {
		fmt.Println(err)
	}
	l := &ResultList{
		L:make([]string,0),
	}

	//ParseMap(m, "href", l)
	return runListQueue(spider, l)
}

func (spider *CollySpider)  runListResult() (error) {
	a,_ := GetJson(spider.Rule.Index)
	m := make(map[string]interface{})
	err := json.Unmarshal(a, &m)
	if err != nil {
		fmt.Println(err)
	}
	fields := make([]string, 0)
	fields = append(fields, "body")
	fields = append(fields, "title")
	fields = append(fields, "author")
	fields = append(fields, "publishDate")
	fields = append(fields, "ad_meta")
	rm := make(map[string]string)
	rm["title"] = "title"
	rm["author"] = "author"
	rm["publishDate"] = "publish_time"
	rm["ad_meta"] = "meta"
	rm["body"] = "content"
	result := &Results{
		Pos: 0,
		ResultMap: rm,
	}
	l := ParseMapsFindList(m, "items")
	ParseList(l, fields, result)
	return  nil
}

func runListQueue(spider *CollySpider, it Iterator) (error) {
	spider.c.OnHTML("#app", func(e *colly.HTMLElement) {
		fmt.Println(e.Request.URL.RequestURI(), e.DOM.Text())
	})
	u,err := url.Parse(spider.Rule.Index)
	if err != nil {
		return err
	}
	q, _ := queue.New(
		1, // Number of consumer threads
		&queue.InMemoryQueueStorage{MaxSize: it.Length()}, // Use default queue storage,
	)
	for ;; {
		if ! it.HasNext() {
			break
		}
		s := it.Next()
		q.AddURL(checkUrl(u, s))
	}
	return q.Run(spider.c)
}

func (spider *CollySpider) runHtmlList() (error) {
	spider.c.OnHTML(spider.Rule.IndexDom, func(e *colly.HTMLElement) {
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
				q.AddURL(checkUrl(e.Request.URL, v))
			}
			// run queue
			q.Run(spider.c)
		}
	})

	spider.c.Visit(spider.Rule.Index)
	return nil
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
	spider.c.SetRequestTimeout(30*time.Second)
	c := spider.c
	rp, err := proxy.RoundRobinProxySwitcher("socks5://127.0.0.1:1086")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	c.SetProxyFunc(rp)

	if spider.Rule.IndexType == "json" {
		spider.runListResult()
	}else{

		c.OnRequest(func(r *colly.Request) {
			fmt.Println("抓取到页面", r.URL)
		})
		c.OnError(func(response *colly.Response, e error) {
			fmt.Println(e)
		})

		spider.runHtmlList()
		spider.runHtmlBody()
	}

	return nil
}

func checkUrl(url *url.URL, v string) string {
	if strings.HasPrefix(v, "//") {
		scheme := url.Scheme
		return scheme+":"+v
	}else if strings.HasPrefix(v, "/") {
		scheme := url.Scheme
		return scheme+"://"+url.Host + v
	}else {
		return v
	}
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