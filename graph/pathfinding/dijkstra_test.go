package pathfinding

import (
	"github.com/stretchr/testify/assert"
	"github.com/tsagae/goalgo/graph"
	"github.com/tsagae/goalgo/graph/directed"
	"testing"
)

// Origin, destination
func testGraph(t *testing.T) (graph.Node[string, int], graph.Node[string, int]) {
	g := directed.NewDirectedMapGraph[string, int]()
	g.AddNodes("A", "B", "C", "D", "E", "F")
	g.AddEdges("A",
		directed.EdgeTuple[string, int]{LabelTo: "B", Weight: 4},
		directed.EdgeTuple[string, int]{LabelTo: "C", Weight: 2},
	)
	g.AddEdges("B",
		directed.EdgeTuple[string, int]{LabelTo: "C", Weight: 5},
		directed.EdgeTuple[string, int]{LabelTo: "D", Weight: 10},
	)
	g.AddEdges("C",
		directed.EdgeTuple[string, int]{LabelTo: "E", Weight: 3},
	)
	g.AddEdges("D",
		directed.EdgeTuple[string, int]{LabelTo: "F", Weight: 11},
	)
	g.AddEdges("E",
		directed.EdgeTuple[string, int]{LabelTo: "D", Weight: 4},
	)

	start, err := g.GetNode("A")
	assert.Equal(t, nil, err)
	destination, err := g.GetNode("F")
	assert.Equal(t, nil, err)

	return start, destination
}

func TestDijkstra(t *testing.T) {
	start, destination := testGraph(t)
	path, distances := Dijkstra(start, destination)

	expectedDistances := map[string]int{
		"A": 0, "B": 4, "C": 2, "D": 9, "E": 5, "F": 20,
	}
	assert.Equal(t, expectedDistances, distances)

	expectedPath := []string{"A", "C", "E", "D", "F"}
	actualPath := make([]string, 0, len(path))
	for _, v := range path {
		actualPath = append(actualPath, v.GetLabel())
	}
	assert.Equal(t, expectedPath, actualPath)

}
