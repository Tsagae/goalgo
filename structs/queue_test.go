package structs

import (
	"testing"
)

func TestQueue(t *testing.T) {
	queue := NewQueue[int]()
	if queue.Size() != 0 {
		t.Errorf("Incorrect starting size")
	}

	queue.Enqueue(10)
	if queue.Peek() != 10 {
		t.Errorf("Incorrect top element")
	}
	queue.Enqueue(20)

	item := queue.Dequeue()
	if item != 10 {
		t.Errorf("Inccorrect element dequeued")
	}

	if queue.Size() != 1 {
		t.Errorf("Incorrect size")
	}

}
