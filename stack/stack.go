package stack

type Stack[T any] struct {
	top *StackNode[T]
	len uint
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		top: nil,
		len: 0,
	}
}

func (s Stack[T]) IsEmpty() bool {
	return s.len == 0
}

func (s Stack[T]) Len() uint {
	return s.len
}

func (s *Stack[T]) Push(val T) {
	node := &StackNode[T]{
		value: val,
		prev:  nil,
	}

	if s.top == nil {
		s.top = node
	} else {
		node.prev = s.top
		s.top = node
	}
	s.len++
}

func (s *Stack[T]) Pop() *StackNode[T] {
	node := s.top
	s.top = s.top.prev
	s.len--
	return node
}

func (s Stack[T]) Top() *StackNode[T] {
	if s.len == 0 {
		return nil
	}

	return s.top
}

type StackNode[T any] struct {
	value T
	prev  *StackNode[T]
}

func (n StackNode[T]) Value() T {
	return n.value
}
