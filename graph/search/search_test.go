package search

import (
	"github.com/stretchr/testify/assert"
	"github.com/tsagae/goalgo/graph"
	"github.com/tsagae/goalgo/graph/directed"
	"github.com/tsagae/goalgo/structs"
	"testing"
)

func getTestDirectedMapGraph() directed.DirectedMapGraph[string, int] {
	g := directed.NewDirectedMapGraph[string, int]()
	g.AddNode("a")
	g.AddNode("b")
	g.AddNode("c")
	g.AddNode("d")
	g.AddNode("e")
	g.AddNode("f")

	g.AddNodes("a", "b", "c", "d", "e", "f")

	g.AddEdges("a",
		directed.EdgeTuple[string, int]{"b", 1},
		directed.EdgeTuple[string, int]{"c", 1},
	)
	g.AddEdges("b",
		directed.EdgeTuple[string, int]{"d", 1},
		directed.EdgeTuple[string, int]{"f", 1},
		directed.EdgeTuple[string, int]{"c", 1},
		directed.EdgeTuple[string, int]{"a", 1},
	)
	g.AddEdges("c",
		directed.EdgeTuple[string, int]{"a", 1},
		directed.EdgeTuple[string, int]{"b", 1},
		directed.EdgeTuple[string, int]{"f", 1},
	)
	g.AddEdges("d",
		directed.EdgeTuple[string, int]{"b", 1},
		directed.EdgeTuple[string, int]{"e", 1},
		directed.EdgeTuple[string, int]{"f", 1},
	)
	g.AddEdges("e",
		directed.EdgeTuple[string, int]{"d", 1},
		directed.EdgeTuple[string, int]{"f", 1},
	)
	g.AddEdges("f",
		directed.EdgeTuple[string, int]{"c", 1},
		directed.EdgeTuple[string, int]{"b", 1},
		directed.EdgeTuple[string, int]{"e", 1},
	)
	return g
}

func TestDFS(t *testing.T) {
	g := getTestDirectedMapGraph()
	visited := structs.NewMapSet[string]()
	firstNode, err := g.GetNode("a")
	assert.Equal(t, nil, err)

	testVisited := structs.NewMapSet[string]()
	testVisitedPtr := &testVisited

	DFS(firstNode, &visited, func(node graph.Node[string, int]) {
		testVisitedPtr.Put(node.GetLabel())
	})

	assert.Equal(t, true, testVisited.Find("a"))
	assert.Equal(t, true, testVisited.Find("b"))
	assert.Equal(t, true, testVisited.Find("c"))
	assert.Equal(t, true, testVisited.Find("d"))
	assert.Equal(t, true, testVisited.Find("e"))
	assert.Equal(t, true, testVisited.Find("f"))
}
func TestBFS(t *testing.T) {
	g := getTestDirectedMapGraph()
	firstNode, err := g.GetNode("a")
	assert.Equal(t, nil, err)

	testVisited := structs.NewMapSet[string]()
	testVisitedPtr := &testVisited

	BFS(firstNode, func(node graph.Node[string, int]) {
		testVisitedPtr.Put(node.GetLabel())
	})

	assert.Equal(t, true, testVisited.Find("a"))
	assert.Equal(t, true, testVisited.Find("b"))
	assert.Equal(t, true, testVisited.Find("c"))
	assert.Equal(t, true, testVisited.Find("d"))
	assert.Equal(t, true, testVisited.Find("e"))
	assert.Equal(t, true, testVisited.Find("f"))
}
