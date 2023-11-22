package pathfinding

import (
	"github.com/tsagae/algoritmi/graph"
	"github.com/tsagae/algoritmi/structs"
	"slices"
)

// Returns path, distances from start to target
func Dijkstra[T comparable, W graph.Weight](start graph.Node[T, W], target graph.Node[T, W]) ([]graph.Node[T, W], map[T]W) {
	distances := make(map[T]W)
	previous := make(map[T]graph.Node[T, W])
	visited := structs.NewMapSet[T]()
	queue := structs.NewPrioQueue[graph.Node[T, W], W]()
	var zeroValue W

	distances[start.GetLabel()] = zeroValue
	queue.Insert(start, 0)

	for !queue.IsEmpty() {
		node := queue.Dequeue()
		if node == target {
			break
		}
		if visited.Find(node.GetLabel()) {
			continue
		}
		visited.Put(node.GetLabel())
		for _, edge := range node.GetEdges() {
			neighbor := edge.GetNodeTo()
			newDistance := distances[node.GetLabel()] + edge.GetWeight()
			if !visited.Find(neighbor.GetLabel()) {
				queue.Insert(neighbor, newDistance)
			}
			if _, ok := distances[neighbor.GetLabel()]; ok {
				oldDistance := distances[neighbor.GetLabel()]
				if newDistance < oldDistance {
					distances[neighbor.GetLabel()] = newDistance
					queue.ChangePriority(node, newDistance)
					previous[edge.GetNodeTo().GetLabel()] = node
				}
			} else {
				distances[neighbor.GetLabel()] = newDistance
				queue.ChangePriority(node, newDistance)
				previous[edge.GetNodeTo().GetLabel()] = node
			}
		}
	}
	return getPathToArr(previous, start, target), distances
}

func getPathToArr[T comparable, W graph.Weight](previuosMap map[T]graph.Node[T, W], start graph.Node[T, W], targetNode graph.Node[T, W]) []graph.Node[T, W] {
	path := make([]graph.Node[T, W], 0) // Change this with linkedlist to append to head and then toArray?
	path = append(path, targetNode)
	startLabel := start.GetLabel()
	currentLabel := targetNode.GetLabel()

	for currentLabel != startLabel {
		currentNode := previuosMap[currentLabel]
		path = append(path, currentNode)
		currentLabel = currentNode.GetLabel()
	}
	slices.Reverse(path)
	return path
}
