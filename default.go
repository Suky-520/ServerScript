//默认的类生成函数,可以简化类的创建

package ss

import (
	"ss/collection"
)

/************** DefaultClass ******************/

// DefaultClass 默认的JavaScript类
type DefaultClass struct {
	name        string
	constructor func(args []JavaScript, ctx *Context) JavaScript
}

var _ JavaScript = &DefaultClass{}
var _ JavaScriptClass = &DefaultClass{}

// NewDefaultClass 新建一个JavaScript类
func NewDefaultClass(name string, constructor func(args []JavaScript, ctx *Context) JavaScript) JavaScript {
	return &DefaultClass{
		name:        name,
		constructor: constructor,
	}
}

func (d *DefaultClass) GetName() string {
	return d.name
}
func (d *DefaultClass) TypeOf() JavaScriptType {
	return JavaScriptType{Name: d.name, Native: true, Type: ClassType}
}

func (d *DefaultClass) Constructor(args []JavaScript, ctx *Context) JavaScript {
	return d.constructor(args, ctx)
}
func (d *DefaultClass) Print() string {
	return printClass(d)
}

/************** DefaultObject ******************/

// DefaultObject 默认的JavaScript对象
type DefaultObject struct {
	name string
	prop JavaScriptProperties
}

var _ JavaScript = &DefaultObject{}

// NewDefaultObject 新建一个JavaScript对象
func NewDefaultObject(name string) *DefaultObject {
	return &DefaultObject{
		name: name,
		prop: collection.NewArrayMap[string, JavaScript](),
	}
}

func (d *DefaultObject) GetName() string {
	return d.name
}

func (d *DefaultObject) AddProperty(name string, value JavaScript) {
	d.prop.Set(name, value)
}

func (d *DefaultObject) GetProperty(name string) JavaScript {
	if v, ok := d.prop.Get(name); ok {
		return v
	}
	return NewUndefined()
}

func (d *DefaultObject) TypeOf() JavaScriptType {
	return JavaScriptType{Name: d.name, Native: true, Type: ObjectType}
}

func (d *DefaultObject) GetProperties() JavaScriptProperties {
	return d.prop
}

/************** DefaultFunction ******************/

// DefaultFunction 默认的JavaScript函数
type DefaultFunction struct {
	name string
	call func(args []JavaScript, ctx *Context) JavaScript
}

var _ JavaScript = &DefaultFunction{}
var _ JavaScriptFunction = &DefaultFunction{}

// NewDefaultFunction 新建一个JavaScript函数
func NewDefaultFunction(name string, call func(args []JavaScript, ctx *Context) JavaScript) JavaScript {
	return &DefaultFunction{
		name: name,
		call: call,
	}
}

func (d *DefaultFunction) GetName() string {
	return d.name
}
func (d *DefaultFunction) Print() string {
	return printFunction(d)
}
func (d *DefaultFunction) TypeOf() JavaScriptType {
	return JavaScriptType{Name: d.name, Native: true, Type: FunctionType}
}

func (d *DefaultFunction) Call(state int, args []JavaScript, ctx *Context) JavaScript {
	return d.call(args, ctx)
}
