package structs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDequeue(t *testing.T) {
	numbers := []int{3, 5, 1, 0, 9, 6, 4, 7, 2, 8}
	queue := NewPrioQueue[int, int]()

	if queue.Size() != 0 {
		t.Errorf("Incorrect queue initial size")
	}

	for _, v := range numbers {
		queue.Insert(v, v)
	}

	for i := 0; i < len(numbers); i++ {
		peeked := queue.Peek()
		dequeued := queue.Dequeue()
		if peeked != dequeued {
			t.Errorf("Incorrect element peeked")
		}

		if i != dequeued {
			t.Errorf("Incorrect element dequeued")
		}
		if queue.Size() != len(numbers)-i-1 {
			t.Errorf("Incorrect queue size: found: %v expected: %v", queue.Size(), len(numbers)-i-1)
		}
	}

}

func TestRemove(t *testing.T) {
	numbers := []int{3, 5, 1, 0, 9, 6, 4, 7, 2, 8}
	removeOrder := []int{9, 3, 8, 0, 1, 6, 7, 2, 5, 4}
	queue := NewPrioQueue[int, int]()

	for _, v := range numbers {
		queue.Insert(v, v)
	}

	for i, v := range removeOrder {
		if v != queue.Remove(v) {
			t.Errorf("Incorrect element removed")
		}
		if queue.Size() != len(numbers)-i-1 {
			t.Errorf("Incorrect queue size: found: %v expected: %v", queue.Size(), len(numbers)-i-1)
		}
	}

}

func TestPrioQueue_ChangePriority(t *testing.T) {
	values := []int{3, 5, 1, 0, 9, 6, 4, 7, 2, 8}
	removeOrder := []int{6, 1, 2, 3, 4, 5, 7, 8, 9, 0}
	queue := NewPrioQueue[int, int]()

	for _, v := range values {
		queue.Insert(v, v)
	}

	queue.ChangePriority(6, -1)
	queue.ChangePriority(0, 100)

	for i, v := range removeOrder {
		assert.Equal(t, v, queue.Dequeue())
		if queue.Size() != len(values)-i-1 {
			t.Errorf("Incorrect queue size: found: %v expected: %v", queue.Size(), len(values)-i-1)
		}
	}
}
