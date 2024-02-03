package list

type Node[T any] struct {
	val  T
	next *Node[T]
	prev *Node[T]
	list *List[T]
}

func (node *Node[T]) Value() T {
	return node.val
}

func (node *Node[T]) AddBefore(val T) {
	if node == node.list.head {
		//prev is null (node is the head)
		newNode := &Node[T]{val, node, nil, node.list}
		node.list.head = newNode
		node.prev = newNode
	} else {
		newNode := &Node[T]{val, node, node.prev, node.list}
		node.prev.next = newNode
		node.prev = newNode
	}
	node.list.size++
}

func (node *Node[T]) AddAfter(val T) {
	if node == node.list.tail {
		//next is null (node is the tail)
		newNode := &Node[T]{val, nil, node, node.list}
		node.next = newNode
		node.list.tail = newNode
	} else {
		next := node.next
		newNode := &Node[T]{val, next, node, node.list}
		next.prev = newNode
		node.next = newNode
	}
	node.list.size++
}

func (node *Node[T]) Remove() {
	if node.list.size == 1 {
		node.list.head = nil
		node.list.tail = nil
	} else if node == node.list.head {
		newHead := node.next
		newHead.prev = nil
		node.list.head = newHead
	} else if node == node.list.tail {
		newTail := node.prev
		newTail.next = nil
		node.list.tail = newTail
	} else {
		node.prev.next = node.next
		node.next.prev = node.prev
	}
	node.list.size--
	node.next = nil
	node.prev = nil
	node.list = nil
}
