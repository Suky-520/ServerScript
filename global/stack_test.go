package global

import (
	"fmt"
	"testing"
	"time"
)

func TestDefaultStack(t *testing.T) {
	s := newDefaultStack[int]()
	start := time.Now()
	for i := 0; i < 10000*10000; i++ {
		s.Push(1)
		s.Pop()
	}
	fmt.Println(time.Since(start))
}

func TestHighStack(t *testing.T) {
	s := newHighStack[int]()
	start := time.Now()
	for i := 0; i < 10000*10000; i++ {
		s.Push(1)
		s.Pop()
	}
	fmt.Println(time.Since(start))
}
