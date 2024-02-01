package structs

import (
	"fmt"
	"strings"
)

type LinkedList[T any] struct {
	head *ListNode[T]
	tail *ListNode[T]
	size int
}

func NewLinkedList[T any]() LinkedList[T] {
	return LinkedList[T]{head: nil, tail: nil}
}

func (list *LinkedList[T]) AddFirst(val T) {
	if list.size == 0 {
		newNode := &ListNode[T]{val, nil, nil, list}
		list.head = newNode
		list.tail = newNode
		list.size++
	} else {
		list.head.AddBefore(val)
	}
}

func (list *LinkedList[T]) AddLast(val T) {
	if list.size == 0 {
		newNode := &ListNode[T]{val, nil, nil, list}
		list.head = newNode
		list.tail = newNode
		list.size++
	} else {
		list.tail.AddAfter(val)
	}
}

func (list *LinkedList[T]) Add(index int, itemToInsert T) {
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
func (list *LinkedList[T]) GetNode(index int) *ListNode[T] {
	return list.getNode(index)
}

// Get Result is undefined if called outside the range of the list
func (list *LinkedList[T]) Get(index int) T {
	var valToRet T
	node := list.getNode(index)
	if node != nil {
		valToRet = node.val
	}
	return valToRet
}

// GetFirst Result is undefined if called on empty list
func (list *LinkedList[T]) GetFirst() T {
	return list.Get(0)
}

// GetLast Result is undefined if called on empty list
func (list *LinkedList[T]) GetLast() T {
	return list.Get(list.size - 1)
}

// Remove Result is undefined if called outside the range of the list
func (list *LinkedList[T]) Remove(index int) T {
	var removedVal T
	node := list.getNode(index)
	if node != nil {
		removedVal = node.val
		node.Remove()
	}
	return removedVal
}

// RemoveFirst Result is undefined if called outside the range of the list
func (list *LinkedList[T]) RemoveFirst() T {
	return list.Remove(0)
}

// RemoveLast Result is undefined if called outside the range of the list
func (list *LinkedList[T]) RemoveLast() T {
	return list.Remove(list.size - 1)
}

func (list *LinkedList[T]) Size() int {
	return list.size
}

func (list *LinkedList[T]) IsEmpty() bool {
	return list.head == nil
}

func (list *LinkedList[T]) ToString() string {
	var sb strings.Builder
	iter := list.Iterator()
	sb.WriteString("[")
	for iter.HasNext() {
		sb.WriteString(fmt.Sprintf("%v", iter.Next()))
		if iter.HasNext() {
			sb.WriteString(", ")
		}
	}
	if sb.Len() == 1 {
		return "[]"
	}
	sb.WriteString("]")
	return sb.String()
}

func (list *LinkedList[T]) getNode(index int) *ListNode[T] {
	if index < 0 || index >= list.size {
		return nil
	}
	if index == 0 {
		return list.head
	} else if index == list.size-1 {
		return list.tail
	}
	var node *ListNode[T]
	if index > list.size/2 {
		node = list.getFromBottom(index)
	} else {
		node = list.getFromTop(index)
	}
	return node
}

func (list *LinkedList[T]) getFromTop(index int) *ListNode[T] {
	curNode := list.head
	for i := 0; i < list.size; i++ {
		if i == index {
			return curNode
		}
		curNode = curNode.next
	}
	return nil
}

func (list *LinkedList[T]) getFromBottom(index int) *ListNode[T] {
	curNode := list.tail
	for i := list.size - 1; i >= 0; i-- {
		if i == index {
			return curNode
		}
		curNode = curNode.prev
	}
	return nil
}
