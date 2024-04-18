package global

type Stack[T any] interface {
	ToSlice() []T               //转为切片
	Get(i int) (item T)         //获取元素
	Push(item T)                //加入队尾
	PushAll(data []T) (ok bool) //批量添加
	PushLeft(item T) (ok bool)  //加入队首
	SetItem(i int, item T) bool //写入元素

	Pop() (item T, ok bool)               //弹出
	PopLeft() (item T, ok bool)           //左侧弹出
	PopMany(size int) (item []T, ok bool) //弹出多个

	Peek() (item T, ok bool)     //窥视
	PeekLeft() (item T, ok bool) //左侧窥视

	Clear()    //清空
	Size() int //大小
}

// NewStack 新建栈对象
func NewStack[T any]() Stack[T] {
	return newDefaultStack[T]()
}
