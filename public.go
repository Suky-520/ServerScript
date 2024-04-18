// 一些工具函数

package ss

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// 全局打印的字体颜色
const (
	GreenStyle  = "\033[32m"
	YellowStyle = "\033[33m"
	BoldStyle   = "\033[1m"
	BlueStyle   = "\033[34m"
	GreyStyle   = "\033[90m"
	RedStyle    = "\033[31m"
	PurpleStyle = "\033[35m"
	Underline   = "\033[4m"
	EndStyle    = "\033[0m"
)

// JsToGo js类型转换为go类
func JsToGo(val JavaScript) any {
	switch v := val.(type) {
	case *String:
		return v.Data
	case *Boolean:
		return v.Data
	case *Number:
		return v.Data
	case *Array:
		ar := make([]any, len(v.Data))
		for i, item := range v.Data {
			ar[i] = JsToGo(item)
		}
		return ar
	case *Object:
		obj := make(map[string]any)
		for _, key := range v.props.Keys() {
			item, _ := v.props.Get(key)
			obj[key] = JsToGo(item)
		}
		return obj
	case *Function:
		return v
	default:
		return nil
	}
}

// GoToJs 将go类型转换为js类型
func GoToJs(val any) JavaScript {
	if j, ok := val.(JavaScript); ok {
		return j
	}
	vv := reflect.ValueOf(val)
	if vv.Kind() == reflect.Slice {
		arr := make([]JavaScript, vv.Len())
		for i := 0; i < vv.Len(); i++ {
			element := vv.Index(i)
			arr[i] = GoToJs(element.Interface())
		}
		return GoToArray(arr)
	}
	switch v := val.(type) {
	case bool:
		return GoToBoolean(v)
	case string:
		return GoToString(v)
	case float32:
		return GoToNumber(float64(v))
	case float64:
		return GoToNumber(v)
	case int:
		return GoToNumber(float64(v))
	case int8:
		return GoToNumber(float64(v))
	case int16:
		return GoToNumber(float64(v))
	case int32:
		return GoToNumber(float64(v))
	case int64:
		return GoToNumber(float64(v))
	case []any:
		arr := make([]JavaScript, 0)
		for _, item := range v {
			arr = append(arr, GoToJs(item))
		}
		return GoToArray(arr)
	case nil:
		return NewNull()
	case map[string]any:
		object := NewObject()
		for key, value := range v {
			object.SetProperty(key, GoToJs(value))
		}
		return object
	case map[string]string:
		object := NewObject()
		for key, value := range v {
			object.SetProperty(key, GoToString(value))
		}
		return object
	case []string:
		arr := make([]JavaScript, 0)
		for _, item := range v {
			arr = append(arr, GoToString(item))
		}
		return GoToArray(arr)
	default:
		return GoToString(fmt.Sprint(v))
	}
}

// JsonToJs 将json字符串转换为js对象
func JsonToJs(str string) JavaScript {
	var data interface{}
	err := json.Unmarshal([]byte(str), &data)
	if err != nil {
		panic(NewSyntaxError("Invalid JSON"))
	}
	return GoToJs(data)
}

// PrintError 打印错误信息
func PrintError(msg string) {
	fmt.Println("\033[31m" + msg + "\033[0m")
}

// PrintErrorStack 打印堆栈
func PrintErrorStack(file int, line, column int, token string) {
	fmt.Printf("	%v:%v:%v %v \r\n", getSourceFilePath(file), line, column, token)
}
