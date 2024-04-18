//Program 表示一个js程序

package ss

import (
	"github.com/antlr4-go/antlr/v4"
	"ss/global"
)

// Program JavaScript程序
type Program struct {
	Code            []Instruct
	file            int   //文件
	ident           int   //0加载,1修改
	delete          bool  //是否删除关键字
	toIndex         []int //跳转序列
	optionalToIndex global.Stack[int]
	optionalCode    global.Stack[int]
	modifier        modifier //修饰符
}

func newProgram(file int) *Program {
	c := make([]Instruct, 0)
	return &Program{
		Code:            c,
		file:            file,
		toIndex:         make([]int, 0),
		optionalToIndex: global.NewStack[int](),
		optionalCode:    global.NewStack[int](),
	}
}

// 将指令压入栈
func (p *Program) push(c Instruct, bpr *antlr.BaseParserRuleContext, file int) {
	if bpr != nil {
		p.Code = append(p.Code, c.setSrc(bpr, file))
	} else {
		p.Code = append(p.Code, c)
	}
}

// 将指令弹出
func (p *Program) pop() (item Instruct, ok bool) {
	size := len(p.Code)
	ok = size > 0
	if ok {
		item = p.Code[size-1]
		p.Code = p.Code[:size-1]
	}
	return
}

// 指令窥视
func (p *Program) peek() (item Instruct, ok bool) {
	size := len(p.Code)
	ok = size > 0
	if ok {
		item = p.Code[size-1]
	}
	return
}

// 将指令压入栈头部
func (p *Program) pushHead(c Instruct, bpr *antlr.BaseParserRuleContext, file int) {
	if bpr != nil {
		p.Code = append([]Instruct{c.setSrc(bpr, file)}, p.Code...)
	} else {
		p.Code = append([]Instruct{c}, p.Code...)
	}
}
