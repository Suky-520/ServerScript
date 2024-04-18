// JavaScript中的Map类。

package ss

import (
	"fmt"
	"ss/collection"
)

// Map 对象实现
type Map struct {
	Data *collection.ArrayMap[any, JavaScript]
}

// MapIterator 迭代器
type MapIterator struct {
	data  *Map
	index int
	keys  []any
}

var _ JavaScript = &Map{}

func init() {
	setGlobalObject("Map", initMap())
}

func initMap() JavaScript {
	return NewDefaultClass("Map", func(args []JavaScript, ctx *Context) JavaScript {
		return NewMap()
	})
}
func NewMap() *Map {
	ma := &Map{Data: collection.NewArrayMap[any, JavaScript]()}
	return ma
}

func mapSet(m *Map) JavaScript {
	return NewDefaultFunction("set", func(args []JavaScript, ctx *Context) JavaScript {
		if len(args) < 2 {
			return NewUndefined()
		}
		if v, ok := JsToBasic(args[0]); ok {
			m.Data.Set(v, args[1])
		} else {
			m.Data.Set(args[0], args[1])
		}
		return m
	})
}

func mapKeys(m *Map) JavaScript {
	return NewDefaultFunction("keys", func(args []JavaScript, ctx *Context) JavaScript {
		return GoToJs(m.Data.Keys())
	})
}

func mapValues(m *Map) JavaScript {
	return NewDefaultFunction("values", func(args []JavaScript, ctx *Context) JavaScript {
		keys := m.Data.Keys()
		arr := make([]JavaScript, len(keys))
		for i, k := range keys {
			arr[i], _ = m.Data.Get(k)
		}
		return GoToArray(arr)
	})
}

func mapGet(m *Map) JavaScript {
	return NewDefaultFunction("get", func(args []JavaScript, ctx *Context) JavaScript {
		if len(args) < 1 {
			return NewUndefined()
		}
		if v, ok := JsToBasic(args[0]); ok {
			v1, e := m.Data.Get(v)
			if !e {
				return NewUndefined()
			}
			return v1
		}
		v, e := m.Data.Get(args[0])
		if !e {
			return NewUndefined()
		}
		return v
	})
}

func mapDelete(m *Map) JavaScript {
	return NewDefaultFunction("delete", func(args []JavaScript, ctx *Context) JavaScript {
		if len(args) < 1 {
			return NewUndefined()
		}
		if v, ok := JsToBasic(args[0]); ok {
			return GoToBoolean(m.Data.Delete(v))
		}
		return GoToBoolean(m.Data.Delete(args[0]))
	})
}

func mapClear(m *Map) JavaScript {
	return NewDefaultFunction("clear", func(args []JavaScript, ctx *Context) JavaScript {
		m.Data.Clear()
		return NewUndefined()
	})
}

func mapHas(m *Map) JavaScript {
	return NewDefaultFunction("get", func(args []JavaScript, ctx *Context) JavaScript {
		if len(args) < 1 {
			return NewUndefined()
		}
		if v1, ok := args[0].(JavaScriptBasic); ok {
			return GoToBoolean(m.Data.Has(v1.Value()))
		}
		if _, ok := args[0].(*Undefined); ok {
			return GoToBoolean(m.Data.Has(Undefined{}))
		}
		return GoToBoolean(m.Data.Has(args[0]))
	})
}

func (m *Map) GetProperty(name string) JavaScript {
	switch name {
	case "size":
		return GoToNumber(m.Data.Size())
	case "set":
		return mapSet(m)
	case "keys":
		return mapKeys(m)
	case "values":
		return mapValues(m)
	case "get":
		return mapGet(m)
	case "has":
		return mapHas(m)
	case "delete":
		return mapDelete(m)
	case "clear":
		return mapClear(m)
	default:
		return NewUndefined()
	}
}
func (m *Map) GetIterator() JavaScriptIterator {
	return &MapIterator{index: 0, data: m, keys: m.Data.Keys()}
}

func (m *Map) Print() string {
	return fmt.Sprintf("Map(%v)", m.Data.Size())
}

func (m *Map) GetName() string {
	return "map"
}

func (m *Map) DeleteProperty(name string) JavaScript {
	if v, ok := m.Data.Get(name); ok {
		m.Data.Delete(name)
		return v
	}
	return NewUndefined()
}

func (m *Map) ToString() string {
	return "map"
}

/**************实现迭代器接口******************/

func (a *MapIterator) Next() (v []JavaScript, done bool) {
	defer func() {
		a.index++
	}()
	if a.index >= len(a.keys) {
		return []JavaScript{NewUndefined(), NewUndefined()}, true
	}
	va, _ := a.data.Data.Get(a.keys[a.index])
	return []JavaScript{GoToJs(a.keys[a.index]), va}, false
}

func (a *MapIterator) GetName() string {
	return "map.iterator"
}
