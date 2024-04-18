// Buffer这是ss中用来表示JavaScript Buffer的结构体
//这是一个引用类型,数据结构是一个字节数组，同时也实现了io.Writer和io.Reader接口

package ss

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
	"ss/collection"
	"strconv"
	"strings"
)

// Buffer 字节流对象,是一个字节数组，这是一个引用类型
type Buffer struct {
	*bytes.Buffer
	prop JavaScriptProperties
}

var _ JavaScript = &Buffer{}

var _ io.Writer = &Buffer{}

var _ io.Reader = &Buffer{}

// 初始化
func init() {
	c := func(args []JavaScript, ctx *Context) JavaScript {
		if len(args) > 0 {
			if v, ok := args[0].(*Number); ok {
				return GoToBuffer(make([]byte, int(v.Data)))
			}
		}
		return GoToBuffer(make([]byte, 0))
	}
	setGlobalObject("Buffer", NewDefaultClass("Buffer", c))
}

// GoToBuffer 将buffer和[]byte转换为js中的Buffer对象
func GoToBuffer[T *bytes.Buffer | []byte](d T) *Buffer {
	var a any = d
	o := &Buffer{prop: collection.NewArrayMap[string, JavaScript]()}
	switch t := a.(type) {
	case *bytes.Buffer:
		o.Buffer = t
	case []byte:
		o.Buffer = bytes.NewBuffer(t)
	}
	//写入数据
	o.prop.Set("reset", o.reset())
	o.prop.Set("size", o.size())
	o.prop.Set("writeUInt8", o.writeUInt8())
	o.prop.Set("toByteString", o.toByteString())
	o.prop.Set("toString", o.toString())
	o.prop.Set("toBase64", o.toBase64())
	SetReaderFunction(o.prop, o, "buffer")
	SetWriterFunction(o.prop, o, "buffer")
	return o
}

func (b *Buffer) GetName() string {
	return "buffer"
}
func (b *Buffer) GetProperties() JavaScriptProperties {
	return b.prop
}
func (b *Buffer) ToString() string {
	return b.Buffer.String()
}
func (b *Buffer) ToNumber() float64 {
	return 0
}
func (b *Buffer) ToBool() bool {
	return true
}
func (b *Buffer) getValue(name string) JavaScript {
	num, err := strconv.Atoi(name)
	if err != nil {
		panic(NewSyntaxError("'" + name + "' is not a number'"))
	}
	if b.Buffer.Len() > num {
		x := b.Buffer.Bytes()[num]
		return GoToNumber(int(x))
	} else {
		return NewUndefined()
	}
}
func (b *Buffer) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.Buffer)
}
func (b *Buffer) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &b.Buffer)
}
func (b *Buffer) MarshalBin() ([]byte, error) {
	data := []byte{8}
	data = append(data, b.Buffer.Bytes()...)
	return data, nil
}
func (b *Buffer) UnmarshalBin(data []byte) (any, error) {
	b.Buffer = bytes.NewBuffer(data)
	return nil, nil
}

func (b *Buffer) writeUInt8() JavaScript {
	return NewDefaultFunction("buffer.writeUInt8", func(args []JavaScript, ctx *Context) JavaScript {
		if len(args) == 0 {
			return NewUndefined()
		}
		index := -1
		if len(args) > 1 {
			if n, ok := args[1].(*Number); ok {
				index = int(n.Data)
			} else {
				panic(NewError("Error", "参数2数据类型错误"))
			}
		}
		if v, ok := args[0].(*Number); ok {
			if v.Data > 255 || v.Data < 0 {
				panic(NewError("Error", "数据超过指定范围"))
			}
			if index < 0 {
				b.Buffer.WriteByte(byte(v.Data))
			} else {
				if b.Buffer.Len() <= index {
					return NewUndefined()
				}
				b.Buffer.Bytes()[index] = byte(v.Data)
			}
			return NewUndefined()
		} else {
			panic(NewError("Error", "参数1数据类型错误"))
		}
	})
}

func (b *Buffer) toByteString() JavaScript {
	return NewDefaultFunction("buffer.toByteString", func(args []JavaScript, ctx *Context) JavaScript {
		return GoToString(fmt.Sprint(b.Buffer.Bytes()))
	})
}

func (b *Buffer) toBase64() JavaScript {
	return NewDefaultFunction("buffer.toBase64", func(args []JavaScript, ctx *Context) JavaScript {
		return GoToString(base64.StdEncoding.EncodeToString(b.Buffer.Bytes()))
	})
}

func (b *Buffer) toString() JavaScript {
	return NewDefaultFunction("buffer.toString", func(args []JavaScript, ctx *Context) JavaScript {
		charset := ""
		if len(args) > 0 {
			if v, ok := args[0].(JavaScriptString); ok {
				charset = v.ToString()
			}
		}
		switch strings.ToLower(charset) {
		case "gbk":
			reader := transform.NewReader(b.Buffer, simplifiedchinese.GBK.NewDecoder())
			all, err := io.ReadAll(reader)
			if err != nil {
				panic(NewError("Error", err.Error()))
			}
			return GoToString(string(all))
		case "gb18030":
			reader := transform.NewReader(b.Buffer, simplifiedchinese.GB18030.NewDecoder())
			all, err := io.ReadAll(reader)
			if err != nil {
				panic(NewError("Error", err.Error()))
			}
			return GoToString(string(all))
		case "gb2312":
			reader := transform.NewReader(b.Buffer, simplifiedchinese.HZGB2312.NewDecoder())
			all, err := io.ReadAll(reader)
			if err != nil {
				panic(NewError("Error", err.Error()))
			}
			return GoToString(string(all))
		default:
			return GoToString(string(b.Buffer.Bytes()))
		}

	})
}

func (b *Buffer) reset() JavaScript {
	return NewDefaultFunction("buffer.reset", func(args []JavaScript, ctx *Context) JavaScript {
		b.Reset()
		return NewUndefined()
	})
}

func (b *Buffer) size() JavaScript {
	return NewDefaultFunction("buffer.size", func(args []JavaScript, ctx *Context) JavaScript {
		return GoToNumber(b.Len())
	})
}
