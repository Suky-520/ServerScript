//用户定义模块

package ss

import (
	"ss/collection"
)

// Module JavaScript中的模块
type Module struct {
	export *collection.ArrayMap[string, *Variable] //导出的模块
	vars   *collection.ArrayMap[string, *Variable] //上下文变量
}

var _ JavaScript = &Module{}

// newModule 新建js中的module模块
func newModule(vars *collection.ArrayMap[string, *Variable], export *collection.ArrayMap[string, *Variable]) *Module {
	m := &Module{
		export: export,
		vars:   vars,
	}
	return m
}

func (m *Module) GetName() string {
	return "module"
}

func (m *Module) TypeOf() JavaScriptType {
	return JavaScriptType{Name: "module", Type: ObjectType, Native: false}
}
