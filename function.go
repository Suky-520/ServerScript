//Function 结构体表示JavaScript中用户自定义函数

package ss

import (
	"fmt"
	"ss/collection"
	"sync"
)

// Function JavaScript用户通过js代码定义的函数
type Function struct {
	name       string               //函数名
	formal     []string             //形参
	body       []Instruct           //函数体
	props      JavaScriptProperties //函数属性
	ctx        *Context             //函数声明的上下文
	this       *Object              //函数的this对象
	sync.Mutex                      //同步
}

var _ JavaScript = &Function{}
var _ JavaScriptFunction = &Function{}

// 新建函数
func newFunction(formal []string, body []Instruct) *Function {
	f := &Function{formal: formal, body: body, props: collection.NewArrayMap[string, JavaScript]()}
	return f
}

// Call JavaScript函数调用
func (f *Function) Call(state int, args []JavaScript, ctx *Context) JavaScript {
	if state == 0 {
		return f.call(args, ctx)
	} else if state == 1 {
		AsyncCall(func() {
			f.call(args, ctx)
		})
	} else if state == 2 {
		f.Lock()
		defer f.Unlock()
		return f.call(args, ctx)
	}
	return NewUndefined()
}

// js函数调用
func (f *Function) call(args []JavaScript, ctx *Context) JavaScript {
	c := ctx.newCtx()
	c.code = f.body
	if f.ctx.vars.Has(f.name) && f.ctx.parent == ctx {
		c.parent = f.ctx.parent
	} else {
		c.parent = f.ctx
	}
	c.size = len(f.body)
	c.statement = Call
	f.initVar(args, c)
	if f.this != nil {
		c.vars.Set("this", &Variable{Value: f.this, Modifier: Const})
	}
	c.exec()
	v, _ := c.value.Pop()
	c = nil
	return v
}

func (f *Function) initVar(args []JavaScript, ctx *Context) {
	for i, v := range f.formal {
		if args == nil || len(args) < i+1 {
			ctx.vars.Set(v, &Variable{Value: NewUndefined(), Modifier: Let})
		} else {
			ctx.vars.Set(v, &Variable{Value: args[i], Modifier: Let})
		}
	}
	if args == nil {
		ctx.vars.Set("arguments", &Variable{Value: &Array{Data: make([]JavaScript, 0)}, Modifier: Let})
	} else {
		ctx.vars.Set("arguments", &Variable{Value: &Array{Data: args}, Modifier: Let})
	}
}

func (f *Function) GetName() string {
	return f.name
}

func (f *Function) ToString() string {
	return fmt.Sprintf("[Function: %v]", f.name)
}

func (f *Function) GetProperties() JavaScriptProperties {
	return f.props
}
func (f *Function) MarshalJSON() ([]byte, error) {
	return []byte("null"), nil
}
