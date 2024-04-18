// BigInt这是ss中用来表示JavaScript BigInt的结构体
//这是一个JavaScript基本类型,所以必须实现JavaScriptBasic接口

package ss

import (
	"encoding/json"
	"fmt"
	"math"
	"math/big"
	"strings"
)

// BigInt JavaScript BigInt类型
type BigInt struct {
	Data *big.Int
}

// 实现JavaScript接口
var _ JavaScript = &BigInt{}

// 实现JavaScriptBasic接口
var _ JavaScriptBasic = &BigInt{}

// 实现CodeCompiler接口
var _ CodeCompiler = &BigInt{}

func (b *BigInt) GetName() string {
	return "bigint"
}
func (b *BigInt) GetProperty(name string) JavaScript {
	switch name {
	case "toString":
		return bigintToStr(b)
	default:
		return NewUndefined()
	}
}
func (b *BigInt) ToString() string {
	return fmt.Sprint(b.Data)
}
func (b *BigInt) ToNumber() float64 {
	return math.NaN()
}
func (b *BigInt) ToBool() bool {
	return b.Data.Cmp(big.NewInt(0)) != 0
}
func (b *BigInt) ToByteArray() []byte {
	return b.Data.Bytes()
}
func (b *BigInt) Print() string {
	return YellowStyle + b.ToString() + "n" + EndStyle
}
func (b *BigInt) Value() any {
	return b.Data
}
func (b *BigInt) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.Data.String())
}
func (b *BigInt) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &b.Data)
}
func (b *BigInt) MarshalBin() ([]byte, error) {
	data := []byte{1}
	data = append(data, b.Data.Bytes()...)
	return data, nil
}
func (b *BigInt) UnmarshalBin(data []byte) (any, error) {
	b.Data.SetBytes(data)
	return nil, nil
}

// GoToBigInt go类型转js的bigint
func GoToBigInt[T *big.Int | int | int64 | string](d T) *BigInt {
	o := &BigInt{}
	var a any = d
	switch t := a.(type) {
	case *big.Int:
		o.Data = t
	case int:
		o.Data = big.NewInt(int64(t))
	case int64:
		o.Data = big.NewInt(t)
	case string:
		str := strings.TrimSuffix(strings.ToLower(t), "n")
		var num big.Int
		_, success := num.SetString(str, 10)
		if !success {
			panic(NewRangeError("The string cannot be converted to a BigInt because it is not an int"))
		}
		o.Data = &num
	default:
		panic("The type is not supported")
	}
	return o
}

func bigintToStr(b *BigInt) JavaScript {
	return NewDefaultFunction("toString", func(args []JavaScript, ctx *Context) JavaScript {
		if len(args) == 0 {
			return GoToString(b.ToString())
		}
		radix := 10
		if j, ok := args[0].(*Number); ok {
			radix = int(j.Data)
		}
		if radix < 2 || radix > 36 {
			panic(NewRangeError("toString() radix argument must be between 2 and 36"))
		}
		if radix == 10 {
			return GoToString(b.ToString())
		}
		return GoToString(b.Data.Text(radix))
	})
}

func hexStrToBigInt(str string) JavaScript {
	s := strings.TrimSuffix(strings.ToLower(str), "n")
	var num big.Int
	_, success := num.SetString(s, 16)
	if !success {
		panic(NewRangeError("The string cannot be converted to a BigInt"))
	}
	return GoToBigInt(&num)
}

func octalStrToBigInt(str string) JavaScript {
	s := strings.TrimSuffix(strings.ToLower(str), "n")
	var num big.Int
	_, success := num.SetString(s, 8)
	if !success {
		panic(NewRangeError("The string cannot be converted to a BigInt"))
	}
	return GoToBigInt(&num)
}

func binaryStrToBigInt(str string) JavaScript {
	s := strings.TrimSuffix(strings.ToLower(str), "n")
	var num big.Int
	_, success := num.SetString(s, 2)
	if !success {
		panic(NewRangeError("The string cannot be converted to a BigInt"))
	}
	return GoToBigInt(&num)
}
