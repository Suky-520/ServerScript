//JavaScript中用户定义的object对象

package ss

import (
	"bytes"
	"encoding/json"
	"ss/collection"
)

// Object JavaScript 对象类型
type Object struct {
	props JavaScriptProperties
}

// ObjectIterator 对象属性迭代器
type ObjectIterator struct {
	key   []string
	value JavaScriptProperties
	index int
}

var _ JavaScript = &Object{}

var _ JavaScriptIterator = &ObjectIterator{}

var _ JavaScriptBinary = &Object{}

// NewObject 新建一个js中的object对象
func NewObject() *Object {
	obj := &Object{props: collection.NewConcurrentArrayMap[string, JavaScript]()}
	return obj
}

func GoToObject(data map[string]any) *Object {
	return GoToJs(data).(*Object)
}

func (o *Object) GetName() string {
	return "object"
}

func (o *Object) GetIterator() JavaScriptIterator {
	return &ObjectIterator{index: 0, key: o.props.Keys(), value: o.props}
}

func (o *Object) GetProperty(name string) JavaScript {
	if v, ok := o.props.Get(name); ok {
		return v
	}
	if name == "constructor" {
		panic(NewTypeError(name + " is not a constructor"))
	} else {
		return NewUndefined()
	}
}
func (o *Object) DeleteProperty(name string) JavaScript {
	if ok := o.props.Has(name); ok {
		return GoToBoolean(o.props.Delete(name))
	} else {
		return GoToBoolean(false)
	}
}
func (o *Object) SetProperty(name string, value JavaScript) {
	o.props.Set(name, value)
}
func (o *Object) GetProperties() JavaScriptProperties {
	return o.props
}

func (o *Object) UnmarshalJSON(data []byte) error {
	m := make(map[string]any)
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	o.props = GoToObject(m).props
	return nil
}

func (o *Object) MarshalJSON() ([]byte, error) {
	var sb bytes.Buffer
	sb.WriteString("{")
	for i, k := range o.props.Keys() {
		mk, err := json.Marshal(k)
		if err != nil {
			continue
		}
		mv, _ := o.props.Get(k)
		t := JsToType(mv).Type
		if t == UndefinedType || t == FunctionType {
			continue
		}
		mvv, err := json.Marshal(mv)
		if err != nil {
			continue
		}
		if i > 0 {
			sb.WriteString(",")
		}
		sb.Write(mk)
		sb.WriteString(":")
		sb.Write(mvv)
	}
	sb.WriteString("}")
	return sb.Bytes(), nil
}

func (o *Object) Binary() []byte {
	c, _ := o.MarshalJSON()
	return c
}

/**************实现迭代器接口******************/

func (a *ObjectIterator) Next() (v []JavaScript, done bool) {
	defer func() {
		a.index++
	}()
	if a.index >= len(a.key) {
		return []JavaScript{NewUndefined(), NewUndefined()}, true
	}
	o := a.key[a.index]
	va, _ := a.value.Get(o)
	return []JavaScript{GoToString(o), va}, false
}

func (a *ObjectIterator) GetName() string {
	return "object.iterator"
}
