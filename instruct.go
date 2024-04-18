// Instruct是ss虚拟机的指令接口

package ss

import "github.com/antlr4-go/antlr/v4"

// Instruct 指令接口
type Instruct interface {
	exec(ctx *Context)
	getSrc() *Src
	setSrc(bpr *antlr.BaseParserRuleContext, file int) Instruct
	CodeCompiler
}

// 语句块 22
type block []Instruct

// 迭代语句块 23
type iteration []Instruct

// 对象创建 24
type createObject []Instruct

// 指令跳转 25
type to int

// 创建数组 26
type array int

// 加载值 27
type loadVal struct {
	value JavaScript
	src   *Src
}

// 加载变量 28
type loadVar struct {
	name string
	src  *Src
}

// 指令 29
type cmd struct {
	c   command
	src *Src
}

// 修改变量 30
type modifyVar struct {
	name string
	src  *Src
}

// 创建变量 31
type createVar struct {
	mod  modifier
	name string
	src  *Src
}

// 修改对象属性 32
type modifyProperty struct {
	name string
	src  *Src
}

// 操作属性 33
type property struct {
	name     string //为空时表示名称从value中取
	delete   bool   //是否删除
	optional bool   //是否可选链判断
	src      *Src
}

// 函数调用 34
type call struct {
	index       int  // 序号,表示函数的实参开始的下标
	state       int  //执行状态:0正常执行,1异步执行,2同步加锁
	constructor bool //是否构造函数调用
	optional    bool //是否可选链判断
	src         *Src
}

// 导入 35
type importModule struct {
	DefaultName string
	Namespace   string
	Items       map[string]string
	From        string
	File        string
	code        []Instruct
	src         *Src
}

// 导出项 36
type exportItem struct {
	name string
	as   string
	src  *Src
}

// 函数声明 37
type function struct {
	formal []string   //形参
	body   []Instruct //函数体
}

// try语句块 38
type try struct {
	try     []Instruct
	finally []Instruct
	catch   catch
}

// catch语句块
type catch struct {
	formal string
	body   []Instruct
}

// 指令跳转 39
type skip int

// forOf循环 40
type forOf struct {
	vars []string
	list bool
}
