package undirected

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func getGraph() UndirectedMapGraph[string, int] {
	g := NewUndirectedMapGraph[string, int]()
	g.AddNodes("A", "B", "C", "D")

	g.AddEdge("A", "B", 1)
	g.AddEdge("A", "C", 3)
	g.AddEdge("A", "D", 5)

	g.AddEdge("B", "C", 2)
	g.AddEdge("B", "D", 3)

	g.AddEdge("C", "D", 4)

	return g
}

func TestUndirectedMapGraph_GetNode(t *testing.T) {
	g := getGraph()
	_, err := g.GetNode("F")
	assert.NotNil(t, err)

	nodeA, err := g.GetNode("A")
	assert.Nil(t, err)

	assert.Equal(t, "A", nodeA.GetLabel())

	edges := nodeA.GetEdges()
	assert.Equal(t, 3, len(edges))
	for i, v := range edges {

		assert.Condition(t, func() (success bool) {
			from := v.GetNodeFrom().GetLabel()
			to := v.GetNodeTo().GetLabel()

			switch i {
			case 0:
				return (from == "A" && to == "B") || (from == "B" && to == "A")
			case 1:
				return (from == "A" && to == "C") || (from == "C" && to == "A")
			case 2:
				return (from == "A" && to == "D") || (from == "D" && to == "A")
			}
			return false
		})
	}

}

func TestUndirectedMapGraph_GetAllEdges(t *testing.T) {
	g := getGraph()
	type edgeForSet struct {
		from   string
		to     string
		weight int
	}
	edgeMap := make(map[edgeForSet]bool)
	edgeMap[edgeForSet{"B", "A", 1}] = false
	edgeMap[edgeForSet{"A", "C", 3}] = false
	edgeMap[edgeForSet{"D", "A", 5}] = false
	edgeMap[edgeForSet{"B", "C", 2}] = false
	edgeMap[edgeForSet{"D", "B", 3}] = false
	edgeMap[edgeForSet{"C", "D", 4}] = false

	edgeList := g.GetAllEdges()
	assert.Equal(t, 6, len(edgeList))
	for _, v := range edgeList {
		curEdge := edgeForSet{v.GetNodeFrom().GetLabel(), v.GetNodeTo().GetLabel(), v.GetWeight()}
		flippedEdge := edgeForSet{curEdge.to, curEdge.from, v.GetWeight()}
		if _, ok := edgeMap[curEdge]; ok {
			existsFlipped := edgeMap[flippedEdge]
			assert.Equal(t, false, existsFlipped)
			assert.Equal(t, false, edgeMap[curEdge])
			edgeMap[curEdge] = true
		} else if _, ok := edgeMap[flippedEdge]; ok {
			existsCur := edgeMap[curEdge]
			assert.Equal(t, false, existsCur)
			assert.Equal(t, false, edgeMap[flippedEdge])
			edgeMap[flippedEdge] = true
		} else {
			assert.Error(t, fmt.Errorf("edge not present in map"))
		}
	}
}
