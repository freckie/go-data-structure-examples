package queue

type ListQueue[T any] struct {
	rear  *ListQueueNode[T]
	front *ListQueueNode[T]
	len   uint
}

func NewListQueue[T any]() *ListQueue[T] {
	return &ListQueue[T]{
		rear:  nil,
		front: nil,
		len:   0,
	}
}

func (q ListQueue[T]) IsEmpty() bool {
	return q.len == 0
}

func (q ListQueue[T]) Len() uint {
	return q.len
}

func (q *ListQueue[T]) Rear() *ListQueueNode[T] {
	return q.rear
}

func (q *ListQueue[T]) Front() *ListQueueNode[T] {
	return q.front
}

func (q *ListQueue[T]) Push(val T) {
	node := &ListQueueNode[T]{
		value: val,
		prev:  nil,
	}

	if q.len == 0 {
		q.front = node
	} else {
		q.rear.prev = node
	}
	q.rear = node
	q.len++
}

func (q *ListQueue[T]) Pop() *ListQueueNode[T] {
	if q.len == 0 {
		return nil
	}

	node := q.front
	q.front = q.front.prev
	q.len--
	if q.len == 0 {
		q.rear = nil
	}
	return node
}

type ListQueueNode[T any] struct {
	value T
	prev  *ListQueueNode[T]
}

func (n ListQueueNode[T]) Value() T {
	return n.value
}
