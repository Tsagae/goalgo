package spanningtree

import (
	"github.com/stretchr/testify/assert"
	"github.com/tsagae/algoritmi/graph/undirected"
	"github.com/tsagae/algoritmi/structs"
	"testing"
)

func getGraph() undirected.UndirectedMapGraph[string, int] {
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

type simpleEdge struct {
	from   string
	to     string
	weight int
}

func checkEdge(expected simpleEdge, actualEdgesSet structs.MapSet[simpleEdge]) bool {
	flippedExpected := simpleEdge{expected.to, expected.from, expected.weight}
	toRet := actualEdgesSet.Find(expected) || actualEdgesSet.Find(flippedExpected)
	actualEdgesSet.Remove(expected)
	actualEdgesSet.Remove(flippedExpected)
	return toRet
}

func TestPrim(t *testing.T) {

	g := getGraph()
	nodeA, err := g.GetNode("A")
	assert.Nil(t, err)
	spanningTree := Prim(nodeA)

	simpleEdgesSet := structs.NewMapSet[simpleEdge]()

	allEdgesFromTree := spanningTree.GetAllEdges()
	for _, v := range allEdgesFromTree {
		simpleEdgesSet.Put(simpleEdge{v.GetNodeFrom().GetLabel(), v.GetNodeTo().GetLabel(), v.GetWeight()})
	}

	assert.True(t, checkEdge(simpleEdge{"A", "B", 1}, simpleEdgesSet))
	assert.True(t, checkEdge(simpleEdge{"C", "B", 2}, simpleEdgesSet))
	assert.True(t, checkEdge(simpleEdge{"B", "D", 3}, simpleEdgesSet))
	assert.Zero(t, simpleEdgesSet.Size())
}
