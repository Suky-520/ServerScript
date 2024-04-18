package ss

import (
	"ss/collection"
)

var (
	// Files 存储所有源文件
	sourceFile = make(map[string]string)
	//全局对象
	globalObject = make(map[string]any)
	//全局模块
	globalModules = make(map[string]any)
	//全局原生模块
	globalNativeModules = make(map[string]*Module)
	//全局文件
	globalFiles = make([]string, 0)
	// version 版本号
	version = []uint8{0, 7, 5}
)

// 清空一下文件缓存
func clearFileCache() {
	sourceFile = make(map[string]string)
	globalFiles = make([]string, 0)
}

// DefaultModule 模块
type DefaultModule struct {
	mod *collection.ArrayMap[string, *Variable]
}

// NewModule 新建模块对象
func NewModule() *DefaultModule {
	return &DefaultModule{mod: collection.NewArrayMap[string, *Variable]()}
}

// RegisterModule 注册模块
func RegisterModule(name string, value *DefaultModule) {
	setGlobalNativeModule(name, &Module{export: value.mod})
}

// SetExport 添加导出项
func (d *DefaultModule) SetExport(name string, value JavaScript) {
	d.mod.Set(name, &Variable{Value: value, Modifier: Const})
}

// SetDefaultExport 添加默认导出项
func (d *DefaultModule) SetDefaultExport(value JavaScript) {
	d.mod.Set("default", &Variable{Value: value, Modifier: Const})
}

// getGlobalObject 获取全局对象
func getGlobalObject(name string) any {
	return globalObject[name]
}

// setGlobalObject 设置全局对象
func setGlobalObject(name string, obj any) {
	if _, ok := globalObject[name]; ok {
		panic("对象" + name + "已存在")
	} else {
		globalObject[name] = obj
	}
}

func SetGlobalModule(name string, obj any) {
	if _, ok := globalModules[name]; ok {
		panic("模块" + name + "已存在")
	} else {
		globalModules[name] = obj
	}
}

func getGlobalModule(name string) any {
	return globalModules[name]
}

func getGlobalNativeModule(name string) *Module {
	return globalNativeModules[name]
}

func setGlobalNativeModule(name string, module *Module) {
	if _, ok := globalNativeModules[name]; ok {
		panic("模块" + name + "已存在")
	} else {
		globalNativeModules[name] = module
	}
}

func getSourceFilePath(index int) string {
	return globalFiles[index]
}

func setSourceFile(path string, content string) int {
	sourceFile[path] = content
	globalFiles = append(globalFiles, path)
	return len(globalFiles) - 1
}

func getGlobalFiles() []string {
	return globalFiles
}
