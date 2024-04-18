//JavaScript中的Set类

package ss

import (
	"fmt"
	"ss/collection"
)

// Set 对象
type Set struct {
	Data *collection.ArrayMap[any, int]
}

// SetIterator 迭代器
type SetIterator struct {
	index int
	keys  []any
}

var _ JavaScript = &Set{}

func init() {
	c := func(args []JavaScript, ctx *Context) JavaScript {
		return NewSet()
	}
	setGlobalObject("Set", NewDefaultClass("Set", c))
}

func NewSet() *Set {
	ma := &Set{Data: collection.NewArrayMap[any, int]()}
	return ma
}

func setAdd(m *Set) JavaScript {
	return NewDefaultFunction("add", func(args []JavaScript, ctx *Context) JavaScript {
		if len(args) < 1 {
			return NewUndefined()
		}
		if v, ok := JsToBasic(args[0]); ok {
			m.Data.Set(v, 1)
		} else {
			m.Data.Set(args[0], 1)
		}
		return m
	})
}

func setDelete(m *Set) JavaScript {
	return NewDefaultFunction("delete", func(args []JavaScript, ctx *Context) JavaScript {
		if len(args) == 0 {
			return NewUndefined()
		}
		if v, ok := JsToBasic(args[0]); ok {
			return GoToBoolean(m.Data.Delete(v))
		} else {
			return GoToBoolean(m.Data.Delete(args[0]))
		}
	})
}

func setClear(m *Set) JavaScript {
	return NewDefaultFunction("clear", func(args []JavaScript, ctx *Context) JavaScript {
		m.Data.Clear()
		return NewUndefined()
	})
}

func setHas(m *Set) JavaScript {
	return NewDefaultFunction("get", func(args []JavaScript, ctx *Context) JavaScript {
		if len(args) == 0 {
			return NewUndefined()
		}
		if _, ok := args[0].(*Undefined); ok {
			return GoToBoolean(m.Data.Has(Undefined{}))
		}
		if v1, ok := args[0].(JavaScriptBasic); ok {
			return GoToBoolean(m.Data.Has(v1.Value()))
		}
		return GoToBoolean(m.Data.Has(args[0]))
	})
}

func (m *Set) GetProperty(name string) JavaScript {
	switch name {
	case "size":
		return GoToNumber(float64(m.Data.Size()))
	case "add":
		return setAdd(m)
	case "has":
		return setHas(m)
	case "delete":
		return setDelete(m)
	case "clear":
		return setClear(m)
	default:
		return NewUndefined()
	}
}

func (m *Set) Print() string {
	return fmt.Sprintf("Set(%v)", m.Data.Size())
}

func (m *Set) GetName() string {
	return "set"
}
func (m *Set) GetIterator() JavaScriptIterator {
	return &SetIterator{index: 0, keys: m.Data.Keys()}
}

/**************实现迭代器接口******************/

func (a *SetIterator) Next() (v []JavaScript, done bool) {
	defer func() {
		a.index++
	}()
	if a.index >= len(a.keys) {
		return []JavaScript{NewUndefined()}, true
	}
	return []JavaScript{GoToJs(a.keys[a.index])}, false
}

func (a *SetIterator) GetName() string {
	return "set.iterator"
}
