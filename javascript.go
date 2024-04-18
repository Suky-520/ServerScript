//定义了各种JavaScript相关接口

package ss

import (
	"encoding/json"
	"fmt"
	"math"
	"ss/collection"
)

// JavaScriptBasic 基本类型接口
type JavaScriptBasic interface {
	//Value 返回当前js对象的值,js基本类型会返回go的基本类型
	Value() any
}

// JavaScriptGetProperty 对象类型接口
type JavaScriptGetProperty interface {
	// GetProperty 属性获取
	GetProperty(name string) JavaScript
}

// JavaScriptDeleteProperty 对象类型接口
type JavaScriptDeleteProperty interface {
	// DeleteProperty 删除属性
	DeleteProperty(name string) JavaScript
}

// JavaScriptSetProperty 对象类型接口
type JavaScriptSetProperty interface {
	// SetProperty 属性设置
	SetProperty(name string, value JavaScript)
}

// JavaScriptGetProperties 对象类型接口
type JavaScriptGetProperties interface {
	//GetProperties 获取所有属性名和函数名列表
	GetProperties() JavaScriptProperties
}

type JavaScriptBool interface {
	// ToBool 在逻辑运算、条件判断等场景下，任何类型都会转换为bool。
	/**
	Boolean: 显而易见，true 保持为 true，false 保持为 false。
	Null: null 转换为 false。
	Undefined: undefined 转换为 false。
	ToNumber:
		0 和 -0 转换为 false。
		NaN（不是一个数字）转换为 false。
		其他所有数字（包括所有正数和负数，除了 0 和 -0）转换为 true。
	BigInt:
		0n（BigInt 类型的 0）转换为 false。
		任何非零的 BigInt 值（包括正数和负数）转换为 true。
	ToString:
		空字符串 "" 转换为 false。
		任何非空字符串转换为 true。
	Object（包括数组、对象、函数等）:
		任何对象（包括空对象 {}、任何数组 []、任何函数）都转换为 true，无论其内容如何。
	*/
	ToBool() bool
}

// JavaScriptNumber 算术运算时会调用对象的Number接口,
type JavaScriptNumber interface {
	// ToNumber 当不同类型的值在需要数值的上下文中使用时（比如算术运算），会按照一定的规则自动转换成 Number 类型。这个过程称为类型强制转换
	/**
	Boolean: true 转换为 1，false 转换为 0。
	Null: null 转换为 0。
	Undefined: undefined 转换为 NaN（Not a ToNumber，不是一个数字）。
	ToNumber: 数字类型的值不需要转换。
	BigInt: 直接将 BigInt 值转换为 ToNumber 可能会导致精度丢失，因为 ToNumber 类型有限的表示范围和精度。在强制转换过程中，如果 BigInt 值在 ToNumber 能表示的范围内，则转换为相应的 ToNumber 值；否则，可能导致精度丢失。
	ToString:
		空字符串 "" 转换为 0。
		如果字符串是纯数字形式（如 "123"），则转换为对应的数字（123）。
		如果字符串包含非数字字符（如 "123abc" 或 "abc"），则转换为 NaN。
		如果字符串是数字前后有空格（如 " 123 "），空格会被忽略，转换为数字（123）。
	Object（包括数组、对象、函数等）:
		对于对象（包括数组和函数），首先会尝试调用对象的 valueOf() 方法，如果结果是原始值，则按照上述规则转换为数字。
		如果 valueOf() 方法不存在或者返回的不是原始值，则调用 toString() 方法，并将得到的字符串按照字符串到数字的转换规则转换。
		空数组 [] 转换为 0，只有一个数字元素的数组（如 [123]）转换为 123，其他情况（如 [1, 2] 或者 ["abc"]）转换为 NaN。
		空对象 {} 和其他不直接转换为原始值的对象转换尝试通常结果为 NaN。
	*/
	ToNumber() float64
}

// JavaScriptComputable 通用的计算接口
type JavaScriptComputable interface {
	// Calculate 计算
	Calculate(v JavaScript, operate Operate) JavaScript
}

// JavaScriptClass 类
type JavaScriptClass interface {
	//Constructor 构造函数
	Constructor(args []JavaScript, ctx *Context) JavaScript
}

// JavaScriptFunction 函数
type JavaScriptFunction interface {
	//Call 函数调用
	Call(state int, args []JavaScript, ctx *Context) JavaScript
}

// JavaScriptBinary 接口
type JavaScriptBinary interface {
	//Binary 返回二进制数据
	Binary() []byte
}

// JavaScriptString 字符串接口
type JavaScriptString interface {
	//ToString 对象的字符串表示
	ToString() string
}

// JavaScriptPrint 打印接口
type JavaScriptPrint interface {
	// Print 打印,带有字体颜色和格式化
	Print() string
}

// JavaScript 对象的基本接口,实现了该接口的go结构体,可以被注册为js对象
// 支持的类型: number，string，boolean，bigint,null,function,undefined
type JavaScript interface {
	//GetName 名称
	GetName() string
}

// JavaScriptType js对象类型
type JavaScriptType struct {
	Name       string   //对象名
	Type       DataType //对象类型
	Native     bool     //是否原生
	TypeString string   //类型字符串表示
}

// JavaScriptProperties 接口,该接口是一个有序map,用来记录对象的属性列表
type JavaScriptProperties interface {
	Set(key string, value JavaScript)
	Get(key string) (JavaScript, bool)
	Delete(key string) bool
	Has(key string) bool
	Size() int
	Keys() []string
}

// JavaScriptGetIterator 遍历一个对象的属性,用于for-in语法
type JavaScriptGetIterator interface {
	GetIterator() JavaScriptIterator
}

// JavaScriptIterator 迭代器接口
type JavaScriptIterator interface {
	Next() (v []JavaScript, done bool)
	JavaScript
}

// JsToBasic 将js对象转为基本类型的值
func JsToBasic(value JavaScript) (any, bool) {
	if v, ok := value.(JavaScriptBasic); ok {
		return v.Value(), ok
	}
	return nil, false
}

// JsToString 将js对象转换为字符串表示形式
func JsToString(js JavaScript) string {
	if v, ok := js.(JavaScriptString); ok {
		return v.ToString()
	}
	return "[object " + js.GetName() + "]"
}

// JsToBool 将js对象转换为bool类型
func JsToBool(js JavaScript) bool {
	if v, ok := js.(JavaScriptBool); ok {
		return v.ToBool()
	}
	return true
}

// JsToNumber 将js对象转换为数字类型
func JsToNumber(js JavaScript) float64 {
	if v, ok := js.(JavaScriptNumber); ok {
		return v.ToNumber()
	}
	return math.NaN()
}

// JsToCalculate 进行计算
func JsToCalculate(v1 JavaScript, v2 JavaScript, operate Operate) JavaScript {
	if v, ok := v1.(JavaScriptComputable); ok {
		return v.Calculate(v2, operate)
	}
	return DefaultCalculate(v1, v2, operate)
}

// JsGetProperty 获取js的属性
func JsGetProperty(js JavaScript, name string) JavaScript {
	if v, ok := js.(JavaScriptGetProperty); ok {
		if x := v.GetProperty(name); !IsUndefined(x) {
			return x
		}
	}
	if v, ok := js.(JavaScriptGetProperties); ok {
		if x, ok := v.GetProperties().Get(name); ok && !IsUndefined(x) {
			return x
		}
	}
	//默认增加toString方法
	if name == "toString" {
		return NewDefaultFunction("toString", func(args []JavaScript, ctx *Context) JavaScript {
			return GoToString(JsToString(js))
		})
	}
	return NewUndefined()
}

// JsSetProperty 获取js的属性
func JsSetProperty(js JavaScript, name string, value JavaScript) {
	if v, ok := js.(JavaScriptSetProperty); ok {
		v.SetProperty(name, value)
	}
}

// JsGetProperties 获取js的属性
func JsGetProperties(js JavaScript) JavaScriptProperties {
	if v, ok := js.(JavaScriptGetProperties); ok {
		return v.GetProperties()
	}
	return collection.NewArrayMap[string, JavaScript]()
}

// JsDeleteProperty 获取js的属性
func JsDeleteProperty(js JavaScript, name string) JavaScript {
	if v, ok := js.(JavaScriptDeleteProperty); ok {
		return v.DeleteProperty(name)
	}
	return NewUndefined()
}

// JsToJson 将js对象转换为json字符串
func JsToJson(js JavaScript) string {
	if v, ok := js.(json.Marshaler); ok {
		marshal, err := json.Marshal(v)
		if err != nil {
			panic(NewSyntaxError(err.Error()))
		}
		return string(marshal)
	} else {
		return "null"
	}
}

// JsToPrint 将js对象转换为打印格式的字符串
func JsToPrint(js JavaScript, indent string) string {
	//打印用户定义对象
	if v, ok := js.(*Object); ok {
		return printObject(JsGetProperties(v), indent, "")
	}
	//打印数组
	if v, ok := js.(*Array); ok {
		return printArray(v)
	}
	//打印用户定义函数
	if v, ok := js.(*Function); ok {
		return printFunction(v)
	}
	//实现打印接口的对象
	if v, ok := js.(JavaScriptPrint); ok {
		return v.Print()
	}
	//没有实现打印接口的
	if v, ok := js.(JavaScriptGetProperties); ok {
		return printObject(v.GetProperties(), indent, printTitle(v.(JavaScript)))
	}
	return JsToString(js)
}

// 类型字符串表示
var types = []string{"string", "number", "boolean", "null", "undefined", "bigint", "object", "array", "function", "buffer", "class"}

// JsToType 将js对象转换为类型
func JsToType(js JavaScript) JavaScriptType {
	jst := JavaScriptType{Name: js.GetName(), Native: true, Type: ObjectType, TypeString: "object"}
	switch js.(type) {
	case *Object:
		jst.Type = ObjectType
		jst.Native = false
	case *Function:
		jst.Type = FunctionType
		jst.Native = false
	case *Boolean:
		jst.Type = BooleanType
	case *Number:
		jst.Type = NumberType
	case *Null:
		jst.Type = NullType
	case *Undefined:
		jst.Type = UndefinedType
	case *String:
		jst.Type = StringType
	case *RegExp:
		jst.Type = ObjectType
	case *Buffer:
		jst.Type = BufferType
	case *BigInt:
		jst.Type = BigIntType
	case *Array:
		jst.Type = ArrayType
	case JavaScriptClass:
		jst.Type = ClassType
	case JavaScriptFunction:
		jst.Type = FunctionType
	default:
		jst.Type = ObjectType
	}
	jst.TypeString = types[jst.Type-1]
	return jst
}

func printTitle(v JavaScript) string {
	t := JsToType(v)
	return fmt.Sprintf("%v[%v %v]%v", PurpleStyle, t.TypeString, t.Name, EndStyle)
}

// printObject 打印对象
func printObject(props JavaScriptProperties, indent string, title string) string {
	if props.Size() == 0 {
		return title + "{}"
	}
	str := title + "{\r\n"
	s := "  "
	if len(indent) > 1 {
		s = indent
	}
	for _, k := range props.Keys() {
		v, _ := props.Get(k)
		str = str + s + k + ": " + JsToPrint(v, s+"  ") + ",\r\n"
	}
	str = str + s[2:] + "}"
	return str
}

// printFunction 打印函数
func printFunction(obj JavaScript) string {
	return fmt.Sprintf("%v[Function: %v]%v", BlueStyle, obj.GetName(), EndStyle)
}

// printClass 打印类
func printClass(obj JavaScript) string {
	return fmt.Sprintf("%v[Class: %v]%v", BlueStyle, obj.GetName(), EndStyle)
}

// printArray 打印数组
func printArray(a *Array) string {
	s := "[ "
	for i, val := range a.Data {
		pri := ""
		if val != nil {
			pri = JsToPrint(val, "")
		} else {
			pri = "undefined"
		}
		if i == 0 {
			s = s + pri
		} else {
			s = s + ", " + pri
		}
	}
	s = s + " ]"
	return s
}

// IsUndefined 是否为Undefined
func IsUndefined(js JavaScript) bool {
	if js == nil {
		return true
	}
	t := JsToType(js)
	if t.Type == UndefinedType {
		return true
	}
	return false
}
