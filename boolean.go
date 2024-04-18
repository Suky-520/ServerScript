// Boolean这是ss中用来表示JavaScript boolean的结构体
//这是一个基本类型,所以必须实现JavaScriptBasic接口

package ss

import (
	"encoding/json"
	"fmt"
)

// Boolean JavaScript布尔类型
type Boolean struct {
	Data bool
}

// 实现JavaScript接口
var _ JavaScript = &Boolean{}

// 实现JavaScriptBasic接口
var _ JavaScriptBasic = &Boolean{}

// 实现CodeCompiler接口
var _ CodeCompiler = &Boolean{}

func (b *Boolean) GetName() string {
	return "boolean"
}
func (b *Boolean) ToString() string {
	return fmt.Sprint(b.Data)
}
func (b *Boolean) ToNumber() float64 {
	if b.Data {
		return 1
	}
	return 0
}
func (b *Boolean) ToBool() bool {
	return b.Data
}
func (b *Boolean) Value() any {
	return b.Data
}
func (b *Boolean) Print() string {
	return YellowStyle + b.ToString() + EndStyle
}
func (b *Boolean) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}
func (b *Boolean) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &b.Data)
}
func (b *Boolean) MarshalBin() ([]byte, error) {
	if b.Data {
		return []byte{2, 1}, nil
	} else {
		return []byte{2, 0}, nil
	}
}
func (b *Boolean) UnmarshalBin(data []byte) (any, error) {
	if data[0] == 1 {
		b.Data = true
	} else {
		b.Data = false
	}
	return nil, nil
}

func GoToBoolean[T bool | string](i T) *Boolean {
	o := &Boolean{Data: false}
	if fmt.Sprint(i) == "true" {
		o = &Boolean{Data: true}
	}
	return o
}
