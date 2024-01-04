package structs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDequeue(t *testing.T) {
	numbers := []int{3, 5, 1, 0, 9, 6, 4, 7, 2, 8}
	queue := NewPrioQueue[int, int]()

	assert.True(t, queue.IsEmpty())

	for _, v := range numbers {
		queue.Insert(v, v)
	}

	for i := 0; i < len(numbers); i++ {
		peeked := queue.Peek()
		dequeued := queue.Dequeue()
		assert.Equal(t, peeked, dequeued)

		assert.Equal(t, i, dequeued)

		assert.Equal(t, len(numbers)-i-1, queue.Size())
	}

}

func TestRemove(t *testing.T) {
	numbers := []int{3, 5, 1, 0, 9, 6, 4, 7, 2, 8}
	removeOrder := []int{9, 3, 8, 0, 1, 6, 7, 2, 5, 4}
	queue := NewPrioQueue[int, int]()

	for _, v := range numbers {
		queue.Insert(v, v)
	}

	for i, v := range removeOrder {

		assert.Equal(t, v, queue.Remove(v))
		assert.Equal(t, len(numbers)-i-1, queue.Size())
	}

}

func TestPrioQueue_ChangePriority(t *testing.T) {
	values := []int{3, 5, 1, 0, 9, 6, 4, 7, 2, 8}
	removeOrder := []int{6, 1, 2, 3, 4, 5, 7, 8, 9, 0}
	queue := NewPrioQueue[int, int]()

	for _, v := range values {
		queue.Insert(v, v)
	}

	queue.ChangePriority(6, -1)
	queue.ChangePriority(0, 100)

	for i, v := range removeOrder {
		assert.Equal(t, v, queue.Dequeue())
		assert.Equal(t, len(values)-i-1, queue.Size())
	}
}

func TestPrioQueue_GetPriority(t *testing.T) {
	values := []rune{'A', 'B', 'C', 'D', 'E'}
	priorities := []int{3, 10, 2, 4, 3}
	queue := NewPrioQueue[rune, int]()
	for i, v := range values {
		queue.Insert(v, priorities[i])
	}

	for i, v := range values {
		p, err := queue.GetPriority(v)
		assert.Nil(t, err)
		assert.Equal(t, priorities[i], p)
	}
}
