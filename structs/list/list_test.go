package list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestList(t *testing.T) {
	list := NewList[int]()
	assert.Zero(t, list.Size(), "Incorrect starting size")
}

func TestList_AddLast(t *testing.T) {
	list := NewList[int]()

	for i := 5; i < 10; i++ {
		list.AddLast(i)
	}

	assert.Equal(t, 5, list.Size(), "Incorrect size after AddLast")

	node := list.GetNode(0)
	index := 5
	for node != nil {
		assert.Equal(t, index, node.Value())
		node = node.Next()
		index++
	}

}

func TestList_Add(t *testing.T) {
	list := NewList[int]()

	for i := 5; i < 10; i++ {
		list.AddLast(i)
	}

	list.Add(2, 99)
	//5,6,99,7,8,9

	assert.Equal(t, list.Get(2), 99, "Incorrect Get or Add")

	list.Add(4, 88)
	//5,6,99,7,88,8,9

	assert.Equal(t, 7, list.Size(), "Incorrect size after Add")
}

func TestList_AddFirst(t *testing.T) {
	list := NewList[int]()
	list.AddFirst(10)
	assert.Equal(t, 10, list.head.val)
	assert.Equal(t, 10, list.tail.val)
}

func TestList_Remove(t *testing.T) {
	list := NewList[int]()

	list.AddLast(10)
	list.AddLast(20)
	list.AddLast(30)

	list.Remove(1)

	assert.Equal(t, 30, list.Get(1), "Incorrect Get or Remove")

	assert.Equal(t, 2, list.Size(), "Incorrect size after Remove")
}

func TestList_GetFirst(t *testing.T) {
	list := NewList[int]()
	list.AddLast(10)
	list.AddLast(20)
	list.AddLast(30)
	assert.Equal(t, 10, list.GetFirst())
	list.Remove(0)
	assert.Equal(t, 20, list.GetFirst())
	list.AddFirst(50)
	assert.Equal(t, 50, list.GetFirst())
}

func TestList_GetLast(t *testing.T) {
	list := NewList[int]()
	list.AddLast(10)
	list.AddLast(20)
	assert.Equal(t, 20, list.GetLast())

	list.AddLast(30)
	assert.Equal(t, 30, list.GetLast())

	list.Remove(0)
	assert.Equal(t, 30, list.GetLast())

	list.Remove(list.Size() - 1)
	assert.Equal(t, 20, list.GetLast())
}

func TestList_RemoveFirst(t *testing.T) {
	list := NewList[int]()
	list.AddLast(10)
	list.AddLast(20)
	list.AddLast(30)
	list.RemoveFirst()
	assert.Equal(t, 20, list.GetFirst())
	list.RemoveFirst()
	assert.Equal(t, 30, list.GetFirst())
}

func TestList_RemoveLast(t *testing.T) {
	list := NewList[int]()
	list.AddLast(10)
	list.AddLast(20)
	list.AddLast(30)
	list.RemoveLast()
	assert.Equal(t, 20, list.GetLast())
	list.RemoveLast()
	assert.Equal(t, 10, list.GetLast())
}
