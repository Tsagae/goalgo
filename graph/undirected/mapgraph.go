package undirected

import (
	"github.com/tsagae/goalgo/graph"
	"github.com/tsagae/goalgo/graph/directed"
	"github.com/tsagae/goalgo/structs/set"
)

type UndirectedMapGraph[T comparable, W graph.Weight] struct {
	innerGraph directed.DirectedMapGraph[T, W]
}

// NewUndirectedMapGraph Constructor
func NewUndirectedMapGraph[T comparable, W graph.Weight]() UndirectedMapGraph[T, W] {
	return UndirectedMapGraph[T, W]{
		innerGraph: directed.NewDirectedMapGraph[T, W](),
	}
}

// Graph metods

// AddNode adds a node to the graph. Does nothing if the node is already present
func (g *UndirectedMapGraph[T, W]) AddNode(label T) {
	g.innerGraph.AddNode(label)
}

func (g *UndirectedMapGraph[T, W]) AddNodes(labels ...T) {
	for _, v := range labels {
		g.AddNode(v)
	}
}

func (g *UndirectedMapGraph[T, W]) RemoveNode(label T) {
	panic("not implemented")
}

// AddEdge adds an edge to the graph. Does nothing if an edge from "from" to "to" or vice versa already exists. If "from" or "to" do not exist, they are added
func (g *UndirectedMapGraph[T, W]) AddEdge(from T, to T, weight W) {
	g.innerGraph.AddEdge(from, to, weight)
	g.innerGraph.AddEdge(to, from, weight)
}

func (g *UndirectedMapGraph[T, W]) RemoveEdge(label T) {
	panic("not implemented")
}

// GetNode returns a node from the graph. Returns an error if the node is not found
func (g *UndirectedMapGraph[T, W]) GetNode(label T) (graph.Node[T, W], error) {
	node, err := g.innerGraph.GetNode(label)
	return node, err
}

// Gets all edges from the graph without duplicates. If the graph contains an edge {a,b,3} it will return either {a,b,3} or {b,a,3} but not both
func (g *UndirectedMapGraph[T, W]) GetAllEdges() []graph.Edge[T, W] {
	type edgeForSet[T comparable, W graph.Weight] struct {
		from   T
		to     T
		weight W
	}
	edgeSet := set.NewMapSet[edgeForSet[T, W]]()
	edgeList := g.innerGraph.GetAllEdges()
	edgeListToRet := make([]graph.Edge[T, W], 0, len(edgeList)/2)
	for _, v := range edgeList {
		edge := edgeForSet[T, W]{
			from:   v.GetNodeFrom().GetLabel(),
			to:     v.GetNodeTo().GetLabel(),
			weight: v.GetWeight(),
		}
		flipped := edgeForSet[T, W]{
			from:   edge.to,
			to:     edge.from,
			weight: edge.weight,
		}
		if edgeSet.Find(edge) || edgeSet.Find(flipped) {
			continue
		}
		edgeSet.Put(edge)
		edgeListToRet = append(edgeListToRet, v)
	}

	return edgeListToRet
}
