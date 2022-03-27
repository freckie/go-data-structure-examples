package stack

import (
	"testing"
)

func TestInSingleGoroutine(t *testing.T) {
	s := NewStack[int]()

	s.Push(1)
	if s.Len() != 1 {
		t.Error("failed to push an element")
	}

	tmp := s.Pop()
	if tmp.Value() != 1 || s.Len() != 0 {
		t.Error("failed to pop from the stack")
	}
}
