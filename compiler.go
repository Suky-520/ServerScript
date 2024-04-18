//将ss语法树转换为ss虚拟机指令
//使用antlr的访问者模式进行遍历

package ss

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"math"
	"os"
	"path/filepath"
	"ss/parser"
	"strconv"
	"strings"
)

// 模块文件的缓存，用于判断是否存在循环引用
var moduleFiles = make(map[string]bool)

// 校验是否出现循环引用
func getModuleFile(file int, from string) string {
	if getGlobalNativeModule(from) != nil {
		return ""
	}
	src := getSourceFilePath(file)
	file1 := filepath.Join(filepath.Dir(src), from)
	if _, ok := moduleFiles[file1]; ok {
		f := ""
		for k, _ := range moduleFiles {
			if f == "" {
				f = k
			} else {
				f = f + "," + k
			}
		}
		panic("循环引用 " + f)
	}
	moduleFiles[file1] = true
	return file1
}

func (p *Program) compile(source parser.ISourceElementsContext) {
	defer func() {
		if err := recover(); err != nil {
			PrintError("编译错误: " + fmt.Sprint(err))
			os.Exit(1)
		}
	}()
	for _, element := range source.AllSourceElement() {
		p.Statement(element.Statement().(*parser.StatementContext))
	}
}

func (p *Program) Statement(element *parser.StatementContext) {
	c := element.GetChild(0)
	switch c.(type) {
	case *parser.BlockContext:
		p.Block(c.(*parser.BlockContext))
	case *parser.VariableStatementContext:
		p.VariableStatement(c.(*parser.VariableStatementContext))
	case *parser.ImportStatementContext:
		p.ImportStatement(c.(*parser.ImportStatementContext))
	case *parser.ExportDeclarationContext:
		p.ExportDeclaration(c.(*parser.ExportDeclarationContext))
	case *parser.ExportItemsContext:
		p.ExportItems(c.(*parser.ExportItemsContext))
	case *parser.ExportBlockContext:
		p.ExportBlock(c.(*parser.ExportBlockContext))
	case *parser.ExportDefaultDeclarationContext:
		p.ExportDefaultDeclaration(c.(*parser.ExportDefaultDeclarationContext))
	case *parser.FunctionDeclarationContext:
		p.FunctionDeclaration(c.(*parser.FunctionDeclarationContext))
	case *parser.ExpressionStatementContext:
		p.ExpressionStatement(c.(*parser.ExpressionStatementContext))
	case *parser.IfStatementContext:
		p.IfStatement(c.(*parser.IfStatementContext))
	case *parser.ContinueStatementContext:
		p.ContinueStatement(c.(*parser.ContinueStatementContext))
	case *parser.BreakStatementContext:
		p.BreakStatement(c.(*parser.BreakStatementContext))
	case *parser.ReturnStatementContext:
		p.ReturnStatement(c.(*parser.ReturnStatementContext))
	case *parser.SwitchStatementContext:
		p.SwitchStatement(c.(*parser.SwitchStatementContext))
	case *parser.ThrowStatementContext:
		p.ThrowStatement(c.(*parser.ThrowStatementContext))
	case *parser.TryStatementContext:
		p.TryStatement(c.(*parser.TryStatementContext))
	case *parser.WhileStatementContext:
		p.WhileStatement(c.(*parser.WhileStatementContext))
	case *parser.ForStatementContext:
		p.ForStatement(c.(*parser.ForStatementContext))
	case *parser.DoStatementContext:
		p.DoStatement(c.(*parser.DoStatementContext))
	case *parser.ForOfStatementContext:
		p.ForOfStatement(c.(*parser.ForOfStatementContext))
	}
}

func (p *Program) Block(ctx *parser.BlockContext) {
	p1 := newProgram(p.file)
	if ctx.StatementList() != nil {
		p1.StatementList(ctx.StatementList().(*parser.StatementListContext))
	}
	p.push(block(p1.Code), &ctx.BaseParserRuleContext, p.file)
}

func (p *Program) VariableStatement(ctx *parser.VariableStatementContext) {
	p.VariableDeclarationList(ctx.VariableDeclarationList().(*parser.VariableDeclarationListContext))
}

func (p *Program) VariableDeclarationList(ctx *parser.VariableDeclarationListContext) {
	m := ctx.VarModifier().GetText()
	switch m {
	case "let":
		p.modifier = Let
	case "var":
		p.modifier = Var
	case "const":
		p.modifier = Const
	}
	for _, decl := range ctx.AllVariableDeclaration() {
		p.VariableDeclaration(decl.(*parser.VariableDeclarationContext))
	}
	p.push(Eos, &ctx.BaseParserRuleContext, p.file)
}

func (p *Program) VariableDeclaration(ctx *parser.VariableDeclarationContext) {
	if ctx.Assign() == nil {
		p.push(createVar{mod: p.modifier, name: ctx.Assignable().GetText()}, &ctx.BaseParserRuleContext, p.file)
	} else {
		if ctx.SingleExpression() != nil {
			p.SingleExpression(ctx.GetChild(2))
		}
		p.push(createVar{mod: p.modifier, name: ctx.Assignable().GetText()}, &ctx.BaseParserRuleContext, p.file)
	}
}

func (p *Program) ImportStatement(ctx *parser.ImportStatementContext) {
	switch ctx.ImportFromBlock().(type) {
	case *parser.ImportFromBlock1Context:
		v := ctx.ImportFromBlock().(*parser.ImportFromBlock1Context)
		text := v.ImportFrom().StringLiteral().GetText()
		from := text[1 : len(text)-1]
		file := getModuleFile(p.file, from)
		ie := importModule{
			From: from,
			File: file,
		}
		if file != "" {
			s := parseFile(file)
			ie.code = s.Code
		}
		if v.ImportDefault() != nil {
			ie.DefaultName = v.ImportDefault().AliasName().IdentifierName(0).Identifier().GetText()
		}
		ie.Namespace = v.ImportNamespace().IdentifierName().Identifier().Identifier().GetText()
		p.push(ie, &v.BaseParserRuleContext, p.file)
	case *parser.ImportFromBlock2Context:
		v := ctx.ImportFromBlock().(*parser.ImportFromBlock2Context)
		text := v.ImportFrom().StringLiteral().GetText()
		from := text[1 : len(text)-1]
		file := getModuleFile(p.file, from)
		ie := importModule{
			From: from,
			File: file,
		}
		if file != "" {
			s := parseFile(file)
			ie.code = s.Code
		}
		if v.ImportDefault() != nil {
			ie.DefaultName = v.ImportDefault().AliasName().IdentifierName(0).Identifier().GetText()
		}
		mp := make(map[string]string)
		for _, an := range v.ImportModuleItems().AllImportAliasName() {
			if an.ImportedBinding() != nil {
				mp[an.ModuleExportName().IdentifierName().GetText()] = an.ImportedBinding().Identifier().GetText()
			} else {
				mp[an.ModuleExportName().IdentifierName().GetText()] = ""
			}
		}
		ie.Items = mp
		p.push(ie, &v.BaseParserRuleContext, p.file)
	case *parser.ImportFromBlock3Context:
		v := ctx.ImportFromBlock().(*parser.ImportFromBlock3Context)
		text := v.ImportFrom().StringLiteral().GetText()
		from := text[1 : len(text)-1]
		file := getModuleFile(p.file, from)
		ie := importModule{
			From:        from,
			File:        file,
			DefaultName: v.IdentifierName().Identifier().Identifier().GetText(),
		}
		if file != "" {
			s := parseFile(file)
			ie.code = s.Code
		}
		p.push(ie, &v.BaseParserRuleContext, p.file)
	}
	p.push(Eos, &ctx.BaseParserRuleContext, p.file)
}

func (p *Program) ExportBlock(ctx *parser.ExportBlockContext) {
	PrintError("SyntaxError: 语法暂不支持(ExportFromBlock)")
	PrintErrorStack(p.file, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), ctx.GetText())
	panic("")
}

func (p *Program) ExportItems(ctx *parser.ExportItemsContext) {
	for _, v := range ctx.ExportModuleItems().AllExportAliasName() {
		m := v.AllModuleExportName()
		if len(m) == 1 {
			e := exportItem{name: m[0].IdentifierName().GetText()}
			p.push(e, &ctx.BaseParserRuleContext, p.file)
		} else {
			e := exportItem{name: m[0].IdentifierName().GetText(), as: m[1].IdentifierName().GetText()}
			p.push(e, &ctx.BaseParserRuleContext, p.file)
		}
	}
}

func (p *Program) ExportDeclaration(ctx *parser.ExportDeclarationContext) {
	p1 := newProgram(p.file)
	p1.Declaration(ctx.Declaration().(*parser.DeclarationContext))
	p1.push(cmd{c: ExportDeclaration}, &ctx.BaseParserRuleContext, p.file)
	p.push(block(p1.Code), &ctx.BaseParserRuleContext, p.file)
}

func (p *Program) ExportDefaultDeclaration(ctx *parser.ExportDefaultDeclarationContext) {
	p1 := newProgram(p.file)
	p1.SingleExpression(ctx.SingleExpression())
	p1.push(cmd{c: ExportDefault}, &ctx.BaseParserRuleContext, p.file)
	p.push(block(p1.Code), &ctx.BaseParserRuleContext, p.file)
	p.push(Eos, &ctx.BaseParserRuleContext, p.file)
}

func (p *Program) FunctionDeclaration(ctx *parser.FunctionDeclarationContext) {
	p.pushHead(Eos, &ctx.BaseParserRuleContext, p.file)
	p.pushHead(createVar{mod: p.modifier, name: ctx.Identifier().GetText()}, &ctx.BaseParserRuleContext, p.file)
	p.createFunc(ctx.FormalParameterList(), ctx.FunctionBody(), &ctx.BaseParserRuleContext)
	f, _ := p.pop()
	p.Code = append([]Instruct{f}, p.Code...)
}

func (p *Program) ExpressionStatement(ctx *parser.ExpressionStatementContext) {
	p.ExpressionSequence(ctx.ExpressionSequence().(*parser.ExpressionSequenceContext))
	p.push(Eos, &ctx.BaseParserRuleContext, p.file)
}

func (p *Program) IfStatement(ctx *parser.IfStatementContext) {
	toIndex := make([]int, 0)
	p.ifBlock(ctx.IfBlock().ExpressionSequence().(*parser.ExpressionSequenceContext), ctx.IfBlock().Statement().(*parser.StatementContext))
	for _, elseIf := range ctx.AllElseIfBlock() {
		p.push(to(0), nil, p.file)
		toIndex = append(toIndex, len(p.Code)-1)
		p.ifBlock(elseIf.ExpressionSequence().(*parser.ExpressionSequenceContext), elseIf.Statement().(*parser.StatementContext))
	}
	if ctx.ElseBlock() != nil {
		p.push(to(0), nil, p.file)
		toIndex = append(toIndex, len(p.Code)-1)
		p.ifBlock(nil, ctx.ElseBlock().Statement().(*parser.StatementContext))
	}
	k := len(p.Code)
	for _, v := range toIndex {
		p.Code[v] = to(k - v)
	}
	p.push(Eos, &ctx.BaseParserRuleContext, p.file)
}

func (p *Program) ifBlock(e *parser.ExpressionSequenceContext, b *parser.StatementContext) {
	if e != nil {
		p.ExpressionSequence(e)
		p.push(If, &b.BaseParserRuleContext, p.file)
		p.push(to(3), nil, p.file)
	}
	p.Statement(b)
	return
}

func (p *Program) ContinueStatement(ctx *parser.ContinueStatementContext) {
	p.push(Continue, &ctx.BaseParserRuleContext, p.file)
}

func (p *Program) BreakStatement(ctx *parser.BreakStatementContext) {
	p.push(cmd{c: Break}, &ctx.BaseParserRuleContext, p.file)
}

func (p *Program) ReturnStatement(ctx *parser.ReturnStatementContext) {
	if ctx.ExpressionSequence() != nil {
		p.ExpressionSequence(ctx.ExpressionSequence().(*parser.ExpressionSequenceContext))
	}
	p.push(Return, &ctx.BaseParserRuleContext, p.file)
}

func (p *Program) SwitchStatement(ctx *parser.SwitchStatementContext) {
	toIndex := make([]int, 0)
	p1 := newProgram(p.file)
	for _, v := range ctx.CaseBlock().AllCaseClauses() {
		for _, v1 := range v.AllCaseClause() {
			p1.ExpressionSequence(ctx.ExpressionSequence().(*parser.ExpressionSequenceContext))
			p1.ExpressionSequence(v1.ExpressionSequence().(*parser.ExpressionSequenceContext))
			p1.push(Equals, &ctx.BaseParserRuleContext, p.file)
			p1.push(If, &ctx.BaseParserRuleContext, p.file)
			p1.push(to(0), nil, p.file)
			i := len(p1.Code) - 1
			p2 := newProgram(p.file)
			if v1.StatementList() != nil {
				p2.StatementList(v1.StatementList().(*parser.StatementListContext))
			}
			p1.push(block(p2.Code), &ctx.BaseParserRuleContext, p.file)
			p1.push(to(0), nil, p.file)
			toIndex = append(toIndex, len(p1.Code)-1)
			p1.Code[i] = to(len(p1.Code) - i)
		}
	}
	if ctx.CaseBlock().DefaultClause() != nil && ctx.CaseBlock().DefaultClause().StatementList() != nil {
		p2 := newProgram(p.file)
		p2.StatementList(ctx.CaseBlock().DefaultClause().StatementList().(*parser.StatementListContext))
		p1.push(block(p2.Code), &ctx.BaseParserRuleContext, p.file)
	}
	k := len(p1.Code)
	for _, v := range toIndex {
		p1.Code[v] = to(k - v)
	}
	p1.push(Eos, &ctx.BaseParserRuleContext, p.file)
	s := block(p1.Code)
	p.push(s, &ctx.BaseParserRuleContext, p.file)
	p.push(Eos, &ctx.BaseParserRuleContext, p.file)
}

func (p *Program) ThrowStatement(ctx *parser.ThrowStatementContext) {
	p.ExpressionSequence(ctx.ExpressionSequence().(*parser.ExpressionSequenceContext))
	p.push(cmd{c: Throw}, &ctx.BaseParserRuleContext, p.file)
}

func (p *Program) TryStatement(ctx *parser.TryStatementContext) {
	t := try{}
	p1 := newProgram(p.file)
	p1.Block(ctx.Block().(*parser.BlockContext))
	t.try = p1.Code[0].(block)
	ch := ctx.CatchProduction()
	if ch != nil {
		ca := catch{}
		if ch.Assignable() != nil {
			ca.formal = ch.Assignable().GetText()
		}
		p2 := newProgram(p.file)
		p2.Block(ch.Block().(*parser.BlockContext))
		ca.body = p2.Code[0].(block)
		t.catch = ca
	}
	fi := ctx.FinallyProduction()
	if fi != nil {
		p2 := newProgram(p.file)
		p2.Block(fi.Block().(*parser.BlockContext))
		t.finally = p2.Code[0].(block)
	}
	p.push(t, &ctx.BaseParserRuleContext, p.file)
	p.push(Eos, &ctx.BaseParserRuleContext, p.file)
}

func (p *Program) WhileStatement(ctx *parser.WhileStatementContext) {
	p1 := newProgram(p.file)
	p1.ExpressionSequence(ctx.ExpressionSequence().(*parser.ExpressionSequenceContext))
	p1.push(While, &ctx.BaseParserRuleContext, p.file)
	p1.Statement(ctx.Statement().(*parser.StatementContext))
	t := to(-len(p1.Code))
	p1.push(t, &ctx.BaseParserRuleContext, p.file)
	p1.push(Eos, &ctx.BaseParserRuleContext, p.file)
	i := iteration(p1.Code)
	p.push(i, &ctx.BaseParserRuleContext, p.file)
	p.push(Eos, &ctx.BaseParserRuleContext, p.file)
}

func (p *Program) DoStatement(ctx *parser.DoStatementContext) {
	p1 := newProgram(p.file)
	p1.Statement(ctx.Statement().(*parser.StatementContext))
	p1.ExpressionSequence(ctx.ExpressionSequence().(*parser.ExpressionSequenceContext))
	p1.push(While, &ctx.BaseParserRuleContext, p.file)
	t := to(-len(p1.Code))
	p1.push(t, &ctx.BaseParserRuleContext, p.file)
	p1.push(Eos, &ctx.BaseParserRuleContext, p.file)
	i := iteration(p1.Code)
	p.push(i, &ctx.BaseParserRuleContext, p.file)
	p.push(Eos, &ctx.BaseParserRuleContext, p.file)
}

func (p *Program) ForOfStatement(ctx *parser.ForOfStatementContext) {
	p1 := newProgram(p.file)
	p1.ExpressionSequence(ctx.ExpressionSequence().(*parser.ExpressionSequenceContext))
	m := ctx.VarModifier().GetText()
	ids := make([]string, 0)
	if ctx.Identifier() != nil {
		ids = append(ids, ctx.Identifier().GetText())
	} else {
		for _, i := range ctx.OfArrayLiteral().IdentifierList().AllIdentifier() {
			ids = append(ids, i.GetText())
		}
	}
	for _, v := range ids {
		p1.push(loadVal{value: NewUndefined()}, nil, p.file)
		if m == "let" {
			p1.push(createVar{mod: Let, name: v}, nil, p.file)
		} else {
			p1.push(createVar{mod: Const, name: v}, nil, p.file)
		}
	}
	p1.push(forOf{vars: ids, list: ctx.OfArrayLiteral() != nil}, nil, p.file)
	t := len(p1.Code) - 1
	p1.Statement(ctx.Statement().(*parser.StatementContext))
	p1.push(to(t-len(p1.Code)), &ctx.BaseParserRuleContext, p.file)
	i := iteration(p1.Code)
	p.push(i, &ctx.BaseParserRuleContext, p.file)
	p.push(Eos, &ctx.BaseParserRuleContext, p.file)
}

func (p *Program) ForStatement(ctx *parser.ForStatementContext) {
	p1 := newProgram(p.file)
	if ctx.ForStatement1() != nil {
		if ctx.ForStatement1().ExpressionSequence() != nil {
			p1.ExpressionSequence(ctx.ForStatement1().ExpressionSequence().(*parser.ExpressionSequenceContext))
		} else if ctx.ForStatement1().VariableDeclarationList() != nil {
			p1.VariableDeclarationList(ctx.ForStatement1().VariableDeclarationList().(*parser.VariableDeclarationListContext))
		}
	}
	t := len(p1.Code)
	if ctx.ForStatement2() != nil {
		p1.ExpressionSequence(ctx.ForStatement2().ExpressionSequence().(*parser.ExpressionSequenceContext))
	} else {
		p1.push(loadVal{value: GoToBoolean("true")}, &ctx.BaseParserRuleContext, p.file)
	}
	p1.push(While, &ctx.BaseParserRuleContext, p.file)
	if ctx.Statement() != nil {
		p1.Statement(ctx.Statement().(*parser.StatementContext))
	}
	if ctx.ForStatement3() != nil {
		p1.ExpressionSequence(ctx.ForStatement3().ExpressionSequence().(*parser.ExpressionSequenceContext))
	}
	p1.push(Eos, &ctx.BaseParserRuleContext, p.file)
	to1 := to(t - len(p1.Code))
	p1.push(to1, &ctx.BaseParserRuleContext, p.file)
	p.push(iteration(p1.Code), &ctx.BaseParserRuleContext, p.file)
	p.push(Eos, &ctx.BaseParserRuleContext, p.file)
}

func (p *Program) StatementList(list *parser.StatementListContext) {
	for _, statement := range list.AllStatement() {
		p.Statement(statement.(*parser.StatementContext))
	}
}

// statement end

func (p *Program) SingleExpression(ctx antlr.Tree) {
	switch ctx.(type) {
	case *parser.FunctionExpressionContext:
		p.FunctionExpression(ctx.(*parser.FunctionExpressionContext))
	case *parser.MemberIndexExpressionContext:
		p.MemberIndexExpression(ctx.(*parser.MemberIndexExpressionContext))
	case *parser.MemberDotExpressionContext:
		p.MemberDotExpression(ctx.(*parser.MemberDotExpressionContext))
	case *parser.NewExpressionContext:
		p.NewExpression1(ctx.(*parser.NewExpressionContext))
	case *parser.ArgumentsExpressionContext:
		p.ArgumentsExpression(ctx.(*parser.ArgumentsExpressionContext))
	case *parser.PrefixArgumentsExpressionContext:
		p.PrefixArgumentsExpression(ctx.(*parser.PrefixArgumentsExpressionContext))
	case *parser.PostIncrementExpressionContext:
		p.PostIncrementExpression(ctx.(*parser.PostIncrementExpressionContext))
	case *parser.PostDecreaseExpressionContext:
		p.PostDecreaseExpression(ctx.(*parser.PostDecreaseExpressionContext))
	case *parser.DeleteExpressionContext:
		p.DeleteExpression(ctx.(*parser.DeleteExpressionContext))
	case *parser.PreIncrementExpressionContext:
		p.PreIncrementExpression(ctx.(*parser.PreIncrementExpressionContext))
	case *parser.PreDecreaseExpressionContext:
		p.PreDecreaseExpression(ctx.(*parser.PreDecreaseExpressionContext))
	case *parser.UnaryPlusExpressionContext:
		p.UnaryPlusExpression(ctx.(*parser.UnaryPlusExpressionContext))
	case *parser.UnaryMinusExpressionContext:
		p.UnaryMinusExpression(ctx.(*parser.UnaryMinusExpressionContext))
	case *parser.BitNotExpressionContext:
		p.BitNotExpression(ctx.(*parser.BitNotExpressionContext))
	case *parser.NotExpressionContext:
		p.NotExpression(ctx.(*parser.NotExpressionContext))
	case *parser.PowerExpressionContext:
		p.PowerExpression(ctx.(*parser.PowerExpressionContext))
	case *parser.MultiplicativeExpressionContext:
		p.MultiplicativeExpression(ctx.(*parser.MultiplicativeExpressionContext))
	case *parser.AdditiveExpressionContext:
		p.AdditiveExpression(ctx.(*parser.AdditiveExpressionContext))
	case *parser.CoalesceExpressionContext:
		p.CoalesceExpression(ctx.(*parser.CoalesceExpressionContext))
	case *parser.BitShiftExpressionContext:
		p.BitShiftExpression(ctx.(*parser.BitShiftExpressionContext))
	case *parser.RelationalExpressionContext:
		p.RelationalExpression(ctx.(*parser.RelationalExpressionContext))
	case *parser.InExpressionContext:
		p.InExpression(ctx.(*parser.InExpressionContext))
	case *parser.EqualityExpressionContext:
		p.EqualityExpression(ctx.(*parser.EqualityExpressionContext))
	case *parser.BitExpressionContext:
		p.BitExpression(ctx.(*parser.BitExpressionContext))
	case *parser.LogicalExpressionContext:
		p.LogicalExpression(ctx.(*parser.LogicalExpressionContext))
	case *parser.TernaryExpressionContext:
		p.TernaryExpression(ctx.(*parser.TernaryExpressionContext))
	case *parser.AssignmentExpressionContext:
		p.AssignmentExpression(ctx.(*parser.AssignmentExpressionContext))
	case *parser.AssignmentOperatorExpressionContext:
		p.AssignmentOperatorExpression(ctx.(*parser.AssignmentOperatorExpressionContext))
	case *parser.TemplateStringExpressionContext:
		p.TemplateStringExpression(ctx.(*parser.TemplateStringExpressionContext))
	case *parser.ThisExpressionContext:
		p.ThisExpression(ctx.(*parser.ThisExpressionContext))
	case *parser.IdentifierExpressionContext:
		p.IdentifierExpression(ctx.(*parser.IdentifierExpressionContext))
	case *parser.LiteralExpressionContext:
		p.LiteralExpression(ctx.(*parser.LiteralExpressionContext))
	case *parser.ArrayLiteralExpressionContext:
		p.ArrayLiteralExpression(ctx.(*parser.ArrayLiteralExpressionContext))
	case *parser.ObjectLiteralExpressionContext:
		p.ObjectLiteralExpression(ctx.(*parser.ObjectLiteralExpressionContext))
	case *parser.ParenthesizedExpressionContext:
		p.ParenthesizedExpression(ctx.(*parser.ParenthesizedExpressionContext))
	}
}

func (p *Program) FunctionExpression(ctx *parser.FunctionExpressionContext) {
	switch ctx.AnonymousFunction().(type) {
	case *parser.AnonymousFunctionDeclContext:
		x := ctx.AnonymousFunction().(*parser.AnonymousFunctionDeclContext)
		p.createFunc(x.FormalParameterList(), x.FunctionBody(), &ctx.BaseParserRuleContext)
	case *parser.ArrowFunctionContext:
		x := ctx.AnonymousFunction().(*parser.ArrowFunctionContext)
		p.createArrowFunc(x.ArrowFunctionParameters(), x.FunctionBody(), &ctx.BaseParserRuleContext)
	case *parser.ArrowSingleFunctionContext:
		x := ctx.AnonymousFunction().(*parser.ArrowSingleFunctionContext)
		p.createArrowSingleFunc(x.ArrowFunctionParameters(), x.SingleExpression(), &ctx.BaseParserRuleContext)
	}
}

func (p *Program) MemberIndexExpression(ctx *parser.MemberIndexExpressionContext) {
	p.SingleExpression(ctx.GetChild(0))
	p.ExpressionSequence(ctx.ExpressionSequence().(*parser.ExpressionSequenceContext))
	p.push(property{optional: ctx.QuestionMarkDot() != nil}, &ctx.BaseParserRuleContext, p.file)
	if ctx.QuestionMarkDot() != nil {
		p.Optional(ctx.GetParent())
	}
}

func (p *Program) MemberDotExpression(ctx *parser.MemberDotExpressionContext) {
	p.SingleExpression(ctx.GetChild(0))
	i := ctx.IdentifierName().GetText()
	p.push(property{name: i, optional: ctx.QuestionMarkDot() != nil}, &ctx.BaseParserRuleContext, p.file)
	if ctx.QuestionMarkDot() != nil {
		p.Optional(ctx.GetParent())
	}
}

func (p *Program) NewExpression1(ctx *parser.NewExpressionContext) {
	i := ctx.Identifier().Identifier().GetText()
	p.push(loadVar{name: i}, &ctx.BaseParserRuleContext, p.file)
	p.Arguments(ctx.Arguments().(*parser.ArgumentsContext))
	p.push(call{index: len(ctx.Arguments().AllArgument()), constructor: true}, &ctx.BaseParserRuleContext, p.file)

}

func (p *Program) IsQuestionMember(ctx antlr.Tree) bool {
	if ctx.GetChildCount() == 0 {
		return false
	}
	c := ctx.GetChild(0)
	if v, ok := c.(*parser.MemberDotExpressionContext); ok && v.QuestionMarkDot() != nil {
		return true
	} else if v, ok := c.(*parser.MemberIndexExpressionContext); ok && v.QuestionMarkDot() != nil {
		return true
	}
	return p.IsQuestionMember(c)
}

func (p *Program) ArgumentsExpression(ctx *parser.ArgumentsExpressionContext) {

	p.SingleExpression(ctx.SingleExpression())

	//用于实现可选链
	isQuestionMember := p.IsQuestionMember(ctx)
	isArgument := len(ctx.Arguments().AllArgument()) > 0
	skipIndex := 0
	if isArgument && isQuestionMember {
		skipIndex = len(p.Code)
		p.push(skip(0), nil, p.file)
	}

	p.Arguments(ctx.Arguments().(*parser.ArgumentsContext))

	callOptional := ctx.QuestionMarkDot() != nil

	p.push(call{index: len(ctx.Arguments().AllArgument()), optional: callOptional}, &ctx.BaseParserRuleContext, p.file)

	//用于实现可选链
	if isQuestionMember {
		if isArgument {
			if callOptional {
				p.Code[skipIndex] = skip(len(p.Code) - skipIndex)
			} else {
				p.Code[skipIndex] = skip(len(p.Code) - skipIndex - 1)
			}
		} else if !isArgument {
			p.push(skip(0), nil, p.file)
		}
	}
	if callOptional {
		p.push(skip(0), nil, p.file)
	}
}

func (p *Program) PrefixArgumentsExpression(ctx *parser.PrefixArgumentsExpressionContext) {
	p.SingleExpression(ctx.SingleExpression())

	//用于实现可选链
	isQuestionMember := p.IsQuestionMember(ctx)
	isArgument := len(ctx.Arguments().AllArgument()) > 0
	skipIndex := 0
	if isArgument && isQuestionMember {
		skipIndex = len(p.Code)
		p.push(skip(0), nil, p.file)
	}

	p.Arguments(ctx.Arguments().(*parser.ArgumentsContext))
	callOptional := ctx.QuestionMarkDot() != nil
	state := 1
	if ctx.Sync() != nil {
		state = 2
	}
	p.push(call{index: len(ctx.Arguments().AllArgument()), state: state, optional: ctx.QuestionMarkDot() != nil}, &ctx.BaseParserRuleContext, p.file)

	//用于实现可选链
	if isQuestionMember {
		if isArgument {
			if callOptional {
				p.Code[skipIndex] = skip(len(p.Code) - skipIndex)
			} else {
				p.Code[skipIndex] = skip(len(p.Code) - skipIndex - 1)
			}
		} else if !isArgument {
			p.push(skip(0), nil, p.file)
		}
	}
	if callOptional {
		p.push(skip(0), nil, p.file)
	}
}

func (p *Program) Optional(ctx antlr.Tree) {
	if _, ok := ctx.(*parser.AssignmentExpressionContext); ok {
		p.push(skip(0), nil, p.file)
	} else if _, ok = ctx.(*parser.ExpressionSequenceContext); ok {
		p.push(skip(0), nil, p.file)
	} else if _, ok = ctx.(*parser.CoalesceExpressionContext); ok {
		p.push(skip(0), nil, p.file)
	} else if _, ok = ctx.(*parser.VariableDeclarationContext); ok {
		p.push(skip(0), nil, p.file)
	} else if _, ok = ctx.(*parser.ArgumentContext); ok {
		p.push(skip(0), nil, p.file)
	}
}

func (p *Program) Arguments(ctx *parser.ArgumentsContext) {
	for _, v := range ctx.AllArgument() {
		p.SingleExpression(v.SingleExpression())
		if v.Ellipsis() != nil {
			p.push(Ellipsis, &ctx.BaseParserRuleContext, p.file)
		}
		p.Optional(v.GetParent())
	}
}

func (p *Program) PostIncrementExpression(ctx *parser.PostIncrementExpressionContext) {
	// a++
	p.ident = 0
	p.SingleExpression(ctx.SingleExpression())
	p.SingleExpression(ctx.SingleExpression())
	p.push(loadVal{value: GoToNumber(1)}, &ctx.BaseParserRuleContext, p.file)
	p.push(Plus, &ctx.BaseParserRuleContext, p.file)
	p.ident = 1
	p.SingleExpression(ctx.SingleExpression())
	p.ident = 0
}

func (p *Program) PostDecreaseExpression(ctx *parser.PostDecreaseExpressionContext) {
	p.ident = 0
	p.SingleExpression(ctx.SingleExpression())
	p.SingleExpression(ctx.SingleExpression())
	p.push(loadVal{value: GoToNumber(1)}, &ctx.BaseParserRuleContext, p.file)
	p.push(Minus, &ctx.BaseParserRuleContext, p.file)
	p.ident = 1
	p.SingleExpression(ctx.SingleExpression())
	p.ident = 0
}

func (p *Program) DeleteExpression(ctx *parser.DeleteExpressionContext) {
	p.SingleExpression(ctx.SingleExpression())
	if v, ok := p.pop(); ok {
		if pro, ok := v.(property); ok {
			pro.delete = true
			p.push(pro, &ctx.BaseParserRuleContext, p.file)
		} else {
			p.push(v, &ctx.BaseParserRuleContext, p.file)
		}
	}
}

func (p *Program) PreIncrementExpression(ctx *parser.PreIncrementExpressionContext) {
	p.ident = 0
	p.SingleExpression(ctx.SingleExpression())
	p.push(loadVal{value: GoToNumber(1)}, &ctx.BaseParserRuleContext, p.file)
	p.push(Plus, &ctx.BaseParserRuleContext, p.file)
	p.ident = 1
	p.SingleExpression(ctx.SingleExpression())
	p.ident = 0
	p.SingleExpression(ctx.SingleExpression())
}
func (p *Program) PreDecreaseExpression(ctx *parser.PreDecreaseExpressionContext) {
	p.ident = 0
	p.SingleExpression(ctx.SingleExpression())
	p.push(loadVal{value: GoToNumber(1)}, &ctx.BaseParserRuleContext, p.file)
	p.push(Minus, &ctx.BaseParserRuleContext, p.file)
	p.ident = 1
	p.SingleExpression(ctx.SingleExpression())
	p.ident = 0
	p.SingleExpression(ctx.SingleExpression())
}

func (p *Program) UnaryPlusExpression(ctx *parser.UnaryPlusExpressionContext) {
	p.ident = 0
	p.SingleExpression(ctx.SingleExpression())
	p.push(UnaryPlus, &ctx.BaseParserRuleContext, p.file)
	p.ident = 0
}

func (p *Program) UnaryMinusExpression(ctx *parser.UnaryMinusExpressionContext) {
	p.ident = 0
	p.SingleExpression(ctx.SingleExpression())
	p.push(UnaryMinus, &ctx.BaseParserRuleContext, p.file)
	p.ident = 0
}
func (p *Program) BitNotExpression(ctx *parser.BitNotExpressionContext) {
	p.ident = 0
	p.SingleExpression(ctx.SingleExpression())
	p.push(BitNot, &ctx.BaseParserRuleContext, p.file)
}

func (p *Program) NotExpression(ctx *parser.NotExpressionContext) {
	p.ident = 0
	p.SingleExpression(ctx.SingleExpression())
	p.push(Not, &ctx.BaseParserRuleContext, p.file)
}

func (p *Program) PowerExpression(ctx *parser.PowerExpressionContext) {
	p.SingleExpression(ctx.SingleExpression(0))
	p.SingleExpression(ctx.SingleExpression(1))
	p.push(Power, nil, p.file)
}

func (p *Program) MultiplicativeExpression(ctx *parser.MultiplicativeExpressionContext) {
	p.SingleExpression(ctx.SingleExpression(0))
	p.SingleExpression(ctx.SingleExpression(1))
	if ctx.Multiply() != nil {
		p.push(Multiply, nil, p.file)
	} else if ctx.Modulus() != nil {
		p.push(Modulus, nil, p.file)
	} else if ctx.Divide() != nil {
		p.push(Divide, nil, p.file)
	}
}

func (p *Program) AdditiveExpression(ctx *parser.AdditiveExpressionContext) {
	p.SingleExpression(ctx.SingleExpression(0))
	p.SingleExpression(ctx.SingleExpression(1))
	if ctx.Plus() != nil {
		p.push(Plus, nil, p.file)
	} else if ctx.Minus() != nil {
		p.push(Minus, nil, p.file)
	}
}

func (p *Program) CoalesceExpression(ctx *parser.CoalesceExpressionContext) {
	p.SingleExpression(ctx.SingleExpression(0))
	p.push(NullCoalesce, &ctx.BaseParserRuleContext, p.file)
	i := len(p.Code)
	p.push(to(0), &ctx.BaseParserRuleContext, p.file)
	p.SingleExpression(ctx.SingleExpression(1))
	p.Code[i] = to(len(p.Code) - i)
}

func (p *Program) BitShiftExpression(ctx *parser.BitShiftExpressionContext) {
	p.SingleExpression(ctx.SingleExpression(0))
	p.SingleExpression(ctx.SingleExpression(1))
	if ctx.RightShiftArithmetic() != nil {
		p.push(RightShiftArithmetic, nil, p.file)
	} else if ctx.LeftShiftArithmetic() != nil {
		p.push(LeftShiftArithmetic, nil, p.file)
	} else if ctx.RightShiftLogical() != nil {
		p.push(RightShiftLogical, nil, p.file)
	}
}

func (p *Program) RelationalExpression(ctx *parser.RelationalExpressionContext) {
	p.ident = 0
	p.SingleExpression(ctx.SingleExpression(0))
	p.SingleExpression(ctx.SingleExpression(1))
	if ctx.LessThan() != nil {
		p.push(LessThan, nil, p.file)
	} else if ctx.MoreThan() != nil {
		p.push(MoreThan, nil, p.file)
	} else if ctx.LessThanEquals() != nil {
		p.push(LessThanEquals, nil, p.file)
	} else if ctx.GreaterThanEquals() != nil {
		p.push(GreaterThanEquals, nil, p.file)
	}
}

func (p *Program) InExpression(ctx *parser.InExpressionContext) {
	p.SingleExpression(ctx.SingleExpression(0))
	p.SingleExpression(ctx.SingleExpression(1))
	p.push(In, &ctx.BaseParserRuleContext, p.file)
}
func (p *Program) EqualityExpression(ctx *parser.EqualityExpressionContext) {
	p.ident = 0
	p.SingleExpression(ctx.SingleExpression(0))
	p.SingleExpression(ctx.SingleExpression(1))
	if ctx.Equals_() != nil {
		p.push(Equals, nil, p.file)
	} else if ctx.NotEquals() != nil {
		p.push(NotEquals, nil, p.file)
	} else if ctx.IdentityEquals() != nil {
		p.push(IdentityEquals, nil, p.file)
	} else if ctx.IdentityNotEquals() != nil {
		p.push(IdentityNotEquals, nil, p.file)
	}
}

func (p *Program) BitExpression(ctx *parser.BitExpressionContext) {
	p.SingleExpression(ctx.SingleExpression(0))
	p.SingleExpression(ctx.SingleExpression(1))
	if ctx.BitAnd() != nil {
		p.push(BitAnd, nil, p.file)
	} else if ctx.BitOr() != nil {
		p.push(BitOr, nil, p.file)
	} else if ctx.BitXOr() != nil {
		p.push(BitXOr, nil, p.file)
	}
}

func (p *Program) LogicalExpression(ctx *parser.LogicalExpressionContext) {
	p.SingleExpression(ctx.SingleExpression(0))
	if ctx.Or() != nil {
		p.push(Or, nil, p.file)
	} else if ctx.And() != nil {
		p.push(And, nil, p.file)
	}
	index := len(p.Code)
	p.push(to(0), nil, p.file)
	p.SingleExpression(ctx.SingleExpression(1))
	p.Code[index] = to(len(p.Code) - index)
}

func (p *Program) TernaryExpression(ctx *parser.TernaryExpressionContext) {
	p.SingleExpression(ctx.SingleExpression(0))
	p.push(If, nil, p.file)
	//第二条的开始
	p.push(to(0), nil, p.file)
	i := len(p.Code) - 1
	p.SingleExpression(ctx.SingleExpression(1))
	p.Code[i] = to(len(p.Code) - i + 1)
	//结尾
	p.push(to(0), nil, p.file)
	i = len(p.Code) - 1
	p.SingleExpression(ctx.SingleExpression(2))
	p.Code[i] = to(len(p.Code) - i)
}

func (p *Program) AssignmentExpression(ctx *parser.AssignmentExpressionContext) {
	p.ident = 0 //0加载,1修改
	p.SingleExpression(ctx.SingleExpression(1))
	if _, ok := ctx.GetChild(0).(*parser.IdentifierExpressionContext); ok {
		p.ident = 1
	}
	p.SingleExpression(ctx.SingleExpression(0))
	p.leftHand(&ctx.BaseParserRuleContext)
}

func (p *Program) AssignmentOperatorExpression(ctx *parser.AssignmentOperatorExpressionContext) {
	s := ctx.AssignmentOperator().GetText()
	if s == "??=" {
		p.SingleExpression(ctx.SingleExpression(0))
		p.push(NullCoalesce, &ctx.BaseParserRuleContext, p.file)
		i := len(p.Code)
		p.push(to(0), &ctx.BaseParserRuleContext, p.file)
		p.SingleExpression(ctx.SingleExpression(1))
		p.Code[i] = to(len(p.Code) - i)
	} else {
		p.SingleExpression(ctx.SingleExpression(0))
		p.SingleExpression(ctx.SingleExpression(1))
		switch s {
		case "*=":
			p.push(Multiply, nil, p.file)
		case "/=":
			p.push(Divide, nil, p.file)
		case "%=":
			p.push(Modulus, nil, p.file)
		case "+=":
			p.push(Plus, nil, p.file)
		case "-=":
			p.push(Minus, nil, p.file)
		case "<<=":
			p.push(LeftShiftArithmetic, nil, p.file)
		case ">>=":
			p.push(RightShiftArithmetic, nil, p.file)
		case ">>>=":
			p.push(RightShiftLogical, nil, p.file)
		case "&=":
			p.push(BitAnd, nil, p.file)
		case "^=":
			p.push(BitXOr, nil, p.file)
		case "|=":
			p.push(BitOr, nil, p.file)
		case "**=":
			p.push(Power, nil, p.file)
		}
	}

	if v, ok := ctx.SingleExpression(0).(*parser.IdentifierExpressionContext); ok {
		p.ident = 1
		p.push(modifyVar{name: v.GetText()}, &ctx.BaseParserRuleContext, p.file)
	} else {
		p.ident = 0
		p.SingleExpression(ctx.SingleExpression(0))
	}
	p.ident = 0
}

func (p *Program) TemplateStringExpression(ctx *parser.TemplateStringExpressionContext) {
	PrintError("SyntaxError: 语法暂不支持(TemplateStringExpression)")
	PrintErrorStack(p.file, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), ctx.GetText())
	panic("")
}

func (p *Program) ThisExpression(ctx *parser.ThisExpressionContext) {
	p.push(cmd{c: This}, &ctx.BaseParserRuleContext, p.file)
}

func (p *Program) IdentifierExpression(ctx *parser.IdentifierExpressionContext) {
	i := ctx.GetText()
	if p.ident == 0 {
		p.push(loadVar{name: i}, &ctx.BaseParserRuleContext, p.file)
	} else {
		p.push(modifyVar{name: i}, &ctx.BaseParserRuleContext, p.file)
	}
	p.ident = 0
}

func (p *Program) LiteralExpression(ctx *parser.LiteralExpressionContext) {
	var v JavaScript
	literal := ctx.Literal()
	switch vt := literal.GetChild(0).(type) {
	case *parser.TemplateStringLiteralContext:
		for i, v2 := range vt.AllTemplateStringAtom() {
			if v2.SingleExpression() != nil {
				p.SingleExpression(v2.SingleExpression())
			} else {
				p.push(loadVal{value: GoToString(v2.GetText())}, &ctx.BaseParserRuleContext, p.file)
			}
			if i > 0 {
				p.push(Plus, &ctx.BaseParserRuleContext, p.file)
			}
		}
		return
	case *parser.NumericLiteralContext:
		if vt.ExponentLiteral() != nil {
			v = GoToNumber(vt.GetText())
		} else if vt.HexIntegerLiteral() != nil {
			v = hexStrToNumber(vt.GetText())
		} else if vt.OctalIntegerLiteral() != nil {
			v = octalStrToNumber(vt.GetText())
		} else if vt.BinaryIntegerLiteral() != nil {
			v = binaryStrToNumber(vt.GetText())
		} else {
			v = GoToNumber(vt.GetText())
		}
	case *parser.BigintLiteralContext:
		text := strings.TrimSuffix(vt.GetText(), "n")
		if vt.BigDecimalIntegerLiteral() != nil {
			v = GoToBigInt(text)
		} else if vt.BigHexIntegerLiteral() != nil {
			v = hexStrToBigInt(text)
		} else if vt.BigOctalIntegerLiteral() != nil {
			v = octalStrToBigInt(text)
		} else if vt.BigBinaryIntegerLiteral() != nil {
			v = binaryStrToBigInt(text)
		}
	default:
		if literal.StringLiteral() != nil {
			str := literal.GetText()
			if str[0] == '\'' {
				str = str[1 : len(str)-1]
				str = strings.ReplaceAll(str, "\\'", "'")
				str = strings.ReplaceAll(str, "\"", "\\\"")
				str = "\"" + str + "\""
			}
			//处理转译字符串
			unquotedString, err := strconv.Unquote(str)
			if err != nil {
				panic(NewSyntaxError(fmt.Sprint("字符串字面量：", str, "解析错误")))
			}
			v = GoToString(unquotedString)
		} else if literal.NullLiteral() != nil {
			v = NewNull()
		} else if literal.UndefinedLiteral() != nil {
			v = NewUndefined()
		} else if literal.BooleanLiteral() != nil {
			v = GoToBoolean(literal.GetText())
		} else if literal.RegularExpressionLiteral() != nil {
			v = GoToRegExp(ctx.GetText())
		} else if literal.InfinityLiteral() != nil {
			v = GoToNumber(math.Inf(1))
		} else if literal.NaNLiteral() != nil {
			v = GoToNumber(math.NaN())
		} else {
			fmt.Println("其他：", ctx.GetText())
		}
	}
	if v != nil {
		c := loadVal{value: v}
		p.push(c, nil, p.file)
	}
}
func (p *Program) ArrayLiteralExpression(ctx *parser.ArrayLiteralExpressionContext) {
	for _, v := range ctx.ArrayLiteral().ElementList().AllArrayElement() {
		p.SingleExpression(v.SingleExpression())
		if v.Ellipsis() != nil {
			p.push(Ellipsis, &ctx.BaseParserRuleContext, p.file)
		}
	}
	l := array(len(ctx.ArrayLiteral().ElementList().AllArrayElement()))
	p.push(l, &ctx.BaseParserRuleContext, p.file)
}
func (p *Program) ObjectLiteralExpression(ctx *parser.ObjectLiteralExpressionContext) {
	p.ObjectLiteral(ctx.ObjectLiteral().(*parser.ObjectLiteralContext))
}

func (p *Program) ParenthesizedExpression(ctx *parser.ParenthesizedExpressionContext) {
	p.ExpressionSequence(ctx.ExpressionSequence().(*parser.ExpressionSequenceContext))
}

func (p *Program) ObjectLiteral(ctx *parser.ObjectLiteralContext) {
	p1 := newProgram(p.file)
	for _, x := range ctx.AllPropertyAssignment() {
		switch x.(type) {
		case *parser.PropertyExpressionAssignmentContext:
			v := x.(*parser.PropertyExpressionAssignmentContext)
			p1.SingleExpression(v.SingleExpression())
			p1.PropertyName(v.PropertyName().(*parser.PropertyNameContext))
		case *parser.ComputedPropertyExpressionAssignmentContext:
			v := x.(*parser.ComputedPropertyExpressionAssignmentContext)
			p1.SingleExpression(v.SingleExpression(1))
			p1.SingleExpression(v.SingleExpression(0))
			//@Todo 需要修改
			p1.push(createVar{mod: p.modifier}, &ctx.BaseParserRuleContext, p.file)
			p1.modifier = 0
		case *parser.FunctionPropertyContext:
			v := x.(*parser.FunctionPropertyContext)
			p1.createFunc(v.FormalParameterList(), v.FunctionBody(), &v.BaseParserRuleContext)
			p1.PropertyName(v.PropertyName().(*parser.PropertyNameContext))
		case *parser.PropertyShorthandContext:
			v := x.(*parser.PropertyShorthandContext)
			if n, ok := v.SingleExpression().(*parser.IdentifierExpressionContext); ok {
				p1.push(loadVar{name: n.GetText()}, &n.BaseParserRuleContext, p.file)
				if v.Ellipsis() != nil {
					p1.push(Ellipsis, &ctx.BaseParserRuleContext, p.file)
				}
			} else {
				PrintError("SyntaxError: Unexpected identifier '" + v.GetText() + "'")
				PrintErrorStack(p.file, v.GetStart().GetLine(), v.GetStart().GetColumn(), v.GetText())
				panic("")
			}
		}
	}
	b := createObject(p1.Code)
	p.push(b, &ctx.BaseParserRuleContext, p.file)
}

func (p *Program) PropertyName(ctx *parser.PropertyNameContext) {
	c := ctx.GetChild(0)
	switch c.(type) {
	case *parser.IdentifierNameContext:
		p.push(createVar{mod: p.modifier, name: c.(*parser.IdentifierNameContext).GetText()}, &ctx.BaseParserRuleContext, p.file)
	case *parser.TemplateStringLiteralContext:
		p.push(createVar{mod: p.modifier, name: c.(*parser.TemplateStringLiteralContext).GetText()}, &ctx.BaseParserRuleContext, p.file)
	case *parser.NumericLiteralContext:
		p.push(createVar{mod: p.modifier, name: c.(*parser.NumericLiteralContext).GetText()}, &ctx.BaseParserRuleContext, p.file)
	case *antlr.TerminalNodeImpl:
		str := c.(*antlr.TerminalNodeImpl).GetText()
		str = str[1 : len(str)-1]
		//处理转译字符串
		unquotedString, err := strconv.Unquote(`"` + str + `"`)
		if err != nil {
			panic(fmt.Sprintf("Unquote failed: %v", err))
		}
		p.push(createVar{mod: p.modifier, name: unquotedString}, &ctx.BaseParserRuleContext, p.file)
	default:
		if ctx.SingleExpression() != nil {
			p.SingleExpression(ctx.SingleExpression())
			//@Todo需要修改
			p.push(createVar{mod: p.modifier}, &ctx.BaseParserRuleContext, p.file)
		}
	}
}

func (p *Program) createArrowFunc(formalParam parser.IArrowFunctionParametersContext, body parser.IFunctionBodyContext, bpr *antlr.BaseParserRuleContext) {
	formal := make([]string, 0)
	if formalParam.FormalParameterList() != nil {
		formal = p.FormalParameterArg(formalParam.FormalParameterList())
	} else if formalParam.Identifier() != nil {
		formal = append(formal, formalParam.Identifier().GetText())
	}
	p2 := newProgram(p.file)
	if body != nil && body.SourceElements() != nil {
		p2.compile(body.SourceElements())
	}
	p.push(function{formal: formal, body: p2.Code}, bpr, p.file)
}

func (p *Program) createArrowSingleFunc(formalParam parser.IArrowFunctionParametersContext, body parser.ISingleExpressionContext, bpr *antlr.BaseParserRuleContext) {
	formal := make([]string, 0)
	if formalParam.FormalParameterList() != nil {
		formal = p.FormalParameterArg(formalParam.FormalParameterList())
	} else if formalParam.Identifier() != nil {
		formal = append(formal, formalParam.Identifier().GetText())
	}
	p2 := newProgram(p.file)
	if body != nil {
		p2.SingleExpression(body)
	}
	p.push(function{formal: formal, body: p2.Code}, bpr, p.file)
}

func (p *Program) createFunc(formalParam parser.IFormalParameterListContext, body parser.IFunctionBodyContext, bpr *antlr.BaseParserRuleContext) {
	formal := make([]string, 0)
	if formalParam != nil {
		formal = p.FormalParameterArg(formalParam)
	}
	p2 := newProgram(p.file)
	if body != nil && body.SourceElements() != nil {
		p2.compile(body.SourceElements())
	}
	p.push(function{formal: formal, body: p2.Code}, bpr, p.file)
}

func (p *Program) FormalParameterArg(params parser.IFormalParameterListContext) []string {
	formal := make([]string, 0)
	for _, v := range params.AllFormalParameterArg() {
		formal = append(formal, v.Assignable().GetText())
	}
	last := params.LastFormalParameterArg()
	if last != nil {
		if v, ok := last.GetChild(1).(*parser.IdentifierExpressionContext); ok {
			formal = append(formal, "..."+v.Identifier().GetText())
		} else {
			PrintError("SyntaxError: Unexpected token '" + last.GetText() + "'")
			PrintErrorStack(p.file, last.GetStart().GetLine(), last.GetStart().GetColumn(), last.GetText())
			panic("")
		}
	}
	return formal
}

func (p *Program) ExpressionSequence(ctx *parser.ExpressionSequenceContext) {
	c := len(ctx.AllSingleExpression())
	for i := c - 1; i >= 0; i-- {
		p.SingleExpression(ctx.SingleExpression(i))
	}
}

func (p *Program) leftHand(bpr *antlr.BaseParserRuleContext) {
	if v, ok := p.pop(); ok {
		if n, ok := v.(property); ok {
			if n.name == "" {
				if v, ok := p.pop(); ok {
					if val, ok := v.(loadVal); ok {
						p.push(modifyProperty{name: JsToString(val.value)}, bpr, p.file)
					}
				}
			} else {
				p.push(modifyProperty{name: n.name}, bpr, p.file)
			}
		} else {
			p.push(v, bpr, p.file)
		}
	} else {
		PrintError("SyntaxError: Invalid left-hand side in assignment")
		PrintErrorStack(p.file, bpr.GetStart().GetLine(), bpr.GetStart().GetColumn(), bpr.GetText())
		panic("")
	}
}

func (p *Program) Declaration(ctx *parser.DeclarationContext) {
	v := ctx.GetChild(0)
	switch v.(type) {
	case *parser.VariableStatementContext:
		p.VariableStatement(v.(*parser.VariableStatementContext))
	case *parser.FunctionDeclarationContext:
		p.FunctionDeclaration(v.(*parser.FunctionDeclarationContext))
	}
}
