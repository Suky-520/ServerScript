//这是源代码解析器,可以将js源文件解析为AST

package ss

import (
	"bufio"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"os"
	"path/filepath"
	"ss/parser"
)

// AstParser AST解析器
type AstParser struct {
	parser.BaseServerScriptParserListener          //父类
	Program                               *Program //最终转换的程序
	ident                                 int      //0加载,1修改
}

// ErrorListener 错误监听器
type ErrorListener struct {
	*antlr.DefaultErrorListener //继承ErrorListener基类
	file                        int
	error                       bool
}

// 将js源文件编译成程序
func parseFile(filePath string) *Program {
	absPath, err := filepath.Abs(filePath)
	file, err := os.Open(absPath)
	if err != nil {
		panic(NewReferenceError("no such File '" + filePath + "'"))
		os.Exit(1)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	is := antlr.NewIoStream(reader)
	fileIndex := setSourceFile(absPath, "")
	p := doParse(is, fileIndex)
	return p
}

// 将js源代码字符串编译成程序
func parseSource(src string, filePath string) *Program {
	fileIndex := setSourceFile(filePath, src)
	is := antlr.NewInputStream(src)
	return doParse(is, fileIndex)
}

// 语法解析
func doParse(is *antlr.InputStream, fileIndex int) *Program {
	defer func() {
		if err := recover(); err != nil {
			switch err.(type) {
			case error:
				PrintError("CompilerError: " + err.(error).Error())
			default:
				PrintError("CompilerError: " + fmt.Sprint(err))
			}
		}
	}()
	lexer := parser.NewServerScriptLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewServerScriptParser(stream)
	p.RemoveErrorListeners()
	err := &ErrorListener{file: fileIndex}
	p.AddErrorListener(err)
	tree := p.Program()
	if err.error {
		return nil
	}
	pro := newProgram(fileIndex)
	if tree.SourceElements() == nil {
		return nil
	}
	pro.compile(tree.SourceElements())
	return pro
}

// GetComment 获取注释,1单行注释，2多行注释
func GetComment(src string, channel int) []string {
	is := antlr.NewInputStream(src)
	//1、创建词法分析器
	lexer := parser.NewServerScriptLexer(is)
	comments := make([]string, 0)
	for {
		token := lexer.NextToken()
		if token.GetTokenType() < 0 {
			break
		}
		if token.GetChannel() == channel {
			comments = append(comments, token.GetText())
		}
	}
	return comments
}

// SyntaxError 编译异常监听
func (d *ErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	p := recognizer.(antlr.Parser)
	token := p.GetCurrentToken().GetText()
	d.error = true
	PrintError("SyntaxError: " + token)
	PrintErrorStack(d.file, line, column, token)
}
