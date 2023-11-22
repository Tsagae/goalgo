package structs

import "cmp"

// Priority queue implemented with an heap. Lowest number -> highest priority
type PrioQueue[T comparable, P cmp.Ordered] struct {
	heap []prioQueueItem[T, P]
}

type prioQueueItem[T comparable, P cmp.Ordered] struct {
	priority P
	item     T
}

func NewPrioQueue[T comparable, P cmp.Ordered]() PrioQueue[T, P] {
	return PrioQueue[T, P]{heap: make([]prioQueueItem[T, P], 0)}
}

//size enqueue dequeue peek

func (q *PrioQueue[T, P]) Size() int {
	return len(q.heap)
}

func (q *PrioQueue[T, P]) Insert(item T, priority P) {
	//adds the item as the last leaf and reorders from bottom
	q.heap = append(q.heap, prioQueueItem[T, P]{item: item, priority: priority})
	q.reorderFromBottom(q.Size() - 1)
}

// Removes and returns the item with the highest priority
func (q *PrioQueue[T, P]) Dequeue() T {
	itemToRet := q.heap[0].item
	q.heap[0], q.heap[q.Size()-1] = q.heap[q.Size()-1], q.heap[0]
	q.heap = q.heap[0 : q.Size()-1]
	q.reorderRootIterative(0)
	return itemToRet
}

func (q *PrioQueue[T, P]) Remove(priority P) T {
	//swaps the element with a certain priority with the last leaf, removes the last one and reorders
	index := q.indexFromKey(priority)
	itemToRet := q.heap[index].item
	q.heap[index], q.heap[q.Size()-1] = q.heap[q.Size()-1], q.heap[index]
	q.heap = q.heap[0 : q.Size()-1]
	if index != len(q.heap) {
		if priority > q.heap[index].priority {
			q.reorderRootIterative(index)
		} else if priority < q.heap[index].priority {
			q.reorderFromBottom(index)
		}
	}
	return itemToRet
}

func (q *PrioQueue[T, P]) Peek() T {
	return q.heap[0].item
}

func (q *PrioQueue[T, P]) ChangePriority(item T, newPriority P) {
	itemIndex := -1
	var foundItem *prioQueueItem[T, P]
	for i := 0; i < len(q.heap); i++ {
		if q.heap[i].item == item {
			itemIndex = i
			foundItem = &q.heap[i]
			break
		}
	}
	if itemIndex == -1 { //item not found
		return
	}
	oldPriority := foundItem.priority
	foundItem.priority = newPriority

	if newPriority > oldPriority {
		q.reorderRootIterative(itemIndex)
	} else {
		q.reorderFromBottom(itemIndex)
	}
}

func (q *PrioQueue[T, P]) reorderRootIterative(index int) {
	for {
		//looking if the root is smaller than the smallest of the two children
		l, r := q.getIndexChildren(index)
		var minIndex int
		if l == -1 && r == -1 {
			return
		} else if l == -1 {
			minIndex = r
		} else if r == -1 {
			minIndex = l
		} else {
			if q.heap[l].priority < q.heap[r].priority {
				minIndex = l
			} else {
				minIndex = r
			}
		}
		//if true swap the root with the smallest children
		if q.heap[index].priority > q.heap[minIndex].priority {
			q.heap[index], q.heap[minIndex] = q.heap[minIndex], q.heap[index]
			index = minIndex
		} else {
			return
		}
	}
}

func (q *PrioQueue[T, P]) getIndexChildren(index int) (int, int) {
	l := index*2 + 1
	r := index*2 + 2
	if l >= len(q.heap) {
		l = -1
	}
	if r >= len(q.heap) {
		r = -1
	}
	return l, r
}

func getIndexParent(index int) int {
	if index%2 == 0 {
		index /= 2
		index--
	} else {
		index /= 2
	}
	return index
}

func (q *PrioQueue[T, P]) reorderFromBottom(index int) {
	for {
		var lowerIndex int = index
		lowerIndex = getIndexParent(index)
		if lowerIndex < 0 {
			return
		}
		if q.heap[lowerIndex].priority > q.heap[index].priority {
			q.heap[lowerIndex], q.heap[index] = q.heap[index], q.heap[lowerIndex]
			index = lowerIndex
		} else {
			return
		}
	}
}

func (q *PrioQueue[T, P]) indexFromKey(key P) int {
	var i int
	var v prioQueueItem[T, P]
	for i, v = range q.heap {
		if v.priority == key {
			break
		}
	}
	return i
}
