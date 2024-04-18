// Error结构体是JavaScript中Error对象的底层实现。
// 该结构体包含错误名称、错误消息和错误堆栈等信息。

package ss

import (
	"ss/collection"
)

// Error JavaScript中的错误类型
type Error struct {
	name    string
	message string
	stack   string
	prop    JavaScriptProperties
}

var _ JavaScript = Error{}

// 初始化Error对象
func init() {
	list := []string{"TypeError", "Error", "SyntaxError", "ReferenceError", "RangeError", "URIError"}
	for _, v := range list {
		initError(v)
	}
}

// 初始化Error
func initError(name string) {
	setGlobalObject(name, NewDefaultClass(name, func(args []JavaScript, ctx *Context) JavaScript {
		if len(args) == 0 {
			return NewError(name, "")
		}
		return NewError(name, JsToString(args[0]))
	}))
}

// NewError 新建Error
func NewError(name string, message string) Error {
	o := Error{
		message: message,
		name:    name,
		prop:    collection.NewArrayMap[string, JavaScript](),
	}
	o.prop.Set("name", GoToString(name))
	o.prop.Set("message", GoToString(message))
	return o
}
func NewTypeError(message string) Error {
	return NewError("TypeError", message)
}
func NewSyntaxError(message string) Error {
	return NewError("SyntaxError", message)
}
func NewReferenceError(message string) Error {
	return NewError("ReferenceError", message)
}
func NewRangeError(message string) Error {
	return NewError("RangeError", message)
}
func NewURIError(message string) Error {
	return NewError("URIError", message)
}

func (a Error) GetName() string {
	return a.name
}

func (a Error) GetProperty(name string) JavaScript {
	if v, ok := a.prop.Get(name); ok {
		return v
	}
	return NewUndefined()
}

func (a Error) GetProperties() JavaScriptProperties {
	return a.prop
}

func (a Error) ToString() string {
	return a.name + " : " + a.message
}
func (a Error) Print() string {
	return RedStyle + a.name + " : " + a.message + EndStyle
}
