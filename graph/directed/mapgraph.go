package directed

import (
	"fmt"

	"github.com/tsagae/goalgo/graph"
)

type DirectedMapGraph[T comparable, W graph.Weight] struct {
	innerMap map[T][]MapGraphEdge[T, W]
}

type MapGraphNode[T comparable, W graph.Weight] struct {
	label T
	graph *DirectedMapGraph[T, W]
}

type MapGraphEdge[T comparable, W graph.Weight] struct {
	labelFrom T
	labelTo   T
	weight    W
	graph     *DirectedMapGraph[T, W]
}

// EdgeTuple Used only for constructing the graph
type EdgeTuple[T comparable, W graph.Weight] struct {
	LabelTo T
	Weight  W
}

// NewDirectedMapGraph Constructor
func NewDirectedMapGraph[T comparable, W graph.Weight]() DirectedMapGraph[T, W] {
	return DirectedMapGraph[T, W]{
		innerMap: make(map[T][]MapGraphEdge[T, W]),
	}
}

// Graph metods

// AddNode adds a node to the graph. Does nothing if the node is already present
func (g *DirectedMapGraph[T, W]) AddNode(label T) {
	if _, ok := g.innerMap[label]; ok {
		return
	}
	g.innerMap[label] = make([]MapGraphEdge[T, W], 0)
}

func (g *DirectedMapGraph[T, W]) AddNodes(labels ...T) {
	for _, v := range labels {
		g.AddNode(v)
	}
}

func (g *DirectedMapGraph[T, W]) RemoveNode(label T) {
	panic("not implemented")
}

// AddEdge adds an edge to the graph. Does nothing if an edge from "from" to "to" already exists. If "from" or "to" do not exist, they are added
func (g *DirectedMapGraph[T, W]) AddEdge(from T, to T, weight W) {
	g.AddNode(from)
	g.AddNode(to)
	originalList := g.innerMap[from]
	for _, v := range originalList {
		if v.labelTo == to && v.weight == weight {
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

func (g *DirectedMapGraph[T, W]) AddEdges(from T, tuples ...EdgeTuple[T, W]) {
	for _, v := range tuples {
		g.AddEdge(from, v.LabelTo, v.Weight)
	}
}

func (g *DirectedMapGraph[T, W]) RemoveEdge(label T) {
	panic("not implemented")
}

// GetNode returns a node from the graph. Returns an error if the node is not found
func (g *DirectedMapGraph[T, W]) GetNode(label T) (graph.Node[T, W], error) {
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

func (g *DirectedMapGraph[T, W]) GetAllEdges() []graph.Edge[T, W] {
	allEdges := make([]graph.Edge[T, W], 0, len(g.innerMap))
	for k := range g.innerMap {
		edgeList := g.innerMap[k]
		for i := 0; i < len(edgeList); i++ {
			allEdges = append(allEdges, &edgeList[i])
		}
	}
	return allEdges
}

// Node methods

func (n *MapGraphNode[T, W]) GetEdges() []graph.Edge[T, W] {
	originalList := n.graph.innerMap[n.label]
	toReturn := make([]graph.Edge[T, W], 0, len(originalList))
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

func (e *MapGraphEdge[T, W]) GetNodeFrom() graph.Node[T, W] {
	return &MapGraphNode[T, W]{
		label: e.labelFrom,
		graph: e.graph,
	}
}

func (e *MapGraphEdge[T, W]) GetNodeTo() graph.Node[T, W] {
	return &MapGraphNode[T, W]{
		label: e.labelTo,
		graph: e.graph,
	}
}

func (e *MapGraphEdge[T, W]) GetWeight() W {
	return e.weight
}
