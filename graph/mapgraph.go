package graph

import (
	"fmt"
)

type MapGraph[T comparable, W Weight] struct {
	innerMap map[T][]MapGraphEdge[T, W]
}

type MapGraphNode[T comparable, W Weight] struct {
	label T
	graph *MapGraph[T, W]
}

type MapGraphEdge[T comparable, W Weight] struct {
	labelFrom T
	labelTo   T
	weight    W
	graph     *MapGraph[T, W]
}

// EdgeTuple Used only for constructing the graph
type EdgeTuple[T comparable, W Weight] struct {
	LabelTo T
	Weight  W
}

// NewMapGraph Constructor
func NewMapGraph[T comparable, W Weight]() MapGraph[T, W] {
	return MapGraph[T, W]{
		innerMap: make(map[T][]MapGraphEdge[T, W]),
	}
}

// Graph metods

// AddNode adds a node to the graph. Does nothing if the node is already present
func (g *MapGraph[T, W]) AddNode(label T) {
	if _, ok := g.innerMap[label]; ok {
		return
	}
	g.innerMap[label] = make([]MapGraphEdge[T, W], 0)
}

func (g *MapGraph[T, W]) AddNodes(labels ...T) {
	for _, v := range labels {
		g.AddNode(v)
	}
}

func (g *MapGraph[T, W]) RemoveNode(label T) {
	panic("not implemented")
}

// AddEdge adds an edge to the graph. Does nothing if an edge from "from" to "to" already exists (Weight is ignored in this case)
func (g *MapGraph[T, W]) AddEdge(from T, to T, weight W) {
	originalList := g.innerMap[from]
	for _, v := range originalList {
		if v.labelTo == to {
			return
		}
	}
	g.innerMap[from] = append(originalList,
		MapGraphEdge[T, W]{
			labelFrom: from,
			labelTo:   to,
			weight:    weight,
			graph:     g,
		})
}

func (g *MapGraph[T, W]) AddEdges(from T, tuples ...EdgeTuple[T, W]) {
	for _, v := range tuples {
		g.AddEdge(from, v.LabelTo, v.Weight)
	}
}

func (g *MapGraph[T, W]) RemoveEdge(label T) {
	panic("not implemented")
}

// GetNode returns a node from the graph. Returns an error if the node is not found
func (g *MapGraph[T, W]) GetNode(label T) (Node[T, W], error) {
	nodeToRet := MapGraphNode[T, W]{
		label: label,
		graph: nil,
	}
	if _, ok := g.innerMap[label]; ok {
		nodeToRet.graph = g
		return &nodeToRet, nil
	}
	return &nodeToRet, fmt.Errorf("node not %v found in graph", label)
}

// Node methods

func (n *MapGraphNode[T, W]) GetEdges() []Edge[T, W] {
	originalList := n.graph.innerMap[n.label]
	toReturn := make([]Edge[T, W], 0, len(originalList))
	for _, v := range originalList {
		toReturn = append(toReturn, &MapGraphEdge[T, W]{
			v.labelFrom,
			v.labelTo,
			v.weight,
			v.graph,
		})
	}
	return toReturn
}

func (n *MapGraphNode[T, W]) GetLabel() T {
	return n.label
}

// Edge methods

func (e *MapGraphEdge[T, W]) GetNodeFrom() Node[T, W] {
	return &MapGraphNode[T, W]{
		label: e.labelFrom,
		graph: e.graph,
	}
}

func (e *MapGraphEdge[T, W]) GetNodeTo() Node[T, W] {
	return &MapGraphNode[T, W]{
		label: e.labelTo,
		graph: e.graph,
	}
}

func (e *MapGraphEdge[T, W]) GetWeight() W {
	return e.weight
}
