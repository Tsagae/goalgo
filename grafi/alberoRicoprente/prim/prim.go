package main

import (
	"fmt"
)

// Grafo non orientato, ogni arco è rappresentato due volte A->B [2] | B->A [2]
type Grafo struct {
	innerGrafo map[string]map[string]int
}

type Arco struct {
	x    string
	y    string
	peso int
}

type MinHeap struct {
	innerMinHeap []Arco
}

func main() {
	g := newGrafo()
	fmt.Println(g)
	g.aggiungi("a", "d", 4)
	g.aggiungi("a", "b", 1)
	g.aggiungi("d", "b", 4)
	g.aggiungi("a", "e", 3)
	g.aggiungi("d", "e", 4)
	g.aggiungi("e", "b", 2)
	g.aggiungi("e", "c", 4)
	g.aggiungi("c", "f", 5)
	g.aggiungi("e", "f", 7)

	alberoRicoprente := g.getAlberoRicoprente()
	fmt.Println(alberoRicoprente.innerGrafo)
}

func (g *Grafo) getAlberoRicoprente() *Grafo {
	albero := newGrafo()
	var archi []Arco = make([]Arco, 0, len(g.innerGrafo)*2)
	for k, v := range g.innerGrafo {

	}
	return albero
}

func newGrafo() *Grafo {
	g := new(Grafo)
	g.innerGrafo = make(map[string]map[string]int)
	return g
}

func (g *Grafo) aggiungi(from string, to string, peso int) {
	if g.innerGrafo[from] == nil {
		g.innerGrafo[from] = make(map[string]int)
	}
	g.innerGrafo[from][to] = peso

	if g.innerGrafo[to] == nil {
		g.innerGrafo[to] = make(map[string]int)
	}
	g.innerGrafo[to][from] = peso
}

func (g *Grafo) rimuovi(from string, to string) {
	delete(g.innerGrafo[from], to)
	delete(g.innerGrafo[to], from)
}

// HEAP
func (h *MinHeap) getIndexFigli(index int) (int, int) {
	sx := 2*index + 1
	dx := 2*index + 2
	if sx >= len(h.innerMinHeap) || sx < 0 {
		sx = -1
	}
	if dx >= len(h.innerMinHeap) || dx < 0 {
		dx = -1
	}
	return sx, dx
}

func (h *MinHeap) getIndexPadre(index int) int {
	index /= 2
	if index%2 == 0 {
		index--
	}
	if index < 0 {
		index = -1
	}
	if index < 0 || index > len(h.innerMinHeap) {
		index = -1
	}
	return index
}

func (h *MinHeap) add(arco *Arco) {

}

func (h *MinHeap) sistemaRadice(radice int) {
	heap := h.innerMinHeap
	for {
		sx, dx := h.getIndexFigli(radice)
		var indexMinore int
		if sx < 0 && dx < 0 {
			return
		}
		if sx < 0 {
			indexMinore = dx
		} else if dx < 0 {
			indexMinore = sx
		} else {
			indexMinore = dx
			if heap[sx].peso < heap[dx].peso {
				indexMinore = sx
			}
		}
		if heap[radice].peso > heap[indexMinore].peso {
			heap[radice], heap[indexMinore] = heap[indexMinore], heap[radice]
		} else {
			return
		}
		radice = indexMinore
	}
}

func (h *MinHeap) sistemaDalBasso(figlio int) {
	heap := h.innerMinHeap
	for {
		padre := h.getIndexPadre(figlio)
		if padre < 0 {
			return
		}
		if heap[padre].peso < heap[figlio].peso {
			return
		}
		heap[padre], heap[figlio] = heap[figlio], heap[padre]
		figlio = padre
	}
}
