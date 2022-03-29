package stack

import (
	"sync"
)

type ThreadSafeStack[T any] struct {
	sync.RWMutex
	top *ThreadSafeStackNode[T]
	len uint
}

func NewThreadSafeStack[T any]() *ThreadSafeStack[T] {
	return &ThreadSafeStack[T]{
		top: nil,
		len: 0,
	}
}

func (s *ThreadSafeStack[T]) IsEmpty() bool {
	return s.len == 0
}

func (s *ThreadSafeStack[T]) Len() uint {
	return s.len
}

func (s *ThreadSafeStack[T]) Push(val T) {
	node := &ThreadSafeStackNode[T]{
		value: val,
		prev:  nil,
	}

	s.Lock()
	defer s.Unlock()

	if s.len != 0 {
		node.prev = s.top
	}
	s.top = node
	s.len++
}

func (s *ThreadSafeStack[T]) Pop() *ThreadSafeStackNode[T] {
	if s.len == 0 {
		return nil
	}

	s.Lock()
	defer s.Unlock()

	node := s.top
	s.top = s.top.prev
	s.len--
	return node
}

func (s *ThreadSafeStack[T]) Top() *ThreadSafeStackNode[T] {
	return s.top
}

type ThreadSafeStackNode[T any] struct {
	value T
	prev  *ThreadSafeStackNode[T]
}

func (n ThreadSafeStackNode[T]) Value() T {
	return n.value
}
