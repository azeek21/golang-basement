package utils

import "errors"

type QueueNode struct {
	next  *QueueNode
	prev  *QueueNode
	value any
}

type Queue struct {
	head *QueueNode
	tail *QueueNode
	size int
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Len() int {
	return q.size
}

func (q *Queue) Enque(value any) *Queue {
	newNode := &QueueNode{
		value: value,
	}

	if q.size == 0 {
		q.head = newNode
		q.tail = newNode
		q.size++
		return q
	}

	newNode.prev = q.tail
	q.tail.next = newNode
	q.tail = newNode
	q.size++
	return q
}

func (q *Queue) Deque() (any, error) {
	res := q.head.value
	if q.size == 0 {
		return nil, errors.New("Queue is empty, tring to deque... Use .Empty() to check if queue is empty")
	}

	if q.size == 1 {
		q.head = nil
		q.tail = nil
		q.size = 0
		return res, nil
	}
	q.head.next.prev = nil
	q.head = q.head.next
	q.size--
	return res, nil
}

func (q *Queue) Empty() bool {
	return q.size == 0
}
