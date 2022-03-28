package stack

import (
	"sync"
	"testing"
)

func TestInMultipleGoroutines(t *testing.T) {
	const maxGoroutines = 7000

	s := NewThreadSafeStack[int]()
	// s := NewStack[int]()

	wg := sync.WaitGroup{}
	for i := 0; i < maxGoroutines; i++ {
		wg.Add(1)
		go func(val int) {
			s.Push(val)
			defer wg.Done()
		}(i)
	}
	wg.Wait()
	t.Logf("length : %d \n", s.Len())

	if s.Len() != maxGoroutines {
		t.Error("failed")
	}

	for i := 0; i < maxGoroutines; i++ {
		s.Pop()
	}
}
