package structs

type Stack[T any] struct {
	innerStack []T
}

func NewStack[T any]() Stack[T] {
	var stack Stack[T]
	stack.innerStack = make([]T, 0)
	return stack
}

func (stack *Stack[T]) Size() int {
	return len(stack.innerStack)
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
