package queue

type Queue[T any] interface {
	IsEmpty() bool
	Len() uint
	Rear() *QueueNode[T]
	Front() *QueueNode[T]
	Pop() *QueueNode[T]
	Push(val T)
}

type QueueNode[T any] interface {
	Value() T
}
