package spanningtree

import (
	"github.com/tsagae/goalgo/graph"
	"github.com/tsagae/goalgo/graph/undirected"
	"github.com/tsagae/goalgo/structs/prioqueue"
	"github.com/tsagae/goalgo/structs/set"
)

// numOfNodes can be set to -1 if the total number of nodes is not known. Affects performance but result is correct
func Prim[T comparable, W graph.Weight](start graph.Node[T, W], numOfNodes int) undirected.UndirectedMapGraph[T, W] {
	spanningTree := undirected.NewUndirectedMapGraph[T, W]()
	spanningTreeNodes := set.NewMapSet[T]()
	edgesQueue := prioqueue.NewPrioQueue[graph.Edge[T, W], W]()

	spanningTreeNodes.Put(start.GetLabel())
	for _, edge := range start.GetEdges() {
		edgesQueue.Insert(edge, edge.GetWeight())
	}

	for edgesQueue.Size() != 0 {
		if spanningTreeNodes.Size() == numOfNodes {
			break
		}
		edge := edgesQueue.Dequeue()
		nodeTo := edge.GetNodeTo()
		nodeFrom := edge.GetNodeFrom()
		if !spanningTreeNodes.Find(nodeTo.GetLabel()) {
			spanningTreeNodes.Put(nodeTo.GetLabel())
			spanningTree.AddEdge(nodeFrom.GetLabel(), nodeTo.GetLabel(), edge.GetWeight())
			for _, nextEdge := range nodeTo.GetEdges() {
				if nextEdge.GetNodeTo().GetLabel() != nodeFrom.GetLabel() {
					edgesQueue.Insert(nextEdge, nextEdge.GetWeight())
				}
			}
		}
	}

	return spanningTree
}
