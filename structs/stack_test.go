package structs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack(t *testing.T) {
	stack := NewStack[int]()
	assert.Equal(t, true, stack.IsEmpty())

	stack.Push(10)
	assert.Equal(t, 10, stack.Peek())

	stack.Push(20)

	assert.Equal(t, 20, stack.Pop())

	assert.Equal(t, 1, stack.Size())
}
