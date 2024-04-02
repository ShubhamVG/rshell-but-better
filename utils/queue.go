package utils

import "errors"

type Node[T any] struct {
	Data T
	Next *Node[T]
}

type Queue[T any] struct {
	First *Node[T]
	Last  *Node[T]
	Len   uint
}

func NewQueue[T any]() Queue[T] {
	return Queue[T]{Len: 0}
}

func (q Queue[T]) Dequeue() (Node[T], error) {
	if q.Len == 0 {
		return Node[T]{}, errors.New("Empty queue.")
	}

	first := *q.First
	q.First = q.First.Next
	q.Len--

	return first, nil
}

func (q *Queue[T]) Enqueue(nodePtr *Node[T]) {
	q.Last.Next = nodePtr
	q.Last = nodePtr
	q.Len++
}
