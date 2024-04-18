//JavaScript中的null类型

package ss

/*************** Null *****************/

// Null JavaScript null类型
type Null struct {
}

var _ JavaScript = &Null{}
var _ JavaScriptBasic = &Null{}
var _ CodeCompiler = &Null{}

// NewNull 新建js中的null对象
func NewNull() *Null {
	return &Null{}
}

func (n *Null) GetName() string {
	return "null"
}

func (n *Null) ToString() string {
	return "null"
}
func (n *Null) Print() string {
	return BoldStyle + "null" + EndStyle
}

func (n *Null) Value() any {
	return Null{}
}
func (n *Null) ToNumber() float64 {
	return 0
}
func (n *Null) ToBool() bool {
	return false
}
func (n *Null) MarshalJSON() ([]byte, error) {
	return []byte("null"), nil
}
func (n *Null) UnmarshalBin(data []byte) (any, error) {
	return NewNull(), nil
}

func (n *Null) MarshalBin() ([]byte, error) {
	return []byte{3}, nil
}

/*************** Undefined *****************/

// Undefined JavaScript undefined类型
type Undefined struct {
}

// NewUndefined 新建js中的Undefined对象
func NewUndefined() *Undefined {
	o := &Undefined{}
	return o
}

var _ JavaScript = &Undefined{}
var _ JavaScriptBasic = &Undefined{}
var _ CodeCompiler = &Undefined{}

func (u *Undefined) GetName() string {
	return "undefined"
}

func (u *Undefined) ToString() string {
	return "undefined"
}
func (u *Undefined) Print() string {
	return GreyStyle + "undefined" + EndStyle
}

func (u *Undefined) Value() any {
	return Undefined{}
}

func (u *Undefined) ToBool() bool {
	return false
}
func (u *Undefined) MarshalJSON() ([]byte, error) {
	return []byte("null"), nil
}
func (u *Undefined) MarshalBin() ([]byte, error) {
	return []byte{4}, nil
}
func (u *Undefined) UnmarshalBin(bytes []byte) (any, error) {
	return NewUndefined(), nil
}
