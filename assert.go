//assert对象,用于单元测试

package ss

import (
	"fmt"
	"math/big"
)

// assert内置对象，用于进行测试
func init() {
	obj := NewDefaultObject("assert")
	obj.AddProperty("ok", NewDefaultFunction("ok", okFunc()))
	obj.AddProperty("deepEqual", deepEqual())
	obj.AddProperty("equal", equal())
	obj.AddProperty("throws", throws())
	mod := NewModule()
	mod.SetExport("default", obj)
	RegisterModule("assert", mod)
}

func okFunc() func(args []JavaScript, ctx *Context) JavaScript {
	return func(args []JavaScript, ctx *Context) JavaScript {
		if len(args) == 0 {
			JsError(fmt.Sprint("AssertionError", "assert.ok() requires one argument"), ctx)
			return NewUndefined()
		}
		if JsToBool(args[0]) {
			return NewUndefined()
		}
		if len(args) > 1 {
			JsError(fmt.Sprint("AssertionError", JsToString(args[1])), ctx)
			return NewUndefined()
		}
		JsError(fmt.Sprint("AssertionError", JsToString(args[0])+"!=true"), ctx)
		return NewUndefined()
	}
}

func deepEqual() JavaScript {
	return NewDefaultFunction("deepEqual", func(args []JavaScript, ctx *Context) JavaScript {
		if len(args) < 2 {
			JsError(fmt.Sprint("AssertionError", "assert.deepEqual() requires two argument"), ctx)
			return NewUndefined()
		}

		if JsToJson(args[0]) == JsToJson(args[1]) {
			return NewUndefined()
		}
		if len(args) > 2 {
			JsError(fmt.Sprint("AssertionError", JsToString(args[2])), ctx)
			return NewUndefined()
		}
		JsError(fmt.Sprint("AssertionError", JsToPrint(args[0], "")+" not deepEqual "+JsToPrint(args[1], "")), ctx)
		return NewUndefined()
	})
}

func equal() JavaScript {
	return NewDefaultFunction("equal", func(args []JavaScript, ctx *Context) JavaScript {
		if len(args) < 2 {
			JsError(fmt.Sprint("AssertionError", "assert.equal() requires two argument"), ctx)
			return NewUndefined()
		}
		v1 := args[0]
		v2 := args[1]
		v1Type := JsToType(v1)
		v2Type := JsToType(v2)
		if v1Type.Type == v1Type.Type {
			if v1Type.Type == NumberType && v2Type.Type == NumberType {
				if IsNaN(v1) && IsNaN(v2) {
					return NewUndefined()
				}
			}
			if v1Type.Type == BigIntType && v2Type.Type == BigIntType {
				result := v1.(JavaScriptBasic).Value().(*big.Int).Cmp(v2.(JavaScriptBasic).Value().(*big.Int))
				if result == 0 {
					return NewUndefined()
				}
			}
			if b1, ok := v1.(JavaScriptBasic); ok {
				if b2, ok := v2.(JavaScriptBasic); ok {
					if b1.Value() == b2.Value() {
						return NewUndefined()
					}
				}
			}
		}
		if len(args) > 2 {
			JsError(fmt.Sprint("AssertionError", JsToString(args[2])), ctx)
			return NewUndefined()
		}
		JsError(fmt.Sprint("AssertionError", JsToPrint(args[0], "")+" not equal "+JsToPrint(args[1], "")), ctx)
		return NewUndefined()
	})
}

func throws() JavaScript {
	return NewDefaultFunction("throws", func(args []JavaScript, ctx *Context) JavaScript {
		defer func() {
			if err := recover(); err != nil {
				var name = ""
				//if v, ok := args[1].(*ss.DefaultJavaScript); ok {
				//	name = v.name
				//}
				//
				//if v, ok := err.(ss.Error); ok && v.name == name {
				//	return
				//}
				if name == "" && fmt.Sprint(err) == JsToString(args[1]) {
					return
				}
				if len(args) > 2 {
					JsError(fmt.Sprint("AssertionError", JsToString(args[2])), ctx)
					return
				}
				JsError(fmt.Sprint("AssertionError", JsToString(args[0])+" not throws "+JsToString(args[1])), ctx)
			}
		}()
		if len(args) < 2 {
			JsError(fmt.Sprint("AssertionError", "assert.throws() requires two argument"), ctx)
			return NewUndefined()
		}
		if f, ok := args[0].(*Function); ok {
			ctx.RunFunction(f)
		} else {
			JsError(fmt.Sprint("AssertionError", "参数类型错误"), ctx)
			return NewUndefined()
		}
		if len(args) > 2 {
			JsError(fmt.Sprint("AssertionError", JsToString(args[2])), ctx)
			return NewUndefined()
		}
		JsError(fmt.Sprint("AssertionError", JsToString(args[0])+" not throws "+JsToString(args[1])), ctx)
		return NewUndefined()
	})
}
