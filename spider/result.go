package spider


type Results struct {
	Pos int `json:"pos"`
	Value []*Result `json:"value"`
	ResultMap   map[string]string      `json:"result_map"`
}
type Result struct {
		Title       interface{}            `json:"title"`
		Metas       map[string]interface{} `json:"metas"`
		Author      interface{}            `json:"author"`
		Content     []interface{}          `json:"content"`
		PublishTime interface{}            `json:"publish_time"`
}

type ResultsInterface interface{
	SetTitle(title interface{})
	AddMeta(name string,value interface{})
	SetAuthor(author interface{})
	SetContent(content interface{})
	SetPublishTime(time interface{})
	SetResultValue(field string, value interface{}, fieldName string)
	GetResultMap(string string) string
}

func (res *Results) SetResultValue(field string, value interface{}, fieldName string) {
	field = res.GetResultMap(field)
	switch field {
		case "title":
			res.SetTitle(value)
		case "meta":
			res.AddMeta(fieldName, value)
		case "author":
			res.SetAuthor(value)
		case "publish_time":
			res.SetPublishTime(value)
		case "content":
			res.SetContent(value)
	}
}

func (res *Results) GetResultMap(string string) string {
	v,ok := res.ResultMap[string]
	if ok {
		return v
	}
	return ""
}

func (res *Results) SetTitle(title interface{}) {
	res.Value[res.Pos].Title = title
}

func (res *Results) AddMeta(name string,value interface{}) {
	res.Value[res.Pos].Metas[name] = value
}

func (res *Results) SetAuthor(author interface{}) {
	res.Value[res.Pos].Author = author
}

func (res *Results) SetContent(content interface{}) {
	res.Value[res.Pos].Content = append(res.Value[res.Pos].Content, content)
}

func (res *Results) SetPublishTime(time interface{}) {
	res.Value[res.Pos].PublishTime = time
}