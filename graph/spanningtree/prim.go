package spanningtree

import (
	"github.com/tsagae/algoritmi/graph"
	"github.com/tsagae/algoritmi/graph/undirected"
)

func Prim[T comparable, W graph.Weight](start graph.Node[T, W]) undirected.UndirectedMapGraph[T, W] {
	finalGraph := undirected.NewUndirectedMapGraph[T, W]()

	return finalGraph
}
