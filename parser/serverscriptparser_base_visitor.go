// Code generated from /Users/lzp/代码/workspace_go/ss/parser/ServerScriptParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // ServerScriptParser

import "github.com/antlr4-go/antlr/v4"

type BaseServerScriptParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseServerScriptParserVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitSourceElements(ctx *SourceElementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitSourceElement(ctx *SourceElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitStatementList(ctx *StatementListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitImportStatement(ctx *ImportStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitImportFromBlock1(ctx *ImportFromBlock1Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitImportFromBlock2(ctx *ImportFromBlock2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitImportFromBlock3(ctx *ImportFromBlock3Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitImportModuleItems(ctx *ImportModuleItemsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitImportAliasName(ctx *ImportAliasNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitModuleExportName(ctx *ModuleExportNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitImportedBinding(ctx *ImportedBindingContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitImportDefault(ctx *ImportDefaultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitImportNamespace(ctx *ImportNamespaceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitImportFrom(ctx *ImportFromContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitAliasName(ctx *AliasNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitExportDeclaration(ctx *ExportDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitExportItems(ctx *ExportItemsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitExportBlock(ctx *ExportBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitExportDefaultDeclaration(ctx *ExportDefaultDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitExportNamespace(ctx *ExportNamespaceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitExportFromBlock(ctx *ExportFromBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitExportModuleItems(ctx *ExportModuleItemsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitExportAliasName(ctx *ExportAliasNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitDeclaration(ctx *DeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitVariableStatement(ctx *VariableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitVariableDeclarationList(ctx *VariableDeclarationListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitVariableDeclaration(ctx *VariableDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitEmptyStatement_(ctx *EmptyStatement_Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitExpressionStatement(ctx *ExpressionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitIfStatement(ctx *IfStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitIfBlock(ctx *IfBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitElseIfBlock(ctx *ElseIfBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitElseBlock(ctx *ElseBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitDoStatement(ctx *DoStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitWhileStatement(ctx *WhileStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitForStatement(ctx *ForStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitForOfStatement(ctx *ForOfStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitOfArrayLiteral(ctx *OfArrayLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitIdentifierList(ctx *IdentifierListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitForStatement1(ctx *ForStatement1Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitForStatement2(ctx *ForStatement2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitForStatement3(ctx *ForStatement3Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitVarModifier(ctx *VarModifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitContinueStatement(ctx *ContinueStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitBreakStatement(ctx *BreakStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitReturnStatement(ctx *ReturnStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitSwitchStatement(ctx *SwitchStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitCaseBlock(ctx *CaseBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitCaseClauses(ctx *CaseClausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitCaseClause(ctx *CaseClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitDefaultClause(ctx *DefaultClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitThrowStatement(ctx *ThrowStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitTryStatement(ctx *TryStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitCatchProduction(ctx *CatchProductionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitFinallyProduction(ctx *FinallyProductionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitFunctionDeclaration(ctx *FunctionDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitFormalParameterList(ctx *FormalParameterListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitFormalParameterArg(ctx *FormalParameterArgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitLastFormalParameterArg(ctx *LastFormalParameterArgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitFunctionBody(ctx *FunctionBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitArrayLiteral(ctx *ArrayLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitElementList(ctx *ElementListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitArrayElement(ctx *ArrayElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitPropertyExpressionAssignment(ctx *PropertyExpressionAssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitComputedPropertyExpressionAssignment(ctx *ComputedPropertyExpressionAssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitFunctionProperty(ctx *FunctionPropertyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitPropertyShorthand(ctx *PropertyShorthandContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitPropertyName(ctx *PropertyNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitArguments(ctx *ArgumentsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitArgument(ctx *ArgumentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitExpressionSequence(ctx *ExpressionSequenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitTemplateStringExpression(ctx *TemplateStringExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitTernaryExpression(ctx *TernaryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitPowerExpression(ctx *PowerExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitPreIncrementExpression(ctx *PreIncrementExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitObjectLiteralExpression(ctx *ObjectLiteralExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitPrefixArgumentsExpression(ctx *PrefixArgumentsExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitInExpression(ctx *InExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitNotExpression(ctx *NotExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitPreDecreaseExpression(ctx *PreDecreaseExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitArgumentsExpression(ctx *ArgumentsExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitThisExpression(ctx *ThisExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitLogicalExpression(ctx *LogicalExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitFunctionExpression(ctx *FunctionExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitUnaryMinusExpression(ctx *UnaryMinusExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitAssignmentExpression(ctx *AssignmentExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitPostDecreaseExpression(ctx *PostDecreaseExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitUnaryPlusExpression(ctx *UnaryPlusExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitDeleteExpression(ctx *DeleteExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitEqualityExpression(ctx *EqualityExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitMultiplicativeExpression(ctx *MultiplicativeExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitBitShiftExpression(ctx *BitShiftExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitParenthesizedExpression(ctx *ParenthesizedExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitAdditiveExpression(ctx *AdditiveExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitRelationalExpression(ctx *RelationalExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitPostIncrementExpression(ctx *PostIncrementExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitBitNotExpression(ctx *BitNotExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitNewExpression(ctx *NewExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitLiteralExpression(ctx *LiteralExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitArrayLiteralExpression(ctx *ArrayLiteralExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitMemberDotExpression(ctx *MemberDotExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitMemberIndexExpression(ctx *MemberIndexExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitIdentifierExpression(ctx *IdentifierExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitBitExpression(ctx *BitExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitAssignmentOperatorExpression(ctx *AssignmentOperatorExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitCoalesceExpression(ctx *CoalesceExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitAssignable(ctx *AssignableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitObjectLiteral(ctx *ObjectLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitAnonymousFunctionDecl(ctx *AnonymousFunctionDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitArrowFunction(ctx *ArrowFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitArrowSingleFunction(ctx *ArrowSingleFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitArrowFunctionParameters(ctx *ArrowFunctionParametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitAssignmentOperator(ctx *AssignmentOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitLiteral(ctx *LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitTemplateStringLiteral(ctx *TemplateStringLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitTemplateStringAtom(ctx *TemplateStringAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitNumericLiteral(ctx *NumericLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitBigintLiteral(ctx *BigintLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitIdentifierName(ctx *IdentifierNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitIdentifier(ctx *IdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitReservedWord(ctx *ReservedWordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitKeyword(ctx *KeywordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseServerScriptParserVisitor) VisitEos(ctx *EosContext) interface{} {
	return v.VisitChildren(ctx)
}
