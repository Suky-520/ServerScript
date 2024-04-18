// Array这是ss中用来表示JavaScript数组的结构体
//这是一个JavaScript引用类型

package ss

import (
	"bytes"
	"encoding/json"
	"math"
	"sort"
	"ss/collection"
	"strconv"
)

// Array JavaScript数组
type Array struct {
	Data []JavaScript
}

// ArrayIterator 数组迭代器
type ArrayIterator struct {
	data  *Array
	index int
}

// 实现JavaScript接口
var _ JavaScript = &Array{}

// 实现JavaScriptBinary接口
var _ JavaScriptBinary = &Array{}

// 实现JavaScriptIterator接口
var _ JavaScriptIterator = &ArrayIterator{}

func (a *Array) GetName() string {
	return "array"
}
func (a *Array) GetIterator() JavaScriptIterator {
	return &ArrayIterator{index: 0, data: a}
}
func (a *Array) GetProperty(name string) JavaScript {
	switch name {
	case "length":
		return GoToNumber(float64(len(a.Data)))
	case "join":
		return arrJoin(a)
	case "push":
		return arrPush(a)
	case "pop":
		return arrPop(a)
	case "shift":
		return arrShift(a)
	case "unshift":
		return arrUnshift(a)
	case "slice":
		return arrSlice(a)
	case "splice":
		return arrSplice(a)
	case "indexOf":
		return arrIndexOf(a)
	case "lastIndexOf":
		return arrLastIndexOf(a)
	case "filter":
		return arrFilter(a)
	case "sort":
		return arrSort(a)
	case "map":
		return arrMap(a)
	default:
		return arrGetValue(a, name)
	}
}
func (a *Array) GetProperties() JavaScriptProperties {
	return collection.NewArrayMap[string, JavaScript]()
}
func (a *Array) ToString() string {
	s := bytes.Buffer{}
	for i, val := range a.Data {
		str := ""
		if val != nil {
			str = JsToString(val)
		}
		if i > 0 {
			s.WriteString(",")
		}
		s.WriteString(str)
	}
	return s.String()
}
func (a *Array) ToNumber() float64 {
	if len(a.Data) == 0 {
		return 0
	}
	if len(a.Data) > 1 {
		return math.NaN()
	}
	if v, ok := a.Data[0].(*Number); ok {
		return v.Data
	}
	if v, ok := a.Data[0].(*String); ok {
		return v.ToNumber()
	}
	return math.NaN()
}
func (a *Array) SetProperty(name string, prop JavaScript) {
	i, e := strconv.Atoi(name)
	if e != nil {
		panic(NewTypeError("'" + name + "' is not a number"))
	}
	if len(a.Data) <= i {
		newSlice := make([]JavaScript, i+1)
		copy(newSlice, a.Data)
		a.Data = newSlice
	}
	a.Data[i] = prop
}
func (a *Array) DeleteProperty(name string) JavaScript {
	num, err := strconv.Atoi(name)
	if err != nil {
		panic(NewSyntaxError("'" + name + "' is not a number'"))
	}
	if len(a.Data) > num {
		a.Data[num] = NewUndefined()
		return GoToBoolean(true)
	} else {
		return GoToBoolean(false)
	}
}
func (a *Array) Binary() []byte {
	var arr []byte
	for _, v := range a.Data {
		if b, ok := v.(JavaScriptBinary); ok {
			arr = append(arr, b.Binary()...)
		}
	}
	return arr
}
func (a *Array) MarshalJSON() ([]byte, error) {
	return json.Marshal(a.Data)
}
func (a *Array) UnmarshalJSON(data []byte) error {
	arr := make([]any, 0)
	err := json.Unmarshal(data, &arr)
	if err != nil {
		return err
	}
	a.Data = GoToArray(arr).Data
	return nil
}

/**************实现迭代器接口******************/

func (a *ArrayIterator) Next() (v []JavaScript, done bool) {
	defer func() {
		a.index++
	}()
	if a.index >= len(a.data.Data) {
		return []JavaScript{NewUndefined(), NewUndefined()}, true
	}
	return []JavaScript{GoToNumber(a.index), a.data.Data[a.index]}, false
}
func (a *ArrayIterator) GetName() string {
	return "array.iterator"
}

// GoToArray 将go切片转换为JavaScript数组对象
func GoToArray[T []JavaScript | []string | []any](items T) *Array {
	if items == nil {
		a := &Array{Data: make([]JavaScript, 0)}
		return a
	}
	var a any = items
	switch t := a.(type) {
	case []JavaScript:
		//深度复制
		copiedArray := make([]JavaScript, len(t))
		copy(copiedArray, t)
		o := &Array{
			Data: copiedArray,
		}
		return o
	case []string:
		return GoToJs(items).(*Array)
	case []any:
		return GoToJs(items).(*Array)
	default:
		arr := &Array{Data: make([]JavaScript, 0)}
		return arr
	}
}

func arrGetValue(a *Array, name string) JavaScript {
	num, err := strconv.Atoi(name)
	if err != nil {
		panic(NewSyntaxError("'" + name + "' is not a number'"))
	}
	if len(a.Data) > num {
		return a.Data[num]
	} else {
		return NewUndefined()
	}
}

func arrJoin(a *Array) JavaScript {
	return NewDefaultFunction("join", func(args []JavaScript, ctx *Context) JavaScript {
		sp := ","
		if len(args) > 0 {
			sp = JsToString(args[0])
		}
		str := ""
		for _, val := range a.Data {
			if str != "" {
				str = str + sp + JsToString(val)
			} else {
				str = JsToString(val)
			}
		}
		return GoToString(str)
	})
}

func arrPush(a *Array) JavaScript {
	return NewDefaultFunction("push", func(args []JavaScript, ctx *Context) JavaScript {
		if len(args) == 0 {
			return GoToNumber(float64(len(a.Data)))
		}
		a.Data = append(a.Data, args...)
		return GoToNumber(float64(len(a.Data)))
	})
}

func arrPop(a *Array) JavaScript {
	return NewDefaultFunction("pop", func(args []JavaScript, ctx *Context) JavaScript {
		if len(a.Data) == 0 {
			return NewUndefined()
		}
		v := a.Data[len(a.Data)-1]
		a.Data = a.Data[:len(a.Data)-1]
		return v
	})
}

func arrUnshift(a *Array) JavaScript {
	return NewDefaultFunction("unshift", func(args []JavaScript, ctx *Context) JavaScript {
		if len(args) == 0 {
			return GoToNumber(float64(len(a.Data)))
		}
		a.Data = append(args, a.Data...)
		return GoToNumber(float64(len(a.Data)))
	})
}

func arrShift(a *Array) JavaScript {
	return NewDefaultFunction("shift", func(args []JavaScript, ctx *Context) JavaScript {
		if len(a.Data) == 0 {
			return NewUndefined()
		}
		v := a.Data[0]
		a.Data = a.Data[1:]
		return v
	})
}

func arrSlice(a *Array) JavaScript {
	return NewDefaultFunction("slice", func(args []JavaScript, ctx *Context) JavaScript {
		if len(args) == 0 {
			return GoToArray(a.Data[:])
		}
		if len(args) == 1 {
			if i, ok := args[0].(*Number); ok {
				if len(a.Data) <= int(i.Data) {
					return GoToArray(make([]JavaScript, 0))
				} else {
					if i.Data > 0 {
						return GoToArray(a.Data[int(i.Data):])
					} else {
						l := len(a.Data) + int(i.Data)
						if l < 0 {
							return GoToArray(a.Data[:])
						} else {
							return GoToArray(a.Data[l:])
						}
					}
				}
			} else {
				return GoToArray(a.Data[:])
			}
		}
		s := 0
		e := 0
		if i, ok := args[0].(*Number); ok {
			s = int(i.Data)
		} else {
			return GoToArray(a.Data[:])
		}
		if i, ok := args[1].(*Number); ok {
			e = int(i.Data)
		} else {
			return GoToArray(a.Data[:])
		}
		if s < 0 {
			s = len(a.Data) + s
		}
		if e < 0 {
			e = len(a.Data) + e
		}
		if s > len(a.Data) {
			s = len(a.Data)
		}
		if e > len(a.Data) {
			e = len(a.Data)
		}
		return GoToArray(a.Data[s:e])
	})
}

func arrGetIndex(a *Array, s int) int {
	if s < 0 {
		s = s + len(a.Data)
	}
	if s < 0 {
		s = 0
	}
	if s > len(a.Data) {
		s = len(a.Data)
	}
	return s
}

func arrSplice(a *Array) JavaScript {
	return NewDefaultFunction("splice", func(args []JavaScript, ctx *Context) JavaScript {
		if len(a.Data) == 0 {
			return GoToArray(make([]JavaScript, 0))
		}
		s := 0
		e := 0
		if len(args) == 1 {
			if i, ok := args[0].(*Number); ok {
				s = arrGetIndex(a, int(i.Data))
			} else {
				//非整数,删除所有元素
				d := a.Data
				a.Data = make([]JavaScript, 0)
				return GoToArray(d)
			}
			d := GoToArray(a.Data[s:])
			a.Data = a.Data[:s]
			return d
		}
		if i, ok := args[0].(*Number); ok {
			s = arrGetIndex(a, int(i.Data))
		} else {
			//非整数,删除所有元素
			d := a.Data
			a.Data = make([]JavaScript, 0)
			return GoToArray(d)
		}

		if i, ok := args[1].(*Number); ok {
			if i.Data < 0 {
				e = 0
			} else {
				e = int(i.Data)
			}
		} else {
			//非整数,删除所有元素
			d := a.Data
			a.Data = make([]JavaScript, 0)
			return GoToArray(d)
		}
		if s+e > len(a.Data) {
			e = len(a.Data)
		} else {
			e = s + e
		}
		d := GoToArray(a.Data[s:e])
		l := make([]JavaScript, len(a.Data[e:]))
		copy(l, a.Data[e:])
		a.Data = append(a.Data[:s], args[2:]...)
		a.Data = append(a.Data, l...)
		return d
	})
}

func arrIndexOf(a *Array) JavaScript {
	return NewDefaultFunction("indexOf", func(args []JavaScript, ctx *Context) JavaScript {
		if len(args) == 0 {
			return GoToNumber(-1)
		}
		for i, v := range a.Data {
			if a1, ok := v.(JavaScriptBasic); ok {
				if a2, ok := args[0].(JavaScriptBasic); ok {
					if a1.Value() == a2.Value() {
						return GoToNumber(float64(i))
					}
				}
			}
			if v == args[0] {
				return GoToNumber(float64(i))
			}
		}
		return GoToNumber(-1)
	})
}

func arrLastIndexOf(a *Array) JavaScript {
	return NewDefaultFunction("lastIndexOf", func(args []JavaScript, ctx *Context) JavaScript {
		if len(args) == 0 {
			return GoToNumber(-1)
		}
		for i := len(a.Data) - 1; i >= 0; i-- {
			if a1, ok := a.Data[i].(JavaScriptBasic); ok {
				if a2, ok := args[0].(JavaScriptBasic); ok {
					if a1.Value() == a2.Value() {
						return GoToNumber(float64(i))
					}
				}
			}
			if a.Data[i] == args[0] {
				return GoToNumber(float64(i))
			}
		}
		return GoToNumber(-1)
	})
}

func arrFilter(a *Array) JavaScript {
	return NewDefaultFunction("filter", func(args []JavaScript, ctx *Context) JavaScript {
		if len(args) == 0 {
			panic(NewTypeError("undefined is not a function"))
			return NewUndefined()
		}
		if f, ok := args[0].(*Function); ok {
			arr := make([]JavaScript, 0)
			for i, v := range a.Data {
				r := ctx.RunFunction(f, v, GoToNumber(float64(i)), a)
				if b, ok := r.(*Boolean); ok && b.Data {
					arr = append(arr, v)
				}
			}
			return GoToArray(arr)
		} else {
			panic(NewTypeError(JsToString(args[0]) + " is not a function"))
			return NewUndefined()
		}
	})
}

func arrSort(a *Array) JavaScript {
	return NewDefaultFunction("sort", func(args []JavaScript, ctx *Context) JavaScript {
		// 检查是否有比较函数作为参数传入
		var compareFunc func(i, j int) bool
		if len(args) > 0 {
			if cmp, ok := args[0].(*Function); ok {
				compareFunc = func(i, j int) bool {
					// 调用比较函数，期望返回一个布尔值
					result := ctx.RunFunction(cmp, a.Data[i], a.Data[j])
					if res, ok := result.(*Number); ok {
						return res.Data < 0
					}
					return false // 或者处理错误
				}
			}
		}
		if compareFunc == nil {
			// 如果没有提供比较函数，则使用默认的字典序排序
			compareFunc = func(i, j int) bool {
				return JsToString(a.Data[i]) < JsToString(a.Data[j])
			}
		}
		sort.SliceStable(a.Data, compareFunc)
		return a
	})
}

func arrMap(a *Array) JavaScript {
	return NewDefaultFunction("map", func(args []JavaScript, ctx *Context) JavaScript {
		if len(args) == 0 {
			panic(NewTypeError("undefined is not a function"))
			return NewUndefined()
		}
		if f, ok := args[0].(*Function); ok {
			arr := make([]JavaScript, 0)
			for i, v := range a.Data {
				r := ctx.RunFunction(f, v, GoToNumber(float64(i)), a)
				arr = append(arr, r)
			}
			return GoToArray(arr)
		} else {
			panic(NewTypeError(JsToString(args[0]) + " is not a function"))
			return NewUndefined()
		}
	})
}
