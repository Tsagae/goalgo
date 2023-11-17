package structs

import (
	"testing"
)

func TestStack(t *testing.T) {
	stack := NewStack[int]()
	if stack.Size() != 0 {
		t.Errorf("Incorrect starting size")
	}

	stack.Push(10)
	if stack.Peek() != 10 {
		t.Errorf("Incorrect top element")
	}
	stack.Push(20)

	item := stack.Pop()
	if item != 20 {
		t.Errorf("Inccorrect element popped")
	}

	if stack.Size() != 1 {
		t.Errorf("Incorrect size")
	}

}
