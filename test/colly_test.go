package test

import (
	"testing"
	"github.com/lwl1989/go-spider/spider"
)

//test add Func
func Test_Spider(t *testing.T) {
	pageDome := &spider.PageDom{
		Article: ".arttext p",
		Content: ".arttext p",
	}
	sp := &spider.CollySpider{
		Rule: &spider.Rule{
			Index:"https://tw.news.yahoo.com/_td-news/api/resource/content;fetchNewAttribution=true;getDetailView=true;getFullLcp=false;imageResizer=null;relatedContent=%7B%22enabled%22%3Atrue%7D;site=news;uuids=%5B%2257dbd72a-17e0-3ff1-a818-9ad4ec3b8131%22%2C%2220ef43fb-174c-373c-bbba-46b7e3dbdc7a%22%2C%226e5e0131-d7e7-3226-926a-971096e75531%22%2C%22468c7667-2449-3315-8e24-711afd980e9b%22%2C%228853d88a-856d-3c87-8377-e1576e73c7e8%22%2C%2232cc09e3-0b1c-3b8e-82d2-087a1a879a20%22%2C%227a408353-c42c-3b05-a38c-415224eb3683%22%2C%2230da0880-3564-3ba5-a5de-548afca94dd7%22%2C%2286a2754c-5a9a-3c5f-a490-49a3aee8bcfb%22%2C%223ad41a25-b995-3d8c-a87b-6a00f044877c%22%2C%224230615e-fd41-385d-bba6-3c746b09d9ae%22%2C%22df93dd7f-c318-3d7c-8212-06d50b6aa1e6%22%2C%22d93531e9-03ce-3cea-a761-7c57b083b915%22%5D?bkt=news-TW-zh-Hant-TW-def&device=desktop&feature=videoDocking&intl=tw&lang=zh-Hant-TW&partner=none&prid=a4ekmqte4ag5k&region=TW&site=news&tz=Asia%2FSeoul&ver=2.0.1121&returnMeta=true",
			IndexDom:".news-list",
			IndexListDom:".cropper a",
			IndexListAttr:"href",
			MaxPage:10000,
			PageReg: "",
			PageDom : pageDome,
		},
	}
	sp.Run()
}
