package spider

type Rule struct {
	Index string `json:"index"`         //入口
	IndexType string `json:"index_type,omitempty"` //列表页类型
	IndexDom string `json:"index_dom"` //首页要抓取的层级,如果是json模式，则是列表的key
	IndexListDom string `json:"index_list_dom,omitempty"` //列表的dom
	IndexListAttr string `json:"index_list_attr,omitempty"` //要获取嘚属性
	MaxPage int `json:"max_page"`    //列表页的最大抓取数量
	PageType string `json:"page_type"` //详情页类型
	PageReg  string `json:"page_reg"`  //列表页规则（正则）
	PageDom *PageDom `json:"page_dom"` //子页面要抓取的层级
	Metas *Metas `json:"metas,omitempty"`      //要抓取的meta
	RemoveTags []string `json:"remove_tags,omitempty"` //要移除的标签 比如 <script <iframe <ul .....
	CallBack ICall `json:"call_back,omitempty"`
}


//meta need set name
type Metas struct {
	Names []string `json:"names,omitempty"`
}
func (m *Metas) GetNames() []string {
	return m.Names
}

//page dom need set object like '.content p'
//content and article must set and other can be empty string
type PageDom struct {
	Article string `json:"article"`
	Title string `json:"title,omitempty"`
	Author string `json:"author,omitempty"`
	Time string `json:"time,omitempty"`
	Content string `json:"content"`
	Others string `json:"others,omitempty"`
}
func (dom *PageDom) GetArticle() string {
	return dom.Article
}
func (dom *PageDom) GetTitle() string {
	return dom.Title
}
func (dom *PageDom) GetAuthor() string {
	return dom.Author
}
func (dom *PageDom) GetTime() string {
	return dom.Time
}
func (dom *PageDom) GetContent() string {
	return dom.Content
}
func (dom *PageDom) GetOthers() string {
	return dom.Others
}