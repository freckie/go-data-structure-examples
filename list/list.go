package list

type List[T any] struct {
	head *ListNode[T]
	tail *ListNode[T]
	len  uint
}

func NewList[T any]() *List[T] {
	return &List[T]{
		head: nil,
		tail: nil,
		len:  0,
	}
}

func (l List[T]) IsEmpty() bool {
	return l.len == 0
}

func (l *List[T]) Begin() *ListNode[T] {
	if l.len == 0 {
		return nil
	}
	return l.head
}

func (l *List[T]) End() *ListNode[T] {
	if l.len == 0 {
		return nil
	}
	return l.tail
}

func (l *List[T]) Clear() {
	if l.len == 0 {
		return
	}

	for it := l.head; it != nil; it = it.next {
		it.prev = nil
	}
	l.head = nil
	l.tail = nil
	l.len = 0
}

func (l *List[T]) Len() uint {
	return l.len
}

func (l List[T]) At(index uint) *ListNode[T] {
	if l.len == 0 {
		return nil
	}

	it := l.head
	for i := uint(0); i < index; i++ {
		it = it.next
	}
	return it
}

func (l *List[T]) PushFront(val T) *ListNode[T] {
	return l.insertAfter(val, nil)
}

func (l *List[T]) PushBack(val T) *ListNode[T] {
	return l.insertAfter(val, l.tail)
}

func (l *List[T]) Insert(val T, target *ListNode[T]) *ListNode[T] {
	return l.insertAfter(val, target)
}

func (l *List[T]) Remove(target *ListNode[T]) {
	if l.len == 1 {
		l.head = nil
		l.tail = nil
	} else if target == l.head {
		target.next.prev = nil
		l.head = target.next
	} else if target == l.tail {
		target.prev.next = nil
		l.tail = target.prev
	} else {
		target.prev.next = target.next
		target.next.prev = target.prev
	}
	target = nil
	l.len--
}

func (l *List[T]) insertAfter(val T, target *ListNode[T]) *ListNode[T] {
	node := &ListNode[T]{value: val}

	if target == nil && l.len != 0 {
		node.prev = nil
		node.next = l.head
		l.head.prev = node
		l.head = node
	} else if target == nil && l.len == 0 {
		node.prev = nil
		node.next = nil
		l.head = node
		l.tail = node
	} else if target == l.tail {
		node.prev = l.tail
		node.next = nil
		l.tail.next = node
		l.tail = node
	} else {
		node.prev = target
		node.next = target.next
		target.next.prev = node
		target.next = node
	}
	l.len++

	return node
}

type ListNode[T any] struct {
	value T
	prev  *ListNode[T]
	next  *ListNode[T]
}

func (n ListNode[T]) Value() T {
	return n.value
}

func (n ListNode[T]) Prev() *ListNode[T] {
	return n.prev
}

func (n ListNode[T]) Next() *ListNode[T] {
	return n.next
}
