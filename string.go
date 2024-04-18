//JavaScript中的字符串

package ss

import (
	"encoding/json"
	"fmt"
	"regexp"
	"ss/collection"
	"strings"
	"unicode"
)

// String JavaScript 字符串类型
type String struct {
	Data string //字符串
}

// StringIterator 字符串迭代器
type StringIterator struct {
	data  *String
	index int
}

var _ JavaScript = &String{}
var _ JavaScriptBasic = &String{}
var _ JavaScriptBinary = &String{}
var _ CodeCompiler = &String{}

// GoToString 将go的string变量转换为js中的string变量
func GoToString[T any](data T) *String {
	o := &String{Data: fmt.Sprint(data)}
	return o
}

func (s *String) GetName() string {
	return "string"
}
func (s *String) GetIterator() JavaScriptIterator {
	return &StringIterator{index: 0, data: s}
}
func (s *String) GetProperty(name string) JavaScript {
	switch name {
	case "length":
		return GoToNumber(len(s.Data))
	case "match":
		return stringMatch(s)
	case "search":
		return stringSearch(s)
	case "replace":
		return stringReplace(s)
	case "split":
		return stringSplit(s)
	case "charAt":
		return stringCharAt(s)
	case "charCodeAt":
		return stringCharCodeAt(s)
	case "concat":
		return stringConcat(s)
	case "toUpperCase":
		return stringToUpperCase(s)
	case "toLowerCase":
		return stringToLowerCase(s)
	case "includes":
		return stringIncludes(s)
	case "indexOf":
		return stringIndexOf(s)
	case "lastIndexOf":
		return stringLastIndexOf(s)
	case "slice":
		return stringSlice(s)
	case "substring":
		return stringSubstring(s)
	case "trim":
		return stringTrim(s)
	case "trimStart":
		return stringTrimStart(s)
	case "trimEnd":
		return stringTrimEnd(s)
	case "padStart":
		return stringPadStart(s)
	case "padEnd":
		return stringPadEnd(s)
	case "localeCompare":
		return stringLocaleCompare(s)
	default:
		return NewUndefined()
	}
}
func (s *String) GetProperties() JavaScriptProperties {
	return collection.NewArrayMap[string, JavaScript]()
}
func (s *String) ToString() string {
	return s.Data
}
func (s *String) ToNumber() float64 {
	if s.Data == "" {
		return 0
	}
	return GoToNumber(s.Data).Data
}
func (s *String) Print() string {
	return GreenStyle + s.ToString() + EndStyle
}
func (s *String) ToBool() bool {
	return s.Data != ""
}
func (s *String) Binary() []byte {
	return []byte(s.Data)
}
func (s *String) Value() any {
	return s.Data
}
func (s *String) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &s.Data)
}
func (s *String) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Data)
}
func (s *String) UnmarshalBin(data []byte) (any, error) {
	s.Data = string(data)
	return nil, nil
}
func (s *String) MarshalBin() ([]byte, error) {
	data := []byte{7}
	data = append(data, []byte(s.Data)...)
	return data, nil
}

func getRegexp(r *RegExp) *regexp.Regexp {
	return regexp.MustCompile(r.pattern)
}

func stringMatch(s *String) JavaScript {
	return NewDefaultFunction("match", func(args []JavaScript, ctx *Context) JavaScript {
		if v, ok := args[0].(*RegExp); ok {
			matches := getRegexp(v).FindAllString(s.Data, -1)
			if len(matches) == 0 {
				return NewNull()
			}
			return GoToJs(matches)
		}
		return NewUndefined()
	})
}

func stringSearch(s *String) JavaScript {
	return NewDefaultFunction("search", func(args []JavaScript, ctx *Context) JavaScript {
		if len(args) == 0 {
			return GoToNumber(0)
		}
		if v, ok := args[0].(*RegExp); ok {
			searchResult := getRegexp(v).FindStringIndex(s.Data)
			return GoToNumber(float64(searchResult[0]))
		} else {
			return GoToNumber(float64(runeIndexOf(s.Data, JsToString(args[0]))))
		}
	})
}

// 返回字符串在s中首次出现的位置
func runeIndexOf(source string, str string) int {
	i := strings.Index(source, str)
	if i < 0 {
		return i
	}
	return len([]rune(source[:i]))
}

func runeLastIndexOf(source string, str string) int {
	i := strings.LastIndex(source, str)
	if i < 0 {
		return i
	}
	return len([]rune(source[:i]))
}

func stringReplace(s *String) JavaScript {
	return NewDefaultFunction("replace", func(args []JavaScript, ctx *Context) JavaScript {
		if v, ok := args[0].(*RegExp); ok {
			replaceResult := getRegexp(v).ReplaceAllString(s.Data, JsToString(args[1]))
			return GoToJs(replaceResult)
		} else {
			return GoToString(strings.Replace(s.Data, JsToString(args[0]), JsToString(args[1]), 1))
		}
	})
}

func stringSplit(s *String) JavaScript {
	return NewDefaultFunction("split", func(args []JavaScript, ctx *Context) JavaScript {
		if len(args) == 0 {
			return GoToJs([]string{s.Data})
		}
		i := -1
		if len(args) > 1 {
			if v, ok := args[1].(*Number); ok {
				i = int(v.Data)
			}
		}
		if i == 0 {
			return GoToArray([]any{})
		}
		if v, ok := args[0].(*RegExp); ok {
			splitResult := getRegexp(v).Split(s.Data, -1)
			if i > 0 {
				return GoToJs(splitResult[:i])
			} else {
				return GoToJs(splitResult)
			}
		} else {
			a := strings.Split(s.Data, JsToString(args[0]))
			if i > 0 {
				return GoToJs(a[:i])
			} else {
				return GoToJs(a)
			}
		}
	})
}

func charLen(s string) int {
	return len([]rune(s))
}

func charRune(s string) []rune {
	return []rune(s)
}

func stringCharAt(s *String) JavaScript {
	return NewDefaultFunction("charAt", func(args []JavaScript, ctx *Context) JavaScript {
		i := 0
		if len(args) > 0 {
			if j, ok := args[0].(*Number); ok {
				i = int(j.Data)
			}
		}
		if i < 0 || i >= charLen(s.Data) {
			return GoToString("")
		} else {
			return GoToString(string(charRune(s.Data)[i]))
		}
	})
}
func stringCharCodeAt(s *String) JavaScript {
	return NewDefaultFunction("charCodeAt", func(args []JavaScript, ctx *Context) JavaScript {
		i := 0
		if len(args) > 0 {
			if j, ok := args[0].(*Number); ok {
				i = int(j.Data)
			}
		}
		if i < 0 || i >= charLen(s.Data) {
			return NewNaN()
		} else {
			return GoToNumber(float64(charRune(s.Data)[i]))
		}
	})
}

func stringConcat(s *String) JavaScript {
	return NewDefaultFunction("concat", func(args []JavaScript, ctx *Context) JavaScript {
		if len(args) == 0 {
			return s
		}
		i := s.Data
		for _, v := range args {
			i = i + JsToString(v)
		}
		return GoToString(i)
	})
}

func stringTrimStart(s *String) JavaScript {
	return NewDefaultFunction("trimStart", func(args []JavaScript, ctx *Context) JavaScript {
		return GoToString(strings.TrimLeftFunc(s.Data, unicode.IsSpace))
	})
}

func stringTrimEnd(s *String) JavaScript {
	return NewDefaultFunction("trimEnd", func(args []JavaScript, ctx *Context) JavaScript {
		return GoToString(strings.TrimRightFunc(s.Data, unicode.IsSpace))
	})
}

func stringTrim(s *String) JavaScript {
	return NewDefaultFunction("trim", func(args []JavaScript, ctx *Context) JavaScript {
		return GoToString(strings.TrimSpace(s.Data))
	})
}

func stringPadStart(s *String) JavaScript {
	return NewDefaultFunction("padStart", func(args []JavaScript, ctx *Context) JavaScript {
		if len(args) == 0 {
			return s
		}
		length := 0
		str := " "
		if j, ok := args[0].(*Number); ok {
			length = int(j.Data)
		}
		if len(args) > 1 {
			str = JsToString(args[1])
		}
		if charLen(s.Data) >= length {
			return s
		}
		paddingNeeded := length - charLen(s.Data)
		padStr := ""
		for len([]rune(padStr)) < paddingNeeded {
			padStr += str
		}
		return GoToString(padStr + s.Data)
	})
}

func stringPadEnd(s *String) JavaScript {
	return NewDefaultFunction("padEnd", func(args []JavaScript, ctx *Context) JavaScript {
		if len(args) == 0 {
			return s
		}
		length := 0
		str := " "
		if j, ok := args[0].(*Number); ok {
			length = int(j.Data)
		}
		if len(args) > 1 {
			str = JsToString(args[1])
		}
		if charLen(s.Data) >= length {
			return s
		}
		paddingNeeded := length - charLen(s.Data)
		padStr := ""
		for len([]rune(padStr)) < paddingNeeded {
			padStr += str
		}
		return GoToString(s.Data + padStr)
	})
}

func stringLocaleCompare(s *String) JavaScript {
	return NewDefaultFunction("localeCompare", func(args []JavaScript, ctx *Context) JavaScript {
		// 确保有足够的参数
		if len(args) < 1 {
			panic(NewTypeError("Missing argument in localeCompare call"))
		}
		// 获取要比较的字符串
		target, ok := args[0].(*String)
		if !ok {
			panic(NewTypeError("Argument must be a string in localeCompare call"))
		}
		// 使用strings.Compare进行比较
		result := strings.Compare(s.Data, target.Data)
		// 根据比较结果返回JavaScript中的Number类型
		return GoToNumber(float64(result))
	})
}

func stringToUpperCase(s *String) JavaScript {
	return NewDefaultFunction("toUpperCase", func(args []JavaScript, ctx *Context) JavaScript {
		return GoToString(strings.ToUpper(s.Data))
	})
}

func stringToLowerCase(s *String) JavaScript {
	return NewDefaultFunction("toLowerCase", func(args []JavaScript, ctx *Context) JavaScript {
		return GoToString(strings.ToLower(s.Data))
	})
}

func stringIncludes(s *String) JavaScript {
	return NewDefaultFunction("includes", func(args []JavaScript, ctx *Context) JavaScript {
		if len(args) == 0 {
			return GoToBoolean(false)
		}
		i := 0
		if len(args) > 1 {
			if j, ok := args[1].(*Number); ok {
				i = int(j.Data)
			}
		}
		if i == 0 {
			return GoToBoolean(strings.Contains(s.Data, JsToString(args[0])))
		}
		return GoToBoolean(strings.Contains(s.Data[i:], JsToString(args[0])))
	})
}

func stringIndexOf(s *String) JavaScript {
	return NewDefaultFunction("indexOf", func(args []JavaScript, ctx *Context) JavaScript {
		if len(args) == 0 {
			return GoToNumber(-1)
		}
		i := 0
		if len(args) > 1 {
			if j, ok := args[1].(*Number); ok {
				i = int(j.Data)
			}
		}
		if i > charLen(s.Data) {
			return GoToNumber(-1)
		}
		if i <= 0 {
			return GoToNumber(float64(runeIndexOf(s.Data, JsToString(args[0]))))
		}
		x := runeIndexOf(string(charRune(s.Data)[i:]), JsToString(args[0]))
		if x < 0 {
			return GoToNumber(x)
		}
		return GoToNumber(float64(x + i))
	})
}

func stringLastIndexOf(s *String) JavaScript {
	return NewDefaultFunction("lastIndexOf", func(args []JavaScript, ctx *Context) JavaScript {
		if len(args) == 0 {
			return GoToNumber(-1)
		}
		i := charLen(s.Data)
		if len(args) > 1 {
			if j, ok := args[1].(*Number); ok {
				i = int(j.Data)
			}
		}
		if i < 0 {
			return GoToNumber(-1)
		}
		if i > charLen(s.Data) {
			return GoToNumber(charLen(s.Data) - 1)
		}
		return GoToNumber(float64(runeLastIndexOf(string(charRune(s.Data)[:i]), JsToString(args[0]))))
	})
}

func stringSlice(s *String) JavaScript {
	return NewDefaultFunction("slice", func(args []JavaScript, ctx *Context) JavaScript {
		if len(args) == 0 {
			return s
		}
		// 获取字符串的长度
		length := len(s.Data)
		start := 0
		end := length
		if len(args) >= 1 {
			if i, ok := args[0].(*Number); ok {
				start = int(i.Data)
			}
		}
		if len(args) > 1 {
			if i, ok := args[1].(*Number); ok {
				end = int(i.Data)
			}
		}
		// 将负索引转换为相对于末尾的正索引
		if start < 0 {
			start = length + start
			if start < 0 {
				start = 0
			}
		}
		if end < 0 {
			end = length + end
		}
		// 确保索引有效
		if start > end {
			return GoToString("")
		}
		if start < 0 || start >= length {
			return GoToString("")
		}
		if end < 0 {
			return GoToString("")
		}
		if end > length {
			end = length
		}
		// 返回子字符串
		return GoToString(s.Data[start:end])
	})
}

func stringSubstring(s *String) JavaScript {
	return NewDefaultFunction("substring", func(args []JavaScript, ctx *Context) JavaScript {
		if len(args) == 0 {
			return s
		}
		// 获取字符串的长度
		length := len([]rune(s.Data))
		start := 0
		end := length
		if len(args) >= 1 {
			if i, ok := args[0].(*Number); ok {
				start = int(i.Data)
			}
		}
		if len(args) > 1 {
			if i, ok := args[1].(*Number); ok {
				end = int(i.Data)
			}
		}
		if start < 0 {
			start = 0
		}
		if end > length || end < 0 {
			end = length
		}
		if start > end {
			start, end = end, start // 交换start和end的值，与JS的substring行为一致
		}
		if start >= length || end > length {
			return GoToString("")
		}
		// 返回子字符串
		return GoToString(string([]rune(s.Data)[start:end]))
	})
}

/**************实现迭代器接口******************/

func (a *StringIterator) Next() (v []JavaScript, done bool) {
	defer func() {
		a.index++
	}()
	if a.index >= len(a.data.Data) {
		return []JavaScript{NewUndefined(), NewUndefined()}, true
	}
	return []JavaScript{GoToNumber(a.index), GoToString(string(a.data.Data[a.index]))}, false
}

func (a *StringIterator) GetName() string {
	return "array.iterator"
}
