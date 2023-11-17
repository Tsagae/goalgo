package structs

import (
	"fmt"
)

type LinkedList[T any] struct {
	head *listNode[T]
	//TODO: keep a pointer to the tail too
	size int
}

type listNode[T any] struct {
	val  T
	next *listNode[T]
}

type ListIterator[T any] struct {
	node *listNode[T]
	//TODO: change iterator implementation with a yield function?
}

func (list *LinkedList[T]) Iterator() ListIterator[T] {
	return ListIterator[T]{node: list.head}
}

func (iter *ListIterator[T]) HasNext() bool {
	return iter.node != nil
}

// Result is undefined if iter.hasNext() is false
func (iter *ListIterator[T]) Next() T {
	valToRet := iter.node.val
	iter.node = iter.node.next
	return valToRet
}

func NewLinkedList[T any]() LinkedList[T] {
	return LinkedList[T]{head: nil}
}

func createNode[T any](val T, next *listNode[T]) *listNode[T] {
	return &listNode[T]{
		val:  val,
		next: next,
	}
}

func (list *LinkedList[T]) Add(val T) {
	lastNode := list.getLastNode()
	newNode := createNode(val, nil)
	if lastNode == nil {
		list.head = newNode
	} else {
		lastNode.next = newNode
	}
	list.size++
}

func (list *LinkedList[T]) Insert(index int, itemToInsert T) {
	if index < 0 {
		return
	}

	nodeToInsert := createNode[T](itemToInsert, nil)
	if index == 0 {
		nodeToInsert.next = list.head
		list.head = nodeToInsert
		list.size++
	} else {
		node := list.getNode(index - 1)
		if node == nil {
			return
		}
		nodeToInsert.next = node.next
		node.next = nodeToInsert
		list.size++
	}
}

func (list *LinkedList[T]) getLastNode() *listNode[T] {
	currentNode := list.head
	if currentNode == nil {
		return nil
	}
	for currentNode.next != nil {
		currentNode = currentNode.next
	}
	return currentNode
}

// Result is undefined if called on empty list
func (list *LinkedList[T]) GetLast() T {
	currentNode := list.head
	if currentNode == nil {
		var zeroVal T
		return zeroVal
	}
	for currentNode.next != nil {
		currentNode = currentNode.next
	}
	return currentNode.val
}

func (list *LinkedList[T]) getNode(index int) *listNode[T] {
	if index < 0 {
		return nil
	}
	currentNode := list.head
	if currentNode == nil {
		return nil
	}
	var i int = 0
	for currentNode != nil {
		if i == index {
			break
		}
		currentNode = currentNode.next
		i++
	}
	return currentNode
}

// Result is undefined if called outside the range of the list
func (list *LinkedList[T]) Get(index int) T {
	node := list.getNode(index)
	if node == nil {
		var zeroVal T
		return zeroVal
	}
	return node.val
}

// Result is undefined if called outside the range of the list
func (list *LinkedList[T]) Remove(index int) T {
	if index < 0 || list.IsEmpty() {
		var zeroVal T
		return zeroVal
	}
	if index == 0 {
		prevHead := list.head
		list.head = list.head.next
		list.size--
		return prevHead.val
	} else {
		node := list.getNode(index - 1)
		nodeToRemove := node.next
		if nodeToRemove != nil {
			node.next = nodeToRemove.next
			list.size--
			return nodeToRemove.val
		}
	}
	var zeroVal T
	return zeroVal
}

func (list *LinkedList[T]) Size() int {
	return list.size
}

func (list *LinkedList[T]) IsEmpty() bool {
	return list.head == nil
}

func (list *LinkedList[T]) ToString() string {
	outStr := ""
	iter := list.Iterator()
	for iter.HasNext() {
		outStr += fmt.Sprintf("%v, ", iter.Next())
	}
	if len(outStr) == 0 {
		return "[]"
	}
	outStr = outStr[:len(outStr)-2]
	return "[" + outStr + "]"
}
