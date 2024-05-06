package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type QueTestCase struct {
	name     string
	arg      int
	expected int
}

func TestEnque(t *testing.T) {
	queue := NewQueue()
	for i := 0; i < 10; i++ {

		t.Run(fmt.Sprintf("%v should equal to  %v\n", i, i), func(t *testing.T) {
			queue.Enque(i)
			res := queue.tail.value.(int)
			assert.Equal(t, i, res)
		})

		t.Run(fmt.Sprintf("Size should be %v\n", i+1), func(t *testing.T) {
			assert.Equal(t, i+1, queue.size)
		})
	}
}

func TestDeque(t *testing.T) {
	queue := NewQueue()
	for i := 0; i < 10; i++ {
		queue.Enque(i)
	}

	for i := 0; i < 10; i++ {
		t.Run(fmt.Sprintf("Deque should return %v\n", i), func(t *testing.T) {
			res, _ := queue.Deque()
			res = res.(int)
			assert.Equal(t, i, res)

		})

		t.Run(fmt.Sprintf("Size should be %v\n", i+1), func(t *testing.T) {
			assert.Equal(t, 10-i-1, queue.size)
		})
	}
}

func TestEmptyQue(t *testing.T) {
	queue := NewQueue()
	var emptyQueNode *QueueNode
	t.Run("Empty que should have head and tail to nil", func(t *testing.T) {
		assert.Equal(t, emptyQueNode, queue.head)
		assert.Equal(t, emptyQueNode, queue.tail)
	})

	t.Run("Empty que should have 0 size", func(t *testing.T) {
		assert.Equal(t, 0, queue.size)
	})
}

func TestSingleItemQue(t *testing.T) {
	queue := NewQueue()
	queue.Enque("Hi")

	t.Run("In single item que both tail and head should point to same node", func(t *testing.T) {
		assert.Equal(t, queue.head, queue.tail)
		assert.Equal(t, queue.head.value, queue.tail.value)
	})
}

func TestTwoItemQueue(t *testing.T) {
	queue := NewQueue()
	queue.Enque("Hi")
	queue.Enque("Bye")

	t.Run("Pointer links test", func(t *testing.T) {
		assert.Equal(t, queue.head.next, queue.tail)
		assert.Equal(t, queue.head.next.value, queue.tail.value)
		assert.Equal(t, queue.tail.prev, queue.head)
		assert.Equal(t, queue.tail.prev.value, queue.head.value)
		assert.Equal(t, queue.head.value, "Hi")
		assert.Equal(t, queue.tail.value, "Bye")
	})
}

func TestSizeAndEmptyness(t *testing.T) {
	queue := NewQueue()
	queue.Enque("Hello")
	queue.Enque("World")
	queue.Enque("Bye")

	t.Run("Sizes and emptyness", func(t *testing.T) {
		assert.Equal(t, 3, queue.Len())

		rest, _ := queue.Deque()
		assert.Equal(t, rest, "Hello")
		assert.Equal(t, 2, queue.Len())
		assert.Equal(t, false, queue.Empty())

		rest, _ = queue.Deque()
		assert.Equal(t, rest, "World")
		assert.Equal(t, 1, queue.Len())
		assert.Equal(t, false, queue.Empty())

		rest, _ = queue.Deque()
		assert.Equal(t, rest, "Bye")
		assert.Equal(t, 0, queue.Len())
		assert.Equal(t, true, queue.Empty())
	})
}

func TestEmptyQueDeque(t *testing.T) {
	queue := NewQueue()
	t.Run("Dequeuing empty que should return error and nil", func(t *testing.T) {
		res, err := queue.Deque()
		assert.Equal(t, nil, res)
		assert.NotEqual(t, nil, err)
	})
}
