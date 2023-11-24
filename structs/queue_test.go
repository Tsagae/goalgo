package structs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueue(t *testing.T) {
	queue := NewQueue[int]()

	assert.True(t, queue.IsEmpty())

	queue.Enqueue(10)
	assert.Equal(t, 10, queue.Peek())

	queue.Enqueue(20)
	assert.Equal(t, 10, queue.Dequeue())

	assert.Equal(t, 1, queue.Size())

}
