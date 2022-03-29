package queue

import (
	"sync"
)

type ThreadSafeListQueue[T any] struct {
	sync.RWMutex
	rear  *ThreadSafeListQueueNode[T]
	front *ThreadSafeListQueueNode[T]
	len   uint
}

func NewThreadSafeListQueue[T any]() *ThreadSafeListQueue[T] {
	return &ThreadSafeListQueue[T]{
		rear:  nil,
		front: nil,
		len:   0,
	}
}

func (q *ThreadSafeListQueue[T]) IsEmpty() bool {
	return q.len == 0
}

func (q *ThreadSafeListQueue[T]) Len() uint {
	return q.len
}

func (q *ThreadSafeListQueue[T]) Rear() *ThreadSafeListQueueNode[T] {
	return q.rear
}

func (q *ThreadSafeListQueue[T]) Front() *ThreadSafeListQueueNode[T] {
	return q.front
}

func (q *ThreadSafeListQueue[T]) Push(val T) {
	node := &ThreadSafeListQueueNode[T]{
		value: val,
		prev:  nil,
	}

	q.Lock()
	defer q.Unlock()

	if q.len == 0 {
		q.front = node
	} else {
		q.rear.prev = node
	}
	q.rear = node
	q.len++
}

func (q *ThreadSafeListQueue[T]) Pop() *ThreadSafeListQueueNode[T] {
	if q.len == 0 {
		return nil
	}

	q.Lock()
	defer q.Unlock()
	node := q.front
	q.front = q.front.prev
	q.len--
	if q.len == 0 {
		q.rear = nil
	}
	return node
}

type ThreadSafeListQueueNode[T any] struct {
	value T
	prev  *ThreadSafeListQueueNode[T]
}

func (n ThreadSafeListQueueNode[T]) Value() T {
	return n.value
}
