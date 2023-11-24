package graph

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func getTestMapGraph() MapGraph[string, int] {
	graph := NewMapGraph[string, int]()
	graph.AddNode("a")
	graph.AddNode("b")
	graph.AddNode("c")
	graph.AddNode("d")
	graph.AddNode("e")
	graph.AddNode("f")

	graph.AddNodes("a", "b", "c", "d", "e", "f")

	graph.AddEdges("a",
		EdgeTuple[string, int]{"b", 1},
		EdgeTuple[string, int]{"c", 1},
	)
	graph.AddEdges("b",
		EdgeTuple[string, int]{"d", 1},
		EdgeTuple[string, int]{"f", 1},
		EdgeTuple[string, int]{"c", 1},
		EdgeTuple[string, int]{"a", 1},
	)
	graph.AddEdges("c",
		EdgeTuple[string, int]{"a", 1},
		EdgeTuple[string, int]{"b", 1},
		EdgeTuple[string, int]{"f", 1},
	)
	graph.AddEdges("d",
		EdgeTuple[string, int]{"b", 1},
		EdgeTuple[string, int]{"e", 1},
		EdgeTuple[string, int]{"f", 1},
	)
	graph.AddEdges("e",
		EdgeTuple[string, int]{"d", 1},
		EdgeTuple[string, int]{"f", 1},
	)
	graph.AddEdges("f",
		EdgeTuple[string, int]{"c", 1},
		EdgeTuple[string, int]{"b", 1},
		EdgeTuple[string, int]{"e", 1},
	)
	return graph
}

func TestMapGraph(t *testing.T) {
	graph := getTestMapGraph()
	nodeD, err := graph.GetNode("d")
	assert.Equal(t, nil, err)
	assert.Equal(t, &MapGraphNode[string, int]{label: "d", graph: &graph}, nodeD)

	dEdges := make([]Edge[string, int], 0)
	dEdges = append(dEdges, &MapGraphEdge[string, int]{
		labelFrom: "d",
		labelTo:   "b",
		weight:    1,
		graph:     &graph,
	})
	dEdges = append(dEdges, &MapGraphEdge[string, int]{
		labelFrom: "d",
		labelTo:   "e",
		weight:    1,
		graph:     &graph,
	})
	dEdges = append(dEdges, &MapGraphEdge[string, int]{
		labelFrom: "d",
		labelTo:   "f",
		weight:    1,
		graph:     &graph,
	})

	actual := nodeD.GetEdges()
	expected := dEdges
	assert.Equal(t, expected, actual)
}

func TestMapGraph_AddEdge_Duplicate(t *testing.T) {
	graph := NewMapGraph[string, int]()
	graph.AddNode("a")
	graph.AddNode("b")
	graph.AddEdge("a", "b", 1)
	graph.AddEdge("a", "b", 2)
	graph.AddEdge("a", "b", 1)
	nodeA, err := graph.GetNode("a")
	assert.Nil(t, err)
	expected := nodeA.GetEdges()
	assert.ElementsMatch(t, expected,
		[]Edge[string, int]{
			&MapGraphEdge[string, int]{"a", "b", 1, &graph},
			&MapGraphEdge[string, int]{"a", "b", 2, &graph},
		})
}

func TestMapGraph_GetNode_NotFound(t *testing.T) {
	graph := NewMapGraph[string, int]()
	graph.AddNode("a")
	graph.AddNode("b")
	_, err := graph.GetNode("c")
	assert.NotNil(t, err)
}

func TestMapGraphEdge_EdgeProperties(t *testing.T) {
	graph := NewMapGraph[string, int]()
	graph.AddNode("a")
	graph.AddNode("b")
	graph.AddEdge("a", "b", 3)
	nodeA, err := graph.GetNode("a")
	assert.Nil(t, err)

	assert.Equal(t, "a", nodeA.GetEdges()[0].GetNodeFrom().GetLabel())
	assert.Equal(t, 3, nodeA.GetEdges()[0].GetWeight())
}
