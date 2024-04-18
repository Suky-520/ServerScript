// Code generated from /Users/lzp/代码/workspace_go/ss/parser/ServerScriptParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // ServerScriptParser

import "github.com/antlr4-go/antlr/v4"

// ServerScriptParserListener is a complete listener for a parse tree produced by ServerScriptParser.
type ServerScriptParserListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterSourceElements is called when entering the sourceElements production.
	EnterSourceElements(c *SourceElementsContext)

	// EnterSourceElement is called when entering the sourceElement production.
	EnterSourceElement(c *SourceElementContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterStatementList is called when entering the statementList production.
	EnterStatementList(c *StatementListContext)

	// EnterImportStatement is called when entering the importStatement production.
	EnterImportStatement(c *ImportStatementContext)

	// EnterImportFromBlock1 is called when entering the ImportFromBlock1 production.
	EnterImportFromBlock1(c *ImportFromBlock1Context)

	// EnterImportFromBlock2 is called when entering the ImportFromBlock2 production.
	EnterImportFromBlock2(c *ImportFromBlock2Context)

	// EnterImportFromBlock3 is called when entering the ImportFromBlock3 production.
	EnterImportFromBlock3(c *ImportFromBlock3Context)

	// EnterImportModuleItems is called when entering the importModuleItems production.
	EnterImportModuleItems(c *ImportModuleItemsContext)

	// EnterImportAliasName is called when entering the importAliasName production.
	EnterImportAliasName(c *ImportAliasNameContext)

	// EnterModuleExportName is called when entering the moduleExportName production.
	EnterModuleExportName(c *ModuleExportNameContext)

	// EnterImportedBinding is called when entering the importedBinding production.
	EnterImportedBinding(c *ImportedBindingContext)

	// EnterImportDefault is called when entering the importDefault production.
	EnterImportDefault(c *ImportDefaultContext)

	// EnterImportNamespace is called when entering the importNamespace production.
	EnterImportNamespace(c *ImportNamespaceContext)

	// EnterImportFrom is called when entering the importFrom production.
	EnterImportFrom(c *ImportFromContext)

	// EnterAliasName is called when entering the aliasName production.
	EnterAliasName(c *AliasNameContext)

	// EnterExportDeclaration is called when entering the ExportDeclaration production.
	EnterExportDeclaration(c *ExportDeclarationContext)

	// EnterExportItems is called when entering the ExportItems production.
	EnterExportItems(c *ExportItemsContext)

	// EnterExportBlock is called when entering the ExportBlock production.
	EnterExportBlock(c *ExportBlockContext)

	// EnterExportDefaultDeclaration is called when entering the ExportDefaultDeclaration production.
	EnterExportDefaultDeclaration(c *ExportDefaultDeclarationContext)

	// EnterExportNamespace is called when entering the exportNamespace production.
	EnterExportNamespace(c *ExportNamespaceContext)

	// EnterExportFromBlock is called when entering the exportFromBlock production.
	EnterExportFromBlock(c *ExportFromBlockContext)

	// EnterExportModuleItems is called when entering the exportModuleItems production.
	EnterExportModuleItems(c *ExportModuleItemsContext)

	// EnterExportAliasName is called when entering the exportAliasName production.
	EnterExportAliasName(c *ExportAliasNameContext)

	// EnterDeclaration is called when entering the declaration production.
	EnterDeclaration(c *DeclarationContext)

	// EnterVariableStatement is called when entering the variableStatement production.
	EnterVariableStatement(c *VariableStatementContext)

	// EnterVariableDeclarationList is called when entering the variableDeclarationList production.
	EnterVariableDeclarationList(c *VariableDeclarationListContext)

	// EnterVariableDeclaration is called when entering the variableDeclaration production.
	EnterVariableDeclaration(c *VariableDeclarationContext)

	// EnterEmptyStatement_ is called when entering the emptyStatement_ production.
	EnterEmptyStatement_(c *EmptyStatement_Context)

	// EnterExpressionStatement is called when entering the expressionStatement production.
	EnterExpressionStatement(c *ExpressionStatementContext)

	// EnterIfStatement is called when entering the ifStatement production.
	EnterIfStatement(c *IfStatementContext)

	// EnterIfBlock is called when entering the ifBlock production.
	EnterIfBlock(c *IfBlockContext)

	// EnterElseIfBlock is called when entering the elseIfBlock production.
	EnterElseIfBlock(c *ElseIfBlockContext)

	// EnterElseBlock is called when entering the elseBlock production.
	EnterElseBlock(c *ElseBlockContext)

	// EnterDoStatement is called when entering the DoStatement production.
	EnterDoStatement(c *DoStatementContext)

	// EnterWhileStatement is called when entering the WhileStatement production.
	EnterWhileStatement(c *WhileStatementContext)

	// EnterForStatement is called when entering the ForStatement production.
	EnterForStatement(c *ForStatementContext)

	// EnterForOfStatement is called when entering the ForOfStatement production.
	EnterForOfStatement(c *ForOfStatementContext)

	// EnterOfArrayLiteral is called when entering the ofArrayLiteral production.
	EnterOfArrayLiteral(c *OfArrayLiteralContext)

	// EnterIdentifierList is called when entering the identifierList production.
	EnterIdentifierList(c *IdentifierListContext)

	// EnterForStatement1 is called when entering the forStatement1 production.
	EnterForStatement1(c *ForStatement1Context)

	// EnterForStatement2 is called when entering the forStatement2 production.
	EnterForStatement2(c *ForStatement2Context)

	// EnterForStatement3 is called when entering the forStatement3 production.
	EnterForStatement3(c *ForStatement3Context)

	// EnterVarModifier is called when entering the varModifier production.
	EnterVarModifier(c *VarModifierContext)

	// EnterContinueStatement is called when entering the continueStatement production.
	EnterContinueStatement(c *ContinueStatementContext)

	// EnterBreakStatement is called when entering the breakStatement production.
	EnterBreakStatement(c *BreakStatementContext)

	// EnterReturnStatement is called when entering the returnStatement production.
	EnterReturnStatement(c *ReturnStatementContext)

	// EnterSwitchStatement is called when entering the switchStatement production.
	EnterSwitchStatement(c *SwitchStatementContext)

	// EnterCaseBlock is called when entering the caseBlock production.
	EnterCaseBlock(c *CaseBlockContext)

	// EnterCaseClauses is called when entering the caseClauses production.
	EnterCaseClauses(c *CaseClausesContext)

	// EnterCaseClause is called when entering the caseClause production.
	EnterCaseClause(c *CaseClauseContext)

	// EnterDefaultClause is called when entering the defaultClause production.
	EnterDefaultClause(c *DefaultClauseContext)

	// EnterThrowStatement is called when entering the throwStatement production.
	EnterThrowStatement(c *ThrowStatementContext)

	// EnterTryStatement is called when entering the tryStatement production.
	EnterTryStatement(c *TryStatementContext)

	// EnterCatchProduction is called when entering the catchProduction production.
	EnterCatchProduction(c *CatchProductionContext)

	// EnterFinallyProduction is called when entering the finallyProduction production.
	EnterFinallyProduction(c *FinallyProductionContext)

	// EnterFunctionDeclaration is called when entering the functionDeclaration production.
	EnterFunctionDeclaration(c *FunctionDeclarationContext)

	// EnterFormalParameterList is called when entering the formalParameterList production.
	EnterFormalParameterList(c *FormalParameterListContext)

	// EnterFormalParameterArg is called when entering the formalParameterArg production.
	EnterFormalParameterArg(c *FormalParameterArgContext)

	// EnterLastFormalParameterArg is called when entering the lastFormalParameterArg production.
	EnterLastFormalParameterArg(c *LastFormalParameterArgContext)

	// EnterFunctionBody is called when entering the functionBody production.
	EnterFunctionBody(c *FunctionBodyContext)

	// EnterArrayLiteral is called when entering the arrayLiteral production.
	EnterArrayLiteral(c *ArrayLiteralContext)

	// EnterElementList is called when entering the elementList production.
	EnterElementList(c *ElementListContext)

	// EnterArrayElement is called when entering the arrayElement production.
	EnterArrayElement(c *ArrayElementContext)

	// EnterPropertyExpressionAssignment is called when entering the PropertyExpressionAssignment production.
	EnterPropertyExpressionAssignment(c *PropertyExpressionAssignmentContext)

	// EnterComputedPropertyExpressionAssignment is called when entering the ComputedPropertyExpressionAssignment production.
	EnterComputedPropertyExpressionAssignment(c *ComputedPropertyExpressionAssignmentContext)

	// EnterFunctionProperty is called when entering the FunctionProperty production.
	EnterFunctionProperty(c *FunctionPropertyContext)

	// EnterPropertyShorthand is called when entering the PropertyShorthand production.
	EnterPropertyShorthand(c *PropertyShorthandContext)

	// EnterPropertyName is called when entering the propertyName production.
	EnterPropertyName(c *PropertyNameContext)

	// EnterArguments is called when entering the arguments production.
	EnterArguments(c *ArgumentsContext)

	// EnterArgument is called when entering the argument production.
	EnterArgument(c *ArgumentContext)

	// EnterExpressionSequence is called when entering the expressionSequence production.
	EnterExpressionSequence(c *ExpressionSequenceContext)

	// EnterTemplateStringExpression is called when entering the TemplateStringExpression production.
	EnterTemplateStringExpression(c *TemplateStringExpressionContext)

	// EnterTernaryExpression is called when entering the TernaryExpression production.
	EnterTernaryExpression(c *TernaryExpressionContext)

	// EnterPowerExpression is called when entering the PowerExpression production.
	EnterPowerExpression(c *PowerExpressionContext)

	// EnterPreIncrementExpression is called when entering the PreIncrementExpression production.
	EnterPreIncrementExpression(c *PreIncrementExpressionContext)

	// EnterObjectLiteralExpression is called when entering the ObjectLiteralExpression production.
	EnterObjectLiteralExpression(c *ObjectLiteralExpressionContext)

	// EnterPrefixArgumentsExpression is called when entering the PrefixArgumentsExpression production.
	EnterPrefixArgumentsExpression(c *PrefixArgumentsExpressionContext)

	// EnterInExpression is called when entering the InExpression production.
	EnterInExpression(c *InExpressionContext)

	// EnterNotExpression is called when entering the NotExpression production.
	EnterNotExpression(c *NotExpressionContext)

	// EnterPreDecreaseExpression is called when entering the PreDecreaseExpression production.
	EnterPreDecreaseExpression(c *PreDecreaseExpressionContext)

	// EnterArgumentsExpression is called when entering the ArgumentsExpression production.
	EnterArgumentsExpression(c *ArgumentsExpressionContext)

	// EnterThisExpression is called when entering the ThisExpression production.
	EnterThisExpression(c *ThisExpressionContext)

	// EnterLogicalExpression is called when entering the LogicalExpression production.
	EnterLogicalExpression(c *LogicalExpressionContext)

	// EnterFunctionExpression is called when entering the FunctionExpression production.
	EnterFunctionExpression(c *FunctionExpressionContext)

	// EnterUnaryMinusExpression is called when entering the UnaryMinusExpression production.
	EnterUnaryMinusExpression(c *UnaryMinusExpressionContext)

	// EnterAssignmentExpression is called when entering the AssignmentExpression production.
	EnterAssignmentExpression(c *AssignmentExpressionContext)

	// EnterPostDecreaseExpression is called when entering the PostDecreaseExpression production.
	EnterPostDecreaseExpression(c *PostDecreaseExpressionContext)

	// EnterUnaryPlusExpression is called when entering the UnaryPlusExpression production.
	EnterUnaryPlusExpression(c *UnaryPlusExpressionContext)

	// EnterDeleteExpression is called when entering the DeleteExpression production.
	EnterDeleteExpression(c *DeleteExpressionContext)

	// EnterEqualityExpression is called when entering the EqualityExpression production.
	EnterEqualityExpression(c *EqualityExpressionContext)

	// EnterMultiplicativeExpression is called when entering the MultiplicativeExpression production.
	EnterMultiplicativeExpression(c *MultiplicativeExpressionContext)

	// EnterBitShiftExpression is called when entering the BitShiftExpression production.
	EnterBitShiftExpression(c *BitShiftExpressionContext)

	// EnterParenthesizedExpression is called when entering the ParenthesizedExpression production.
	EnterParenthesizedExpression(c *ParenthesizedExpressionContext)

	// EnterAdditiveExpression is called when entering the AdditiveExpression production.
	EnterAdditiveExpression(c *AdditiveExpressionContext)

	// EnterRelationalExpression is called when entering the RelationalExpression production.
	EnterRelationalExpression(c *RelationalExpressionContext)

	// EnterPostIncrementExpression is called when entering the PostIncrementExpression production.
	EnterPostIncrementExpression(c *PostIncrementExpressionContext)

	// EnterBitNotExpression is called when entering the BitNotExpression production.
	EnterBitNotExpression(c *BitNotExpressionContext)

	// EnterNewExpression is called when entering the NewExpression production.
	EnterNewExpression(c *NewExpressionContext)

	// EnterLiteralExpression is called when entering the LiteralExpression production.
	EnterLiteralExpression(c *LiteralExpressionContext)

	// EnterArrayLiteralExpression is called when entering the ArrayLiteralExpression production.
	EnterArrayLiteralExpression(c *ArrayLiteralExpressionContext)

	// EnterMemberDotExpression is called when entering the MemberDotExpression production.
	EnterMemberDotExpression(c *MemberDotExpressionContext)

	// EnterMemberIndexExpression is called when entering the MemberIndexExpression production.
	EnterMemberIndexExpression(c *MemberIndexExpressionContext)

	// EnterIdentifierExpression is called when entering the IdentifierExpression production.
	EnterIdentifierExpression(c *IdentifierExpressionContext)

	// EnterBitExpression is called when entering the BitExpression production.
	EnterBitExpression(c *BitExpressionContext)

	// EnterAssignmentOperatorExpression is called when entering the AssignmentOperatorExpression production.
	EnterAssignmentOperatorExpression(c *AssignmentOperatorExpressionContext)

	// EnterCoalesceExpression is called when entering the CoalesceExpression production.
	EnterCoalesceExpression(c *CoalesceExpressionContext)

	// EnterAssignable is called when entering the assignable production.
	EnterAssignable(c *AssignableContext)

	// EnterObjectLiteral is called when entering the objectLiteral production.
	EnterObjectLiteral(c *ObjectLiteralContext)

	// EnterAnonymousFunctionDecl is called when entering the AnonymousFunctionDecl production.
	EnterAnonymousFunctionDecl(c *AnonymousFunctionDeclContext)

	// EnterArrowFunction is called when entering the ArrowFunction production.
	EnterArrowFunction(c *ArrowFunctionContext)

	// EnterArrowSingleFunction is called when entering the ArrowSingleFunction production.
	EnterArrowSingleFunction(c *ArrowSingleFunctionContext)

	// EnterArrowFunctionParameters is called when entering the arrowFunctionParameters production.
	EnterArrowFunctionParameters(c *ArrowFunctionParametersContext)

	// EnterAssignmentOperator is called when entering the assignmentOperator production.
	EnterAssignmentOperator(c *AssignmentOperatorContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterTemplateStringLiteral is called when entering the templateStringLiteral production.
	EnterTemplateStringLiteral(c *TemplateStringLiteralContext)

	// EnterTemplateStringAtom is called when entering the templateStringAtom production.
	EnterTemplateStringAtom(c *TemplateStringAtomContext)

	// EnterNumericLiteral is called when entering the numericLiteral production.
	EnterNumericLiteral(c *NumericLiteralContext)

	// EnterBigintLiteral is called when entering the bigintLiteral production.
	EnterBigintLiteral(c *BigintLiteralContext)

	// EnterIdentifierName is called when entering the identifierName production.
	EnterIdentifierName(c *IdentifierNameContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// EnterReservedWord is called when entering the reservedWord production.
	EnterReservedWord(c *ReservedWordContext)

	// EnterKeyword is called when entering the keyword production.
	EnterKeyword(c *KeywordContext)

	// EnterEos is called when entering the eos production.
	EnterEos(c *EosContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitSourceElements is called when exiting the sourceElements production.
	ExitSourceElements(c *SourceElementsContext)

	// ExitSourceElement is called when exiting the sourceElement production.
	ExitSourceElement(c *SourceElementContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitStatementList is called when exiting the statementList production.
	ExitStatementList(c *StatementListContext)

	// ExitImportStatement is called when exiting the importStatement production.
	ExitImportStatement(c *ImportStatementContext)

	// ExitImportFromBlock1 is called when exiting the ImportFromBlock1 production.
	ExitImportFromBlock1(c *ImportFromBlock1Context)

	// ExitImportFromBlock2 is called when exiting the ImportFromBlock2 production.
	ExitImportFromBlock2(c *ImportFromBlock2Context)

	// ExitImportFromBlock3 is called when exiting the ImportFromBlock3 production.
	ExitImportFromBlock3(c *ImportFromBlock3Context)

	// ExitImportModuleItems is called when exiting the importModuleItems production.
	ExitImportModuleItems(c *ImportModuleItemsContext)

	// ExitImportAliasName is called when exiting the importAliasName production.
	ExitImportAliasName(c *ImportAliasNameContext)

	// ExitModuleExportName is called when exiting the moduleExportName production.
	ExitModuleExportName(c *ModuleExportNameContext)

	// ExitImportedBinding is called when exiting the importedBinding production.
	ExitImportedBinding(c *ImportedBindingContext)

	// ExitImportDefault is called when exiting the importDefault production.
	ExitImportDefault(c *ImportDefaultContext)

	// ExitImportNamespace is called when exiting the importNamespace production.
	ExitImportNamespace(c *ImportNamespaceContext)

	// ExitImportFrom is called when exiting the importFrom production.
	ExitImportFrom(c *ImportFromContext)

	// ExitAliasName is called when exiting the aliasName production.
	ExitAliasName(c *AliasNameContext)

	// ExitExportDeclaration is called when exiting the ExportDeclaration production.
	ExitExportDeclaration(c *ExportDeclarationContext)

	// ExitExportItems is called when exiting the ExportItems production.
	ExitExportItems(c *ExportItemsContext)

	// ExitExportBlock is called when exiting the ExportBlock production.
	ExitExportBlock(c *ExportBlockContext)

	// ExitExportDefaultDeclaration is called when exiting the ExportDefaultDeclaration production.
	ExitExportDefaultDeclaration(c *ExportDefaultDeclarationContext)

	// ExitExportNamespace is called when exiting the exportNamespace production.
	ExitExportNamespace(c *ExportNamespaceContext)

	// ExitExportFromBlock is called when exiting the exportFromBlock production.
	ExitExportFromBlock(c *ExportFromBlockContext)

	// ExitExportModuleItems is called when exiting the exportModuleItems production.
	ExitExportModuleItems(c *ExportModuleItemsContext)

	// ExitExportAliasName is called when exiting the exportAliasName production.
	ExitExportAliasName(c *ExportAliasNameContext)

	// ExitDeclaration is called when exiting the declaration production.
	ExitDeclaration(c *DeclarationContext)

	// ExitVariableStatement is called when exiting the variableStatement production.
	ExitVariableStatement(c *VariableStatementContext)

	// ExitVariableDeclarationList is called when exiting the variableDeclarationList production.
	ExitVariableDeclarationList(c *VariableDeclarationListContext)

	// ExitVariableDeclaration is called when exiting the variableDeclaration production.
	ExitVariableDeclaration(c *VariableDeclarationContext)

	// ExitEmptyStatement_ is called when exiting the emptyStatement_ production.
	ExitEmptyStatement_(c *EmptyStatement_Context)

	// ExitExpressionStatement is called when exiting the expressionStatement production.
	ExitExpressionStatement(c *ExpressionStatementContext)

	// ExitIfStatement is called when exiting the ifStatement production.
	ExitIfStatement(c *IfStatementContext)

	// ExitIfBlock is called when exiting the ifBlock production.
	ExitIfBlock(c *IfBlockContext)

	// ExitElseIfBlock is called when exiting the elseIfBlock production.
	ExitElseIfBlock(c *ElseIfBlockContext)

	// ExitElseBlock is called when exiting the elseBlock production.
	ExitElseBlock(c *ElseBlockContext)

	// ExitDoStatement is called when exiting the DoStatement production.
	ExitDoStatement(c *DoStatementContext)

	// ExitWhileStatement is called when exiting the WhileStatement production.
	ExitWhileStatement(c *WhileStatementContext)

	// ExitForStatement is called when exiting the ForStatement production.
	ExitForStatement(c *ForStatementContext)

	// ExitForOfStatement is called when exiting the ForOfStatement production.
	ExitForOfStatement(c *ForOfStatementContext)

	// ExitOfArrayLiteral is called when exiting the ofArrayLiteral production.
	ExitOfArrayLiteral(c *OfArrayLiteralContext)

	// ExitIdentifierList is called when exiting the identifierList production.
	ExitIdentifierList(c *IdentifierListContext)

	// ExitForStatement1 is called when exiting the forStatement1 production.
	ExitForStatement1(c *ForStatement1Context)

	// ExitForStatement2 is called when exiting the forStatement2 production.
	ExitForStatement2(c *ForStatement2Context)

	// ExitForStatement3 is called when exiting the forStatement3 production.
	ExitForStatement3(c *ForStatement3Context)

	// ExitVarModifier is called when exiting the varModifier production.
	ExitVarModifier(c *VarModifierContext)

	// ExitContinueStatement is called when exiting the continueStatement production.
	ExitContinueStatement(c *ContinueStatementContext)

	// ExitBreakStatement is called when exiting the breakStatement production.
	ExitBreakStatement(c *BreakStatementContext)

	// ExitReturnStatement is called when exiting the returnStatement production.
	ExitReturnStatement(c *ReturnStatementContext)

	// ExitSwitchStatement is called when exiting the switchStatement production.
	ExitSwitchStatement(c *SwitchStatementContext)

	// ExitCaseBlock is called when exiting the caseBlock production.
	ExitCaseBlock(c *CaseBlockContext)

	// ExitCaseClauses is called when exiting the caseClauses production.
	ExitCaseClauses(c *CaseClausesContext)

	// ExitCaseClause is called when exiting the caseClause production.
	ExitCaseClause(c *CaseClauseContext)

	// ExitDefaultClause is called when exiting the defaultClause production.
	ExitDefaultClause(c *DefaultClauseContext)

	// ExitThrowStatement is called when exiting the throwStatement production.
	ExitThrowStatement(c *ThrowStatementContext)

	// ExitTryStatement is called when exiting the tryStatement production.
	ExitTryStatement(c *TryStatementContext)

	// ExitCatchProduction is called when exiting the catchProduction production.
	ExitCatchProduction(c *CatchProductionContext)

	// ExitFinallyProduction is called when exiting the finallyProduction production.
	ExitFinallyProduction(c *FinallyProductionContext)

	// ExitFunctionDeclaration is called when exiting the functionDeclaration production.
	ExitFunctionDeclaration(c *FunctionDeclarationContext)

	// ExitFormalParameterList is called when exiting the formalParameterList production.
	ExitFormalParameterList(c *FormalParameterListContext)

	// ExitFormalParameterArg is called when exiting the formalParameterArg production.
	ExitFormalParameterArg(c *FormalParameterArgContext)

	// ExitLastFormalParameterArg is called when exiting the lastFormalParameterArg production.
	ExitLastFormalParameterArg(c *LastFormalParameterArgContext)

	// ExitFunctionBody is called when exiting the functionBody production.
	ExitFunctionBody(c *FunctionBodyContext)

	// ExitArrayLiteral is called when exiting the arrayLiteral production.
	ExitArrayLiteral(c *ArrayLiteralContext)

	// ExitElementList is called when exiting the elementList production.
	ExitElementList(c *ElementListContext)

	// ExitArrayElement is called when exiting the arrayElement production.
	ExitArrayElement(c *ArrayElementContext)

	// ExitPropertyExpressionAssignment is called when exiting the PropertyExpressionAssignment production.
	ExitPropertyExpressionAssignment(c *PropertyExpressionAssignmentContext)

	// ExitComputedPropertyExpressionAssignment is called when exiting the ComputedPropertyExpressionAssignment production.
	ExitComputedPropertyExpressionAssignment(c *ComputedPropertyExpressionAssignmentContext)

	// ExitFunctionProperty is called when exiting the FunctionProperty production.
	ExitFunctionProperty(c *FunctionPropertyContext)

	// ExitPropertyShorthand is called when exiting the PropertyShorthand production.
	ExitPropertyShorthand(c *PropertyShorthandContext)

	// ExitPropertyName is called when exiting the propertyName production.
	ExitPropertyName(c *PropertyNameContext)

	// ExitArguments is called when exiting the arguments production.
	ExitArguments(c *ArgumentsContext)

	// ExitArgument is called when exiting the argument production.
	ExitArgument(c *ArgumentContext)

	// ExitExpressionSequence is called when exiting the expressionSequence production.
	ExitExpressionSequence(c *ExpressionSequenceContext)

	// ExitTemplateStringExpression is called when exiting the TemplateStringExpression production.
	ExitTemplateStringExpression(c *TemplateStringExpressionContext)

	// ExitTernaryExpression is called when exiting the TernaryExpression production.
	ExitTernaryExpression(c *TernaryExpressionContext)

	// ExitPowerExpression is called when exiting the PowerExpression production.
	ExitPowerExpression(c *PowerExpressionContext)

	// ExitPreIncrementExpression is called when exiting the PreIncrementExpression production.
	ExitPreIncrementExpression(c *PreIncrementExpressionContext)

	// ExitObjectLiteralExpression is called when exiting the ObjectLiteralExpression production.
	ExitObjectLiteralExpression(c *ObjectLiteralExpressionContext)

	// ExitPrefixArgumentsExpression is called when exiting the PrefixArgumentsExpression production.
	ExitPrefixArgumentsExpression(c *PrefixArgumentsExpressionContext)

	// ExitInExpression is called when exiting the InExpression production.
	ExitInExpression(c *InExpressionContext)

	// ExitNotExpression is called when exiting the NotExpression production.
	ExitNotExpression(c *NotExpressionContext)

	// ExitPreDecreaseExpression is called when exiting the PreDecreaseExpression production.
	ExitPreDecreaseExpression(c *PreDecreaseExpressionContext)

	// ExitArgumentsExpression is called when exiting the ArgumentsExpression production.
	ExitArgumentsExpression(c *ArgumentsExpressionContext)

	// ExitThisExpression is called when exiting the ThisExpression production.
	ExitThisExpression(c *ThisExpressionContext)

	// ExitLogicalExpression is called when exiting the LogicalExpression production.
	ExitLogicalExpression(c *LogicalExpressionContext)

	// ExitFunctionExpression is called when exiting the FunctionExpression production.
	ExitFunctionExpression(c *FunctionExpressionContext)

	// ExitUnaryMinusExpression is called when exiting the UnaryMinusExpression production.
	ExitUnaryMinusExpression(c *UnaryMinusExpressionContext)

	// ExitAssignmentExpression is called when exiting the AssignmentExpression production.
	ExitAssignmentExpression(c *AssignmentExpressionContext)

	// ExitPostDecreaseExpression is called when exiting the PostDecreaseExpression production.
	ExitPostDecreaseExpression(c *PostDecreaseExpressionContext)

	// ExitUnaryPlusExpression is called when exiting the UnaryPlusExpression production.
	ExitUnaryPlusExpression(c *UnaryPlusExpressionContext)

	// ExitDeleteExpression is called when exiting the DeleteExpression production.
	ExitDeleteExpression(c *DeleteExpressionContext)

	// ExitEqualityExpression is called when exiting the EqualityExpression production.
	ExitEqualityExpression(c *EqualityExpressionContext)

	// ExitMultiplicativeExpression is called when exiting the MultiplicativeExpression production.
	ExitMultiplicativeExpression(c *MultiplicativeExpressionContext)

	// ExitBitShiftExpression is called when exiting the BitShiftExpression production.
	ExitBitShiftExpression(c *BitShiftExpressionContext)

	// ExitParenthesizedExpression is called when exiting the ParenthesizedExpression production.
	ExitParenthesizedExpression(c *ParenthesizedExpressionContext)

	// ExitAdditiveExpression is called when exiting the AdditiveExpression production.
	ExitAdditiveExpression(c *AdditiveExpressionContext)

	// ExitRelationalExpression is called when exiting the RelationalExpression production.
	ExitRelationalExpression(c *RelationalExpressionContext)

	// ExitPostIncrementExpression is called when exiting the PostIncrementExpression production.
	ExitPostIncrementExpression(c *PostIncrementExpressionContext)

	// ExitBitNotExpression is called when exiting the BitNotExpression production.
	ExitBitNotExpression(c *BitNotExpressionContext)

	// ExitNewExpression is called when exiting the NewExpression production.
	ExitNewExpression(c *NewExpressionContext)

	// ExitLiteralExpression is called when exiting the LiteralExpression production.
	ExitLiteralExpression(c *LiteralExpressionContext)

	// ExitArrayLiteralExpression is called when exiting the ArrayLiteralExpression production.
	ExitArrayLiteralExpression(c *ArrayLiteralExpressionContext)

	// ExitMemberDotExpression is called when exiting the MemberDotExpression production.
	ExitMemberDotExpression(c *MemberDotExpressionContext)

	// ExitMemberIndexExpression is called when exiting the MemberIndexExpression production.
	ExitMemberIndexExpression(c *MemberIndexExpressionContext)

	// ExitIdentifierExpression is called when exiting the IdentifierExpression production.
	ExitIdentifierExpression(c *IdentifierExpressionContext)

	// ExitBitExpression is called when exiting the BitExpression production.
	ExitBitExpression(c *BitExpressionContext)

	// ExitAssignmentOperatorExpression is called when exiting the AssignmentOperatorExpression production.
	ExitAssignmentOperatorExpression(c *AssignmentOperatorExpressionContext)

	// ExitCoalesceExpression is called when exiting the CoalesceExpression production.
	ExitCoalesceExpression(c *CoalesceExpressionContext)

	// ExitAssignable is called when exiting the assignable production.
	ExitAssignable(c *AssignableContext)

	// ExitObjectLiteral is called when exiting the objectLiteral production.
	ExitObjectLiteral(c *ObjectLiteralContext)

	// ExitAnonymousFunctionDecl is called when exiting the AnonymousFunctionDecl production.
	ExitAnonymousFunctionDecl(c *AnonymousFunctionDeclContext)

	// ExitArrowFunction is called when exiting the ArrowFunction production.
	ExitArrowFunction(c *ArrowFunctionContext)

	// ExitArrowSingleFunction is called when exiting the ArrowSingleFunction production.
	ExitArrowSingleFunction(c *ArrowSingleFunctionContext)

	// ExitArrowFunctionParameters is called when exiting the arrowFunctionParameters production.
	ExitArrowFunctionParameters(c *ArrowFunctionParametersContext)

	// ExitAssignmentOperator is called when exiting the assignmentOperator production.
	ExitAssignmentOperator(c *AssignmentOperatorContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitTemplateStringLiteral is called when exiting the templateStringLiteral production.
	ExitTemplateStringLiteral(c *TemplateStringLiteralContext)

	// ExitTemplateStringAtom is called when exiting the templateStringAtom production.
	ExitTemplateStringAtom(c *TemplateStringAtomContext)

	// ExitNumericLiteral is called when exiting the numericLiteral production.
	ExitNumericLiteral(c *NumericLiteralContext)

	// ExitBigintLiteral is called when exiting the bigintLiteral production.
	ExitBigintLiteral(c *BigintLiteralContext)

	// ExitIdentifierName is called when exiting the identifierName production.
	ExitIdentifierName(c *IdentifierNameContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)

	// ExitReservedWord is called when exiting the reservedWord production.
	ExitReservedWord(c *ReservedWordContext)

	// ExitKeyword is called when exiting the keyword production.
	ExitKeyword(c *KeywordContext)

	// ExitEos is called when exiting the eos production.
	ExitEos(c *EosContext)
}
