package spider

import (
	"net/http"
	"io"
	"time"
	"net/url"
	"fmt"
)

type JsonBody struct {
	IndexProperty string `json:"index_property"`
	ListProperty  string `json:"list_property"`
}

type Iterator interface{
	HasNext() bool
	Current() string
	Next() string
	Length() int
}

func GetHttpClient() *http.Client {
	c := http.DefaultClient
	c.Timeout = time.Duration(20 * time.Second)
	l,_:=url.Parse("https://127.0.0.1:1087")
	c.Transport = &http.Transport{
		DisableKeepAlives: false,//关闭连接复用，因为后台连接过多最后会造成端口耗尽
		MaxIdleConns: -1,  //最大空闲连接数量
		IdleConnTimeout: time.Duration(20 * time.Second),  //空闲连接超时时间
		Proxy: http.ProxyURL(l), //设置http代理地址
	}
	return c
}

func GetJson(url string) ([]byte,error) {
	res,err := GetHttpClient().Get(url)
	if err != nil {
		return nil,err
	}

	var all = make([]byte, 0)
	for ;; {
		var bt [1024]byte
		n,err  := res.Body.Read(bt[:])


		if err != nil {
			if err == io.EOF {
				all = append(all, bt[:n]...)
				break
			}
			fmt.Println(err)
			return nil,err
		}
		all = append(all, bt[:n]...)
	}

	return all[:],nil
}

type ResultList struct {
	L []string
	pos int
}
func (result *ResultList) append(s string) {
	result.L = append(result.L, s)
}
func (result *ResultList) HasNext() bool {
	return (result.pos + 1) < len(result.L)
}
func (result *ResultList) Current() string {
	return result.L[result.pos]
}
func (result *ResultList) Length() int {
	return len(result.L)
}

func (result *ResultList) Next() string {
	if result.HasNext() {
		result.pos ++
		return result.L[result.pos]
	}
	return ""
}

func ParseMapsFindList(m map[string]interface{}, listField string) []interface{} {
	if len(m) > 0 {
		for k,v := range m {
			switch v.(type) {
			case map[string]interface{}:
				return ParseMapsFindList(v.(map[string]interface{}), listField)
			case []interface{}:
				if k == listField {
					return v.([]interface{})
				}else{
					for _,v1 :=range v.([]interface{}) {
						switch v1.(type) {
							case string:
								return nil
						case map[string]interface{}:
							return ParseMapsFindList(v1.(map[string]interface{}), listField)
						default:
							return nil
						}
					}
				}
			default:
				return nil
			}
			if k == listField {
				switch  v.(type) {
					case map[string]interface{}:
						return ParseMapsFindList(m, listField)
					case []interface{}:
						return v.([]interface{})
					default:
						return nil
				}
			}else{

			}
		}
	}
	return  nil
}

func ParseMap(m map[string]interface{}, fields []string, results *Results) {
	if len(m) > 0 {
		for k,v := range m {
			for _, field := range fields {
				if field == k {
					results.SetResultValue(field, v, "")
				}
			}
			switch  v.(type) {
			case string:

			case map[string]interface{}:
				ParseMap(v.(map[string]interface{}), fields, results)
			case []interface{}:
				for _, v := range v.([]interface{}) {
					switch v.(type) {
					case map[string]interface{}:
						ParseMap(v.(map[string]interface{}), fields, results)
					default:

					}
				}
			}
		}
	}
}
func ParseList(l []interface{}, finds []string, results *Results) {
	if len(l) > 0 {
		for k,v := range l {

			results.Pos = k
			results.Value = append(results.Value, &Result{
				Content: make([]interface{},0),
				Metas: make(map[string]interface{}),
			})

			switch v.(type) {
			case map[string]interface{}:
				ParseMap(v.(map[string]interface{}), finds, results)
			default:

			}

		}
	}
}
