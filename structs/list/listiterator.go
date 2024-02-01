package list

type Iterator[T any] struct {
	node *Node[T]
	//TODO: change iterator implementation with a yield function?
}

func (list *List[T]) Iterator() Iterator[T] {
	return Iterator[T]{node: list.head}
}

func (iter *Iterator[T]) HasNext() bool {
	return iter.node != nil
}

// Next Result is undefined if iter.hasNext() is false. Behavior is undefined if the list is modified while using the iterator
func (iter *Iterator[T]) Next() T {
	valToRet := iter.node.val
	iter.node = iter.node.next
	return valToRet
}
