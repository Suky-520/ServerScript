package global

// DefaultStack 栈对象
type DefaultStack[T any] struct {
	items []T
}

func newDefaultStack[T any]() *DefaultStack[T] {
	return &DefaultStack[T]{items: make([]T, 0)}
}
func (s *DefaultStack[T]) ToSlice() []T {
	return s.items
}

func (s *DefaultStack[T]) Get(i int) (item T) {
	return s.items[i]
}

func (s *DefaultStack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *DefaultStack[T]) Pop() (item T, ok bool) {
	if len(s.items) == 0 {
		return item, false
	}
	index := len(s.items) - 1
	item = s.items[index]
	s.items = s.items[:index]
	return item, true
}
func (s *DefaultStack[T]) SetItem(i int, item T) bool {
	s.items[i] = item
	return true
}

func (s *DefaultStack[T]) PopMany(size int) (item []T, ok bool) {
	n := len(s.items)
	if n == 0 || size <= 0 {
		return nil, false
	}
	if size > n {
		size = n
	}
	items := make([]T, size)
	copy(items, s.items[n-size:])
	s.items = s.items[:n-size]
	return items, true
}

func (s *DefaultStack[T]) PushAll(data []T) (ok bool) {
	s.items = append(s.items, data...)
	return true
}

func (s *DefaultStack[T]) PushLeft(item T) (ok bool) {
	s.items = append([]T{item}, s.items...)
	return true
}

func (s *DefaultStack[T]) Clear() {
	s.items = nil
	s.items = make([]T, 0)
}

func (s *DefaultStack[T]) Size() int {
	return len(s.items)
}

func (s *DefaultStack[T]) PopLeft() (item T, ok bool) {
	n := len(s.items)
	if n == 0 {
		var zero T // T的零值
		return zero, false
	}
	item = s.items[0]
	s.items = s.items[1:]
	return item, true
}

func (s *DefaultStack[T]) Peek() (item T, ok bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	return s.items[len(s.items)-1], true
}

func (s *DefaultStack[T]) PeekLeft() (item T, ok bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	return s.items[0], true
}
