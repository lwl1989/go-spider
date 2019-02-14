package spider

import (
	"errors"
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/queue"
	"regexp"
	"strings"
	"github.com/lwl1989/go-spider/config"
	"github.com/gocolly/colly/proxy"
	"encoding/json"
	"net/url"
	"time"
	"github.com/PuerkitoBio/goquery"
)

var Cf *config.Config

type CollySpider struct {
	c *colly.Collector
	Rule *Rule `json:"rule"`
	Times *TaskTimes `json:"times"`
	hash string
}
type TaskTimes   struct {
	RunTime int64  `json:"run_time"`
	Spacing int64  `json:"spacing_time"`
	EndTime int64  `json:"end_time,omitempty"`
	Number  int  `json:"number,omitempty"`
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
		var result = &Result{
			Metas:make(map[string]interface{}),
		}
		l  := e.DOM.Parents().Find("meta")

		l.Each(func(i int, selection *goquery.Selection) {
			name,exists := selection.Attr("name")
			if exists {
				h,_ := selection.Attr("content")
				result.Metas[name] = h
			}else{
				name,exists = selection.Attr("property")
				if exists {
					h,_ := selection.Attr("content")
					result.Metas[name] = h
				}
			}

		})
		var s = ""
		var err error
		if spider.Rule.PageDom.GetArticle() == spider.Rule.PageDom.GetContent() {
			s,err = e.DOM.Html()
			if err != nil {
				fmt.Println(err)
			}
		}else{
			s,_ = e.DOM.Find(spider.Rule.PageDom.GetContent()).Html()
		}
		if len(spider.Rule.RemoveTags) > 0 {
			result.SetContent(RemoveScript(RemoveSpace(s), spider.Rule.RemoveTags...))
		}else{
			result.SetContent(RemoveScript(RemoveSpace(s)))
		}
		result.Title,_ = e.DOM.Find(spider.Rule.PageDom.GetTitle()).Html()
		result.PublishTime,_ = e.DOM.Find(spider.Rule.PageDom.GetTime()).Html()
		result.Author,_ = e.DOM.Find(spider.Rule.PageDom.GetAuthor()).Html()
		DoCallBack(result)
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
	//这里用pagedom做了一个映射
	//metas暂时不处理
	pd := spider.Rule.PageDom
	fields := make([]string, 0)
	fields = append(fields,  pd.GetAuthor())
	//fields = append(fields,  pd.GetOthers())
	fields = append(fields,  pd.GetTitle())
	fields = append(fields,  pd.GetContent())
	fields = append(fields,  pd.GetTime())
	rm := make(map[string]string)
	rm[pd.GetTitle()] = "title"
	rm[pd.GetAuthor()] = "author"
	rm[pd.GetTime()] = "publish_time"
	//rm[pd.GetOthers()] = "meta"
	rm[pd.GetContent()] = "content"
	result := &Results{
		Pos: 0,
		ResultMap: rm,
	}
	l := ParseMapsFindList(m, spider.Rule.IndexDom)
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

func (spider *CollySpider) Run() {
	if spider.Rule.IndexType == "json" {
		spider.runListResult()
	}else{
		spider.getOneNewColly()
		spider.runHtmlBody()
		spider.runHtmlList()
		spider.c.OnRequest(func(r *colly.Request) {
			//todo: this way need send error to log
			fmt.Println("抓取到页面", r.URL)
		})
		spider.c.OnError(func(response *colly.Response, e error) {
			//todo: this way need send error to log
			fmt.Println(e)
		})
		spider.c.OnScraped(func(response *colly.Response) {
			//todo: the page over ! this way need send error to log
			fmt.Println(response)
		})
	}

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

func (spider *CollySpider) getOneNewColly() {
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

	if Cf.HttpConfig.Proxy != "" {
		rp, err := proxy.RoundRobinProxySwitcher(Cf.HttpConfig.Proxy)
		if err != nil {
			panic(err)
		}
		spider.c.SetProxyFunc(rp)
	}
}

func NewSpider(content []byte) (*CollySpider,error) {
	t := &CollySpider{
			Rule:&Rule{
				PageDom:&PageDom{

				},
			},
			Times:&TaskTimes{

			},
	}
	err := json.Unmarshal(content, t)
	return t,err
}
