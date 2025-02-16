package shared

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	t.Run("Single Item Operations", func(t *testing.T) {
		q := NewQueue[string]()

		// Test empty state
		assert.Equal(t, 0, q.Size(), "Initial queue size should be 0")

		// Test Enque
		q.Enque("a")
		assert.Equal(t, 1, q.Size(), "Size should be 1 after enque")

		// Test Deque
		err, item := q.Deque()
		assert.NoError(t, err, "Should not return error")
		assert.Equal(t, "a", item, "Dequeued item should match")
		assert.Equal(t, 0, q.Size(), "Size should be 0 after deque")
	})

	t.Run("Empty Queue Operations", func(t *testing.T) {
		q := NewQueue[int]()

		// Test Deque on empty queue
		err, item := q.Deque()
		assert.Error(t, err, "Should return error when dequeuing empty queue")
		assert.Equal(t, 0, item, "Should return zero value for type")
	})

	t.Run("Multiple Items Operations", func(t *testing.T) {
		q := NewQueue[int]()

		// Test EnqueMany
		q.EnqueMany(1, 2, 3)
		assert.Equal(t, 3, q.Size(), "Size should be 3 after bulk insert")

		// Verify FIFO order
		err, item := q.Deque()
		assert.NoError(t, err)
		assert.Equal(t, 1, item, "First item should be 1")

		err, item = q.Deque()
		assert.NoError(t, err)
		assert.Equal(t, 2, item, "Second item should be 2")

		q.Enque(4)
		assert.Equal(t, 2, q.Size(), "Size should be 2 after mixed operations")

		err, item = q.Deque()
		assert.NoError(t, err)
		assert.Equal(t, 3, item, "Third item should be 3")

		err, item = q.Deque()
		assert.NoError(t, err)
		assert.Equal(t, 4, item, "Fourth item should be 4")
	})

	t.Run("Edge Cases", func(t *testing.T) {
		t.Run("Empty EnqueMany", func(t *testing.T) {
			q := NewQueue[float64]()
			q.EnqueMany() // Should handle empty arguments
			assert.Equal(t, 0, q.Size(), "Size should remain 0")
		})

		t.Run("Mixed Type Operations", func(t *testing.T) {
			q := NewQueue[any]()
			q.Enque("string")
			q.Enque(42)
			q.Enque(true)

			err, item := q.Deque()
			assert.NoError(t, err)
			assert.Equal(t, "string", item)

			err, item = q.Deque()
			assert.NoError(t, err)
			assert.Equal(t, 42, item)

			err, item = q.Deque()
			assert.NoError(t, err)
			assert.Equal(t, true, item)
		})

		t.Run("Large Number of Items", func(t *testing.T) {
			q := NewQueue[int]()
			const count = 100000
			var items []int
			for i := 0; i < count; i++ {
				items = append(items, i)
			}

			q.EnqueMany(items...)
			assert.Equal(t, count, q.Size(), "Should handle bulk insert")

			for i := 0; i < count; i++ {
				err, item := q.Deque()
				assert.NoError(t, err)
				assert.Equal(t, i, item, "Should maintain order")
			}
		})
	})

	t.Run("Zero Value Handling", func(t *testing.T) {
		q := NewQueue[*string]()

		// Test nil pointer
		q.Enque(nil)
		err, item := q.Deque()
		assert.NoError(t, err)
		assert.Nil(t, item, "Should handle nil values")

		// Test zero value for struct
		type custom struct{ field int }
		q2 := NewQueue[custom]()
		q2.Enque(custom{})
		err2, item2 := q2.Deque()
		assert.NoError(t, err2)
		assert.Equal(t, custom{}, item2, "Should handle zero-value structs")
	})
}
