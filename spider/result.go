package spider


//It's only use in Json Result
type Results struct {
	Pos int `json:"pos"`
	Value []*Result `json:"value"`
	ResultMap   map[string]string      `json:"result_map"`
}
//It use to All Response
type Result struct {
		Title       interface{}            `json:"title"`
		Metas       map[string]interface{} `json:"metas"`
		Author      interface{}            `json:"author"`
		Content     []interface{}          `json:"content"`
		PublishTime interface{}            `json:"publish_time"`
}
//the result Interface
type ResultInterface interface{
	SetTitle(title interface{})
	AddMeta(name string,value interface{})
	SetAuthor(author interface{})
	SetContent(content interface{})
	SetPublishTime(time interface{})
}

type ResultsInterface interface{
	SetResultValue(field string, value interface{}, fieldName string)
	GetResultMap(string string) string
}

func (allRes *Results) SetResultValue(field string, value interface{}, fieldName string) {
	field = allRes.GetResultMap(field)
	switch field {
		case "title":
			allRes.SetTitle(value)
		case "meta":
			allRes.AddMeta(fieldName, value)
		case "author":
			allRes.SetAuthor(value)
		case "publish_time":
			allRes.SetPublishTime(value)
		case "content":
			allRes.SetContent(value)
	}
}

func (allRes *Results) GetResultMap(string string) string {
	v,ok := allRes.ResultMap[string]
	if ok {
		return v
	}
	return ""
}

func (allRes *Results) SetTitle(title interface{}) {
	allRes.Value[allRes.Pos].Title = title
}

func (allRes *Results) AddMeta(name string,value interface{}) {
	allRes.Value[allRes.Pos].Metas[name] = value
}

func (allRes *Results) SetAuthor(author interface{}) {
	allRes.Value[allRes.Pos].Author = author
}

func (allRes *Results) SetContent(content interface{}) {
	allRes.Value[allRes.Pos].Content = append(allRes.Value[allRes.Pos].Content, content)
}

func (allRes *Results) SetPublishTime(time interface{}) {
	allRes.Value[allRes.Pos].PublishTime = time
}

func (res *Result) SetTitle(title interface{}) {
	res.Title = title
}

func (res *Result) AddMeta(name string,value interface{}) {
	res.Metas[name] = value
}

func (res *Result) SetAuthor(author interface{}) {
	res.Author = author
}

func (res *Result) SetContent(content interface{}) {
	res.Content = append(res.Content, content)
}

func (res *Result) SetPublishTime(time interface{}) {
	res.PublishTime = time
}


func DoCallBack(res *Result) {
	if Cf.CallBack.Method == "redis" {
		rc := &RedisCall{
			Params:res,
		}
		rc.Do()
	}else{

	}

}