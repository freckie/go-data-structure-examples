package list

import (
	"testing"
)

func TestInSingleGoroutine(t *testing.T) {
	l := NewList[int]()

	l.PushFront(0)
	if l.head != l.tail {
		t.Error("head != tail when the length is 0")
	}

	l.Remove(l.Begin())
	if l.head != nil || l.tail != nil || l.len != 0 {
		t.Error("failed to remove the head element")
	}

	values := []int{1, 2, 3}
	for _, val := range values {
		l.PushFront(val)
	}
	if l.Len() != 3 {
		t.Error("inaccurate length")
	}
	if l.Begin().Value() != 3 || l.End().Value() != 1 {
		t.Error("inaccurate values")
	}

	l.Clear()
	if l.head != nil || l.tail != nil || l.len != 0 {
		t.Error("failed to clear the list")
	}
}
