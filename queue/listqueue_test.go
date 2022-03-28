package queue

import (
	"testing"
)

func TestListQueue(t *testing.T) {
	q := NewListQueue[int]()

	q.Push(0)
	tmp := q.Pop()
	if tmp.Value() != 0 || q.Len() != 0 {
		t.Error("failed to pop from the queue")
	}

	for i := 0; i < 5; i++ {
		q.Push(i)
	}
	if q.Len() != 5 {
		t.Error("failed to push elements")
	}

	for !q.IsEmpty() {
		q.Pop().Value()
	}
}
