package main

import (
	"github.com/lwl1989/go-spider/spider"
	"github.com/lwl1989/go-spider/config"
	"github.com/lwl1989/go-spider/http"
)


func main()  {
	return
	pageDom := &spider.PageDom{
		Article: ".thoracis",
		Content: ".ndArticle_margin p",
	}
	//ageDom = append(pageDom, ".thoracis .ndArticle_margin p")
	sp5 :=  &spider.CollySpider{
		Rule: &spider.Rule{
			Index:"https://tw.finance.appledaily.com/daily/",
			IndexDom:".abdominis",
			IndexListDom:"ul li a",
			IndexListAttr:"href",
			MaxPage:10000,
			PageReg: "",
			PageDom :pageDom,
		},
	}
	sp5.Run()
	return
	//pageDom = append(pageDom, ".Cf article")
	pageDom.Article = ".Cf"
	pageDom.Content = "article"
	sp4 := &spider.CollySpider{
		Rule: &spider.Rule{
			Index:"https://tw.news.yahoo.com/finance",
			IndexDom:"#app",
			IndexListDom:"div",
			IndexListAttr:"href",
			MaxPage:10000,
			PageReg: "",
			PageDom :pageDom,
		},
	}
	sp4.Run()
	return
	pageDom.Article = ".main"
	pageDom.Content = "article"
	pageDom.Title = "h1"
	//pageDom = append(pageDom, ".main h1")
	//pageDom = append(pageDom, ".main article")
	sp3 := &spider.CollySpider{
		Rule: &spider.Rule{
			Index:"https://news.cnyes.com/news/cat/headline?exp=a",
			IndexDom:".theme-list",
			IndexListDom:"div a",
			IndexListAttr:"href",
			MaxPage:10000,
			PageReg: "",
			PageDom :pageDom,
		},
	}
	sp3.Run()
	return
	pageDom.Title = ""
	pageDom.Content = "#story_body"
	pageDom.Article = pageDom.Content
	sp2 := &spider.CollySpider{
		Rule: &spider.Rule{
			Index:"https://money.udn.com/money/index",
			IndexDom:".tabs_box_wrapper",
			IndexListDom:"#tab1 a",
			IndexListAttr:"href",
			MaxPage:10000,
			PageReg: "",
			PageDom :pageDom,
		},
	}
	sp2.Run()
	return
	//pageDom = append(pageDom, ".entry-main")
	pageDom.Content = ".entry-main"
	pageDom.Article = pageDom.Content
	sp1 := &spider.CollySpider{
		Rule: &spider.Rule{
			Index:"https://ctee.com.tw/",
			IndexDom:".vc_tta-container",
			IndexListDom:".vc_tta-panel .wpb_wrapper a",
			IndexListAttr:"href",
			MaxPage:10000,
			PageReg: "",
			PageDom :pageDom,
		},
	}
	sp1.Run()
	return
	//pageDom = append(pageDom, ".arttext")
	pageDom.Content = ".arttext"
	pageDom.Article = pageDom.Content
	sp := &spider.CollySpider{
		Rule: &spider.Rule{
			Index:"https://www.chinatimes.com/money/",
			IndexDom:".news-list",
			IndexListDom:".cropper a",
			IndexListAttr:"href",
			MaxPage:10000,
			PageReg: "https://www.chinatimes\\.com/realtimenews/([a-zA-Z0-9-_])",
			PageDom :pageDom,
		},
	}
	sp.Run()
	return

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