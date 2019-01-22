package spider

import (
	"strings"
)

//, match ...string
func RemoveScript(content string, match ...string)  string {
	matches := make([]string,0)
	if len(match) == 0 {
		matches = append(matches, `<script`)
		matches = append(matches, "<iframe")
		matches = append(matches, "<ul")
		matches = append(matches, "<style")
	}else{
		matches = match
	}
	return removeScript(content, matches...)
}

//remove tag in content
//return value is string
func removeScript(content string, matches ...string) string {
	if content == "" {
		return content
	}
	if len(matches) == 0 {
		return content
	}
	match := make([]string, 0)
	for k, v := range matches {
		if strings.Index(v, "<") == -1 {
			v = "<"+v
		}
		// if str len eq match len, go other way
		if len(content) < len(v) {
			match := matches[k+1:]
			return removeScript(content, match...)
		}

		start := unicodeIndex(content, v)

		// if not found the start tag, go other way
		if start == -1 {
			match := matches[k+1:]
			return removeScript(content, match...)
		}

		//normal pos and cut len need add
		tagName := strings.TrimLeft(v, "<")
		endTag := "</"+tagName+">"
		cutLen := len(endTag)
		end := unicodeIndex(content, endTag)

		if end == -1 {
			match = matches[k+1:]
			return removeScript(content, match...)
		}else {
			content = removePosition(content, start, end+cutLen)
			start = unicodeIndex(content, v)
			if start == -1 {
				match = matches[k+1:]
			}else{
				match = matches
			}
			return removeScript(content, match...)
		}
	}
	return content
}
//the byte code pos cover to the unicode pos
func unicodeIndex(str, substr string) int {
	// 子串在字符串的字节位置
	result := strings.Index(str,substr)
	if result >= 0 {
		// 获得子串之前的字符串并转换成[]byte
		prefix := []byte(str)[0:result]
		// 将子串之前的字符串转换成[]rune
		rs := []rune(string(prefix))
		// 获得子串之前的字符串的长度，便是子串在字符串的字符位置
		result = len(rs)
	}

	return result
}

//remove the unicode code string with start and end
func removePosition(source string, start int, end int) string {
	var r = []rune(source)
	length := len(r)

	//if condition error
	if end > length || start < 0 || (start == 0 && end == length) {
		return source
	}

	r1 := make([]rune,0)
	r2 := make([]rune,0)
	if start != 0 {
		r1 = r[0:start]
	}
	if end < length {
		r2 = r[end:length]
	}
	return string(append(r1, r2...)[:])
}

func RemoveSpace(content string) string {
	bts := []string{"\r","\n","\t"}
	for _,str := range bts {
		content = strings.Replace(content, str, "", -1)
	}
	return content
}