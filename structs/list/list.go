package list

import (
	"fmt"
	"strings"
)

type List[T any] struct {
	head *Node[T]
	tail *Node[T]
	size int
}

func NewList[T any]() List[T] {
	return List[T]{head: nil, tail: nil}
}

func (list *List[T]) AddFirst(val T) {
	if list.size == 0 {
		newNode := &Node[T]{val, nil, nil, list}
		list.head = newNode
		list.tail = newNode
		list.size++
	} else {
		list.head.AddBefore(val)
	}
}

func (list *List[T]) AddLast(val T) {
	if list.size == 0 {
		newNode := &Node[T]{val, nil, nil, list}
		list.head = newNode
		list.tail = newNode
		list.size++
	} else {
		list.tail.AddAfter(val)
	}
}

func (list *List[T]) Add(index int, itemToInsert T) {
	if index < 0 || index > list.size {
		return
	} else if index == 0 {
		list.AddFirst(itemToInsert)
		return
	} else if index == list.size {
		list.AddLast(itemToInsert)
	} else {
		node := list.getNode(index)
		node.AddBefore(itemToInsert)
	}

}

// GetNode Result is undefined if called outside the range of the list
func (list *List[T]) GetNode(index int) *Node[T] {
	return list.getNode(index)
}

// Get Result is undefined if called outside the range of the list
func (list *List[T]) Get(index int) T {
	var valToRet T
	node := list.getNode(index)
	if node != nil {
		valToRet = node.val
	}
	return valToRet
}

// GetFirst Result is undefined if called on empty list
func (list *List[T]) GetFirst() T {
	return list.Get(0)
}

// GetLast Result is undefined if called on empty list
func (list *List[T]) GetLast() T {
	return list.Get(list.size - 1)
}

// Remove Result is undefined if called outside the range of the list
func (list *List[T]) Remove(index int) T {
	var removedVal T
	node := list.getNode(index)
	if node != nil {
		removedVal = node.val
		node.Remove()
	}
	return removedVal
}

// RemoveFirst Result is undefined if called outside the range of the list
func (list *List[T]) RemoveFirst() T {
	return list.Remove(0)
}

// RemoveLast Result is undefined if called outside the range of the list
func (list *List[T]) RemoveLast() T {
	return list.Remove(list.size - 1)
}

func (list *List[T]) Size() int {
	return list.size
}

func (list *List[T]) IsEmpty() bool {
	return list.head == nil
}

func (list *List[T]) ToString() string {
	var sb strings.Builder
	node := list.head
	sb.WriteString("[")
	for node != nil {
		sb.WriteString(fmt.Sprintf("%v", node.val))
		if node.next != nil {
			sb.WriteString(", ")
		}
		node = node.next
	}
	if sb.Len() == 1 {
		return "[]"
	}
	sb.WriteString("]")
	return sb.String()
}

func (list *List[T]) getNode(index int) *Node[T] {
	if index < 0 || index >= list.size {
		return nil
	}
	if index == 0 {
		return list.head
	} else if index == list.size-1 {
		return list.tail
	}
	var node *Node[T]
	if index > list.size/2 {
		node = list.getFromBottom(index)
	} else {
		node = list.getFromTop(index)
	}
	return node
}

func (list *List[T]) getFromTop(index int) *Node[T] {
	curNode := list.head
	for i := 0; i < list.size; i++ {
		if i == index {
			return curNode
		}
		curNode = curNode.next
	}
	return nil
}

func (list *List[T]) getFromBottom(index int) *Node[T] {
	curNode := list.tail
	for i := list.size - 1; i >= 0; i-- {
		if i == index {
			return curNode
		}
		curNode = curNode.prev
	}
	return nil
}
