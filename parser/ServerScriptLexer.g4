
lexer grammar ServerScriptLexer;

channels { ERROR }

options { superClass=ServerScriptLexerBase; }

HashBangLine:                   { p.IsStartOfFile()}? '#!' ~[\r\n\u2028\u2029]*;
MultiLineComment:               '/*' .*? '*/'             -> channel(HIDDEN);
SingleLineComment:              '//' ~[\r\n\u2028\u2029]* -> channel(HIDDEN);
RegularExpressionLiteral:       '/' RegularExpressionFirstChar RegularExpressionChar* {p.IsRegexPossible()}? '/' IdentifierPart*;

OpenBracket:                    '[';
CloseBracket:                   ']';
OpenParen:                      '(';
CloseParen:                     ')';
OpenBrace:                      '{' {l.ProcessOpenBrace();};
TemplateCloseBrace:             {p.IsInTemplateString()}? '}' -> popMode;
CloseBrace:                     '}' {l.ProcessCloseBrace();};
SemiColon:                      ';';
Comma:                          ',';
Assign:                         '=';
QuestionMark:                   '?';
QuestionMarkDot:                '?.';
Colon:                          ':';
Ellipsis:                       '...';
Dot:                            '.';
PlusPlus:                       '++';
MinusMinus:                     '--';
Plus:                           '+';
Minus:                          '-';
BitNot:                         '~';
Not:                            '!';
Multiply:                       '*';
Divide:                         '/';
Modulus:                        '%';
Power:                          '**';
NullCoalesce:                   '??';
Hashtag:                        '#';
RightShiftArithmetic:           '>>';
LeftShiftArithmetic:            '<<';
RightShiftLogical:              '>>>';
LessThan:                       '<';
MoreThan:                       '>';
LessThanEquals:                 '<=';
GreaterThanEquals:              '>=';
Equals_:                        '==';
NotEquals:                      '!=';
IdentityEquals:                 '===';
IdentityNotEquals:              '!==';
BitAnd:                         '&';
BitXOr:                         '^';
BitOr:                          '|';
And:                            '&&';
Or:                             '||';
MultiplyAssign:                 '*=';
DivideAssign:                   '/=';
ModulusAssign:                  '%=';
PlusAssign:                     '+=';
MinusAssign:                    '-=';
LeftShiftArithmeticAssign:      '<<=';
RightShiftArithmeticAssign:     '>>=';
RightShiftLogicalAssign:        '>>>=';
BitAndAssign:                   '&=';
BitXorAssign:                   '^=';
BitOrAssign:                    '|=';
PowerAssign:                    '**=';
NullishCoalescingAssign:        '??=';
ARROW:                          '=>';

/// Null Literals

NullLiteral:                    'null';
NaNLiteral:                      'NaN';
UndefinedLiteral:                'undefined';

/// Boolean Literals

BooleanLiteral:                 'true'
              |                 'false';

InfinityLiteral:                'Infinity';

/// Numeric Literals

DecimalLiteral:                 DecimalIntegerLiteral '.' [0-9] [0-9_]*
              |                 '.' [0-9] [0-9_]*
              |                 DecimalIntegerLiteral
              ;

ExponentLiteral:                 DecimalIntegerLiteral '.' [0-9] [0-9_]* ExponentPart
              |                 '.' [0-9] [0-9_]* ExponentPart
              |                 DecimalIntegerLiteral ExponentPart
              ;
/// Numeric Literals

HexIntegerLiteral:              '0' [xX] [0-9a-fA-F] HexDigit*;
OctalIntegerLiteral:           '0' [oO] [0-7] [_0-7]*;
BinaryIntegerLiteral:           '0' [bB] [01] [_01]*;

BigHexIntegerLiteral:           '0' [xX] [0-9a-fA-F] HexDigit* 'n';
BigOctalIntegerLiteral:         '0' [oO] [0-7] [_0-7]* 'n';
BigBinaryIntegerLiteral:        '0' [bB] [01] [_01]* 'n';
BigDecimalIntegerLiteral:       DecimalIntegerLiteral 'n';

/// Keywords

Break:                          'break';
Do:                             'do';
Instanceof:                     'instanceof';
Typeof:                         'typeof';
Case:                           'case';
Else:                           'else';
New:                            'new';
Var:                            'var';
Let:                            'let';
Catch:                          'catch';
Finally:                        'finally';
Return:                         'return';
Void:                           'void';
Continue:                       'continue';
For:                            'for';
Switch:                         'switch';
While:                          'while';
Debugger:                       'debugger';
Function_:                       'function';
This:                           'this';
With:                           'with';
Default:                        'default';
If:                             'if';
Throw:                          'throw';
Delete:                         'delete';
In:                             'in';
Try:                            'try';
As:                             'as';
From:                           'from';
Of:                             'of';
Class:                          'class';
Enum:                           'enum';
Extends:                        'extends';
Super:                          'super';
Const:                          'const';
Export:                         'export';
Import:                         'import';

Async:                          'async';
Sync:                           'sync';
Await:                          'await';
Yield:                          'yield';

Implements:                     'implements' {p.IsStrictMode()}?;
Private:                        'private' {p.IsStrictMode()}?;
Public:                         'public' {p.IsStrictMode()}?;
Interface:                      'interface' {p.IsStrictMode()}?;
Package:                        'package' {p.IsStrictMode()}?;
Protected:                      'protected' {p.IsStrictMode()}?;
Static:                         'static' {p.IsStrictMode()}?;



Identifier:                     IdentifierStart IdentifierPart*;

StringLiteral:                 ('"' DoubleStringCharacter* '"'
             |                  '\'' SingleStringCharacter* '\'') {l.ProcessStringLiteral();}
             ;

BackTick:                       '`' {l.IncreaseTemplateDepth();} -> pushMode(TEMPLATE);

WhiteSpaces:                    [\t\u000B\u000C\u0020\u00A0]+ -> channel(HIDDEN);

LineTerminator:                 [\r\n\u2028\u2029] -> channel(HIDDEN);

/// Comments


HtmlComment:                    '<!--' .*? '-->' -> channel(HIDDEN);
CDataComment:                   '<![CDATA[' .*? ']]>' -> channel(HIDDEN);
UnexpectedCharacter:            . -> channel(ERROR);

mode TEMPLATE;

BackTickInside:                 '`' {l.DecreaseTemplateDepth();} -> type(BackTick), popMode;
TemplateStringStartExpression:  '${' -> pushMode(DEFAULT_MODE);
TemplateStringAtom:             ~[`];



fragment DoubleStringCharacter
    : ~["\\\r\n]
    | '\\' EscapeSequence
    | LineContinuation
    ;

fragment SingleStringCharacter
    : ~['\\\r\n]
    | '\\' EscapeSequence
    | LineContinuation
    ;

fragment EscapeSequence
    : CharacterEscapeSequence
    | '0' // no digit ahead! TODO
    | HexEscapeSequence
    | UnicodeEscapeSequence
    | ExtendedUnicodeEscapeSequence
    ;

fragment CharacterEscapeSequence
    : SingleEscapeCharacter
    | NonEscapeCharacter
    ;

fragment HexEscapeSequence
    : 'x' HexDigit HexDigit
    ;

fragment UnicodeEscapeSequence
    : 'u' HexDigit HexDigit HexDigit HexDigit
    | 'u' '{' HexDigit HexDigit+ '}'
    ;

fragment ExtendedUnicodeEscapeSequence
    : 'u' '{' HexDigit+ '}'
    ;

fragment SingleEscapeCharacter
    : ['"\\bfnrtv]
    ;

fragment NonEscapeCharacter
    : ~['"\\bfnrtv0-9xu\r\n]
    ;

fragment EscapeCharacter
    : SingleEscapeCharacter
    | [0-9]
    | [xu]
    ;

fragment LineContinuation
    : '\\' [\r\n\u2028\u2029]+
    ;

fragment HexDigit
    : [_0-9a-fA-F]
    ;

fragment DecimalIntegerLiteral
    : '0'
    | [1-9] [0-9_]*
    ;

fragment ExponentPart
    : [eE] [+-]? [0-9_]+
    ;

fragment IdentifierPart
    : IdentifierStart
    | [\p{Mn}]
    | [\p{Nd}]
    | [\p{Pc}]
    | '\u200C'
    | '\u200D'
    ;

fragment IdentifierStart
    : [\p{L}]
    | [$_]
    | '\\' UnicodeEscapeSequence
    ;

fragment RegularExpressionFirstChar
    : ~[*\r\n\u2028\u2029\\/[]
    | RegularExpressionBackslashSequence
    | '[' RegularExpressionClassChar* ']'
    ;

fragment RegularExpressionChar
    : ~[\r\n\u2028\u2029\\/[]
    | RegularExpressionBackslashSequence
    | '[' RegularExpressionClassChar* ']'
    ;

fragment RegularExpressionClassChar
    : ~[\r\n\u2028\u2029\]\\]
    | RegularExpressionBackslashSequence
    ;

fragment RegularExpressionBackslashSequence
    : '\\' ~[\r\n\u2028\u2029]
    ;
