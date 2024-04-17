package datastructs

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

func (q *Queue[T]) Dequeue() (T, error) {
	var data T

	switch q.Len {
	case 0:
		return data, errors.New("empty queue")
	case 1:
		data = q.First.Data
		q.First = nil
		q.Last = nil
	default:
		data = q.First.Data
		q.First = q.First.Next
	}

	q.Len--

	return data, nil
}

func (q *Queue[T]) Enqueue(elem T) {
	node := Node[T]{Data: elem, Next: nil}

	switch q.Len {
	case 0:
		q.First = &node
	case 1:
		q.First.Next = &node
	default:
		q.Last.Next = &node
	}

	q.Last = &node
	q.Len++
}
