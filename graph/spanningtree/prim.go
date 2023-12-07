package spanningtree

import (
	"github.com/tsagae/algoritmi/graph"
	"github.com/tsagae/algoritmi/graph/undirected"
	"github.com/tsagae/algoritmi/structs"
)

func Prim[T comparable, W graph.Weight](start graph.Node[T, W]) undirected.UndirectedMapGraph[T, W] {
	finalGraph := undirected.NewUndirectedMapGraph[T, W]()
	visited := structs.NewMapSet[T]()
	added := structs.NewMapSet[T]()
	edgesQueue := structs.NewPrioQueue[graph.Edge[T, W], W]()
	nodesQueue := structs.NewQueue[graph.Node[T, W]]()

	nodesQueue.Enqueue(start)

	for !edgesQueue.IsEmpty() {
		cur := edgesQueue.Dequeue()
		from := cur.GetNodeFrom().GetLabel()
		to := cur.GetNodeTo().GetLabel()

		if visited.Find(from) {
			continue
		}
		visited.Put(from)

		if !added.Find(to) {
			finalGraph.AddNodes(from, to)
			finalGraph.AddEdge(from, to, cur.GetWeight())
		}

		for _, edge := range cur.GetNodeTo().GetEdges() {
			if !visited.Find(edge.GetNodeTo().GetLabel()) {
				edgesQueue.Insert(edge, edge.GetWeight())
			}
		}
	}

	return finalGraph
}
