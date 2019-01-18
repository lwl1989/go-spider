package main

import (
	"github.com/lwl1989/go-spider/spider"
	"github.com/lwl1989/go-spider/config"
	"github.com/gocolly/colly"
	"strconv"
	"os"
	"encoding/json"
	"strings"
	"log"
	"fmt"
)


func main()  {
	return
	pageDom := make([]string, 0)
	pageDom = append(pageDom, ".thoracis .ndArticle_margin p")
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
	pageDom = append(pageDom, ".Cf article")
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
	pageDom = append(pageDom, ".main h1")
	pageDom = append(pageDom, ".main article")
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
	pageDom = append(pageDom, "#story_body")
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
	pageDom = append(pageDom, ".entry-main")
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
	spider.Cf = config.GetConfig("/etc/go-spider.json")
	spider.GetLog()
	pageDom = append(pageDom, ".arttext")
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
}
type comment struct {
	Author  string `selector:"a.hnuser"`
	URL     string `selector:".age a[href]" attr:"href"`
	Comment string `selector:".comment"`
	Replies []*comment
	depth   int
}
func json_T() {



		var itemID string
		itemID = "18936581"
		//flag.StringVar(&itemID, "id", "", "hackernews post id")
		//flag.Parse()

		if itemID == "" {
			log.Println("Hackernews post id required")
			os.Exit(1)
		}

		comments := make([]*comment, 0)

		// Instantiate default collector
		c := colly.NewCollector()

		// Extract comment
		c.OnHTML("body", func(e *colly.HTMLElement) {
			fmt.Println(e)
			width, err := strconv.Atoi(e.ChildAttr("td.ind img", "width"))
			if err != nil {
				return
			}
			// hackernews uses 40px spacers to indent comment replies,
			// so we have to divide the width with it to get the depth
			// of the comment
			depth := width / 40
			c := &comment{
				Replies: make([]*comment, 0),
				depth:   depth,
			}
			e.Unmarshal(c)
			c.Comment = strings.TrimSpace(c.Comment[:len(c.Comment)-5])
			if depth == 0 {
				comments = append(comments, c)
				return
			}
			parent := comments[len(comments)-1]
			// append comment to its parent
			for i := 0; i < depth-1; i++ {
				parent = parent.Replies[len(parent.Replies)-1]
			}
			parent.Replies = append(parent.Replies, c)
		})

		c.Visit("https://news.ycombinator.com/item?id=" + itemID)

		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")

		// Dump json to the standard output
		enc.Encode(comments)

}