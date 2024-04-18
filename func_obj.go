// JavaScript中的内置函数和对象

package ss

import (
	"encoding/json"
	"fmt"
	"github.com/gammazero/workerpool"
	"github.com/google/uuid"
	"math"
	"math/rand"
	"net/url"
	"os"
	"runtime"
	"strings"
	"time"
)

func init() {
	setGlobalObject("eval", NewDefaultFunction("eval", eval))
	setGlobalObject("isNaN", NewDefaultFunction("isNaN", isNaN))
	setGlobalObject("parseFloat", NewParseFloat())
	setGlobalObject("parseInt", NewParseInt())
	setGlobalObject("sleep", NewDefaultFunction("sleep", sleep))
	setGlobalObject("setInterval", NewDefaultFunction("setInterval", setInterval))
	setGlobalObject("setTimeout", NewDefaultFunction("setTimeout", setTimeout))
	setGlobalObject("clearInterval", NewDefaultFunction("clearInterval", clearInterval))
	setGlobalObject("encodeURI", encodeURI)
	setGlobalObject("decodeURI", decodeURI)
	setGlobalObject("typeOf", NewDefaultFunction("typeOf", typeOf))
	setGlobalObject("wait", NewDefaultFunction("wait", wait))
	setGlobalObject("Math", initMath())
	setGlobalObject("JSON", initJSON())
	setGlobalObject("Object", initObject())
	setGlobalObject("console", initConsole())
}

var (
	fx            = []string{"sqrt", "abs", "acos", "asin", "atan", "ceil", "cos", "exp", "floor", "log", "round", "sin", "tan", "tanh", "trunc"}
	timeIds       = make(map[string]bool)
	defThreadPool = workerpool.New(math.MaxInt)
	_version      = "0.0.1"
	eval          = func(args []JavaScript, r *Context) JavaScript {
		if len(args) == 0 {
			return NewUndefined()
		}
		if js, ok := args[0].(*String); ok {
			vm1 := NewVM()
			vm1.RunJsSource(js.ToString(), "eval")
			if len(vm1.GetValue()) > 0 {
				return vm1.GetValue()[0]
			} else {
				return NewUndefined()
			}
		}
		return NewUndefined()
	}
	isNaN = func(args []JavaScript, r *Context) JavaScript {
		if len(args) == 0 {
			return GoToBoolean(true)
		}
		return GoToBoolean(math.IsNaN(JsToNumber(args[0])))
	}
	sleep = func(args []JavaScript, r *Context) JavaScript {
		if len(args) == 0 {
			select {}
		}
		if JsToNumber(args[0]) <= 0 {
			select {}
		} else {
			time.Sleep(time.Millisecond * time.Duration(JsToNumber(args[0])))
		}
		return NewUndefined()
	}
	setInterval = func(args []JavaScript, r *Context) JavaScript {
		if len(args) < 2 {
			panic(NewTypeError("参数类型错误"))
		}
		if v, ok := args[0].(*Function); ok {
			if t, ok := args[1].(*Number); ok {
				id := uuid.New().String()
				timeIds[id] = true
				AsyncCall(func() {
					for {
						time.Sleep(time.Millisecond * time.Duration(t.Data))
						if timeIds[id] {
							r.RunFunction(v)
						} else {
							break
						}
					}
					delete(timeIds, id)
				})
				return GoToString(id)
			}
		}
		panic(NewTypeError("参数类型错误"))
	}
	setTimeout = func(args []JavaScript, r *Context) JavaScript {
		if len(args) < 2 {
			panic(NewTypeError("参数类型错误"))
		}
		if v, ok := args[0].(*Function); ok {
			if t, ok := args[1].(*Number); ok {
				id := uuid.New().String()
				timeIds[id] = true
				AsyncCall(func() {
					time.Sleep(time.Millisecond * time.Duration(t.Data))
					if timeIds[id] {
						r.RunFunction(v)
					}
					delete(timeIds, id)
				})
				return GoToString(id)
			}
		}
		panic(NewTypeError("参数类型错误"))
	}
	clearInterval = func(args []JavaScript, r *Context) JavaScript {
		if len(args) == 0 {
			panic(NewTypeError("参数类型错误"))
		}
		timeIds[JsToString(args[0])] = false
		return NewUndefined()
	}
	decodeURI = func(args []JavaScript, r *Context) JavaScript {
		if len(args) == 0 {
			return NewUndefined()
		}
		decodedURI, err := url.QueryUnescape(JsToString(args[0]))
		if err != nil {
			panic(err)
		}
		return GoToString(decodedURI)
	}

	encodeURI = func(args []JavaScript, r *Context) JavaScript {
		if len(args) == 0 {
			return NewUndefined()
		}
		encodedURI := url.QueryEscape(JsToString(args[0]))
		return GoToString(encodedURI)
	}
	typeOf = func(args []JavaScript, r *Context) JavaScript {
		if len(args) == 0 {
			return NewUndefined()
		}
		t := JsToType(args[0])
		obj := NewObject()
		obj.SetProperty("type", GoToNumber(uint8(t.Type)))
		obj.SetProperty("typeString", GoToString(types[t.Type-1]))
		obj.SetProperty("name", GoToString(t.Name))
		obj.SetProperty("native", GoToBoolean(t.Native))
		return obj
	}
	wait = func(args []JavaScript, r *Context) JavaScript {
		defThreadPool.StopWait()
		return NewUndefined()
	}
)

func initMath() JavaScript {
	c := NewDefaultObject("Math")
	c.AddProperty("PI", GoToJs(math.Pi))
	c.AddProperty("E", GoToJs(math.E))
	c.AddProperty("LN2", GoToJs(math.Ln2))
	c.AddProperty("LN10", GoToJs(math.Ln10))
	c.AddProperty("LOG2E", GoToJs(math.Log2E))
	c.AddProperty("LOG10E", GoToJs(math.Log10E))
	c.AddProperty("SQRT2", GoToJs(math.Sqrt2))
	c.AddProperty("SQRT1_2", GoToJs(math.Sqrt(1/2)))
	c.AddProperty("PI", GoToJs(math.Pi))
	c.AddProperty("random", NewDefaultFunction("random", func(args []JavaScript, ctx *Context) JavaScript {
		rand.New(rand.NewSource(time.Now().UnixNano()))
		return GoToJs(rand.Float64())
	}))
	c.AddProperty("pow", NewDefaultFunction("pow", func(args []JavaScript, ctx *Context) JavaScript {
		if len(args) < 2 {
			return NewNaN()
		}
		return GoToNumber(math.Pow(JsToNumber(args[0]), JsToNumber(args[1])))
	}))
	for _, name := range fx {
		c.AddProperty(name, mathFx(name))
	}
	return c
}

func mathFx(name string) JavaScript {
	return NewDefaultFunction(name, func(args []JavaScript, ctx *Context) JavaScript {
		if len(args) > 0 {
			v := JsToNumber(args[0])
			switch name {
			case "sqrt":
				return GoToJs(math.Sqrt(v))
			case "abs":
				return GoToJs(math.Abs(v))
			case "acos":
				return GoToJs(math.Acos(v))
			case "asin":
				return GoToJs(math.Asin(v))
			case "atan":
				return GoToJs(math.Atan(v))
			case "ceil":
				return GoToJs(math.Ceil(v))
			case "cos":
				return GoToJs(math.Cos(v))
			case "exp":
				return GoToJs(math.Exp(v))
			case "floor":
				return GoToJs(math.Floor(v))
			case "log":
				return GoToJs(math.Log(v))
			case "round":
				return GoToJs(math.Round(v))
			case "sin":
				return GoToJs(math.Sin(v))
			case "tan":
				return GoToJs(math.Tan(v))
			case "tanh":
				return GoToJs(math.Tanh(v))
			case "trunc":
				return GoToJs(math.Trunc(v))
			}
		}
		return NewNaN()
	})
}

func initJSON() JavaScript {
	c := NewDefaultObject("JSON")
	c.AddProperty("stringify", NewDefaultFunction("stringify", func(args []JavaScript, ctx *Context) JavaScript {
		if len(args) == 0 {
			return NewUndefined()
		}
		t := JsToType(args[0]).Type
		if t == UndefinedType || t == FunctionType {
			return NewUndefined()
		}
		if _, ok := args[0].(json.Marshaler); ok {
			return GoToString(JsToJson(args[0]))
		} else {
			return GoToString("{}")
		}
	}))
	c.AddProperty("parse", NewDefaultFunction("parse", func(args []JavaScript, ctx *Context) JavaScript {
		if len(args) == 0 {
			panic(NewSyntaxError(`"undefined" is not valid JSON`))
		}
		return JsonToJs(JsToString(args[0]))
	}))
	return c
}

func initObject() JavaScript {
	c := NewDefaultObject("Object")
	c.AddProperty("keys", objectKeys())
	c.AddProperty("values", objectValues())
	c.AddProperty("entries", objectEntries())
	return c
}
func objectKeys() JavaScript {
	return NewDefaultFunction("keys", func(args []JavaScript, ctx *Context) JavaScript {
		if len(args) == 0 {
			return NewUndefined()
		}
		keys := JsGetProperties(args[0]).Keys()
		return GoToJs(keys)
	})
}
func objectValues() JavaScript {
	return NewDefaultFunction("values", func(args []JavaScript, ctx *Context) JavaScript {
		if len(args) == 0 {
			return NewUndefined()
		}
		keys := JsGetProperties(args[0]).Keys()
		array := make([]JavaScript, len(keys))
		for i, k := range keys {
			array[i], _ = JsGetProperties(args[0]).Get(k)
		}
		return GoToArray(array)
	})
}

func objectEntries() JavaScript {
	return NewDefaultFunction("entries", func(args []JavaScript, ctx *Context) JavaScript {
		if len(args) == 0 {
			return NewUndefined()
		}
		keys := JsGetProperties(args[0]).Keys()
		array := make([]JavaScript, len(keys))
		for i, k := range keys {
			val := make([]JavaScript, 2)
			val[0] = GoToString(k)
			val[1], _ = JsGetProperties(args[0]).Get(k)
			array[i] = GoToArray(val)
		}
		return GoToArray(array)
	})
}

func initConsole() JavaScript {
	c := NewDefaultObject("console")
	c.AddProperty("log", NewDefaultFunction("log", func(args []JavaScript, ctx *Context) JavaScript {
		for _, v := range args {
			fmt.Print(JsToPrint(v, ""), " ")
		}
		fmt.Println()
		return NewUndefined()
	}))
	return c
}

// AsyncCall 异步执行函数
func AsyncCall(fun func()) {
	defThreadPool.Submit(fun)
}

// Process 对象实现
func initProcess() JavaScript {
	p := NewDefaultObject("process")
	p.AddProperty("exit", processExit())
	p.AddProperty("cwd", processCwd())
	p.AddProperty("memoryUsage", processMemoryUsage())
	p.AddProperty("pid", GoToNumber(float64(os.Getpid())))
	p.AddProperty("platform", GoToString(runtime.GOOS))
	p.AddProperty("version", GoToString(_version))
	p.AddProperty("versions", processVersions())
	p.AddProperty("env", processEnv())
	return p
}

func processExit() JavaScript {
	return NewDefaultFunction("exit", func(args []JavaScript, ctx *Context) JavaScript {
		if len(args) == 0 {
			os.Exit(0)
			return NewUndefined()
		}
		os.Exit(int(JsToNumber(args[0])))
		return NewUndefined()
	})
}

func processCwd() JavaScript {
	return NewDefaultFunction("cwd", func(args []JavaScript, ctx *Context) JavaScript {
		dir, err := os.Getwd()
		if err != nil {
			panic(NewError("Error", err.Error()))
			return NewUndefined()
		}
		return GoToString(dir)
	})
}
func processMemoryUsage() JavaScript {
	return NewDefaultFunction("memoryUsage", func(args []JavaScript, ctx *Context) JavaScript {
		var m runtime.MemStats
		runtime.ReadMemStats(&m)
		obj := NewObject()
		obj.SetProperty("alloc", GoToNumber(float64(m.Alloc)))
		obj.SetProperty("totalAlloc", GoToNumber(float64(m.TotalAlloc)))
		obj.SetProperty("sys", GoToNumber(float64(m.Sys)))
		obj.SetProperty("numGC", GoToNumber(float64(m.NumGC)))
		return obj
	})
}

func processVersions() JavaScript {
	o := NewObject()
	o.SetProperty("vino", GoToString(_version))
	o.SetProperty("ss", GoToString(fmt.Sprint(version[0], ".", version[1], ".", version[2])))
	return o
}
func processEnv() JavaScript {
	o := NewObject()
	for _, v := range os.Environ() {
		val := strings.Split(v, "=")
		o.SetProperty(val[0], GoToString(val[1]))
	}
	return o
}
