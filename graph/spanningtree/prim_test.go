package spanningtree

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tsagae/algoritmi/graph"
	"github.com/tsagae/algoritmi/graph/undirected"
	"github.com/tsagae/algoritmi/structs"
)

func getGraph1() undirected.UndirectedMapGraph[string, int] {
	g := undirected.NewUndirectedMapGraph[string, int]()
	g.AddNodes("A", "B", "C", "D")

	g.AddEdge("A", "B", 1)
	g.AddEdge("A", "C", 3)
	g.AddEdge("A", "D", 5)

	g.AddEdge("B", "C", 2)
	g.AddEdge("B", "D", 3)

	g.AddEdge("C", "D", 4)

	return g
}

func getGraph2() undirected.UndirectedMapGraph[int, int] {
	g := undirected.NewUndirectedMapGraph[int, int]()
	g.AddEdge(0, 1, 4)
	g.AddEdge(0, 7, 8)

	g.AddEdge(1, 7, 11)
	g.AddEdge(1, 2, 8)

	g.AddEdge(2, 8, 2)
	g.AddEdge(2, 3, 7)
	g.AddEdge(2, 5, 4)

	g.AddEdge(3, 4, 9)
	g.AddEdge(3, 5, 14)

	g.AddEdge(4, 5, 10)

	g.AddEdge(5, 6, 2)

	g.AddEdge(6, 8, 6)
	g.AddEdge(6, 7, 1)

	g.AddEdge(7, 8, 7)

	return g
}

type simpleEdge[T comparable, W graph.Weight] struct {
	from   T
	to     T
	weight W
}

func checkEdge[T comparable, W graph.Weight](expected simpleEdge[T, W], actualEdgesSet structs.MapSet[simpleEdge[T, W]]) bool {
	flippedExpected := simpleEdge[T, W]{expected.to, expected.from, expected.weight}
	toRet := actualEdgesSet.Find(expected) || actualEdgesSet.Find(flippedExpected)
	actualEdgesSet.Remove(expected)
	actualEdgesSet.Remove(flippedExpected)
	return toRet
}

func TestPrim1(t *testing.T) {

	g := getGraph1()
	nodeA, err := g.GetNode("A")
	assert.Nil(t, err)
	spanningTree := Prim(nodeA, -1)

	simpleEdgesSet := structs.NewMapSet[simpleEdge[string, int]]()

	allEdgesFromTree := spanningTree.GetAllEdges()
	for _, v := range allEdgesFromTree {
		simpleEdgesSet.Put(simpleEdge[string, int]{v.GetNodeFrom().GetLabel(), v.GetNodeTo().GetLabel(), v.GetWeight()})
	}

	assert.True(t, checkEdge(simpleEdge[string, int]{"A", "B", 1}, simpleEdgesSet))
	assert.True(t, checkEdge(simpleEdge[string, int]{"C", "B", 2}, simpleEdgesSet))
	assert.True(t, checkEdge(simpleEdge[string, int]{"B", "D", 3}, simpleEdgesSet))
	assert.Zero(t, simpleEdgesSet.Size())

	spanningTree = Prim(nodeA, 3)

	simpleEdgesSet = structs.NewMapSet[simpleEdge[string, int]]()

	allEdgesFromTree = spanningTree.GetAllEdges()
	for _, v := range allEdgesFromTree {
		simpleEdgesSet.Put(simpleEdge[string, int]{v.GetNodeFrom().GetLabel(), v.GetNodeTo().GetLabel(), v.GetWeight()})
	}

	assert.True(t, checkEdge(simpleEdge[string, int]{"A", "B", 1}, simpleEdgesSet))
	assert.True(t, checkEdge(simpleEdge[string, int]{"C", "B", 2}, simpleEdgesSet))
	assert.False(t, checkEdge(simpleEdge[string, int]{"B", "D", 3}, simpleEdgesSet))
	assert.Equal(t, 0, simpleEdgesSet.Size())
}

func TestPrim2(t *testing.T) {

	g := getGraph2()
	nodeA, err := g.GetNode(0)
	assert.Nil(t, err)
	spanningTree := Prim(nodeA, 9)

	simpleEdgesSet := structs.NewMapSet[simpleEdge[int, int]]()

	allEdgesFromTree := spanningTree.GetAllEdges()
	for _, v := range allEdgesFromTree {
		simpleEdgesSet.Put(simpleEdge[int, int]{v.GetNodeFrom().GetLabel(), v.GetNodeTo().GetLabel(), v.GetWeight()})
	}

	assert.True(t, checkEdge(simpleEdge[int, int]{0, 1, 4}, simpleEdgesSet))
	assert.True(t, checkEdge(simpleEdge[int, int]{0, 7, 8}, simpleEdgesSet))
	assert.True(t, checkEdge(simpleEdge[int, int]{7, 6, 1}, simpleEdgesSet))
	assert.True(t, checkEdge(simpleEdge[int, int]{6, 5, 2}, simpleEdgesSet))
	assert.True(t, checkEdge(simpleEdge[int, int]{5, 2, 4}, simpleEdgesSet))
	assert.True(t, checkEdge(simpleEdge[int, int]{2, 8, 2}, simpleEdgesSet))
	assert.True(t, checkEdge(simpleEdge[int, int]{2, 3, 7}, simpleEdgesSet))
	assert.True(t, checkEdge(simpleEdge[int, int]{3, 4, 9}, simpleEdgesSet))
	assert.Zero(t, simpleEdgesSet.Size())

	spanningTree = Prim(nodeA, 9)

	simpleEdgesSet = structs.NewMapSet[simpleEdge[int, int]]()

	allEdgesFromTree = spanningTree.GetAllEdges()
	for _, v := range allEdgesFromTree {
		simpleEdgesSet.Put(simpleEdge[int, int]{v.GetNodeFrom().GetLabel(), v.GetNodeTo().GetLabel(), v.GetWeight()})
	}

	assert.True(t, checkEdge(simpleEdge[int, int]{0, 1, 4}, simpleEdgesSet))
	assert.True(t, checkEdge(simpleEdge[int, int]{0, 7, 8}, simpleEdgesSet))
	assert.True(t, checkEdge(simpleEdge[int, int]{7, 6, 1}, simpleEdgesSet))
	assert.True(t, checkEdge(simpleEdge[int, int]{6, 5, 2}, simpleEdgesSet))
	assert.True(t, checkEdge(simpleEdge[int, int]{5, 2, 4}, simpleEdgesSet))
	assert.True(t, checkEdge(simpleEdge[int, int]{2, 8, 2}, simpleEdgesSet))
	assert.True(t, checkEdge(simpleEdge[int, int]{2, 3, 7}, simpleEdgesSet))
	assert.True(t, checkEdge(simpleEdge[int, int]{3, 4, 9}, simpleEdgesSet))
	assert.Zero(t, simpleEdgesSet.Size())
}
