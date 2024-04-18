//JavaScript中的Reader/Writer

package ss

import (
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
	"ss/collection"
	"strings"
)

// SetReaderFunction 将reader的方法添加到对象的属性中
func SetReaderFunction(prop JavaScriptProperties, reader io.Reader, name string) {
	prop.Set("read", ioRead(reader, name+".read"))
	prop.Set("close", ioClose(reader, name+".close"))
}

// SetWriterFunction 将writer的方法添加到对象的属性中
func SetWriterFunction(prop JavaScriptProperties, writer io.Writer, name string) {
	prop.Set("write", ioWrite(writer, name+".write"))
	prop.Set("close", ioClose(writer, name+".close"))
}

func ioRead(reader io.Reader, name string) JavaScript {
	return NewDefaultFunction(name, func(args []JavaScript, ctx *Context) JavaScript {
		if len(args) > 0 {
			if v, ok := args[0].(*Number); ok {
				b := make([]byte, int(v.Data))
				_, err := reader.Read(b)
				if err != nil {
					return NewUndefined()
				}
				return GoToBuffer(b)
			}
			panic(NewError("Error", "参数类型错误"))
		} else {
			all, err := io.ReadAll(reader)
			if err != nil {
				panic(NewError("Error", err.Error()))
			}
			return GoToBuffer(all)
		}
	})
}

func ioClose(closer any, name string) JavaScript {
	return NewDefaultFunction(name, func(args []JavaScript, ctx *Context) JavaScript {
		if v, ok := closer.(io.Closer); ok {
			v.Close()
		}
		return NewUndefined()
	})
}

func ioWrite(writer io.Writer, name string) JavaScript {
	return NewDefaultFunction(name, func(args []JavaScript, ctx *Context) JavaScript {
		if len(args) < 1 {
			panic(NewError("Error", "参数错误"))
		}

		if v, ok := args[0].(*Buffer); ok {
			//写入buffer
			i, err := writer.Write(v.Bytes())
			if err != nil {
				return NewUndefined()
			}
			return GoToNumber(i)
		} else if v, ok := args[0].(*String); ok {
			//写入字符串
			charset := ""
			if len(args) > 1 {
				if v, ok := args[1].(JavaScriptString); ok {
					charset = v.ToString()
				}
			}
			var i int
			var err error
			switch strings.ToLower(charset) {
			case "gbk":
				wr := transform.NewWriter(writer, simplifiedchinese.GBK.NewEncoder())
				i, err = wr.Write([]byte(v.Data))
			case "gb18030":
				wr := transform.NewWriter(writer, simplifiedchinese.GB18030.NewEncoder())
				i, err = wr.Write([]byte(v.Data))
			case "gb2312":
				wr := transform.NewWriter(writer, simplifiedchinese.HZGB2312.NewEncoder())
				i, err = wr.Write([]byte(v.Data))
			default:
				i, err = writer.Write([]byte(v.Data))
			}
			if err != nil {
				return NewUndefined()
			}
			return GoToNumber(i)
		} else if v, ok := args[0].(io.Reader); ok {
			//写入输入流
			i, err := io.Copy(writer, v)
			if err != nil {
				panic(NewError("Error", err.Error()))
			}
			return GoToNumber(i)
		} else if v, ok := args[0].(io.Reader); ok {
			//写入字节
			i, err := io.Copy(writer, v)
			if err != nil {
				panic(NewError("Error", err.Error()))
			}
			return GoToNumber(i)
		}
		panic(NewError("Error", "参数类型错误"))
	})
}

/**********Reader和Writer的定义************/

type Reader struct {
	io.Reader
	prop JavaScriptProperties
}

var _ JavaScript = &Reader{}

type Writer struct {
	io.Writer
	prop JavaScriptProperties
}

var _ JavaScript = &Writer{}

func GoToReader(reader io.Reader) JavaScript {
	r := &Reader{
		Reader: reader,
		prop:   collection.NewArrayMap[string, JavaScript](),
	}
	SetReaderFunction(r.prop, r, "reader")
	return r
}

func (r *Reader) GetName() string {
	return "reader"
}
func (r *Reader) GetProperties() JavaScriptProperties {
	return r.prop
}

func GoToWriter(writer io.Writer) JavaScript {
	r := &Writer{
		Writer: writer,
		prop:   collection.NewArrayMap[string, JavaScript](),
	}
	SetWriterFunction(r.prop, r, "writer")
	return r
}

func (w *Writer) GetName() string {
	return "writer"
}
func (w *Writer) GetProperties() JavaScriptProperties {
	return w.prop
}
