package structs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestListNode_AddAfter(t *testing.T) {
	list := NewLinkedList[int]()
	list.head = &ListNode[int]{0, nil, nil, &list}
	list.tail = &ListNode[int]{1, nil, nil, &list}
	list.head.next = list.tail
	list.tail.prev = list.head
	list.size = 2

	list.head.AddAfter(2)
	list.tail.AddAfter(3)
	//0,2,1,3
	assert.Equal(t, 0, list.head.val)
	assert.Equal(t, 2, list.head.next.val)
	assert.Equal(t, 1, list.tail.prev.val)
	assert.Equal(t, 3, list.tail.val)
	assert.Equal(t, 4, list.Size())
}

func TestListNode_AddBefore(t *testing.T) {
	list := NewLinkedList[int]()
	list.head = &ListNode[int]{0, nil, nil, &list}
	list.tail = &ListNode[int]{1, nil, nil, &list}
	list.head.next = list.tail
	list.tail.prev = list.head
	list.size = 2

	list.head.AddBefore(2)
	list.head.next.AddBefore(3)
	//2,3,0,1

	assert.Equal(t, 2, list.head.val)
	assert.Equal(t, 3, list.head.next.val)
	assert.Equal(t, 0, list.tail.prev.val)
	assert.Equal(t, 1, list.tail.val)
	assert.Equal(t, 4, list.Size())
}

func TestListNode_Remove(t *testing.T) {
	list := NewLinkedList[int]()
	list.head = &ListNode[int]{0, nil, nil, &list}
	list.tail = &ListNode[int]{1, nil, nil, &list}
	list.head.next = list.tail
	list.tail.prev = list.head
	list.size = 2

	list.head.AddBefore(2)
	list.head.next.AddBefore(3)

	list.head.next.Remove()
	//2,0,1

	assert.Equal(t, 2, list.head.val)
	assert.Equal(t, 0, list.head.next.val)
	assert.Equal(t, 1, list.tail.val)
	assert.Equal(t, 3, list.Size())

	list.head.Remove()
	//0,1

	assert.Equal(t, 0, list.head.val)
	assert.Equal(t, 1, list.tail.val)
	assert.Equal(t, 2, list.Size())

	list.tail.Remove()
	//0

	assert.Equal(t, 0, list.tail.val)
	assert.Equal(t, 0, list.head.val)
	assert.Equal(t, 1, list.Size())

	list.head.Remove()
	assert.Zero(t, list.size)
	assert.Nil(t, list.head)
	assert.Nil(t, list.tail)

}
