//JavaScript中的number类型和Number对象

package ss

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"math"
	"ss/collection"
	"strconv"
	"strings"
	"unicode"
)

// Number JavaScript Number类型
type Number struct {
	Data float64
}

var _ JavaScript = &Number{}
var _ JavaScriptBasic = &Number{}
var _ CodeCompiler = &Number{}

func init() {
	//注册一个全局对象Number
	o := NewDefaultObject("Number")
	o.AddProperty("isFinite", isFinite())
	o.AddProperty("isInteger", isInteger())
	o.AddProperty("isNaN", isNaN1())
	o.AddProperty("isSafeInteger", isSafeInteger())
	o.AddProperty("parseFloat", NewParseFloat())
	o.AddProperty("parseInt", NewParseInt())
	o.AddProperty("MAX_VALUE", GoToNumber(math.MaxFloat64))
	o.AddProperty("MIN_VALUE", GoToNumber(math.SmallestNonzeroFloat64))
	setGlobalObject("Number", o)
}

// GoToNumber 将go语言中的float64转换为js中的number
func GoToNumber[T float64 | float32 | int | int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64 | string](d T) *Number {
	o := &Number{}
	var a any = d
	switch t := a.(type) {
	case float64:
		o.Data = t
	case int:
		o.Data = float64(t)
	case float32:
		o.Data = float64(t)
	case int8:
		o.Data = float64(t)
	case int16:
		o.Data = float64(t)
	case int32:
		o.Data = float64(t)
	case int64:
		o.Data = float64(t)
	case uint8:
		o.Data = float64(t)
	case uint16:
		o.Data = float64(t)
	case uint32:
		o.Data = float64(t)
	case uint64:
		o.Data = float64(t)
	case string:
		i, e := strconv.ParseFloat(t, 64)
		if e != nil {
			o.Data = math.NaN()
		} else {
			o.Data = i
		}
	default:
		panic("not support type ")
	}
	return o
}

// NewNaN 新建NaN
func NewNaN() *Number {
	return GoToNumber(math.NaN())
}

// IsNaN 判断js变量是否为NaN
func IsNaN(s JavaScript) bool {
	t := JsToType(s)
	if t.Type != NumberType {
		return false
	}
	return math.IsNaN(s.(*Number).Data)
}

// 将16进制表示的字符串转换为js中的number
func hexStrToNumber(hexStr string) JavaScript {
	cleanHexStr := strings.TrimPrefix(strings.ToLower(hexStr), "0x")
	i, e := strconv.ParseInt(cleanHexStr, 16, 64)
	if e != nil {
		return NewNaN()
	}
	return GoToNumber(float64(i))
}

// 将8进制表示的字符串转换为js中的number
func octalStrToNumber(octalStrToInt string) JavaScript {
	nStr := strings.TrimPrefix(strings.ToLower(octalStrToInt), "0o")
	i, e := strconv.ParseInt(nStr, 8, 64)
	if e != nil {
		return NewNaN()
	}
	return GoToNumber(float64(i))
}

// 将2进制表示的字符串转换为js中的number
func binaryStrToNumber(octalStrToInt string) JavaScript {
	nStr := strings.TrimPrefix(strings.ToLower(octalStrToInt), "0b")
	i, e := strconv.ParseInt(nStr, 2, 64)
	if e != nil {
		return NewNaN()
	}
	return GoToNumber(float64(i))
}

func floatToString(number float64, radix int) string {
	if radix < 2 || radix > 36 {
		panic(NewRangeError(fmt.Sprintf("radix out of range: %d", radix)))
	}
	// 对于十进制，使用 FormatFloat 并保留原始的浮点数精度。
	if radix == 10 {
		// -1 表示使用最短必要的小数位数，'f' 表示不使用科学计数法。
		return strconv.FormatFloat(number, 'f', -1, 64)
	}
	// 对于其他进制，仅处理整数部分。
	intPart := int64(math.Floor(math.Abs(number)))
	formatted := strconv.FormatInt(intPart, radix)
	if number < 0 {
		formatted = "-" + formatted
	}
	return formatted
}

func numberToStr(n *Number) JavaScript {
	return NewDefaultFunction("toString", func(args []JavaScript, ctx *Context) JavaScript {
		if math.IsInf(n.Data, 1) {
			return GoToString("Infinity")
		}
		if math.IsInf(n.Data, -1) {
			return GoToString("-Infinity")
		}
		if math.IsNaN(n.Data) {
			return GoToString("NaN")
		}
		if len(args) == 0 {
			return GoToString(n.ToString())
		}
		radix := 10
		if j, ok := args[0].(*Number); ok {
			radix = int(j.Data)
		}
		if radix == 10 {
			return GoToString(n.ToString())
		}
		return GoToString(floatToString(n.Data, radix))
	})
}

func numberToExponential(n *Number) JavaScript {
	return NewDefaultFunction("toExponential", func(args []JavaScript, ctx *Context) JavaScript {
		if math.IsInf(n.Data, 1) {
			return GoToString("Infinity")
		}
		if math.IsInf(n.Data, -1) {
			return GoToString("-Infinity")
		}
		if math.IsNaN(n.Data) {
			return GoToString("NaN")
		}
		d := 0
		if len(args) > 0 {
			if j, ok := args[0].(*Number); ok {
				d = int(j.Data)
			}
		}
		if d < 0 || d > 100 {
			panic(NewRangeError("toExponential() digits argument must be between 0 and 100"))
		}
		format := fmt.Sprintf("%%.%de", d)
		if len(args) == 0 {
			format = "%e"
		}
		formatted := fmt.Sprintf(format, n.Data)
		return GoToString(formatted)
	})
}

func isFinite() JavaScript {
	return NewDefaultFunction("isFinite", func(args []JavaScript, ctx *Context) JavaScript {
		if len(args) == 0 {
			return GoToBoolean(false)
		}
		if d, ok := args[0].(*Number); ok {
			if math.IsInf(d.Data, 1) || math.IsInf(d.Data, -1) || math.IsNaN(d.Data) {
				return GoToBoolean(false)
			}
			return GoToBoolean(true)
		} else {
			return GoToBoolean(false)
		}
	})
}

func isInteger() JavaScript {
	return NewDefaultFunction("isInteger", func(args []JavaScript, ctx *Context) JavaScript {
		if len(args) == 0 {
			return GoToBoolean(false)
		}
		if d, ok := args[0].(*Number); ok {
			if math.IsInf(d.Data, 1) || math.IsInf(d.Data, -1) || math.IsNaN(d.Data) {
				return GoToBoolean(false)
			}
			return GoToBoolean(d.Data == math.Floor(d.Data))
		} else {
			return GoToBoolean(false)
		}
	})
}

func isNaN1() JavaScript {
	return NewDefaultFunction("isNaN", func(args []JavaScript, ctx *Context) JavaScript {
		if len(args) == 0 {
			return GoToBoolean(false)
		}
		if d, ok := args[0].(*Number); ok {
			return GoToBoolean(math.IsNaN(d.Data))
		} else {
			return GoToBoolean(false)
		}
	})
}

func NewParseFloat() JavaScript {
	c := func(args []JavaScript, r *Context) JavaScript {
		if len(args) == 0 {
			return NewUndefined()
		}
		str := strings.TrimLeftFunc(JsToString(args[0]), unicode.IsSpace)
		return GoToNumber(getLeadingFloat(str))
	}
	return NewDefaultFunction("parseFloat", c)
}

func getLeadingFloat(s string) string {
	var result []rune
	for _, r := range s {
		if unicode.IsDigit(r) || '.' == r {
			result = append(result, r)
		} else {
			break
		}
	}
	return string(result)
}
func isSafeInteger() JavaScript {
	return NewDefaultFunction("isSafeInteger", func(args []JavaScript, ctx *Context) JavaScript {
		if len(args) == 0 {
			return GoToBoolean(false)
		}
		if d, ok := args[0].(*Number); ok {
			const MaxInt53 = 1<<53 - 1
			return GoToBoolean(d.Data == math.Floor(d.Data) && d.Data >= -MaxInt53 && d.Data <= MaxInt53)
		} else {
			return GoToBoolean(false)
		}
	})
}

func NewParseInt() JavaScript {
	c := func(args []JavaScript, r *Context) JavaScript {
		if len(args) == 0 {
			return NewUndefined()
		}
		base := 10
		if len(args) > 1 {
			if a, ok := args[1].(*Number); ok {
				base = int(a.Data)
			} else {
				return NewNaN()
			}
		}
		s := JsToString(args[0])
		if strings.HasPrefix(s, "0x") {
			if len(args) == 1 {
				base = 16
			}
			f, e := strconv.ParseInt(s[2:], base, 64)
			if e != nil {
				return NewNaN()
			}
			return GoToNumber(float64(f))
		}
		str := getLeadingDigits(strings.TrimLeftFunc(s, unicode.IsSpace), base)
		if str == "" {
			return NewNaN()
		}
		f, e := strconv.ParseInt(str, base, 64)
		if e != nil {
			return NewNaN()
		}
		return GoToNumber(float64(f))
	}
	return NewDefaultFunction("parseInt", c)
}

func getLeadingDigits(s string, base int) string {
	var result []rune
	for _, r := range s {
		if !unicode.IsDigit(r) && !('a' <= r && r <= 'f') && !('A' <= r && r <= 'F') {
			break
		}
		switch base {
		case 2:
			if r == '0' || r == '1' {
				result = append(result, r)
			} else {
				break
			}
		case 8:
			if '0' <= r && r <= '7' {
				result = append(result, r)
			} else {
				break
			}
		case 10:
			if unicode.IsDigit(r) {
				result = append(result, r)
			} else {
				break
			}
		case 16:
			if unicode.IsDigit(r) || ('a' <= r && r <= 'f') || ('A' <= r && r <= 'F') {
				result = append(result, r)
			} else {
				break
			}
		}
		if base != 2 && base != 8 && base != 10 && base != 16 {
			break
		}
	}
	return string(result)
}

func numberToFixed(n *Number) JavaScript {
	return NewDefaultFunction("toFixed", func(args []JavaScript, ctx *Context) JavaScript {
		d := 0
		if len(args) > 0 {
			if j, ok := args[0].(*Number); ok {
				d = int(j.Data)
			}
		}
		if d < 0 || d > 100 {
			panic(NewRangeError("toFixed() digits argument must be between 0 and 100"))
		}
		formatted := fmt.Sprintf("%.*f", d, n.Data)
		return GoToString(formatted)
	})
}

func numberToPrecision(n *Number) JavaScript {
	return NewDefaultFunction("toPrecision", func(args []JavaScript, ctx *Context) JavaScript {
		d := 0
		if len(args) > 0 {
			if j, ok := args[0].(*Number); ok {
				d = int(j.Data)
			}
		} else {
			return GoToString(n.ToString())
		}
		if d < 0 || d > 100 {
			panic(NewRangeError("toPrecision() digits argument must be between 0 and 100"))
		}
		// 特殊情况处理：0 和非数值
		if n.Data == 0 {
			return GoToString("0." + strings.Repeat("0", d-1))
		}
		order := int(math.Log10(math.Abs(n.Data)))
		var formatStr string
		if order >= d || order < -4 {
			// 使用科学记数法格式
			formatStr = fmt.Sprintf("%%.%de", d-1)
		} else {
			// 使用固定小数点格式
			decimalPlaces := d - order - 1
			if n.Data < 1 {
				decimalPlaces = d - order
			}
			if decimalPlaces < 0 {
				decimalPlaces = 0
			}
			formatStr = fmt.Sprintf("%%.%df", decimalPlaces)
		}
		return GoToString(fmt.Sprintf(formatStr, n.Data))
	})
}

/*****************JavaScript接口实现******************/

func (n *Number) GetName() string {
	return "number"
}

func (n *Number) GetProperty(name string) JavaScript {
	switch name {
	case "toFixed":
		return numberToFixed(n)
	case "toPrecision":
		return numberToPrecision(n)
	case "toString":
		return numberToStr(n)
	case "toExponential":
		return numberToExponential(n)
	default:
		return NewUndefined()
	}

}

func (n *Number) GetProperties() JavaScriptProperties {
	return collection.NewArrayMap[string, JavaScript]()
}

func (n *Number) ToString() string {
	if math.IsInf(n.Data, 1) {
		return "Infinity"
	}
	if math.IsInf(n.Data, -1) {
		return "-Infinity"
	}
	if math.IsNaN(n.Data) {
		return "NaN"
	}
	//打印
	var formatStr string
	intPartLength := int(math.Log10(math.Abs(n.Data))) + 1
	if math.Abs(n.Data) < 1e-6 {
		//当绝对值小于1e-6（不包括1e-6本身）时，也会使用科学计数法表示。
		formatStr = "%v"
	} else if n.Data > 0 {
		//对于正数
		if n.Data >= 1e+21 {
			//数字大于或等于1e+21时，JavaScript会自动使用科学计数法表示。
			formatStr = "%v"
		} else {
			//正常显示
			if intPartLength < 17 {
				x := getFractionalPartDigits(n.Data)
				if x > 0 {
					formatStr = fmt.Sprintf("%%.%df", x)
				} else {
					formatStr = "%" + strconv.Itoa(intPartLength) + ".0f"
				}
			} else {
				formatStr = "%" + strconv.Itoa(intPartLength) + ".0f"
			}
		}
	} else {
		//对于正数
		if n.Data <= -1e+21 {
			//小于或等于-1e+21的值，JavaScript会自动使用科学计数法表示。
			formatStr = "%v"
		} else {
			//正常显示
			if intPartLength < 17 {
				x := getFractionalPartDigits(n.Data)
				if x > 0 {
					formatStr = fmt.Sprintf("%%.%df", x)
				} else {
					formatStr = "%" + strconv.Itoa(intPartLength) + ".0f"
				}
			} else {
				formatStr = "%" + strconv.Itoa(intPartLength) + ".0f"
			}
		}
	}
	return fmt.Sprintf(formatStr, n.Data)
}

func getFractionalPartDigits(value float64) int {
	// 将float64转换为字符串
	strValue := strconv.FormatFloat(value, 'f', -1, 64)
	// 分割整数部分和小数部分
	parts := strings.Split(strValue, ".")
	// 如果没有小数部分，返回0
	if len(parts) == 1 {
		return 0
	}
	// 返回小数部分的长度
	return len(parts[1])
}

func (n *Number) UnmarshalBin(data []byte) (any, error) {
	var f float64
	str := string(data)
	switch str {
	case `"+Inf"`:
		f = math.Inf(1)
	case `"-Inf"`:
		f = math.Inf(-1)
	case `"NaN"`:
		f = math.NaN()
	default:
		// 对于正常的float64值，直接解析
		value, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return nil, err
		}
		f = value
	}
	n.Data = f
	return nil, nil
}

func (n *Number) MarshalBin() ([]byte, error) {
	// 检查是否为正负无穷大或NaN
	var da []byte
	if math.IsInf(n.Data, 1) {
		da = []byte(`"+Inf"`)
	} else if math.IsInf(n.Data, -1) {
		da = []byte(`"-Inf"`)
	} else if math.IsNaN(n.Data) {
		da = []byte(`"NaN"`)
	} else {
		da = []byte(strconv.FormatFloat(n.Data, 'f', -1, 64))
	}
	// 对于正常的float64值，使用标准格式
	data := []byte{5}
	data = append(data, da...)
	return data, nil
}

func (n *Number) ToBool() bool {
	if IsNaN(n) {
		return false
	}
	return n.Data != 0
}

func (n *Number) ToNumber() float64 {
	return n.Data
}
func (n *Number) Print() string {
	if math.IsInf(n.Data, 1) {
		return YellowStyle + "Infinity" + EndStyle
	}
	if math.IsInf(n.Data, -1) {
		return YellowStyle + "-Infinity" + EndStyle
	}
	if math.IsNaN(n.Data) {
		return YellowStyle + "NaN" + EndStyle
	}
	return YellowStyle + n.ToString() + EndStyle
}

func (n *Number) Value() any {
	return n.Data
}

func (n *Number) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &n.Data)
}
func (n *Number) MarshalJSON() ([]byte, error) {
	if math.IsInf(n.Data, +1) || math.IsInf(n.Data, -1) || math.IsNaN(n.Data) {
		return []byte(`null`), nil
	}
	return json.Marshal(n.Data)
}

func (n *Number) Binary() []byte {
	var buf bytes.Buffer
	err := binary.Write(&buf, binary.LittleEndian, n.Data)
	if err != nil {
		panic(NewError("Error", err.Error()))
	}
	return buf.Bytes()
}
