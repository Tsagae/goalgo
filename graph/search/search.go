package search

import (
	"github.com/tsagae/algoritmi/graph"
	"github.com/tsagae/algoritmi/structs"
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

func BFS[T comparable, W graph.Weight](startingNode graph.Node[T, W], f func(node2 graph.Node[T, W])) {
	queue := structs.NewQueue[graph.Node[T, W]]()
	visited := structs.NewMapSet[T]()

	queue.Enqueue(startingNode)
	visited.Put(startingNode.GetLabel())

	for !queue.IsEmpty() {
		currentNode := queue.Dequeue()
		f(currentNode)
		for _, v := range currentNode.GetEdges() {
			if !visited.Find(v.GetNodeTo().GetLabel()) {
				queue.Enqueue(v.GetNodeTo())
				visited.Put(v.GetNodeTo().GetLabel())
			}
		}
	}
}
