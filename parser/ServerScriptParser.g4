parser grammar ServerScriptParser;

// 1. ? 匹配某个规则之前，该规则可以出现的次数为0或者1次。如果匹配成功，则继续匹配之后的规则；如果匹配失败，则跳过该规则
// 2. + 匹配某个规则之前，该规则可以出现的次数为1次或者多次。
// 3. * 匹配某个规则之前，该规则可以出现的次数为0次或者多次。

options {
    tokenVocab=ServerScriptLexer;
    superClass=ServerScriptParserBase;
}

program
    : HashBangLine? sourceElements? EOF
    ;

sourceElements
    : sourceElement+
    ;
    
sourceElement
    : statement
    ;

statement
    : block
    | variableStatement //变量声明语句
    | importStatement
    | exportStatement
    | emptyStatement_
    | functionDeclaration
    | expressionStatement
    | ifStatement
    | iterationStatement
    | continueStatement
    | breakStatement
    | returnStatement
    | switchStatement
    | throwStatement
    | tryStatement
    ;

block
    : '{' statementList? '}'
    ;

statementList
    : statement+
    ;

importStatement
    : Import importFromBlock
    ;

importFromBlock
    : importDefault? importNamespace  importFrom eos                         #ImportFromBlock1
    | importDefault? importModuleItems importFrom eos                        #ImportFromBlock2
    | identifierName importFrom eos                                          #ImportFromBlock3
    ;

importModuleItems
    : '{' (importAliasName ',')* (importAliasName ','?)? '}'
    ;

importAliasName
    : moduleExportName (As importedBinding)?
    ;

moduleExportName
    : identifierName
    | StringLiteral
    ;

importedBinding
    : Identifier
    | Yield
    | Await
    ;

importDefault
    : aliasName ','
    ;

importNamespace
    : '*' As identifierName
    ;

importFrom
    : From StringLiteral
    ;

aliasName
    : identifierName (As identifierName)?
    ;

exportStatement
    : Export  declaration   eos                         # ExportDeclaration
    | Export  exportModuleItems   eos                   # ExportItems
    | Export  exportFromBlock   eos                     # ExportBlock
    | Export Default singleExpression eos               # ExportDefaultDeclaration
    ;

exportNamespace
    :identifierName (As identifierName)?
    ;

exportFromBlock
    : exportNamespace importFrom eos
    ;

exportModuleItems
    : '{' (exportAliasName ',')* (exportAliasName ','?)? '}'
    ;

exportAliasName
    : moduleExportName (As moduleExportName)?
    ;

declaration
    : variableStatement
    | functionDeclaration
    ;

variableStatement
    : variableDeclarationList eos
    ;

variableDeclarationList
    : varModifier variableDeclaration (',' variableDeclaration)*
    ;

variableDeclaration
    : assignable ('=' singleExpression)?
    ;

emptyStatement_
    : SemiColon
    ;

expressionStatement
    : {p.notOpenBraceAndNotFunction()}? expressionSequence eos
    ;

ifStatement
    :  ifBlock elseIfBlock* elseBlock?
    ;
ifBlock
    : If '(' expressionSequence ')' statement
    ;
elseIfBlock
    : Else If '(' expressionSequence ')' statement
    ;
elseBlock
    : Else statement
    ;

iterationStatement
    : Do statement While '(' expressionSequence ')' eos                                                              # DoStatement
    | While '(' expressionSequence ')' statement                                                                     # WhileStatement
    | For '(' forStatement1? ';' forStatement2? ';' forStatement3? ')' statement                                     # ForStatement
    //| For '(' varModifier identifier In expressionSequence ')' statement                                             # ForInStatement
    | For '(' varModifier (identifier|ofArrayLiteral) Of expressionSequence ')' statement                            # ForOfStatement
    ;

ofArrayLiteral
    : ('[' identifierList ']')
    ;

identifierList
    : ','* identifier? (','+ identifier)* ','*
    ;


forStatement1
    :expressionSequence
    |variableDeclarationList
    ;
forStatement2
    :expressionSequence
    ;
forStatement3
    :expressionSequence
    ;

varModifier
    : Let
    | Const
    ;

continueStatement
    : Continue ({p.notLineTerminator()}? identifier)? eos
    ;

breakStatement
    : Break ({p.notLineTerminator()}? identifier)? eos
    ;

returnStatement
    : Return ({p.notLineTerminator()}? expressionSequence)? eos
    ;

switchStatement
    : Switch '(' expressionSequence ')' caseBlock
    ;

caseBlock
    : '{' caseClauses? (defaultClause caseClauses?)? '}'
    ;

caseClauses
    : caseClause+
    ;

caseClause
    : Case expressionSequence ':' statementList?
    ;

defaultClause
    : Default ':' statementList?
    ;

throwStatement
    : Throw {p.notLineTerminator()}? expressionSequence eos
    ;

tryStatement
    : Try block (catchProduction finallyProduction? | finallyProduction)
    ;

catchProduction
    : Catch ('(' assignable? ')')? block
    ;

finallyProduction
    : Finally block
    ;

functionDeclaration
    : Function_ identifier '(' formalParameterList? ')' functionBody
    ;

formalParameterList
    : formalParameterArg (',' formalParameterArg)* (',' lastFormalParameterArg)?
    | lastFormalParameterArg
    ;

formalParameterArg
    : assignable ('=' singleExpression)?
    ;

lastFormalParameterArg
    : Ellipsis singleExpression
    ;

functionBody
    : '{' sourceElements? '}'
    ;



arrayLiteral
    : ('[' elementList ']')
    ;

elementList
    : ','* arrayElement? (','+ arrayElement)* ','*
    ;

arrayElement
    : Ellipsis? singleExpression
    ;

propertyAssignment
    : propertyName ':' singleExpression                                             # PropertyExpressionAssignment
    | '[' singleExpression ']' ':' singleExpression                                 # ComputedPropertyExpressionAssignment
    | '*'? propertyName '(' formalParameterList?  ')'  functionBody                 # FunctionProperty
    | Ellipsis? singleExpression                                                    # PropertyShorthand
    ;

propertyName
    : identifierName
    | StringLiteral
    | numericLiteral
    | '[' singleExpression ']'
    ;

arguments
    : '('(argument (',' argument)* ','?)?')'
    ;

argument
    : Ellipsis? (singleExpression | identifier)
    ;

expressionSequence
    : singleExpression (',' singleExpression)*
    ;

singleExpression
    : anonymousFunction                                                     # FunctionExpression
    | singleExpression '?.'? '[' expressionSequence ']'                     # MemberIndexExpression
    | singleExpression ('.' | '?.') identifierName                          # MemberDotExpression
    | singleExpression '?.'? arguments                                      # ArgumentsExpression
    | (Sync | Async) singleExpression '?.'? arguments                       # PrefixArgumentsExpression //带有前缀的函数调用
    | New identifier arguments                                              # NewExpression
    | singleExpression {p.notLineTerminator()}? '++'                        # PostIncrementExpression
    | singleExpression {p.notLineTerminator()}? '--'                        # PostDecreaseExpression
    | Delete singleExpression                                               # DeleteExpression
    | '++' singleExpression                                                 # PreIncrementExpression
    | '--' singleExpression                                                 # PreDecreaseExpression
    | '+' singleExpression                                                  # UnaryPlusExpression
    | '-' singleExpression                                                  # UnaryMinusExpression
    | '~' singleExpression                                                  # BitNotExpression
    | '!' singleExpression                                                  # NotExpression
    | <assoc=right> singleExpression '**' singleExpression                  # PowerExpression
    | singleExpression ('*' | '/' | '%') singleExpression                   # MultiplicativeExpression
    | singleExpression ('+' | '-') singleExpression                         # AdditiveExpression
    | singleExpression '??' singleExpression                                # CoalesceExpression
    | singleExpression ('<<' | '>>' | '>>>') singleExpression               # BitShiftExpression
    | singleExpression ('<' | '>' | '<=' | '>=') singleExpression           # RelationalExpression
    | singleExpression In singleExpression                                  # InExpression
    | singleExpression ('==' | '!=' | '===' | '!==') singleExpression       # EqualityExpression
    | singleExpression ('&' | '^' |'|') singleExpression                    # BitExpression
    | singleExpression ('&&' |'||' )  singleExpression                      # LogicalExpression
    | singleExpression '?' singleExpression ':' singleExpression            # TernaryExpression
    | <assoc=right> singleExpression '=' singleExpression                   # AssignmentExpression
    | <assoc=right> singleExpression assignmentOperator singleExpression    # AssignmentOperatorExpression
    | singleExpression templateStringLiteral                                # TemplateStringExpression
    | This                                                                  # ThisExpression
    | identifier                                                            # IdentifierExpression
    | literal                                                               # LiteralExpression
    | arrayLiteral                                                          # ArrayLiteralExpression
    | objectLiteral                                                         # ObjectLiteralExpression
    | '(' expressionSequence ')'                                            # ParenthesizedExpression
    ;


assignable
    : identifier
    | arrayLiteral
    | objectLiteral
    ;

objectLiteral
    : '{' (propertyAssignment (',' propertyAssignment)* ','?)? '}'
    ;

anonymousFunction
    : Function_ '*'? '(' formalParameterList? ')' functionBody    # AnonymousFunctionDecl
    | arrowFunctionParameters '=>' functionBody                   # ArrowFunction
    | arrowFunctionParameters '=>' singleExpression               # ArrowSingleFunction
    ;

arrowFunctionParameters
    : identifier
    | '(' formalParameterList? ')'
    ;

assignmentOperator
    : '*='
    | '/='
    | '%='
    | '+='
    | '-='
    | '<<='
    | '>>='
    | '>>>='
    | '&='
    | '^='
    | '|='
    | '**='
    | '??='
    ;

literal
    : NullLiteral
    | BooleanLiteral
    | StringLiteral
    | UndefinedLiteral
    | NaNLiteral
    | InfinityLiteral
    | templateStringLiteral
    | bigintLiteral
    | RegularExpressionLiteral
    | numericLiteral
    ;

templateStringLiteral
    : BackTick templateStringAtom* BackTick
    ;

templateStringAtom
    : TemplateStringAtom
    | TemplateStringStartExpression singleExpression TemplateCloseBrace
    ;

numericLiteral
    : DecimalLiteral //数字
    | ExponentLiteral //指数表示法
    | HexIntegerLiteral //整数的16进制表示法
    | OctalIntegerLiteral //八进制表示法
    | BinaryIntegerLiteral //二进制表示法
    ;

bigintLiteral
    : BigDecimalIntegerLiteral
    | BigHexIntegerLiteral
    | BigOctalIntegerLiteral
    | BigBinaryIntegerLiteral
    ;

identifierName
    : identifier
    | reservedWord
    | 'NaN'
    ;

identifier
    : Identifier
    | Async
    | As
    | From
    | Yield
    | Of
    ;

reservedWord
    : keyword
    | NullLiteral
    | BooleanLiteral
    ;

keyword
    : Break
    | Do
    | Instanceof
    | Typeof
    | Case
    | Else
    | New
    | Var
    | Catch
    | Finally
    | Return
    | Void
    | Continue
    | For
    | Switch
    | While
    | Debugger
    | Function_
    | This
    | With
    | Default
    | If
    | Throw
    | Delete
    | In
    | Try

    | Class
    | Enum
    | Extends
    | Super
    | Const
    | Export
    | Import
    | Implements
    | Let
    | Private
    | Public
    | Interface
    | Package
    | Protected
    | Static
    | Yield
    | Async
    | Await
    | Sync
    | From
    | As
    | Of
    ;

eos
    : SemiColon
    | EOF
    | {p.lineTerminatorAhead()}?
    | {p.closeBrace()}?
    ;
