package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/queue"
	"github.com/lwl1989/go-spider/spider"
	"strings"

	//	"github.com/gocolly/colly/queue"
)

func main()  {
	sp := &spider.CollySpider{
		Rule: &spider.Rule{
			Index:"https://www.chinatimes.com/money/",
			IndexDom:".news-list",
			IndexListDom:".cropper a",
			IndexListAttr:"href",
			MaxPage:10000,
			PageReg: "",
			PageDom :".arttext",
		},
	}
	sp.Run()
	return
	//spider.
	//t := "https://www.chinatimes.com/realtimenews/20190110004436-260412"
	//a,b := regexp.Compile("https://www.chinatimes\\.com/realtimenews/([a-zA-Z0-9-_])")
	//fmt.Print(a,b)
	//bb:=a.MatchString(t)
	//fmt.Print(bb)
	//return
	url := "https://www.chinatimes.com/money/"

	// Instantiate default collector
	c := colly.NewCollector(
	//	colly.URLFilters(
	//	regexp.MustCompile("https://www.chinatimes\\.com/realtimenews/([a-zA-Z0-9-_])"),
	//),
	)
	//c.
	//c.MaxDepth = 2;
	// create a request queue with 2 consumer threads
	//q, _ := queue.New(
	//	2, // Number of consumer threads
	//	&queue.InMemoryQueueStorage{MaxSize: 10000}, // Use default queue storage,
	//)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("visiting", r.URL)
	})
	c.OnHTML(".news-list", func(e *colly.HTMLElement) {
		q, _ := queue.New(
			2, // Number of consumer threads
			&queue.InMemoryQueueStorage{MaxSize: 10000}, // Use default queue storage,
		)
		s2 := e.ChildAttrs(".cropper a","href")
		//e.ForEach(".cropper", func(e *colly.HTMLElement) {
		//	//.Attr("href")
		//})
		fmt.Println(e.Request.URL.Scheme)
		for _,v := range s2 {
			if strings.HasPrefix(v, "//") {
				q.AddURL("https:"+v)
			}
		}
		q.Run(c)
		//fmt.Println(s2)
		//s,e1 := e.DOM.Html()
		//fmt.Println(s,e1)
	})
	c.OnHTML(".arttext p", func(e *colly.HTMLElement) {
		fmt.Println(e.Name)
		//q, _ := queue.New(
		//	2, // Number of consumer threads
		//	&queue.InMemoryQueueStorage{MaxSize: 10000}, // Use default queue storage,
		//)

		s,e1 := e.DOM.Html()
		fmt.Println(s,e1)
	})
	c.Visit(url)
	//for i := 0; i < 5; i++ {
	//	// Add URLs to the queue
	//	q.AddURL(fmt.Sprintf("%s?n=%d", url, i))
	//}
	// Consume URLs
	//q.Run(c)
}