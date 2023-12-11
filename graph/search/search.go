package search

import (
	"github.com/tsagae/goalgo/graph"
	"github.com/tsagae/goalgo/structs"
)

func DFS[T comparable, W graph.Weight](startingNode graph.Node[T, W], visited structs.Set[T], f func(node2 graph.Node[T, W])) {
	//TODO: implement iteratively
	if visited.Find(startingNode.GetLabel()) {
		return
	}
	f(startingNode)
	visited.Put(startingNode.GetLabel())

	for _, v := range startingNode.GetEdges() {
		n := v.GetNodeTo()
		if !visited.Find(n.GetLabel()) {
			DFS(n, visited, f)
		}
	}
}

func BFS[T comparable, W graph.Weight](startingNode graph.Node[T, W], f func(nodeFrom graph.Node[T, W], nodeTo graph.Node[T, W])) {
	queue := structs.NewQueue[graph.Node[T, W]]()
	visited := structs.NewMapSet[T]()

	queue.Enqueue(startingNode)
	visited.Put(startingNode.GetLabel())

	for !queue.IsEmpty() {
		currentNode := queue.Dequeue()
		for _, v := range currentNode.GetEdges() {
			nodeTo := v.GetNodeTo()
			f(currentNode, nodeTo)
			if !visited.Find(nodeTo.GetLabel()) {
				queue.Enqueue(nodeTo)
				f(currentNode, nodeTo)
				visited.Put(nodeTo.GetLabel())
			}
		}
	}
}
