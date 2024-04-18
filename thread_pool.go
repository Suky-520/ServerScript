//JavaScript中的ThreadPool类

package ss

import (
	"github.com/gammazero/workerpool"
	"runtime"
	"ss/collection"
)

// ThreadPool 线程池
type ThreadPool struct {
	maxWorkers int
	pool       *workerpool.WorkerPool
	prop       JavaScriptProperties
	workerSize int
}

func init() {
	setGlobalObject("ThreadPool", NewDefaultClass("ThreadPool", threadPool))
}

var threadPool = func(args []JavaScript, ctx *Context) JavaScript {
	workerSize := runtime.NumCPU()
	if len(args) > 0 {
		if s, ok := args[0].(*Number); ok {
			workerSize = int(s.Data)
		}
	}
	th := &ThreadPool{pool: workerpool.New(workerSize), workerSize: workerSize, prop: collection.NewArrayMap[string, JavaScript]()}
	obj := NewDefaultObject("threadPool")
	obj.AddProperty("submit", th.submit())
	obj.AddProperty("size", th.size())
	obj.AddProperty("workerSize", GoToNumber(th.workerSize))
	obj.AddProperty("wait", th.wait())
	obj.AddProperty("stop", th.stop())
	return obj
}

func (t *ThreadPool) submit() JavaScript {
	return NewDefaultFunction("submit", func(args []JavaScript, ctx *Context) JavaScript {
		if len(args) == 0 {
			return NewUndefined()
		}
		if f, ok := args[0].(*Function); ok {
			t.pool.Submit(func() {
				ctx.RunFunction(f, args[1:]...)
			})
			return NewUndefined()
		} else {
			panic(NewTypeError("参数错误"))
		}
	})
}

func (t *ThreadPool) stop() JavaScript {
	return NewDefaultFunction("stop", func(args []JavaScript, ctx *Context) JavaScript {
		t.pool.Stop()
		return NewUndefined()
	})
}

func (t *ThreadPool) size() JavaScript {
	return NewDefaultFunction("size", func(args []JavaScript, ctx *Context) JavaScript {
		return GoToNumber(float64(t.pool.WaitingQueueSize()))
	})
}

func (t *ThreadPool) wait() JavaScript {
	return NewDefaultFunction("wait", func(args []JavaScript, ctx *Context) JavaScript {
		t.pool.StopWait()
		return NewUndefined()
	})
}
