package binarytree

type TreeNode[T any] struct {
	Item T
	L    *TreeNode[T]
	R    *TreeNode[T]
}

func NewNode[T any](item T, l *TreeNode[T], r *TreeNode[T]) *TreeNode[T] {
	return &TreeNode[T]{
		Item: item,
		L:    l,
		R:    r,
	}
}

func (node *TreeNode[T]) RecursiveInfixSearch(f func(*T)) {
	if node != nil {
		node.L.RecursiveInfixSearch(f)
		f(&node.Item)
		node.R.RecursiveInfixSearch(f)
	}
}

func (node *TreeNode[T]) RecursivePrefixSearch(f func(*T)) {
	if node != nil {
		f(&node.Item)
		node.L.RecursivePrefixSearch(f)
		node.R.RecursivePrefixSearch(f)
	}
}

func (node *TreeNode[T]) RecursivePostfixSearch(f func(*T)) {
	if node != nil {
		node.L.RecursivePostfixSearch(f)
		node.R.RecursivePostfixSearch(f)
		f(&node.Item)
	}
}
