package spider

import (
	"strings"
	"fmt"
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


func removeScript(content string, matches ...string) string {
	if content == "" {
		return content
	}
	if len(matches) == 0 {
		fmt.Println("sss")
		return content
	}
	match := make([]string, 0)
	for k, v := range matches {
		start := unicodeIndex(content, v)

		if start == -1 {
			match := matches[k+1:]
			return removeScript(content, match...)
		}

		end := -1
		cutLen := 9

		if strings.Index(v, "script") >= 0 {
			end = unicodeIndex(content, "</script>")
		}
		if strings.Index(v, "iframe") >= 0 {
			end = unicodeIndex(content, "</iframe>")
		}
		if strings.Index(v, "ul") >= 0 {
			end = unicodeIndex(content, "</ul>")
			cutLen = 5
		}
		if strings.Index(v, "style") >= 0 {
			end = unicodeIndex(content, "</style>")
			cutLen = 8
		}
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

func removePosition(source string, start int, end int) string {
	var r = []rune(source)
	length := len(r)

	if end > length {
		end = length
	}

	if start < 0  || start > end {
		return ""
	}

	if start == 0 && end == length {
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
