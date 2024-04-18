//ss的虚拟机,用于执行指令

package ss

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"ss/collection"
	"ss/global"
	"strings"
	"time"
)

// VM 虚拟机
type VM struct {
	ctx   *Context
	error *Context //错误上下文
}

// Context 上下文
type Context struct {
	code          []Instruct                              //指令
	size          int                                     //code长度
	index         int                                     //当前指令执行的索引
	vars          *collection.ArrayMap[string, *Variable] //当前上下文中存储的变量
	value         global.Stack[JavaScript]                //执行过程中产生的值堆栈,该堆栈要求性能极高
	parent        *Context                                //当前上下文的父级上下文
	exports       *collection.ArrayMap[string, *Variable] //当前上下文export导出的模块
	imports       []*Module                               //当前上下文import导入的模块
	vm            *VM                                     //当前虚拟机
	statement     command                                 //当前上下文执行的语句类型 Call,While
	ellipsisCount int                                     //用来标记解构总数
	isSkip        bool                                    //是否触发跳转,用于跳过中间某些指令的执行,直到遇到跳转指令skip为止
	propertyNames global.Stack[string]                    //记录属性名称,用于报错提示
}

// Variable 变量
type Variable struct {
	Value    JavaScript
	Modifier modifier
}

// Start 开始
func Start(vinoVersion string) {
	_version = vinoVersion
	//初始化process对象
	setGlobalObject("process", initProcess())
	if len(os.Args) == 2 && !strings.HasPrefix(os.Args[1], "-") {
		//运行文件
		v := NewVM()
		file := os.Args[1]
		if strings.HasSuffix(file, ".ss") {
			v.RunBinFile(os.Args[1])
		} else {
			v.RunJsFile(os.Args[1])
		}
	} else if len(os.Args) == 2 && os.Args[1] == "-v" {
		//打印版本号
		fmt.Println(GreenStyle+"v"+vinoVersion+EndStyle, GreyStyle+"[ss v"+fmt.Sprint(version[0], ".", version[1], ".", version[2])+"]"+EndStyle)
	} else if len(os.Args) >= 3 && os.Args[1] == "build" {
		fmt.Println("正在编译...")
		start := time.Now()
		jsFile := os.Args[2]
		ssFile := ""
		if len(os.Args) > 3 {
			ssFile = os.Args[3]
		}
		v := NewVM()
		v.Build(jsFile, ssFile)
		duration := time.Since(start).Milliseconds() // 计算耗时
		fmt.Println("编译完成,耗时:", duration, "ms")
	}
}

func newCtx() *Context {
	ctx := &Context{
		vars:          collection.NewArrayMap[string, *Variable](),
		value:         global.NewStack[JavaScript](),
		exports:       collection.NewArrayMap[string, *Variable](),
		imports:       make([]*Module, 0),
		propertyNames: global.NewStack[string](),
	}
	return ctx
}

// NewVM 新建虚拟机
func NewVM() *VM {
	vm := &VM{}
	ctx := newCtx()
	ctx.vm = vm
	vm.ctx = ctx
	return vm
}

// RunBinFile 运行字节码文件
func (vm *VM) RunBinFile(binFile string) {
	defer func() {
		if err := recover(); err != nil {
			switch err.(type) {
			case Error:
				e := err.(Error)
				JsError(e.name+":"+e.message, nil)
			default:
				panic(err)
			}
		}
	}()
	// 打开文件
	file, err := os.Open(binFile)
	if err != nil {
		panic(NewError("Error", err.Error()))
		return
	}
	defer file.Close()
	// 读取文件内容
	content, err := io.ReadAll(file)
	if err != nil {
		panic(NewError("Error", err.Error()))
	}
	b := unBuild(content)
	for _, v := range b.files {
		setSourceFile(v, "")
	}
	getConfigFile(binFile)
	vm.run(0, b.code)
}

// Build 编译文件
func (vm *VM) Build(jsFile string, binFile string) {
	defer func() {
		if err := recover(); err != nil {
			switch err.(type) {
			case Error:
				e := err.(Error)
				JsError(e.name+":"+e.message, nil)
			default:
				panic(err)
			}
		}
	}()
	if binFile == "" {
		base := strings.TrimSuffix(jsFile, filepath.Ext(jsFile))
		binFile = base + ".ss"
	}
	p := parseFile(jsFile)
	if p != nil && p.Code != nil && len(p.Code) > 0 {
		bin := build(p.Code)
		// 检查目录是否存在，如果不存在则创建
		if err := os.MkdirAll(filepath.Dir(binFile), 0755); err != nil {
			panic(NewError("Error", err.Error()))
			return
		}
		file, err := os.Create(binFile)
		if err != nil {
			panic(NewError("Error", err.Error()))
			return
		}
		defer file.Close()
		if _, err = file.Write(bin); err != nil {
			panic(NewError("Error", err.Error()))
			return
		}
	} else {
		panic(NewSyntaxError("编译错误"))
	}
}

// RunJsFile 运行文件
func (vm *VM) RunJsFile(file string) {
	fmt.Println("正在编译...")
	start := time.Now()
	p := parseFile(file)
	getConfigFile(file)
	duration := time.Since(start).Milliseconds() // 计算耗时
	fmt.Println("编译完成,耗时:", duration, "ms")
	if p != nil && p.Code != nil && len(p.Code) > 0 {
		vm.run(p.file, p.Code)
	}
}

// 获取当前默认配置文件
func getConfigFile(main string) {
	absPath, err := filepath.Abs(main)
	if err != nil {
		panic(err)
	}
	file1 := filepath.Join(filepath.Dir(absPath), "./config.js")
	_, err = os.Open(file1)
	if err != nil {
		if getGlobalObject("config") == nil {
			setGlobalObject("config", NewObject())
		}
		return
	}
	p := parseFile(file1)
	vm1 := NewVM()
	if p != nil && p.Code != nil && len(p.Code) > 0 {
		vm1.run(p.file, p.Code)
	}
	c, ok := vm1.ctx.exports.Get("default")
	if !ok {
		//空配置文件
		if getGlobalObject("config") == nil {
			setGlobalObject("config", NewObject())
		}
		return
	} else {
		if getGlobalObject("config") == nil {
			setGlobalObject("config", c.Value)
		}
	}
}

// RunJsSource 运行代码
func (vm *VM) RunJsSource(source string, file string) {
	p := parseSource(source, file)
	if p != nil && p.Code != nil && len(p.Code) > 0 {
		vm.run(p.file, p.Code)
	}
}

// 错误堆栈打印
func (ctx *Context) printStack(err any) {
	switch err.(type) {
	case Error:
		e := err.(Error)
		JsError(e.name+":"+e.message, ctx)
	default:
		panic(err)
	}
	c := ctx
	for {
		if c.size > c.index {
			i := c.code[c.index]
			src := i.getSrc()
			if src != nil {
				PrintErrorStack(i.getSrc().File, src.StartLine, src.StartColumn, src.token)
			}
		}
		if c.parent != nil {
			c = c.parent
		} else {
			break
		}
	}
}

// run 运行
func (vm *VM) run(file int, code []Instruct) {
	defer func() {
		if err := recover(); err != nil {
			if vm.error != nil {
				vm.error.printStack(err)
			} else {
				fmt.Println("不存在error")
			}
		}
	}()
	vm.ctx.code = code
	vm.ctx.size = len(code)
	vm.ctx.index = 0
	vm.ctx.exec()
}

// GetVar 返回当前变量
func (vm *VM) GetVar() map[string]JavaScript {
	return vm.ctx.getVar()
}

// GetValue 返回当前值栈
func (vm *VM) GetValue() []JavaScript {
	return vm.ctx.value.ToSlice()
}

// 虚拟机执行指令
func (ctx *Context) exec() {
	defer func() {
		if err := recover(); err != nil {
			if ctx.vm.error == nil {
				ctx.vm.error = ctx
			}
			panic(err)
		}
	}()
	for {
		if ctx.index < ctx.size {
			c := ctx.code[ctx.index]
			//可选链跳转
			if ctx.isSkip {
				if _, ok := c.(skip); !ok {
					ctx.next()
					continue
				}
			}
			c.exec(ctx)
		} else {
			break
		}
	}
}

// 下一条指令
func (ctx *Context) next() {

	ctx.index++
}

// 新建上下文
func (ctx *Context) newCtx() *Context {
	ctx1 := newCtx()
	ctx1.vm = ctx.vm
	ctx1.parent = ctx
	return ctx1
}

// 获取变量
func (ctx *Context) getVar() map[string]JavaScript {
	vars := make(map[string]JavaScript)
	for _, k := range ctx.vars.Keys() {
		v, _ := ctx.vars.Get(k)
		vars[k] = v.Value
	}
	return vars
}

// RunFunction 执行回调函数
func (ctx *Context) RunFunction(fun *Function, args ...JavaScript) JavaScript {
	c := ctx.newCtx()
	defer func() {
		if err := recover(); err != nil {
			if c.vm.error != nil {
				c.vm.error.printStack(err)
			}
			panic(err)
		}
	}()
	for _, a := range args {
		c.code = append(c.code, loadVal{value: a})
	}
	c.code = append(c.code, call{index: len(args)})
	c.size = len(c.code)
	c.value.Push(fun)
	c.exec()
	v, _ := c.value.Pop()
	//销毁对象,可以快速释放内存
	c = nil
	return v
}

// GetSrcMap 获取源码信息
func (ctx *Context) GetSrcMap() (*Src, string) {
	if len(ctx.code) <= ctx.index {
		return nil, ""
	}
	c := ctx.code[ctx.index]
	src := c.getSrc()
	if src != nil {
		return src, getSourceFilePath(src.File)
	} else {
		return nil, ""
	}
}
