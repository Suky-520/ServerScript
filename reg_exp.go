//JavaScript中的正则表达式

package ss

import (
	"encoding/json"
	"github.com/dlclark/regexp2"
	"ss/collection"
	"strings"
)

// RegExp 正则表达式对象
type RegExp struct {
	pattern string
	str     string
	prop    JavaScriptProperties
}

var _ JavaScript = &RegExp{}
var _ JavaScriptBasic = &RegExp{}
var _ CodeCompiler = &RegExp{}

// GoToRegExp 新建正则表达式对象
func GoToRegExp(str string) *RegExp {
	o := &RegExp{pattern: convertRegex(str), str: str, prop: collection.NewArrayMap[string, JavaScript]()}
	o.prop.Set("test", o.test())
	o.prop.Set("exec", o.exec())
	return o
}

// 将js中正则表达式字面量转换为go语言中的正则表达式
func convertRegex(jsRegex string) string {
	regexBodyAndFlags := strings.ReplaceAll(jsRegex[1:len(jsRegex)-1], "\\/", "/")
	// 分离正则表达式主体和标志
	var pattern, flags string
	lastSlashIndex := strings.LastIndex(regexBodyAndFlags, "/")
	if lastSlashIndex != -1 {
		// 找到了标志的分隔斜杠
		pattern = regexBodyAndFlags[:lastSlashIndex]
		flags = regexBodyAndFlags[lastSlashIndex+1:]
	} else {
		// 没有标志
		pattern = regexBodyAndFlags
	}
	goPattern := pattern
	if strings.Contains(flags, "index") {
		goPattern = "(?index)" + goPattern
	} else if strings.Contains(flags, "m") {
		goPattern = "(?m)" + goPattern
	} else if strings.Contains(flags, "s") {
		goPattern = "(?s)" + goPattern
	} else if flags != "" {
		goPattern = goPattern + "/" + flags
	}
	return goPattern
}

func (r *RegExp) Value() any {
	return r.str
}

func (r *RegExp) GetName() string {
	return "regexp"
}

func (r *RegExp) GetProperty(name string) JavaScript {
	if v, ok := r.prop.Get(name); ok {
		return v
	}
	return NewUndefined()
}

func (r *RegExp) GetProperties() JavaScriptProperties {
	return r.prop
}

func (r *RegExp) ToString() string {
	return r.str
}

func (r *RegExp) exec() JavaScript {
	return NewDefaultFunction("exec", func(args []JavaScript, ctx *Context) JavaScript {
		re := regexp2.MustCompile(r.pattern, 0)
		// 执行匹配操作
		match, e := re.FindStringMatch(JsToString(args[0]))
		if e != nil {
			panic(NewError("Error", e.Error()))
		}
		return GoToJs(match)
	})
}

func (r *RegExp) test() JavaScript {
	return NewDefaultFunction("test", func(args []JavaScript, ctx *Context) JavaScript {
		re := regexp2.MustCompile(r.pattern, 0)
		result, err := re.MatchString(JsToString(args[0]))
		if err != nil {
			panic(NewError("Error", "regular expression error"))
		}
		return GoToBoolean(result)
	})
}

func (r *RegExp) Print() string {
	return RedStyle + r.str + EndStyle
}
func (r *RegExp) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.str)
}
func (r *RegExp) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.str)
}
func (r *RegExp) UnmarshalBin(data []byte) (any, error) {
	r.str = string(data)
	r.pattern = convertRegex(string(data))
	return nil, nil
}

func (r *RegExp) MarshalBin() ([]byte, error) {
	data := []byte{6}
	data = append(data, []byte(r.str)...)
	return data, nil
}
