// Code generated from /Users/lzp/代码/workspace_go/ss/parser/ServerScriptParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // ServerScriptParser

import "github.com/antlr4-go/antlr/v4"

// BaseServerScriptParserListener is a complete listener for a parse tree produced by ServerScriptParser.
type BaseServerScriptParserListener struct{}

var _ ServerScriptParserListener = &BaseServerScriptParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseServerScriptParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseServerScriptParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseServerScriptParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseServerScriptParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseServerScriptParserListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseServerScriptParserListener) ExitProgram(ctx *ProgramContext) {}

// EnterSourceElements is called when production sourceElements is entered.
func (s *BaseServerScriptParserListener) EnterSourceElements(ctx *SourceElementsContext) {}

// ExitSourceElements is called when production sourceElements is exited.
func (s *BaseServerScriptParserListener) ExitSourceElements(ctx *SourceElementsContext) {}

// EnterSourceElement is called when production sourceElement is entered.
func (s *BaseServerScriptParserListener) EnterSourceElement(ctx *SourceElementContext) {}

// ExitSourceElement is called when production sourceElement is exited.
func (s *BaseServerScriptParserListener) ExitSourceElement(ctx *SourceElementContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseServerScriptParserListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseServerScriptParserListener) ExitStatement(ctx *StatementContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseServerScriptParserListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseServerScriptParserListener) ExitBlock(ctx *BlockContext) {}

// EnterStatementList is called when production statementList is entered.
func (s *BaseServerScriptParserListener) EnterStatementList(ctx *StatementListContext) {}

// ExitStatementList is called when production statementList is exited.
func (s *BaseServerScriptParserListener) ExitStatementList(ctx *StatementListContext) {}

// EnterImportStatement is called when production importStatement is entered.
func (s *BaseServerScriptParserListener) EnterImportStatement(ctx *ImportStatementContext) {}

// ExitImportStatement is called when production importStatement is exited.
func (s *BaseServerScriptParserListener) ExitImportStatement(ctx *ImportStatementContext) {}

// EnterImportFromBlock1 is called when production ImportFromBlock1 is entered.
func (s *BaseServerScriptParserListener) EnterImportFromBlock1(ctx *ImportFromBlock1Context) {}

// ExitImportFromBlock1 is called when production ImportFromBlock1 is exited.
func (s *BaseServerScriptParserListener) ExitImportFromBlock1(ctx *ImportFromBlock1Context) {}

// EnterImportFromBlock2 is called when production ImportFromBlock2 is entered.
func (s *BaseServerScriptParserListener) EnterImportFromBlock2(ctx *ImportFromBlock2Context) {}

// ExitImportFromBlock2 is called when production ImportFromBlock2 is exited.
func (s *BaseServerScriptParserListener) ExitImportFromBlock2(ctx *ImportFromBlock2Context) {}

// EnterImportFromBlock3 is called when production ImportFromBlock3 is entered.
func (s *BaseServerScriptParserListener) EnterImportFromBlock3(ctx *ImportFromBlock3Context) {}

// ExitImportFromBlock3 is called when production ImportFromBlock3 is exited.
func (s *BaseServerScriptParserListener) ExitImportFromBlock3(ctx *ImportFromBlock3Context) {}

// EnterImportModuleItems is called when production importModuleItems is entered.
func (s *BaseServerScriptParserListener) EnterImportModuleItems(ctx *ImportModuleItemsContext) {}

// ExitImportModuleItems is called when production importModuleItems is exited.
func (s *BaseServerScriptParserListener) ExitImportModuleItems(ctx *ImportModuleItemsContext) {}

// EnterImportAliasName is called when production importAliasName is entered.
func (s *BaseServerScriptParserListener) EnterImportAliasName(ctx *ImportAliasNameContext) {}

// ExitImportAliasName is called when production importAliasName is exited.
func (s *BaseServerScriptParserListener) ExitImportAliasName(ctx *ImportAliasNameContext) {}

// EnterModuleExportName is called when production moduleExportName is entered.
func (s *BaseServerScriptParserListener) EnterModuleExportName(ctx *ModuleExportNameContext) {}

// ExitModuleExportName is called when production moduleExportName is exited.
func (s *BaseServerScriptParserListener) ExitModuleExportName(ctx *ModuleExportNameContext) {}

// EnterImportedBinding is called when production importedBinding is entered.
func (s *BaseServerScriptParserListener) EnterImportedBinding(ctx *ImportedBindingContext) {}

// ExitImportedBinding is called when production importedBinding is exited.
func (s *BaseServerScriptParserListener) ExitImportedBinding(ctx *ImportedBindingContext) {}

// EnterImportDefault is called when production importDefault is entered.
func (s *BaseServerScriptParserListener) EnterImportDefault(ctx *ImportDefaultContext) {}

// ExitImportDefault is called when production importDefault is exited.
func (s *BaseServerScriptParserListener) ExitImportDefault(ctx *ImportDefaultContext) {}

// EnterImportNamespace is called when production importNamespace is entered.
func (s *BaseServerScriptParserListener) EnterImportNamespace(ctx *ImportNamespaceContext) {}

// ExitImportNamespace is called when production importNamespace is exited.
func (s *BaseServerScriptParserListener) ExitImportNamespace(ctx *ImportNamespaceContext) {}

// EnterImportFrom is called when production importFrom is entered.
func (s *BaseServerScriptParserListener) EnterImportFrom(ctx *ImportFromContext) {}

// ExitImportFrom is called when production importFrom is exited.
func (s *BaseServerScriptParserListener) ExitImportFrom(ctx *ImportFromContext) {}

// EnterAliasName is called when production aliasName is entered.
func (s *BaseServerScriptParserListener) EnterAliasName(ctx *AliasNameContext) {}

// ExitAliasName is called when production aliasName is exited.
func (s *BaseServerScriptParserListener) ExitAliasName(ctx *AliasNameContext) {}

// EnterExportDeclaration is called when production ExportDeclaration is entered.
func (s *BaseServerScriptParserListener) EnterExportDeclaration(ctx *ExportDeclarationContext) {}

// ExitExportDeclaration is called when production ExportDeclaration is exited.
func (s *BaseServerScriptParserListener) ExitExportDeclaration(ctx *ExportDeclarationContext) {}

// EnterExportItems is called when production ExportItems is entered.
func (s *BaseServerScriptParserListener) EnterExportItems(ctx *ExportItemsContext) {}

// ExitExportItems is called when production ExportItems is exited.
func (s *BaseServerScriptParserListener) ExitExportItems(ctx *ExportItemsContext) {}

// EnterExportBlock is called when production ExportBlock is entered.
func (s *BaseServerScriptParserListener) EnterExportBlock(ctx *ExportBlockContext) {}

// ExitExportBlock is called when production ExportBlock is exited.
func (s *BaseServerScriptParserListener) ExitExportBlock(ctx *ExportBlockContext) {}

// EnterExportDefaultDeclaration is called when production ExportDefaultDeclaration is entered.
func (s *BaseServerScriptParserListener) EnterExportDefaultDeclaration(ctx *ExportDefaultDeclarationContext) {
}

// ExitExportDefaultDeclaration is called when production ExportDefaultDeclaration is exited.
func (s *BaseServerScriptParserListener) ExitExportDefaultDeclaration(ctx *ExportDefaultDeclarationContext) {
}

// EnterExportNamespace is called when production exportNamespace is entered.
func (s *BaseServerScriptParserListener) EnterExportNamespace(ctx *ExportNamespaceContext) {}

// ExitExportNamespace is called when production exportNamespace is exited.
func (s *BaseServerScriptParserListener) ExitExportNamespace(ctx *ExportNamespaceContext) {}

// EnterExportFromBlock is called when production exportFromBlock is entered.
func (s *BaseServerScriptParserListener) EnterExportFromBlock(ctx *ExportFromBlockContext) {}

// ExitExportFromBlock is called when production exportFromBlock is exited.
func (s *BaseServerScriptParserListener) ExitExportFromBlock(ctx *ExportFromBlockContext) {}

// EnterExportModuleItems is called when production exportModuleItems is entered.
func (s *BaseServerScriptParserListener) EnterExportModuleItems(ctx *ExportModuleItemsContext) {}

// ExitExportModuleItems is called when production exportModuleItems is exited.
func (s *BaseServerScriptParserListener) ExitExportModuleItems(ctx *ExportModuleItemsContext) {}

// EnterExportAliasName is called when production exportAliasName is entered.
func (s *BaseServerScriptParserListener) EnterExportAliasName(ctx *ExportAliasNameContext) {}

// ExitExportAliasName is called when production exportAliasName is exited.
func (s *BaseServerScriptParserListener) ExitExportAliasName(ctx *ExportAliasNameContext) {}

// EnterDeclaration is called when production declaration is entered.
func (s *BaseServerScriptParserListener) EnterDeclaration(ctx *DeclarationContext) {}

// ExitDeclaration is called when production declaration is exited.
func (s *BaseServerScriptParserListener) ExitDeclaration(ctx *DeclarationContext) {}

// EnterVariableStatement is called when production variableStatement is entered.
func (s *BaseServerScriptParserListener) EnterVariableStatement(ctx *VariableStatementContext) {}

// ExitVariableStatement is called when production variableStatement is exited.
func (s *BaseServerScriptParserListener) ExitVariableStatement(ctx *VariableStatementContext) {}

// EnterVariableDeclarationList is called when production variableDeclarationList is entered.
func (s *BaseServerScriptParserListener) EnterVariableDeclarationList(ctx *VariableDeclarationListContext) {
}

// ExitVariableDeclarationList is called when production variableDeclarationList is exited.
func (s *BaseServerScriptParserListener) ExitVariableDeclarationList(ctx *VariableDeclarationListContext) {
}

// EnterVariableDeclaration is called when production variableDeclaration is entered.
func (s *BaseServerScriptParserListener) EnterVariableDeclaration(ctx *VariableDeclarationContext) {}

// ExitVariableDeclaration is called when production variableDeclaration is exited.
func (s *BaseServerScriptParserListener) ExitVariableDeclaration(ctx *VariableDeclarationContext) {}

// EnterEmptyStatement_ is called when production emptyStatement_ is entered.
func (s *BaseServerScriptParserListener) EnterEmptyStatement_(ctx *EmptyStatement_Context) {}

// ExitEmptyStatement_ is called when production emptyStatement_ is exited.
func (s *BaseServerScriptParserListener) ExitEmptyStatement_(ctx *EmptyStatement_Context) {}

// EnterExpressionStatement is called when production expressionStatement is entered.
func (s *BaseServerScriptParserListener) EnterExpressionStatement(ctx *ExpressionStatementContext) {}

// ExitExpressionStatement is called when production expressionStatement is exited.
func (s *BaseServerScriptParserListener) ExitExpressionStatement(ctx *ExpressionStatementContext) {}

// EnterIfStatement is called when production ifStatement is entered.
func (s *BaseServerScriptParserListener) EnterIfStatement(ctx *IfStatementContext) {}

// ExitIfStatement is called when production ifStatement is exited.
func (s *BaseServerScriptParserListener) ExitIfStatement(ctx *IfStatementContext) {}

// EnterIfBlock is called when production ifBlock is entered.
func (s *BaseServerScriptParserListener) EnterIfBlock(ctx *IfBlockContext) {}

// ExitIfBlock is called when production ifBlock is exited.
func (s *BaseServerScriptParserListener) ExitIfBlock(ctx *IfBlockContext) {}

// EnterElseIfBlock is called when production elseIfBlock is entered.
func (s *BaseServerScriptParserListener) EnterElseIfBlock(ctx *ElseIfBlockContext) {}

// ExitElseIfBlock is called when production elseIfBlock is exited.
func (s *BaseServerScriptParserListener) ExitElseIfBlock(ctx *ElseIfBlockContext) {}

// EnterElseBlock is called when production elseBlock is entered.
func (s *BaseServerScriptParserListener) EnterElseBlock(ctx *ElseBlockContext) {}

// ExitElseBlock is called when production elseBlock is exited.
func (s *BaseServerScriptParserListener) ExitElseBlock(ctx *ElseBlockContext) {}

// EnterDoStatement is called when production DoStatement is entered.
func (s *BaseServerScriptParserListener) EnterDoStatement(ctx *DoStatementContext) {}

// ExitDoStatement is called when production DoStatement is exited.
func (s *BaseServerScriptParserListener) ExitDoStatement(ctx *DoStatementContext) {}

// EnterWhileStatement is called when production WhileStatement is entered.
func (s *BaseServerScriptParserListener) EnterWhileStatement(ctx *WhileStatementContext) {}

// ExitWhileStatement is called when production WhileStatement is exited.
func (s *BaseServerScriptParserListener) ExitWhileStatement(ctx *WhileStatementContext) {}

// EnterForStatement is called when production ForStatement is entered.
func (s *BaseServerScriptParserListener) EnterForStatement(ctx *ForStatementContext) {}

// ExitForStatement is called when production ForStatement is exited.
func (s *BaseServerScriptParserListener) ExitForStatement(ctx *ForStatementContext) {}

// EnterForOfStatement is called when production ForOfStatement is entered.
func (s *BaseServerScriptParserListener) EnterForOfStatement(ctx *ForOfStatementContext) {}

// ExitForOfStatement is called when production ForOfStatement is exited.
func (s *BaseServerScriptParserListener) ExitForOfStatement(ctx *ForOfStatementContext) {}

// EnterOfArrayLiteral is called when production ofArrayLiteral is entered.
func (s *BaseServerScriptParserListener) EnterOfArrayLiteral(ctx *OfArrayLiteralContext) {}

// ExitOfArrayLiteral is called when production ofArrayLiteral is exited.
func (s *BaseServerScriptParserListener) ExitOfArrayLiteral(ctx *OfArrayLiteralContext) {}

// EnterIdentifierList is called when production identifierList is entered.
func (s *BaseServerScriptParserListener) EnterIdentifierList(ctx *IdentifierListContext) {}

// ExitIdentifierList is called when production identifierList is exited.
func (s *BaseServerScriptParserListener) ExitIdentifierList(ctx *IdentifierListContext) {}

// EnterForStatement1 is called when production forStatement1 is entered.
func (s *BaseServerScriptParserListener) EnterForStatement1(ctx *ForStatement1Context) {}

// ExitForStatement1 is called when production forStatement1 is exited.
func (s *BaseServerScriptParserListener) ExitForStatement1(ctx *ForStatement1Context) {}

// EnterForStatement2 is called when production forStatement2 is entered.
func (s *BaseServerScriptParserListener) EnterForStatement2(ctx *ForStatement2Context) {}

// ExitForStatement2 is called when production forStatement2 is exited.
func (s *BaseServerScriptParserListener) ExitForStatement2(ctx *ForStatement2Context) {}

// EnterForStatement3 is called when production forStatement3 is entered.
func (s *BaseServerScriptParserListener) EnterForStatement3(ctx *ForStatement3Context) {}

// ExitForStatement3 is called when production forStatement3 is exited.
func (s *BaseServerScriptParserListener) ExitForStatement3(ctx *ForStatement3Context) {}

// EnterVarModifier is called when production varModifier is entered.
func (s *BaseServerScriptParserListener) EnterVarModifier(ctx *VarModifierContext) {}

// ExitVarModifier is called when production varModifier is exited.
func (s *BaseServerScriptParserListener) ExitVarModifier(ctx *VarModifierContext) {}

// EnterContinueStatement is called when production continueStatement is entered.
func (s *BaseServerScriptParserListener) EnterContinueStatement(ctx *ContinueStatementContext) {}

// ExitContinueStatement is called when production continueStatement is exited.
func (s *BaseServerScriptParserListener) ExitContinueStatement(ctx *ContinueStatementContext) {}

// EnterBreakStatement is called when production breakStatement is entered.
func (s *BaseServerScriptParserListener) EnterBreakStatement(ctx *BreakStatementContext) {}

// ExitBreakStatement is called when production breakStatement is exited.
func (s *BaseServerScriptParserListener) ExitBreakStatement(ctx *BreakStatementContext) {}

// EnterReturnStatement is called when production returnStatement is entered.
func (s *BaseServerScriptParserListener) EnterReturnStatement(ctx *ReturnStatementContext) {}

// ExitReturnStatement is called when production returnStatement is exited.
func (s *BaseServerScriptParserListener) ExitReturnStatement(ctx *ReturnStatementContext) {}

// EnterSwitchStatement is called when production switchStatement is entered.
func (s *BaseServerScriptParserListener) EnterSwitchStatement(ctx *SwitchStatementContext) {}

// ExitSwitchStatement is called when production switchStatement is exited.
func (s *BaseServerScriptParserListener) ExitSwitchStatement(ctx *SwitchStatementContext) {}

// EnterCaseBlock is called when production caseBlock is entered.
func (s *BaseServerScriptParserListener) EnterCaseBlock(ctx *CaseBlockContext) {}

// ExitCaseBlock is called when production caseBlock is exited.
func (s *BaseServerScriptParserListener) ExitCaseBlock(ctx *CaseBlockContext) {}

// EnterCaseClauses is called when production caseClauses is entered.
func (s *BaseServerScriptParserListener) EnterCaseClauses(ctx *CaseClausesContext) {}

// ExitCaseClauses is called when production caseClauses is exited.
func (s *BaseServerScriptParserListener) ExitCaseClauses(ctx *CaseClausesContext) {}

// EnterCaseClause is called when production caseClause is entered.
func (s *BaseServerScriptParserListener) EnterCaseClause(ctx *CaseClauseContext) {}

// ExitCaseClause is called when production caseClause is exited.
func (s *BaseServerScriptParserListener) ExitCaseClause(ctx *CaseClauseContext) {}

// EnterDefaultClause is called when production defaultClause is entered.
func (s *BaseServerScriptParserListener) EnterDefaultClause(ctx *DefaultClauseContext) {}

// ExitDefaultClause is called when production defaultClause is exited.
func (s *BaseServerScriptParserListener) ExitDefaultClause(ctx *DefaultClauseContext) {}

// EnterThrowStatement is called when production throwStatement is entered.
func (s *BaseServerScriptParserListener) EnterThrowStatement(ctx *ThrowStatementContext) {}

// ExitThrowStatement is called when production throwStatement is exited.
func (s *BaseServerScriptParserListener) ExitThrowStatement(ctx *ThrowStatementContext) {}

// EnterTryStatement is called when production tryStatement is entered.
func (s *BaseServerScriptParserListener) EnterTryStatement(ctx *TryStatementContext) {}

// ExitTryStatement is called when production tryStatement is exited.
func (s *BaseServerScriptParserListener) ExitTryStatement(ctx *TryStatementContext) {}

// EnterCatchProduction is called when production catchProduction is entered.
func (s *BaseServerScriptParserListener) EnterCatchProduction(ctx *CatchProductionContext) {}

// ExitCatchProduction is called when production catchProduction is exited.
func (s *BaseServerScriptParserListener) ExitCatchProduction(ctx *CatchProductionContext) {}

// EnterFinallyProduction is called when production finallyProduction is entered.
func (s *BaseServerScriptParserListener) EnterFinallyProduction(ctx *FinallyProductionContext) {}

// ExitFinallyProduction is called when production finallyProduction is exited.
func (s *BaseServerScriptParserListener) ExitFinallyProduction(ctx *FinallyProductionContext) {}

// EnterFunctionDeclaration is called when production functionDeclaration is entered.
func (s *BaseServerScriptParserListener) EnterFunctionDeclaration(ctx *FunctionDeclarationContext) {}

// ExitFunctionDeclaration is called when production functionDeclaration is exited.
func (s *BaseServerScriptParserListener) ExitFunctionDeclaration(ctx *FunctionDeclarationContext) {}

// EnterFormalParameterList is called when production formalParameterList is entered.
func (s *BaseServerScriptParserListener) EnterFormalParameterList(ctx *FormalParameterListContext) {}

// ExitFormalParameterList is called when production formalParameterList is exited.
func (s *BaseServerScriptParserListener) ExitFormalParameterList(ctx *FormalParameterListContext) {}

// EnterFormalParameterArg is called when production formalParameterArg is entered.
func (s *BaseServerScriptParserListener) EnterFormalParameterArg(ctx *FormalParameterArgContext) {}

// ExitFormalParameterArg is called when production formalParameterArg is exited.
func (s *BaseServerScriptParserListener) ExitFormalParameterArg(ctx *FormalParameterArgContext) {}

// EnterLastFormalParameterArg is called when production lastFormalParameterArg is entered.
func (s *BaseServerScriptParserListener) EnterLastFormalParameterArg(ctx *LastFormalParameterArgContext) {
}

// ExitLastFormalParameterArg is called when production lastFormalParameterArg is exited.
func (s *BaseServerScriptParserListener) ExitLastFormalParameterArg(ctx *LastFormalParameterArgContext) {
}

// EnterFunctionBody is called when production functionBody is entered.
func (s *BaseServerScriptParserListener) EnterFunctionBody(ctx *FunctionBodyContext) {}

// ExitFunctionBody is called when production functionBody is exited.
func (s *BaseServerScriptParserListener) ExitFunctionBody(ctx *FunctionBodyContext) {}

// EnterArrayLiteral is called when production arrayLiteral is entered.
func (s *BaseServerScriptParserListener) EnterArrayLiteral(ctx *ArrayLiteralContext) {}

// ExitArrayLiteral is called when production arrayLiteral is exited.
func (s *BaseServerScriptParserListener) ExitArrayLiteral(ctx *ArrayLiteralContext) {}

// EnterElementList is called when production elementList is entered.
func (s *BaseServerScriptParserListener) EnterElementList(ctx *ElementListContext) {}

// ExitElementList is called when production elementList is exited.
func (s *BaseServerScriptParserListener) ExitElementList(ctx *ElementListContext) {}

// EnterArrayElement is called when production arrayElement is entered.
func (s *BaseServerScriptParserListener) EnterArrayElement(ctx *ArrayElementContext) {}

// ExitArrayElement is called when production arrayElement is exited.
func (s *BaseServerScriptParserListener) ExitArrayElement(ctx *ArrayElementContext) {}

// EnterPropertyExpressionAssignment is called when production PropertyExpressionAssignment is entered.
func (s *BaseServerScriptParserListener) EnterPropertyExpressionAssignment(ctx *PropertyExpressionAssignmentContext) {
}

// ExitPropertyExpressionAssignment is called when production PropertyExpressionAssignment is exited.
func (s *BaseServerScriptParserListener) ExitPropertyExpressionAssignment(ctx *PropertyExpressionAssignmentContext) {
}

// EnterComputedPropertyExpressionAssignment is called when production ComputedPropertyExpressionAssignment is entered.
func (s *BaseServerScriptParserListener) EnterComputedPropertyExpressionAssignment(ctx *ComputedPropertyExpressionAssignmentContext) {
}

// ExitComputedPropertyExpressionAssignment is called when production ComputedPropertyExpressionAssignment is exited.
func (s *BaseServerScriptParserListener) ExitComputedPropertyExpressionAssignment(ctx *ComputedPropertyExpressionAssignmentContext) {
}

// EnterFunctionProperty is called when production FunctionProperty is entered.
func (s *BaseServerScriptParserListener) EnterFunctionProperty(ctx *FunctionPropertyContext) {}

// ExitFunctionProperty is called when production FunctionProperty is exited.
func (s *BaseServerScriptParserListener) ExitFunctionProperty(ctx *FunctionPropertyContext) {}

// EnterPropertyShorthand is called when production PropertyShorthand is entered.
func (s *BaseServerScriptParserListener) EnterPropertyShorthand(ctx *PropertyShorthandContext) {}

// ExitPropertyShorthand is called when production PropertyShorthand is exited.
func (s *BaseServerScriptParserListener) ExitPropertyShorthand(ctx *PropertyShorthandContext) {}

// EnterPropertyName is called when production propertyName is entered.
func (s *BaseServerScriptParserListener) EnterPropertyName(ctx *PropertyNameContext) {}

// ExitPropertyName is called when production propertyName is exited.
func (s *BaseServerScriptParserListener) ExitPropertyName(ctx *PropertyNameContext) {}

// EnterArguments is called when production arguments is entered.
func (s *BaseServerScriptParserListener) EnterArguments(ctx *ArgumentsContext) {}

// ExitArguments is called when production arguments is exited.
func (s *BaseServerScriptParserListener) ExitArguments(ctx *ArgumentsContext) {}

// EnterArgument is called when production argument is entered.
func (s *BaseServerScriptParserListener) EnterArgument(ctx *ArgumentContext) {}

// ExitArgument is called when production argument is exited.
func (s *BaseServerScriptParserListener) ExitArgument(ctx *ArgumentContext) {}

// EnterExpressionSequence is called when production expressionSequence is entered.
func (s *BaseServerScriptParserListener) EnterExpressionSequence(ctx *ExpressionSequenceContext) {}

// ExitExpressionSequence is called when production expressionSequence is exited.
func (s *BaseServerScriptParserListener) ExitExpressionSequence(ctx *ExpressionSequenceContext) {}

// EnterTemplateStringExpression is called when production TemplateStringExpression is entered.
func (s *BaseServerScriptParserListener) EnterTemplateStringExpression(ctx *TemplateStringExpressionContext) {
}

// ExitTemplateStringExpression is called when production TemplateStringExpression is exited.
func (s *BaseServerScriptParserListener) ExitTemplateStringExpression(ctx *TemplateStringExpressionContext) {
}

// EnterTernaryExpression is called when production TernaryExpression is entered.
func (s *BaseServerScriptParserListener) EnterTernaryExpression(ctx *TernaryExpressionContext) {}

// ExitTernaryExpression is called when production TernaryExpression is exited.
func (s *BaseServerScriptParserListener) ExitTernaryExpression(ctx *TernaryExpressionContext) {}

// EnterPowerExpression is called when production PowerExpression is entered.
func (s *BaseServerScriptParserListener) EnterPowerExpression(ctx *PowerExpressionContext) {}

// ExitPowerExpression is called when production PowerExpression is exited.
func (s *BaseServerScriptParserListener) ExitPowerExpression(ctx *PowerExpressionContext) {}

// EnterPreIncrementExpression is called when production PreIncrementExpression is entered.
func (s *BaseServerScriptParserListener) EnterPreIncrementExpression(ctx *PreIncrementExpressionContext) {
}

// ExitPreIncrementExpression is called when production PreIncrementExpression is exited.
func (s *BaseServerScriptParserListener) ExitPreIncrementExpression(ctx *PreIncrementExpressionContext) {
}

// EnterObjectLiteralExpression is called when production ObjectLiteralExpression is entered.
func (s *BaseServerScriptParserListener) EnterObjectLiteralExpression(ctx *ObjectLiteralExpressionContext) {
}

// ExitObjectLiteralExpression is called when production ObjectLiteralExpression is exited.
func (s *BaseServerScriptParserListener) ExitObjectLiteralExpression(ctx *ObjectLiteralExpressionContext) {
}

// EnterPrefixArgumentsExpression is called when production PrefixArgumentsExpression is entered.
func (s *BaseServerScriptParserListener) EnterPrefixArgumentsExpression(ctx *PrefixArgumentsExpressionContext) {
}

// ExitPrefixArgumentsExpression is called when production PrefixArgumentsExpression is exited.
func (s *BaseServerScriptParserListener) ExitPrefixArgumentsExpression(ctx *PrefixArgumentsExpressionContext) {
}

// EnterInExpression is called when production InExpression is entered.
func (s *BaseServerScriptParserListener) EnterInExpression(ctx *InExpressionContext) {}

// ExitInExpression is called when production InExpression is exited.
func (s *BaseServerScriptParserListener) ExitInExpression(ctx *InExpressionContext) {}

// EnterNotExpression is called when production NotExpression is entered.
func (s *BaseServerScriptParserListener) EnterNotExpression(ctx *NotExpressionContext) {}

// ExitNotExpression is called when production NotExpression is exited.
func (s *BaseServerScriptParserListener) ExitNotExpression(ctx *NotExpressionContext) {}

// EnterPreDecreaseExpression is called when production PreDecreaseExpression is entered.
func (s *BaseServerScriptParserListener) EnterPreDecreaseExpression(ctx *PreDecreaseExpressionContext) {
}

// ExitPreDecreaseExpression is called when production PreDecreaseExpression is exited.
func (s *BaseServerScriptParserListener) ExitPreDecreaseExpression(ctx *PreDecreaseExpressionContext) {
}

// EnterArgumentsExpression is called when production ArgumentsExpression is entered.
func (s *BaseServerScriptParserListener) EnterArgumentsExpression(ctx *ArgumentsExpressionContext) {}

// ExitArgumentsExpression is called when production ArgumentsExpression is exited.
func (s *BaseServerScriptParserListener) ExitArgumentsExpression(ctx *ArgumentsExpressionContext) {}

// EnterThisExpression is called when production ThisExpression is entered.
func (s *BaseServerScriptParserListener) EnterThisExpression(ctx *ThisExpressionContext) {}

// ExitThisExpression is called when production ThisExpression is exited.
func (s *BaseServerScriptParserListener) ExitThisExpression(ctx *ThisExpressionContext) {}

// EnterLogicalExpression is called when production LogicalExpression is entered.
func (s *BaseServerScriptParserListener) EnterLogicalExpression(ctx *LogicalExpressionContext) {}

// ExitLogicalExpression is called when production LogicalExpression is exited.
func (s *BaseServerScriptParserListener) ExitLogicalExpression(ctx *LogicalExpressionContext) {}

// EnterFunctionExpression is called when production FunctionExpression is entered.
func (s *BaseServerScriptParserListener) EnterFunctionExpression(ctx *FunctionExpressionContext) {}

// ExitFunctionExpression is called when production FunctionExpression is exited.
func (s *BaseServerScriptParserListener) ExitFunctionExpression(ctx *FunctionExpressionContext) {}

// EnterUnaryMinusExpression is called when production UnaryMinusExpression is entered.
func (s *BaseServerScriptParserListener) EnterUnaryMinusExpression(ctx *UnaryMinusExpressionContext) {
}

// ExitUnaryMinusExpression is called when production UnaryMinusExpression is exited.
func (s *BaseServerScriptParserListener) ExitUnaryMinusExpression(ctx *UnaryMinusExpressionContext) {}

// EnterAssignmentExpression is called when production AssignmentExpression is entered.
func (s *BaseServerScriptParserListener) EnterAssignmentExpression(ctx *AssignmentExpressionContext) {
}

// ExitAssignmentExpression is called when production AssignmentExpression is exited.
func (s *BaseServerScriptParserListener) ExitAssignmentExpression(ctx *AssignmentExpressionContext) {}

// EnterPostDecreaseExpression is called when production PostDecreaseExpression is entered.
func (s *BaseServerScriptParserListener) EnterPostDecreaseExpression(ctx *PostDecreaseExpressionContext) {
}

// ExitPostDecreaseExpression is called when production PostDecreaseExpression is exited.
func (s *BaseServerScriptParserListener) ExitPostDecreaseExpression(ctx *PostDecreaseExpressionContext) {
}

// EnterUnaryPlusExpression is called when production UnaryPlusExpression is entered.
func (s *BaseServerScriptParserListener) EnterUnaryPlusExpression(ctx *UnaryPlusExpressionContext) {}

// ExitUnaryPlusExpression is called when production UnaryPlusExpression is exited.
func (s *BaseServerScriptParserListener) ExitUnaryPlusExpression(ctx *UnaryPlusExpressionContext) {}

// EnterDeleteExpression is called when production DeleteExpression is entered.
func (s *BaseServerScriptParserListener) EnterDeleteExpression(ctx *DeleteExpressionContext) {}

// ExitDeleteExpression is called when production DeleteExpression is exited.
func (s *BaseServerScriptParserListener) ExitDeleteExpression(ctx *DeleteExpressionContext) {}

// EnterEqualityExpression is called when production EqualityExpression is entered.
func (s *BaseServerScriptParserListener) EnterEqualityExpression(ctx *EqualityExpressionContext) {}

// ExitEqualityExpression is called when production EqualityExpression is exited.
func (s *BaseServerScriptParserListener) ExitEqualityExpression(ctx *EqualityExpressionContext) {}

// EnterMultiplicativeExpression is called when production MultiplicativeExpression is entered.
func (s *BaseServerScriptParserListener) EnterMultiplicativeExpression(ctx *MultiplicativeExpressionContext) {
}

// ExitMultiplicativeExpression is called when production MultiplicativeExpression is exited.
func (s *BaseServerScriptParserListener) ExitMultiplicativeExpression(ctx *MultiplicativeExpressionContext) {
}

// EnterBitShiftExpression is called when production BitShiftExpression is entered.
func (s *BaseServerScriptParserListener) EnterBitShiftExpression(ctx *BitShiftExpressionContext) {}

// ExitBitShiftExpression is called when production BitShiftExpression is exited.
func (s *BaseServerScriptParserListener) ExitBitShiftExpression(ctx *BitShiftExpressionContext) {}

// EnterParenthesizedExpression is called when production ParenthesizedExpression is entered.
func (s *BaseServerScriptParserListener) EnterParenthesizedExpression(ctx *ParenthesizedExpressionContext) {
}

// ExitParenthesizedExpression is called when production ParenthesizedExpression is exited.
func (s *BaseServerScriptParserListener) ExitParenthesizedExpression(ctx *ParenthesizedExpressionContext) {
}

// EnterAdditiveExpression is called when production AdditiveExpression is entered.
func (s *BaseServerScriptParserListener) EnterAdditiveExpression(ctx *AdditiveExpressionContext) {}

// ExitAdditiveExpression is called when production AdditiveExpression is exited.
func (s *BaseServerScriptParserListener) ExitAdditiveExpression(ctx *AdditiveExpressionContext) {}

// EnterRelationalExpression is called when production RelationalExpression is entered.
func (s *BaseServerScriptParserListener) EnterRelationalExpression(ctx *RelationalExpressionContext) {
}

// ExitRelationalExpression is called when production RelationalExpression is exited.
func (s *BaseServerScriptParserListener) ExitRelationalExpression(ctx *RelationalExpressionContext) {}

// EnterPostIncrementExpression is called when production PostIncrementExpression is entered.
func (s *BaseServerScriptParserListener) EnterPostIncrementExpression(ctx *PostIncrementExpressionContext) {
}

// ExitPostIncrementExpression is called when production PostIncrementExpression is exited.
func (s *BaseServerScriptParserListener) ExitPostIncrementExpression(ctx *PostIncrementExpressionContext) {
}

// EnterBitNotExpression is called when production BitNotExpression is entered.
func (s *BaseServerScriptParserListener) EnterBitNotExpression(ctx *BitNotExpressionContext) {}

// ExitBitNotExpression is called when production BitNotExpression is exited.
func (s *BaseServerScriptParserListener) ExitBitNotExpression(ctx *BitNotExpressionContext) {}

// EnterNewExpression is called when production NewExpression is entered.
func (s *BaseServerScriptParserListener) EnterNewExpression(ctx *NewExpressionContext) {}

// ExitNewExpression is called when production NewExpression is exited.
func (s *BaseServerScriptParserListener) ExitNewExpression(ctx *NewExpressionContext) {}

// EnterLiteralExpression is called when production LiteralExpression is entered.
func (s *BaseServerScriptParserListener) EnterLiteralExpression(ctx *LiteralExpressionContext) {}

// ExitLiteralExpression is called when production LiteralExpression is exited.
func (s *BaseServerScriptParserListener) ExitLiteralExpression(ctx *LiteralExpressionContext) {}

// EnterArrayLiteralExpression is called when production ArrayLiteralExpression is entered.
func (s *BaseServerScriptParserListener) EnterArrayLiteralExpression(ctx *ArrayLiteralExpressionContext) {
}

// ExitArrayLiteralExpression is called when production ArrayLiteralExpression is exited.
func (s *BaseServerScriptParserListener) ExitArrayLiteralExpression(ctx *ArrayLiteralExpressionContext) {
}

// EnterMemberDotExpression is called when production MemberDotExpression is entered.
func (s *BaseServerScriptParserListener) EnterMemberDotExpression(ctx *MemberDotExpressionContext) {}

// ExitMemberDotExpression is called when production MemberDotExpression is exited.
func (s *BaseServerScriptParserListener) ExitMemberDotExpression(ctx *MemberDotExpressionContext) {}

// EnterMemberIndexExpression is called when production MemberIndexExpression is entered.
func (s *BaseServerScriptParserListener) EnterMemberIndexExpression(ctx *MemberIndexExpressionContext) {
}

// ExitMemberIndexExpression is called when production MemberIndexExpression is exited.
func (s *BaseServerScriptParserListener) ExitMemberIndexExpression(ctx *MemberIndexExpressionContext) {
}

// EnterIdentifierExpression is called when production IdentifierExpression is entered.
func (s *BaseServerScriptParserListener) EnterIdentifierExpression(ctx *IdentifierExpressionContext) {
}

// ExitIdentifierExpression is called when production IdentifierExpression is exited.
func (s *BaseServerScriptParserListener) ExitIdentifierExpression(ctx *IdentifierExpressionContext) {}

// EnterBitExpression is called when production BitExpression is entered.
func (s *BaseServerScriptParserListener) EnterBitExpression(ctx *BitExpressionContext) {}

// ExitBitExpression is called when production BitExpression is exited.
func (s *BaseServerScriptParserListener) ExitBitExpression(ctx *BitExpressionContext) {}

// EnterAssignmentOperatorExpression is called when production AssignmentOperatorExpression is entered.
func (s *BaseServerScriptParserListener) EnterAssignmentOperatorExpression(ctx *AssignmentOperatorExpressionContext) {
}

// ExitAssignmentOperatorExpression is called when production AssignmentOperatorExpression is exited.
func (s *BaseServerScriptParserListener) ExitAssignmentOperatorExpression(ctx *AssignmentOperatorExpressionContext) {
}

// EnterCoalesceExpression is called when production CoalesceExpression is entered.
func (s *BaseServerScriptParserListener) EnterCoalesceExpression(ctx *CoalesceExpressionContext) {}

// ExitCoalesceExpression is called when production CoalesceExpression is exited.
func (s *BaseServerScriptParserListener) ExitCoalesceExpression(ctx *CoalesceExpressionContext) {}

// EnterAssignable is called when production assignable is entered.
func (s *BaseServerScriptParserListener) EnterAssignable(ctx *AssignableContext) {}

// ExitAssignable is called when production assignable is exited.
func (s *BaseServerScriptParserListener) ExitAssignable(ctx *AssignableContext) {}

// EnterObjectLiteral is called when production objectLiteral is entered.
func (s *BaseServerScriptParserListener) EnterObjectLiteral(ctx *ObjectLiteralContext) {}

// ExitObjectLiteral is called when production objectLiteral is exited.
func (s *BaseServerScriptParserListener) ExitObjectLiteral(ctx *ObjectLiteralContext) {}

// EnterAnonymousFunctionDecl is called when production AnonymousFunctionDecl is entered.
func (s *BaseServerScriptParserListener) EnterAnonymousFunctionDecl(ctx *AnonymousFunctionDeclContext) {
}

// ExitAnonymousFunctionDecl is called when production AnonymousFunctionDecl is exited.
func (s *BaseServerScriptParserListener) ExitAnonymousFunctionDecl(ctx *AnonymousFunctionDeclContext) {
}

// EnterArrowFunction is called when production ArrowFunction is entered.
func (s *BaseServerScriptParserListener) EnterArrowFunction(ctx *ArrowFunctionContext) {}

// ExitArrowFunction is called when production ArrowFunction is exited.
func (s *BaseServerScriptParserListener) ExitArrowFunction(ctx *ArrowFunctionContext) {}

// EnterArrowSingleFunction is called when production ArrowSingleFunction is entered.
func (s *BaseServerScriptParserListener) EnterArrowSingleFunction(ctx *ArrowSingleFunctionContext) {}

// ExitArrowSingleFunction is called when production ArrowSingleFunction is exited.
func (s *BaseServerScriptParserListener) ExitArrowSingleFunction(ctx *ArrowSingleFunctionContext) {}

// EnterArrowFunctionParameters is called when production arrowFunctionParameters is entered.
func (s *BaseServerScriptParserListener) EnterArrowFunctionParameters(ctx *ArrowFunctionParametersContext) {
}

// ExitArrowFunctionParameters is called when production arrowFunctionParameters is exited.
func (s *BaseServerScriptParserListener) ExitArrowFunctionParameters(ctx *ArrowFunctionParametersContext) {
}

// EnterAssignmentOperator is called when production assignmentOperator is entered.
func (s *BaseServerScriptParserListener) EnterAssignmentOperator(ctx *AssignmentOperatorContext) {}

// ExitAssignmentOperator is called when production assignmentOperator is exited.
func (s *BaseServerScriptParserListener) ExitAssignmentOperator(ctx *AssignmentOperatorContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseServerScriptParserListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseServerScriptParserListener) ExitLiteral(ctx *LiteralContext) {}

// EnterTemplateStringLiteral is called when production templateStringLiteral is entered.
func (s *BaseServerScriptParserListener) EnterTemplateStringLiteral(ctx *TemplateStringLiteralContext) {
}

// ExitTemplateStringLiteral is called when production templateStringLiteral is exited.
func (s *BaseServerScriptParserListener) ExitTemplateStringLiteral(ctx *TemplateStringLiteralContext) {
}

// EnterTemplateStringAtom is called when production templateStringAtom is entered.
func (s *BaseServerScriptParserListener) EnterTemplateStringAtom(ctx *TemplateStringAtomContext) {}

// ExitTemplateStringAtom is called when production templateStringAtom is exited.
func (s *BaseServerScriptParserListener) ExitTemplateStringAtom(ctx *TemplateStringAtomContext) {}

// EnterNumericLiteral is called when production numericLiteral is entered.
func (s *BaseServerScriptParserListener) EnterNumericLiteral(ctx *NumericLiteralContext) {}

// ExitNumericLiteral is called when production numericLiteral is exited.
func (s *BaseServerScriptParserListener) ExitNumericLiteral(ctx *NumericLiteralContext) {}

// EnterBigintLiteral is called when production bigintLiteral is entered.
func (s *BaseServerScriptParserListener) EnterBigintLiteral(ctx *BigintLiteralContext) {}

// ExitBigintLiteral is called when production bigintLiteral is exited.
func (s *BaseServerScriptParserListener) ExitBigintLiteral(ctx *BigintLiteralContext) {}

// EnterIdentifierName is called when production identifierName is entered.
func (s *BaseServerScriptParserListener) EnterIdentifierName(ctx *IdentifierNameContext) {}

// ExitIdentifierName is called when production identifierName is exited.
func (s *BaseServerScriptParserListener) ExitIdentifierName(ctx *IdentifierNameContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BaseServerScriptParserListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BaseServerScriptParserListener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterReservedWord is called when production reservedWord is entered.
func (s *BaseServerScriptParserListener) EnterReservedWord(ctx *ReservedWordContext) {}

// ExitReservedWord is called when production reservedWord is exited.
func (s *BaseServerScriptParserListener) ExitReservedWord(ctx *ReservedWordContext) {}

// EnterKeyword is called when production keyword is entered.
func (s *BaseServerScriptParserListener) EnterKeyword(ctx *KeywordContext) {}

// ExitKeyword is called when production keyword is exited.
func (s *BaseServerScriptParserListener) ExitKeyword(ctx *KeywordContext) {}

// EnterEos is called when production eos is entered.
func (s *BaseServerScriptParserListener) EnterEos(ctx *EosContext) {}

// ExitEos is called when production eos is exited.
func (s *BaseServerScriptParserListener) ExitEos(ctx *EosContext) {}
