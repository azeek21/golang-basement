package shared

import (
	"errors"
)

type Node[T any] struct {
	value T
	next  *Node[T]
}

type Queue[T any] struct {
	head *Node[T]
	tail *Node[T]
	size int
}

type IQueue[T any] interface {
	Enque(item T)
	Deque() (error, T)
	EnqueMany(items ...T)
	Size() int
}

var ERR_EMPTY_QUE = errors.New("Queue is empty")

func NewQueue[T any](initial ...T) IQueue[T] {
	newQueue := Queue[T]{
		head: nil,
		tail: nil,
		size: 0,
	}

	newQueue.EnqueMany(initial...)

	return &newQueue
}

func (q *Queue[T]) Enque(item T) {
	node := Node[T]{
		value: item,
	}

	if q.size == 0 {
		q.head = &node
		q.tail = &node
	} else {
		q.tail.next = &node
		q.tail = q.tail.next
	}
	q.size++
}

func (q *Queue[T]) EnqueMany(items ...T) {
	if items != nil && len(items) > 0 {
		for _, item := range items {
			q.Enque(item)
		}
	}
}

func (q *Queue[T]) Deque() (error, T) {
	var res T
	if q.size == 0 {
		return ERR_EMPTY_QUE, res
	}

	res = q.head.value
	if q.size == 1 {
		q.head = nil
		q.tail = nil
		q.size = 0
		return nil, res
	}

	q.head = q.head.next
	q.size--
	return nil, res
}

func (q *Queue[T]) Size() int {
	return q.size
}
