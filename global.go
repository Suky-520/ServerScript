package ss

import (
	"ss/collection"
)

var _ JavaScript = &Global{}

type Global struct {
	data *collection.ArrayMap[any, JavaScript]
	prop JavaScriptProperties
}

func init() {
	setGlobalObject("global", initGlobal())
}

func initGlobal() JavaScript {
	d := &Global{data: collection.NewArrayMap[any, JavaScript](), prop: collection.NewArrayMap[string, JavaScript]()}
	d.prop.Set("set", d.set())
	d.prop.Set("get", d.get())
	d.prop.Set("delete", d.delete())
	return d
}

func (m *Global) set() JavaScript {
	return NewDefaultFunction("set", func(args []JavaScript, ctx *Context) JavaScript {
		if len(args) < 2 {
			return NewUndefined()
		}
		if v, ok := args[0].(JavaScriptBasic); ok {
			m.data.Set(v.Value(), args[1])
		} else if _, ok := args[0].(*Undefined); ok {
			m.data.Set(Undefined{}, args[1])
		} else {
			m.data.Set(args[0], args[1])
		}
		return m
	})
}

func (m *Global) get() JavaScript {
	return NewDefaultFunction("get", func(args []JavaScript, ctx *Context) JavaScript {
		if len(args) < 1 {
			return NewUndefined()
		}
		if v1, ok := args[0].(JavaScriptBasic); ok {
			v, e := m.data.Get(v1.Value())
			if !e {
				return NewUndefined()
			}
			return v
		}
		if _, ok := args[0].(*Undefined); ok {
			v, e := m.data.Get(Undefined{})
			if !e {
				return NewUndefined()
			}
			return v
		}
		v, e := m.data.Get(args[0])
		if !e {
			return NewUndefined()
		}
		return v
	})
}

func (m *Global) delete() JavaScript {
	return NewDefaultFunction("delete", func(args []JavaScript, ctx *Context) JavaScript {
		if len(args) < 1 {
			return NewUndefined()
		}
		if v1, ok := args[0].(JavaScriptBasic); ok {
			return GoToBoolean(m.data.Delete(v1.Value()))
		}
		if _, ok := args[0].(*Undefined); ok {
			key := Undefined{}
			return GoToBoolean(m.data.Delete(key))
		}
		key := args[0]
		return GoToBoolean(m.data.Delete(key))
	})
}

func (m *Global) forEach() JavaScript {
	return NewDefaultFunction("forEach", func(args []JavaScript, ctx *Context) JavaScript {
		if len(args) < 1 {
			return NewUndefined()
		}
		if f, ok := args[0].(*Function); ok {
			for _, key := range m.data.Keys() {
				v, _ := m.data.Get(key)
				ctx.RunFunction(f, GoToJs(v), GoToJs(key))
			}
		}
		return NewUndefined()
	})
}

func (m *Global) GetName() string {
	return "global"
}

func (m *Global) GetProperty(name string) JavaScript {
	if v, ok := m.prop.Get(name); ok {
		return v
	}
	return NewUndefined()
}

func (m *Global) GetProperties() JavaScriptProperties {
	return m.prop
}
