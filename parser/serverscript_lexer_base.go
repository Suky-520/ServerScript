package parser

import (
	"github.com/antlr4-go/antlr/v4"
)

// ServerScriptLexerBase state
type ServerScriptLexerBase struct {
	*antlr.BaseLexer

	scopeStrictModes []bool
	stackLength      int
	stackIx          int

	lastToken        antlr.Token
	useStrictDefault bool
	useStrictCurrent bool
	templateDepth    int
}

func (l *ServerScriptLexerBase) IsStartOfFile() bool {
	return l.lastToken == nil
}

func (l *ServerScriptLexerBase) pushStrictModeScope(v bool) {
	if l.stackIx == l.stackLength {
		l.scopeStrictModes = append(l.scopeStrictModes, v)
		l.stackLength++
	} else {
		l.scopeStrictModes[l.stackIx] = v
	}
	l.stackIx++
}

func (l *ServerScriptLexerBase) popStrictModeScope() bool {
	l.stackIx--
	v := l.scopeStrictModes[l.stackIx]
	l.scopeStrictModes[l.stackIx] = false
	return v
}

// IsStrictMode is self explanatory.
func (l *ServerScriptLexerBase) IsStrictMode() bool {
	return l.useStrictCurrent
}

// NextToken from the character stream.
func (l *ServerScriptLexerBase) NextToken() antlr.Token {
	next := l.BaseLexer.NextToken() // Get next token
	if next.GetChannel() == antlr.TokenDefaultChannel {
		// Keep track of the last token on def channel
		l.lastToken = next
	}
	return next
}

// ProcessOpenBrace is called when a { is encountered during
// lexing, we push a new scope everytime.
func (l *ServerScriptLexerBase) ProcessOpenBrace() {
	l.useStrictCurrent = l.useStrictDefault
	if l.stackIx > 0 && l.scopeStrictModes[l.stackIx-1] {
		l.useStrictCurrent = true
	}
	l.pushStrictModeScope(l.useStrictCurrent)
}

// ProcessCloseBrace is called when a } is encountered during
// lexing, we pop a scope unless we're inside global scope.
func (l *ServerScriptLexerBase) ProcessCloseBrace() {
	l.useStrictCurrent = l.useStrictDefault
	if l.stackIx > 0 {
		l.useStrictCurrent = l.popStrictModeScope()
	}
}

// ProcessStringLiteral is called when lexing a string literal.
func (l *ServerScriptLexerBase) ProcessStringLiteral() {
	if l.lastToken == nil || l.lastToken.GetTokenType() == ServerScriptLexerOpenBrace {
		if l.GetText() == `"use strict"` || l.GetText() == "'use strict'" {
			if l.stackIx > 0 {
				l.popStrictModeScope()
			}
			l.useStrictCurrent = true
			l.pushStrictModeScope(l.useStrictCurrent)
		}
	}
}

// IsRegexPossible returns true if the lexer can match a
// regex literal.
func (l *ServerScriptLexerBase) IsRegexPossible() bool {
	if l.lastToken == nil {
		return true
	}
	switch l.lastToken.GetTokenType() {
	//这里优化了在bigint除法时,出现和正则表达式冲突
	case ServerScriptLexerIdentifier, ServerScriptLexerNullLiteral,
		ServerScriptLexerBooleanLiteral, ServerScriptLexerThis,
		ServerScriptLexerCloseBracket, ServerScriptLexerCloseParen,
		ServerScriptLexerDecimalLiteral, ServerScriptLexerExponentLiteral,
		ServerScriptLexerHexIntegerLiteral, ServerScriptLexerOctalIntegerLiteral,
		ServerScriptLexerBinaryIntegerLiteral, ServerScriptLexerStringLiteral,
		ServerScriptLexerPlusPlus, ServerScriptLexerMinusMinus,
		ServerScriptLexerBigDecimalIntegerLiteral, ServerScriptLexerBigHexIntegerLiteral,
		ServerScriptLexerBigOctalIntegerLiteral, ServerScriptLexerBigBinaryIntegerLiteral:
		return false
	default:
		return true
	}
}

func (l *ServerScriptLexerBase) IncreaseTemplateDepth() {
	l.templateDepth++
}

func (l *ServerScriptLexerBase) DecreaseTemplateDepth() {
	l.templateDepth--
}

func (l *ServerScriptLexerBase) IsInTemplateString() bool {
	return l.templateDepth > 0
}

func (l *ServerScriptLexerBase) Reset() {
	l.scopeStrictModes = nil
	l.stackLength = 0
	l.stackIx = 0
	l.lastToken = nil
	l.useStrictDefault = false
	l.useStrictCurrent = false
	l.templateDepth = 0
	l.BaseLexer.Reset()
}
