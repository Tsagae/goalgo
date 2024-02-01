package stack

import (
	"fmt"
	"strings"
)

type Stack[T any] struct {
	innerStack []T
}

func NewStack[T any]() Stack[T] {
	return Stack[T]{innerStack: make([]T, 0)}
}

func (stack *Stack[T]) Size() int {
	return len(stack.innerStack)
}

func (stack *Stack[T]) IsEmpty() bool {
	return stack.Size() == 0
}

func (stack *Stack[T]) Push(item T) {
	stack.innerStack = append(stack.innerStack, item)
}

func (stack *Stack[T]) Pop() T {
	nodoToRet := stack.Peek()
	stack.innerStack = stack.innerStack[:stack.Size()-1]
	return nodoToRet
}

func (stack *Stack[T]) Peek() T {
	return stack.innerStack[stack.Size()-1]
}

func (stack *Stack[T]) ToString() string {
	var sb strings.Builder
	sb.WriteString("| ")
	for _, v := range stack.innerStack {
		sb.WriteString(fmt.Sprintf("%v ", v))
	}
	sb.WriteString("<- Top")
	return sb.String()
}
