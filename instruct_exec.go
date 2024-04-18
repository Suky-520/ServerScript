//虚拟机指令的执行过程

package ss

import (
	"fmt"
	"slices"
	"strconv"
)

//本文件用于执行虚拟机指令

// js语句执行结束
func eos(ctx *Context) {
	ctx.value.Clear()
	ctx.propertyNames.Clear()
	ctx.ellipsisCount = 0
}

// 执行运算
func (o Operate) exec(ctx *Context) {
	if o == NullCoalesce {
		nullCoalesce(ctx)
		return
	} else if o == Or {
		v, _ := ctx.value.Peek()
		if !JsToBool(v) {
			ctx.value.Pop()
			ctx.index++
		}
		ctx.next()
		return
	} else if o == And {
		v, _ := ctx.value.Peek()
		if JsToBool(v) {
			ctx.value.Pop()
			ctx.index++
		}
		ctx.next()
		return
	}

	if slices.Contains(unary, o) {
		//单目计算
		v1, _ := ctx.value.Pop()
		v := JsToCalculate(v1, nil, o)
		ctx.value.Push(v)
	} else {
		//双目计算
		v1, _ := ctx.value.Pop()
		v2, _ := ctx.value.Pop()
		v := JsToCalculate(v2, v1, o)
		ctx.value.Push(v)
	}
	ctx.next()
}

// 空值合并运算
func nullCoalesce(ctx *Context) {
	v, _ := ctx.value.Peek()
	t := JsToType(v)
	if t.Type == UndefinedType || t.Type == NullType {
		ctx.index = ctx.index + 1
		ctx.value.Pop()
	}
	ctx.next()
}

// 执行命令,带有src错误提示
func (c cmd) exec(ctx *Context) {
	switch c.c {
	case Throw:
		throw(ctx)
	case Break:
		breakCmd(ctx)
	case ExportDefault:
		exportDefault(ctx)
	case ExportDeclaration:
		exportDeclaration(ctx)
	case This:
		this(ctx)
	}
	ctx.next()
}

// 执行命令
func (c command) exec(ctx *Context) {
	switch c {
	case Eos:
		eos(ctx)
	case While:
		while(ctx)
	case If: //条件判断
		ifCmd(ctx)
	case Continue:
		continueCmd(ctx)
	case Ellipsis:
		ellipsis(ctx)
	case Return:
		ret(ctx)
	default:
		panic(NewSyntaxError(fmt.Sprint("不支持的运算指令:", c)))
	}
	ctx.next()
}

// 加载字面量值
func (l loadVal) exec(ctx *Context) {
	ctx.value.Push(l.value)
	ctx.next()
}

// 导入模块
func (m importModule) exec(ctx *Context) {
	if m.code != nil && len(m.code) > 0 {
		//用户自定义模块
		md := loadUserModule(m, ctx)
		initModule(m, md, ctx)
	} else {
		//本地原生模块
		md := loadNativeModule(m.From, ctx)
		initModule(m, md, ctx)
	}
	ctx.next()
}

// 加载原生模块
func loadNativeModule(from string, ctx *Context) *Module {
	module := getGlobalNativeModule(from)
	if module == nil {
		panic(NewSyntaxError("module " + from + " not found"))
	}
	return module
}

// 加载用户模块
func loadUserModule(m importModule, ctx *Context) *Module {
	if mod := getGlobalModule(m.File); mod != nil {
		md := mod.(*Module)
		return md
	}
	ctx1 := newCtx()
	ctx1.vm = ctx.vm
	c := execBlock(m.code, ctx1, true)
	mod := newModule(c.vars, c.exports)
	//销毁
	c = nil
	SetGlobalModule(m.File, mod)
	//清空
	moduleFiles = make(map[string]bool)
	return mod
}

// 初始化模块
func initModule(m importModule, mod *Module, ctx *Context) {
	if m.DefaultName != "" {
		if v, ok := mod.export.Get("default"); ok {
			ctx.vars.Set(m.DefaultName, v)
		} else {
			panic(NewSyntaxError("The requested module '" + m.From + "' does not provide an exports named 'def'"))
		}
	}
	for k, v := range m.Items {
		v1, ok := mod.export.Get(k)
		if !ok {
			panic(NewSyntaxError("The requested module '" + m.From + "' does not provide an exports named '" + k + "'"))
		}
		if v != "" {
			ctx.vars.Set(v, v1)
		} else {
			ctx.vars.Set(k, v1)
		}
	}
}

// 函数调用
func (c call) exec(ctx *Context) {
	args, _ := ctx.value.PopMany(c.index + ctx.ellipsisCount)
	ctx.ellipsisCount = 0
	if f, ok := ctx.value.Pop(); ok {
		if c.constructor {
			if f1, ok := f.(JavaScriptClass); ok {
				//构造函数
				if r := f1.Constructor(args, ctx); r != nil {
					ctx.value.Push(r)
				} else {
					ctx.value.Push(NewUndefined())
				}
			} else {
				panic(NewTypeError(" is not a class"))
			}
		} else {
			if f1, ok := f.(JavaScriptFunction); ok {
				//普通函数
				if r := f1.Call(c.state, args, ctx); r != nil {
					ctx.value.Push(r)
				} else {
					ctx.value.Push(NewUndefined())
				}
			} else if c.optional {
				ctx.isSkip = true
				ctx.value.Push(NewUndefined())
			} else {
				panic(NewTypeError(" is not a function"))
			}
		}
	}
	ctx.next()
}

// 执行try-catch-finally语句
func (t try) exec(ctx *Context) {
	if err := execTry(t.try, ctx); err != nil {
		c := ctx.newCtx()
		c.code = t.catch.body
		c.size = len(t.catch.body)
		c.vars.Set(t.catch.formal, &Variable{Value: err.(JavaScript), Modifier: Let})
		c.exec()
		c = nil
	}
	execBlock(t.finally, ctx, false)
	ctx.next()
}

// 执行try语句块
func execTry(code []Instruct, ctx *Context) (r any) {
	defer func() {
		if r = recover(); r != nil {
			ctx.vm.error = nil
		}
	}()
	execBlock(code, ctx, false)
	return
}

// 创建数组
func (t array) exec(ctx *Context) {
	c, _ := ctx.value.PopMany(int(t) + ctx.ellipsisCount)
	a := GoToArray(c)
	ctx.value.Push(a)
	ctx.ellipsisCount = 0
	ctx.next()
}

// 创建函数
func (t function) exec(ctx *Context) {
	f := newFunction(t.formal, t.body)
	f.ctx = ctx
	ctx.value.Push(f)
	ctx.next()
}

// 执行js语句块
func (t block) exec(ctx *Context) {
	execBlock(t, ctx, false)
	ctx.next()
}

// 执行js语句块
func execBlock(code []Instruct, ctx *Context, ret bool) *Context {
	c := ctx.newCtx()
	c.code = code
	c.size = len(code)
	c.exec()
	if ret {
		return c
	}
	c = nil
	return nil
}

// 创建对象
func (c createObject) exec(ctx *Context) {
	ctx1 := execBlock(c, ctx, true)
	obj := NewObject()
	for _, k := range ctx1.vars.Keys() {
		v, _ := ctx1.vars.Get(k)
		if t, ok := v.Value.(*Function); ok {
			t.this = obj
		}
		obj.SetProperty(k, v.Value)
	}
	ctx.value.Push(obj)
	ctx1 = nil
	ctx.next()
}

// 导出元素
func (t exportItem) exec(ctx *Context) {
	v, _ := ctx.vars.Get(t.name)
	if v == nil {
		panic(NewSyntaxError("Export '" + t.name + "' is not defined in module"))
	}
	name := t.name
	if t.as != "" {
		name = t.as
	}
	if ctx.exports.Has(name) {
		panic(NewSyntaxError("Duplicate exports of '" + name + "'"))
	}
	ctx.exports.Set(name, v)
	ctx.next()
}

// 迭代重定向
func (t to) exec(ctx *Context) {
	ctx.index = ctx.index + (int)(t)
}

// 执行迭代语句
func (t iteration) exec(ctx *Context) {
	c := ctx.newCtx()
	c.code = t
	c.size = len(t)
	c.statement = While
	c.exec()
	c = nil
	ctx.next()
}

// 创建变量
func (c createVar) exec(ctx *Context) {
	value, _ := ctx.value.Pop()
	if value == nil {
		value = NewUndefined()
	}
	v := &Variable{Value: value, Modifier: c.mod}
	if _, ok := ctx.vars.Get(c.name); ok && ctx.ellipsisCount == 0 {
		panic(NewSyntaxError("Identifier '" + c.name + "' has already been declared"))
	} else {
		if f, ok := value.(*Function); ok {
			f.name = c.name
			v.Value = f
		}
		ctx.vars.Set(c.name, v)
	}
	ctx.next()
}

// 修改变量
func (m modifyVar) exec(ctx *Context) {
	value, _ := ctx.value.Pop()
	modifyCtxVar(m.name, value, ctx)
	ctx.next()
}

// 修改对象属性
func (m modifyProperty) exec(ctx *Context) {
	obj, _ := ctx.value.Pop()
	value, _ := ctx.value.Pop()
	t := JsToType(obj)
	if value == nil {
		panic(NewReferenceError(m.name + " is not defined"))
	} else if _, ok := obj.(*Undefined); ok {
		//不能给Undefined设置属性
		panic(NewTypeError("Cannot set properties of undefined (setting '" + m.name + "')"))
	} else if t.Type == ObjectType || t.Type == ArrayType || t.Type == BufferType {
		JsSetProperty(obj, m.name, value)
	} else {
		//不能给基本类型设置属性
		panic(NewTypeError("Cannot set properties of " + strconv.Itoa(int(t.Type)) + " (setting '" + m.name + "')"))
	}
	ctx.next()
}

// 修改上下文变量
func modifyCtxVar(name string, value JavaScript, ctx *Context) {
	c := ctx
	for {
		if v, ok := c.vars.Get(name); ok {
			if v.Modifier == Const {
				panic(NewTypeError("Assignment to constant Variable."))
			}
			v.Value = value
			return
		} else {
			if c.parent != nil {
				c = c.parent
			} else {
				panic(NewReferenceError(name + " is not defined"))
				return
			}
		}
	}
}

// 加载变量
func loadCtxVar(name string, ctx *Context) bool {
	c := ctx
	for {
		if v, ok := c.vars.Get(name); ok {
			ctx.value.Push(v.Value)
			return true
		} else {
			if c.parent != nil {
				c = c.parent
			} else {
				return false
			}
		}
	}
}

// 获取变量
func (l loadVar) exec(ctx *Context) {
	if loadCtxVar(l.name, ctx) {
		ctx.next()
		return
	}
	v := getGlobalObject(l.name)
	if v != nil {
		ctx.value.Push(v.(JavaScript))
		ctx.next()
		return
	}
	panic(NewReferenceError(l.name + " is not defined"))
}

// 指令跳转
func (t skip) exec(ctx *Context) {
	if ctx.isSkip {
		ctx.index = int(t) + ctx.index
		ctx.isSkip = false
	}
	ctx.next()
}

// 执行可选链校验
func optional(value JavaScript, ctx *Context) bool {
	if _, ok := value.(*Undefined); ok {
		ctx.value.Push(NewUndefined())
		ctx.isSkip = true
		return true
	}
	return false
}

// 获取/删除对象属性
func (l property) exec(ctx *Context) {
	name := l.name
	if name == "" {
		v, _ := ctx.value.Pop()
		name = JsToString(v)
	}
	ctx.propertyNames.Push(name)
	value, _ := ctx.value.Pop()
	t := JsToType(value)
	if t.Type == UndefinedType || t.Type == NullType {
		//可选链校验
		if l.optional {
			ctx.isSkip = true
			ctx.value.Push(NewUndefined())
			ctx.next()
			return
		}
		panic(NewTypeError("Cannot read properties of " + JsToString(value) + " (reading '" + l.name + "')"))
	}
	if l.delete {
		o := JsDeleteProperty(value, name)
		ctx.value.Push(o)
	} else {
		o := JsGetProperty(value, name)
		//可选链校验
		if l.optional && optional(o, ctx) {
			ctx.next()
			return
		}
		ctx.value.Push(o)
	}
	ctx.next()
}

// continue 实现
func continueCmd(ctx *Context) {
	index := ctx.index
	c := ctx
	for {
		if c.statement == While {
			break
		} else {
			c.index = c.size
			if c.parent != nil {
				c = c.parent
			} else {
				//恢复
				ctx.index = index
				panic(NewSyntaxError("Illegal continue statement: no surrounding iteration statement"))
			}
		}
	}
}

// break 实现
func breakCmd(ctx *Context) {
	index := ctx.index
	c := ctx
	for {
		if c.statement == While {
			c.index = c.size
			break
		} else {
			c.index = c.size
			if c.parent != nil {
				c = c.parent
			} else {
				//恢复
				ctx.index = index
				panic(NewSyntaxError("Illegal break statement"))
			}
		}
	}
}

// 循环指令
func while(ctx *Context) {
	value, _ := ctx.value.Pop()
	if v, ok := value.(JavaScriptBool); !ok || !v.ToBool() {
		ctx.index = ctx.size
	}
}

// if指令
func ifCmd(ctx *Context) {
	v, _ := ctx.value.Pop()
	if JsToBool(v) {
		ctx.index = ctx.index + 1
	}
}

// throw指令
func throw(ctx *Context) {
	v, _ := ctx.value.Pop()
	if e, ok := v.(*Error); ok {
		panic(e)
	}
	panic(v)
}

// 默认导出
func exportDefault(ctx *Context) {
	v, _ := ctx.value.Pop()
	if ctx.parent.exports.Has("default") {
		panic(NewSyntaxError("Duplicate exports of 'def'"))
	}
	ctx.parent.exports.Set("default", &Variable{Value: v, Modifier: Const})
}

// 导出变量
func exportDeclaration(ctx *Context) {
	for _, k := range ctx.vars.Keys() {
		v, _ := ctx.vars.Get(k)
		if ctx.parent.vars.Has(k) {
			panic(NewSyntaxError("Identifier '" + k + "' has already been declared"))
		} else {
			ctx.parent.vars.Set(k, v)
			if ctx.parent.exports.Has(k) {
				panic(NewSyntaxError("exports Identifier '" + k + "' has already been declared"))
			} else {
				ctx.parent.exports.Set(k, v)
			}
		}
	}
}

// 解构操作
func ellipsis(ctx *Context) {
	va, _ := ctx.value.Pop()
	switch va.(type) {
	case *Object:
		obj := va.(*Object)
		for _, key := range obj.props.Keys() {
			v, _ := obj.props.Get(key)
			ctx.vars.Set(key, &Variable{Value: v, Modifier: Let})
		}
		ctx.ellipsisCount = ctx.ellipsisCount + len(obj.props.Keys())
	case *Array:
		arr := va.(*Array)
		for _, v := range arr.Data {
			ctx.value.Push(v)
		}
		ctx.ellipsisCount = ctx.ellipsisCount + len(arr.Data) - 1
	default:
		panic(NewRangeError("该类型不支持解构"))
	}
}

// this指令
func this(ctx *Context) {
	if loadCtxVar("this", ctx) {
		return
	}
	panic(NewReferenceError("this is not defined"))
}

// return 指令
func ret(ctx *Context) {
	v, _ := ctx.value.Peek()
	c := ctx
	for {
		//将上下文结束
		c.index = c.size
		if c.statement == Call {
			if c != ctx {
				c.value.Push(v)
			}
			break
		} else {
			c = c.parent
		}
		if c == nil {
			break
		}
	}
	ctx.next()
}

// for of 指令
func (f forOf) exec(ctx *Context) {
	if v, ok := ctx.value.Peek(); ok {
		var jsi JavaScriptIterator
		if i, ok := v.(JavaScriptIterator); ok {
			jsi = i
		} else {
			if k, ok := v.(JavaScriptGetIterator); ok {
				ctx.value.Pop()
				jsi = k.GetIterator()
				ctx.value.Push(jsi)
			} else {
				ctx.index = ctx.size
				ctx.next()
				return
			}
		}
		if values, ok := jsi.Next(); ok {
			ctx.index = ctx.size
			ctx.next()
			return
		} else {
			if len(f.vars) == 0 {
				ctx.index = ctx.size
				ctx.next()
				return
			}
			if f.list {
				for i := range f.vars {
					vx, _ := ctx.vars.Get(f.vars[i])
					if i < len(values) {
						vx.Value = values[i]
					} else {
						vx.Value = NewUndefined()
					}
				}
			} else {
				//不是数组
				vx, _ := ctx.vars.Get(f.vars[0])
				if len(values) == 0 {
					vx.Value = NewUndefined()
				} else if len(values) == 1 {
					vx.Value = values[0]
				} else {
					vx.Value = values[1]
				}
				//vx.Value = GoToArray(values)
			}
		}
	}
	ctx.next()
}
