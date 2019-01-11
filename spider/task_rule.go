package spider

type Rule struct {
	Index string `json:"index"`         //入口
	IndexDom string `json:"index_Dom"` //首页要抓取的层级
	IndexListDom string `json:"index_list_dom"` //列表的dom
	IndexListAttr string `json:"index_list_attr"` //要获取嘚属性
	MaxPage int `json:"max_page"`    //列表页的最大抓取数量
	PageReg   string `json:"page_reg"`  //列表页规则（正则）
	PageDom string `json:"page_dom"` //子页面要抓取的层级
}