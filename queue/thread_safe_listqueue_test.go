package queue

import (
	"sync"
	"testing"
)

func TestListQueueInMultipleGoroutines(t *testing.T) {
	const maxGoroutines = 7000

	q := NewThreadSafeListQueue[int]()
	// q := NewListQueue[int]()

	wg := sync.WaitGroup{}
	for i := 0; i < maxGoroutines; i++ {
		wg.Add(1)
		go func(val int) {
			q.Push(val)
			defer wg.Done()
		}(i)
	}
	wg.Wait()

	t.Logf("length : %d\n", q.Len())

	if q.Len() != maxGoroutines {
		t.Error("failed")
	}

	for i := 0; i < maxGoroutines; i++ {
		q.Pop()
	}

	t.Error("temp")
}
