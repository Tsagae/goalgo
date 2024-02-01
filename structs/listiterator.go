package structs

type ListIterator[T any] struct {
	node *ListNode[T]
	//TODO: change iterator implementation with a yield function?
}

func (list *LinkedList[T]) Iterator() ListIterator[T] {
	return ListIterator[T]{node: list.head}
}

func (iter *ListIterator[T]) HasNext() bool {
	return iter.node != nil
}

// Next Result is undefined if iter.hasNext() is false. Behavior is undefined if the list is modified while using the iterator
func (iter *ListIterator[T]) Next() T {
	valToRet := iter.node.val
	iter.node = iter.node.next
	return valToRet
}
