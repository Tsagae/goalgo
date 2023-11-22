package graph

import (
	"github.com/stretchr/testify/assert"
	"github.com/tsagae/algoritmi/structs"
	"testing"
)

func TestDFS(t *testing.T) {
	graph := getTestMapGraph()
	visited := structs.NewMapSet[string]()
	firstNode, err := graph.GetNode("a")
	assert.Equal(t, nil, err)

	testVisited := structs.NewMapSet[string]()
	testVisitedPtr := &testVisited

	DFS(firstNode, &visited, func(node Node[string, int]) {
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
	graph := getTestMapGraph()
	firstNode, err := graph.GetNode("a")
	assert.Equal(t, nil, err)

	testVisited := structs.NewMapSet[string]()
	testVisitedPtr := &testVisited

	BFS(firstNode, func(node Node[string, int]) {
		testVisitedPtr.Put(node.GetLabel())
	})

	assert.Equal(t, true, testVisited.Find("a"))
	assert.Equal(t, true, testVisited.Find("b"))
	assert.Equal(t, true, testVisited.Find("c"))
	assert.Equal(t, true, testVisited.Find("d"))
	assert.Equal(t, true, testVisited.Find("e"))
	assert.Equal(t, true, testVisited.Find("f"))
}
