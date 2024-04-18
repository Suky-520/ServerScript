//JavaScript中的BufferReader对象

package ss

import (
	"bufio"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
	"ss/collection"
	"strings"
)

type BufferReader struct {
	*bufio.Reader
	prop *collection.ArrayMap[string, JavaScript]
}

var _ io.Reader = &BufferReader{}

func init() {
	setGlobalObject("BufferReader", NewDefaultClass("BufferReader", newBufferReader()))
}

func NewBufferReader(reader io.Reader) *BufferReader {
	c := &BufferReader{Reader: bufio.NewReader(reader), prop: collection.NewArrayMap[string, JavaScript]()}
	c.prop.Set("readString", c.readString())
	c.prop.Set("readChar", c.readChar())
	SetReaderFunction(c.prop, c, "bufferReader")
	return c
}

func (b *BufferReader) GetName() string {
	return "bufferReader"
}

func (b *BufferReader) GetProperties() JavaScriptProperties {
	return b.prop
}

// readChar 用于按照字符进行读取
func (b *BufferReader) readChar() JavaScript {
	return NewDefaultFunction("net.bufferReader.readChar", func(args []JavaScript, ctx *Context) JavaScript {
		r, _, err := b.ReadRune()
		if err == io.EOF {
			return NewUndefined()
		}
		if err != nil {
			// 处理读取错误
			panic(NewError("Error", err.Error()))
		}
		// 读取的Unicode字符
		return GoToString(string(r))
	})
}

// 提供readString方法,用于快速解析文本文件
func (b *BufferReader) readString() JavaScript {
	return NewDefaultFunction("net.bufferReader.readString", func(args []JavaScript, ctx *Context) JavaScript {
		if len(args) < 1 {
			panic(NewError("Error", "参数错误"))
		}
		if v, ok := args[0].(*String); ok {
			if len(v.Data) > 0 {
				readString, err := b.Reader.ReadString('\n')
				if err != nil && readString == "" {
					return NewUndefined()
				}
				return GoToString(readString)
			} else {
				panic(NewError("Error", "数据长度不够"))
			}
		} else if v, ok := args[0].(*Number); ok {
			if v.Data >= 0 || v.Data <= 255 {
				readString, err := b.Reader.ReadString(byte(v.Data))
				if err != nil && readString == "" {
					return NewUndefined()
				}
				return GoToString(readString)
			} else {
				panic(NewError("Error", "数据不是0-255之间"))
			}
		}
		panic(NewError("Error", "参数类型错误"))
	})
}

func newBufferReader() func(args []JavaScript, ctx *Context) JavaScript {
	return func(args []JavaScript, ctx *Context) JavaScript {
		if len(args) < 1 {
			panic(NewError("Error", "参数错误"))
		}
		var utf8Reader *transform.Reader
		charset := ""
		if len(args) > 1 {
			if v, ok := args[1].(JavaScriptString); ok {
				charset = v.ToString()
			}
		}
		var reader io.Reader
		if r, ok := args[0].(io.Reader); ok {
			reader = r
		} else {
			panic(NewError("Error", "参数类型错误"))
		}

		switch strings.ToLower(charset) {
		case "gbk":
			utf8Reader = transform.NewReader(reader, simplifiedchinese.GBK.NewDecoder())
			return NewBufferReader(utf8Reader)
		case "gb18030":
			utf8Reader = transform.NewReader(reader, simplifiedchinese.GB18030.NewDecoder())
			return NewBufferReader(utf8Reader)
		case "gb2312":
			utf8Reader = transform.NewReader(reader, simplifiedchinese.HZGB2312.NewDecoder())
			return NewBufferReader(utf8Reader)
		default:
			return NewBufferReader(reader)
		}
	}
}
