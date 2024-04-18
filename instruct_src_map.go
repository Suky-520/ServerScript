//每一个虚拟机指令都可以包含一个SrcMap对象，用于记录错误位置

package ss

import (
	"github.com/antlr4-go/antlr/v4"
)

// SrcMap 的接口
type SrcMap interface {
	getSrcMap() Src
}

// Src 源码映射
type Src struct {
	File        int    //源码对应的js文件序号
	StartLine   int    //开始行数
	StartColumn int    //开始列数
	token       string //关键代码
	stopLine    int    //结束行数
	stopColumn  int    //结束列数
	err         string //错误信息
}

var _ CodeCompiler = &Src{}

// 新建一个src对象
func createSrc(bpr *antlr.BaseParserRuleContext, file int) *Src {
	s := Src{File: file, StartColumn: bpr.GetStart().GetColumn(), StartLine: bpr.GetStart().GetLine(), token: bpr.GetText(), stopColumn: bpr.GetStop().GetColumn(), stopLine: bpr.GetStop().GetLine()}
	return &s
}

func (m importModule) getSrc() *Src {
	return m.src
}
func (l loadVar) getSrc() *Src {
	return l.src
}
func (o Operate) getSrc() *Src {
	return nil
}
func (l loadVal) getSrc() *Src {
	return l.src
}
func (c call) getSrc() *Src {
	return c.src
}
func (t try) getSrc() *Src {
	return nil
}
func (t array) getSrc() *Src {
	return nil
}
func (t function) getSrc() *Src {
	return nil
}
func (t block) getSrc() *Src {
	return nil
}
func (m modifyProperty) getSrc() *Src {
	return m.src
}
func (c createObject) getSrc() *Src {
	return nil
}
func (t exportItem) getSrc() *Src {
	return t.src
}
func (t to) getSrc() *Src {
	return nil
}
func (t skip) getSrc() *Src {
	return nil
}
func (f forOf) getSrc() *Src {
	return nil
}
func (t iteration) getSrc() *Src {
	return nil
}
func (c command) getSrc() *Src {
	return nil
}
func (c createVar) getSrc() *Src {
	return c.src
}
func (m modifyVar) getSrc() *Src {
	return m.src
}
func (c cmd) getSrc() *Src {
	return c.src
}
func (l property) getSrc() *Src {
	return l.src
}
func (m importModule) setSrc(bpr *antlr.BaseParserRuleContext, file int) Instruct {
	m.src = createSrc(bpr, file)
	return m
}
func (l loadVar) setSrc(bpr *antlr.BaseParserRuleContext, file int) Instruct {
	l.src = createSrc(bpr, file)
	return l
}
func (o Operate) setSrc(bpr *antlr.BaseParserRuleContext, file int) Instruct {
	return o
}
func (l loadVal) setSrc(bpr *antlr.BaseParserRuleContext, file int) Instruct {
	l.src = createSrc(bpr, file)
	return l
}
func (c call) setSrc(bpr *antlr.BaseParserRuleContext, file int) Instruct {
	c.src = createSrc(bpr, file)
	return c
}
func (t try) setSrc(bpr *antlr.BaseParserRuleContext, file int) Instruct {
	return t
}
func (t array) setSrc(bpr *antlr.BaseParserRuleContext, file int) Instruct {
	return t
}
func (t function) setSrc(bpr *antlr.BaseParserRuleContext, file int) Instruct {
	return t
}
func (t block) setSrc(bpr *antlr.BaseParserRuleContext, file int) Instruct {
	return t
}
func (c createObject) setSrc(bpr *antlr.BaseParserRuleContext, file int) Instruct {
	return c
}
func (t exportItem) setSrc(bpr *antlr.BaseParserRuleContext, file int) Instruct {
	t.src = createSrc(bpr, file)
	return t
}
func (c createVar) setSrc(bpr *antlr.BaseParserRuleContext, file int) Instruct {
	c.src = createSrc(bpr, file)
	return c
}
func (f forOf) setSrc(bpr *antlr.BaseParserRuleContext, file int) Instruct {
	return f
}
func (t skip) setSrc(bpr *antlr.BaseParserRuleContext, file int) Instruct {
	return t
}
func (t to) setSrc(bpr *antlr.BaseParserRuleContext, file int) Instruct {
	return t
}
func (t iteration) setSrc(bpr *antlr.BaseParserRuleContext, file int) Instruct {
	return t
}
func (m modifyVar) setSrc(bpr *antlr.BaseParserRuleContext, file int) Instruct {
	m.src = createSrc(bpr, file)
	return m
}
func (l property) setSrc(bpr *antlr.BaseParserRuleContext, file int) Instruct {
	l.src = createSrc(bpr, file)
	return l
}
func (c command) setSrc(bpr *antlr.BaseParserRuleContext, file int) Instruct {
	return c
}
func (c cmd) setSrc(bpr *antlr.BaseParserRuleContext, file int) Instruct {
	c.src = createSrc(bpr, file)
	return c
}
func (m modifyProperty) setSrc(bpr *antlr.BaseParserRuleContext, file int) Instruct {
	m.src = createSrc(bpr, file)
	return m
}
