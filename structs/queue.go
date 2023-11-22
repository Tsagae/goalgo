package structs

type Queue[T any] struct {
	innerList LinkedList[T]
}

func NewQueue[T any]() Queue[T] {
	return Queue[T]{NewLinkedList[T]()}
}

func (q *Queue[T]) Enqueue(item T) {
	q.innerList.Add(item)
}

// Result is undefined if called on empty queue
func (q *Queue[T]) Dequeue() T {
	return q.innerList.Remove(0)
}

// Result is undefined if called on empty queue
func (q *Queue[T]) Peek() T {
	return q.innerList.Get(0)
}

func (q *Queue[T]) Size() int {
	return q.innerList.Size()
}

func (q *Queue[T]) IsEmpty() bool {
	return q.innerList.IsEmpty()
}

func (q *Queue[T]) ToString() string {
	return q.innerList.ToString()
}