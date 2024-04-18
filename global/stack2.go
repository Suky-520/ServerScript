package global

import (
	"sync"
)

var initLength = 100 //初始大小

type HighStack[T any] struct {
	items  []T
	size   int //实际存储的数据量大小
	length int //切片的实际长度
	// 读写锁，用于保护共享资源
	lock sync.RWMutex
}

func newHighStack[T any]() *HighStack[T] {
	return &HighStack[T]{items: make([]T, initLength), size: 0, length: initLength}
}

func (s *HighStack[T]) ToSlice() []T {
	return s.items[0:s.size]
}

func (s *HighStack[T]) Get(i int) (item T) {
	if i < 0 {
		return item
	}
	if i >= s.size {
		return item
	}
	return s.items[i]
}

func (s *HighStack[T]) Push(item T) {
	s.items[s.size] = item
	s.size++
}

func (s *HighStack[T]) expand(size int) bool {
	if size+s.size >= s.length {
		//开始扩容
		if size+s.size-s.length > 20 {

		}
		return true
	} else {
		return true
	}
}

func (s *HighStack[T]) Pop() (item T, ok bool) {
	if s.size == 0 {
		return item, false
	}
	s.size = s.size - 1
	item = s.items[s.size]
	return item, true
}

func (s *HighStack[T]) SetItem(i int, item T) bool {
	if i >= s.size {
		return false
	}
	s.items[i] = item
	return true
}

func (s *HighStack[T]) PopMany(size int) (item []T, ok bool) {
	n := s.size
	if n == 0 || size <= 0 {
		return nil, false
	}
	if size > n {
		size = n
	}
	items := make([]T, size)
	copy(items, s.items[n-size:])
	s.size = s.size - size
	return items, true
}

func (s *HighStack[T]) PushAll(data []T) (ok bool) {
	copy(s.items[s.size:s.size+len(data)], data)
	s.size = s.size + len(data)
	return true
}

func (s *HighStack[T]) PushLeft(item T) (ok bool) {
	copy(s.items[1:s.size+1], s.items)
	s.items[0] = item
	s.size++
	return true
}

func (s *HighStack[T]) Clear() {
	s.items = make([]T, initLength)
	s.size = 0
	s.length = initLength
}

func (s *HighStack[T]) Size() int {
	return s.size
}

func (s *HighStack[T]) PopLeft() (item T, ok bool) {
	n := len(s.items)
	if n == 0 {
		var zero T // T的零值
		return zero, false
	}
	item = s.items[0]
	copy(s.items[0:s.size-1], s.items[1:])
	s.size--
	return item, true
}

func (s *HighStack[T]) Peek() (item T, ok bool) {
	if s.size == 0 {
		var zero T
		return zero, false
	}
	return s.items[s.size-1], true
}

func (s *HighStack[T]) PeekLeft() (item T, ok bool) {
	if s.size == 0 {
		var zero T
		return zero, false
	}
	return s.items[0], true
}
