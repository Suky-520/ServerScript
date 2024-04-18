// Code generated from /Users/lzp/代码/workspace_go/ss/parser/ServerScriptParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // ServerScriptParser

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by ServerScriptParser.
type ServerScriptParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by ServerScriptParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#sourceElements.
	VisitSourceElements(ctx *SourceElementsContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#sourceElement.
	VisitSourceElement(ctx *SourceElementContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#statementList.
	VisitStatementList(ctx *StatementListContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#importStatement.
	VisitImportStatement(ctx *ImportStatementContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#ImportFromBlock1.
	VisitImportFromBlock1(ctx *ImportFromBlock1Context) interface{}

	// Visit a parse tree produced by ServerScriptParser#ImportFromBlock2.
	VisitImportFromBlock2(ctx *ImportFromBlock2Context) interface{}

	// Visit a parse tree produced by ServerScriptParser#ImportFromBlock3.
	VisitImportFromBlock3(ctx *ImportFromBlock3Context) interface{}

	// Visit a parse tree produced by ServerScriptParser#importModuleItems.
	VisitImportModuleItems(ctx *ImportModuleItemsContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#importAliasName.
	VisitImportAliasName(ctx *ImportAliasNameContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#moduleExportName.
	VisitModuleExportName(ctx *ModuleExportNameContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#importedBinding.
	VisitImportedBinding(ctx *ImportedBindingContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#importDefault.
	VisitImportDefault(ctx *ImportDefaultContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#importNamespace.
	VisitImportNamespace(ctx *ImportNamespaceContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#importFrom.
	VisitImportFrom(ctx *ImportFromContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#aliasName.
	VisitAliasName(ctx *AliasNameContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#ExportDeclaration.
	VisitExportDeclaration(ctx *ExportDeclarationContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#ExportItems.
	VisitExportItems(ctx *ExportItemsContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#ExportBlock.
	VisitExportBlock(ctx *ExportBlockContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#ExportDefaultDeclaration.
	VisitExportDefaultDeclaration(ctx *ExportDefaultDeclarationContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#exportNamespace.
	VisitExportNamespace(ctx *ExportNamespaceContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#exportFromBlock.
	VisitExportFromBlock(ctx *ExportFromBlockContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#exportModuleItems.
	VisitExportModuleItems(ctx *ExportModuleItemsContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#exportAliasName.
	VisitExportAliasName(ctx *ExportAliasNameContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#declaration.
	VisitDeclaration(ctx *DeclarationContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#variableStatement.
	VisitVariableStatement(ctx *VariableStatementContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#variableDeclarationList.
	VisitVariableDeclarationList(ctx *VariableDeclarationListContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#variableDeclaration.
	VisitVariableDeclaration(ctx *VariableDeclarationContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#emptyStatement_.
	VisitEmptyStatement_(ctx *EmptyStatement_Context) interface{}

	// Visit a parse tree produced by ServerScriptParser#expressionStatement.
	VisitExpressionStatement(ctx *ExpressionStatementContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#ifStatement.
	VisitIfStatement(ctx *IfStatementContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#ifBlock.
	VisitIfBlock(ctx *IfBlockContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#elseIfBlock.
	VisitElseIfBlock(ctx *ElseIfBlockContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#elseBlock.
	VisitElseBlock(ctx *ElseBlockContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#DoStatement.
	VisitDoStatement(ctx *DoStatementContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#WhileStatement.
	VisitWhileStatement(ctx *WhileStatementContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#ForStatement.
	VisitForStatement(ctx *ForStatementContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#ForOfStatement.
	VisitForOfStatement(ctx *ForOfStatementContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#ofArrayLiteral.
	VisitOfArrayLiteral(ctx *OfArrayLiteralContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#identifierList.
	VisitIdentifierList(ctx *IdentifierListContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#forStatement1.
	VisitForStatement1(ctx *ForStatement1Context) interface{}

	// Visit a parse tree produced by ServerScriptParser#forStatement2.
	VisitForStatement2(ctx *ForStatement2Context) interface{}

	// Visit a parse tree produced by ServerScriptParser#forStatement3.
	VisitForStatement3(ctx *ForStatement3Context) interface{}

	// Visit a parse tree produced by ServerScriptParser#varModifier.
	VisitVarModifier(ctx *VarModifierContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#continueStatement.
	VisitContinueStatement(ctx *ContinueStatementContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#breakStatement.
	VisitBreakStatement(ctx *BreakStatementContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#returnStatement.
	VisitReturnStatement(ctx *ReturnStatementContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#switchStatement.
	VisitSwitchStatement(ctx *SwitchStatementContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#caseBlock.
	VisitCaseBlock(ctx *CaseBlockContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#caseClauses.
	VisitCaseClauses(ctx *CaseClausesContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#caseClause.
	VisitCaseClause(ctx *CaseClauseContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#defaultClause.
	VisitDefaultClause(ctx *DefaultClauseContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#throwStatement.
	VisitThrowStatement(ctx *ThrowStatementContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#tryStatement.
	VisitTryStatement(ctx *TryStatementContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#catchProduction.
	VisitCatchProduction(ctx *CatchProductionContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#finallyProduction.
	VisitFinallyProduction(ctx *FinallyProductionContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#functionDeclaration.
	VisitFunctionDeclaration(ctx *FunctionDeclarationContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#formalParameterList.
	VisitFormalParameterList(ctx *FormalParameterListContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#formalParameterArg.
	VisitFormalParameterArg(ctx *FormalParameterArgContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#lastFormalParameterArg.
	VisitLastFormalParameterArg(ctx *LastFormalParameterArgContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#functionBody.
	VisitFunctionBody(ctx *FunctionBodyContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#arrayLiteral.
	VisitArrayLiteral(ctx *ArrayLiteralContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#elementList.
	VisitElementList(ctx *ElementListContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#arrayElement.
	VisitArrayElement(ctx *ArrayElementContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#PropertyExpressionAssignment.
	VisitPropertyExpressionAssignment(ctx *PropertyExpressionAssignmentContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#ComputedPropertyExpressionAssignment.
	VisitComputedPropertyExpressionAssignment(ctx *ComputedPropertyExpressionAssignmentContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#FunctionProperty.
	VisitFunctionProperty(ctx *FunctionPropertyContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#PropertyShorthand.
	VisitPropertyShorthand(ctx *PropertyShorthandContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#propertyName.
	VisitPropertyName(ctx *PropertyNameContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#arguments.
	VisitArguments(ctx *ArgumentsContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#argument.
	VisitArgument(ctx *ArgumentContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#expressionSequence.
	VisitExpressionSequence(ctx *ExpressionSequenceContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#TemplateStringExpression.
	VisitTemplateStringExpression(ctx *TemplateStringExpressionContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#TernaryExpression.
	VisitTernaryExpression(ctx *TernaryExpressionContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#PowerExpression.
	VisitPowerExpression(ctx *PowerExpressionContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#PreIncrementExpression.
	VisitPreIncrementExpression(ctx *PreIncrementExpressionContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#ObjectLiteralExpression.
	VisitObjectLiteralExpression(ctx *ObjectLiteralExpressionContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#PrefixArgumentsExpression.
	VisitPrefixArgumentsExpression(ctx *PrefixArgumentsExpressionContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#InExpression.
	VisitInExpression(ctx *InExpressionContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#NotExpression.
	VisitNotExpression(ctx *NotExpressionContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#PreDecreaseExpression.
	VisitPreDecreaseExpression(ctx *PreDecreaseExpressionContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#ArgumentsExpression.
	VisitArgumentsExpression(ctx *ArgumentsExpressionContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#ThisExpression.
	VisitThisExpression(ctx *ThisExpressionContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#LogicalExpression.
	VisitLogicalExpression(ctx *LogicalExpressionContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#FunctionExpression.
	VisitFunctionExpression(ctx *FunctionExpressionContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#UnaryMinusExpression.
	VisitUnaryMinusExpression(ctx *UnaryMinusExpressionContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#AssignmentExpression.
	VisitAssignmentExpression(ctx *AssignmentExpressionContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#PostDecreaseExpression.
	VisitPostDecreaseExpression(ctx *PostDecreaseExpressionContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#UnaryPlusExpression.
	VisitUnaryPlusExpression(ctx *UnaryPlusExpressionContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#DeleteExpression.
	VisitDeleteExpression(ctx *DeleteExpressionContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#EqualityExpression.
	VisitEqualityExpression(ctx *EqualityExpressionContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#MultiplicativeExpression.
	VisitMultiplicativeExpression(ctx *MultiplicativeExpressionContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#BitShiftExpression.
	VisitBitShiftExpression(ctx *BitShiftExpressionContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#ParenthesizedExpression.
	VisitParenthesizedExpression(ctx *ParenthesizedExpressionContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#AdditiveExpression.
	VisitAdditiveExpression(ctx *AdditiveExpressionContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#RelationalExpression.
	VisitRelationalExpression(ctx *RelationalExpressionContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#PostIncrementExpression.
	VisitPostIncrementExpression(ctx *PostIncrementExpressionContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#BitNotExpression.
	VisitBitNotExpression(ctx *BitNotExpressionContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#NewExpression.
	VisitNewExpression(ctx *NewExpressionContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#LiteralExpression.
	VisitLiteralExpression(ctx *LiteralExpressionContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#ArrayLiteralExpression.
	VisitArrayLiteralExpression(ctx *ArrayLiteralExpressionContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#MemberDotExpression.
	VisitMemberDotExpression(ctx *MemberDotExpressionContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#MemberIndexExpression.
	VisitMemberIndexExpression(ctx *MemberIndexExpressionContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#IdentifierExpression.
	VisitIdentifierExpression(ctx *IdentifierExpressionContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#BitExpression.
	VisitBitExpression(ctx *BitExpressionContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#AssignmentOperatorExpression.
	VisitAssignmentOperatorExpression(ctx *AssignmentOperatorExpressionContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#CoalesceExpression.
	VisitCoalesceExpression(ctx *CoalesceExpressionContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#assignable.
	VisitAssignable(ctx *AssignableContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#objectLiteral.
	VisitObjectLiteral(ctx *ObjectLiteralContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#AnonymousFunctionDecl.
	VisitAnonymousFunctionDecl(ctx *AnonymousFunctionDeclContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#ArrowFunction.
	VisitArrowFunction(ctx *ArrowFunctionContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#ArrowSingleFunction.
	VisitArrowSingleFunction(ctx *ArrowSingleFunctionContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#arrowFunctionParameters.
	VisitArrowFunctionParameters(ctx *ArrowFunctionParametersContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#assignmentOperator.
	VisitAssignmentOperator(ctx *AssignmentOperatorContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#templateStringLiteral.
	VisitTemplateStringLiteral(ctx *TemplateStringLiteralContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#templateStringAtom.
	VisitTemplateStringAtom(ctx *TemplateStringAtomContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#numericLiteral.
	VisitNumericLiteral(ctx *NumericLiteralContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#bigintLiteral.
	VisitBigintLiteral(ctx *BigintLiteralContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#identifierName.
	VisitIdentifierName(ctx *IdentifierNameContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#identifier.
	VisitIdentifier(ctx *IdentifierContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#reservedWord.
	VisitReservedWord(ctx *ReservedWordContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#keyword.
	VisitKeyword(ctx *KeywordContext) interface{}

	// Visit a parse tree produced by ServerScriptParser#eos.
	VisitEos(ctx *EosContext) interface{}
}
