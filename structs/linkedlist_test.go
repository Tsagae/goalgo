package structs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewLinkedList(t *testing.T) {
	list := NewLinkedList[int]()
	assert.Zero(t, list.Size(), "Incorrect starting size")
}

func TestLinkedList_AddLast(t *testing.T) {
	list := NewLinkedList[int]()

	for i := 5; i < 10; i++ {
		list.AddLast(i)
	}

	assert.Equal(t, 5, list.Size(), "Incorrect size after AddLast")

	iterator := list.Iterator()
	index := 5
	for iterator.HasNext() {
		listItem := iterator.Next()
		assert.Equal(t, index, listItem)
		index++
	}

}
func TestLinkedList_Insert(t *testing.T) {
	list := NewLinkedList[int]()

	for i := 5; i < 10; i++ {
		list.AddLast(i)
	}

	list.Insert(2, 99)

	assert.Equal(t, list.Get(2), 99, "Incorrect Get or Insert")

	if list.Size() != 6 {
		t.Errorf("Incorrect size after Insert")
	}
}

func TestLinkedList_Remove(t *testing.T) {
	list := NewLinkedList[int]()

	list.AddLast(10)
	list.AddLast(20)
	list.AddLast(30)

	list.Remove(1)

	assert.Equal(t, 30, list.Get(1), "Incorrect Get or Remove")

	assert.Equal(t, 2, list.Size(), "Incorrect size after Remove")
}
