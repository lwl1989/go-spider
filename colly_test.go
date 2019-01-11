package main

import (
	"github.com/lwl1989/go-spider/spider"
	"testing"
)

//test add Func
func Test_Spider(t *testing.T) {
	sp := &spider.CollySpider{
		Rule: &spider.Rule{
			Index:"https://www.chinatimes.com/money/",
			IndexDom:".news-list",
			IndexListDom:".cropper a",
			IndexListAttr:"href",
			MaxPage:10000,
			PageReg: "",
			PageDom :".arttext p",
		},
	}
	sp.Run()
}
