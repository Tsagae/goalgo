package main

import "fmt"

type Grafo struct {
	innerGrafo map[string][]Arco
}

type Arco struct {
	to   string
	peso int
}

func main() {
	fmt.Println("prova")
	g := newGrafo()
	g.add("a", "b", 3)
	g.add("a", "c", 2)
	g.add("b", "c", 4)
	g.add("b", "e", 5)
	g.add("c", "e", 5)
	g.add("e", "f", 1)
	g.add("f", "d", 6)
	g.add("d", "b", 6)

	fmt.Println(g)
	fmt.Println(g.dijkstra("f"))
}

func (g *Grafo) dijkstra(start string) map[string]int {
	distanze := make(map[string]int)
	distanze[start] = 0
	visitati := make(map[string]bool)
	daVisitare := make([]string, 0, len(g.innerGrafo))
	daVisitare = append(daVisitare, start)
	for len(daVisitare) != 0 {
		currentNodo := daVisitare[0]
		daVisitare = daVisitare[1:]
		if visitati[currentNodo] {
			continue
		}
		for _, v := range g.innerGrafo[currentNodo] {
			if !visitati[v.to] {
				daVisitare = append(daVisitare, v.to)
			}
			_, ok := distanze[v.to]
			if !ok || distanze[v.to] > distanze[currentNodo]+v.peso {
				distanze[v.to] = distanze[currentNodo] + v.peso
			}
		}
		visitati[currentNodo] = true
		//fmt.Println(visitati)
	}

	return distanze
}

func newGrafo() *Grafo {
	g := new(Grafo)
	g.innerGrafo = make(map[string][]Arco)
	return g
}

func (g *Grafo) add(from string, to string, peso int) {
	if g.innerGrafo[from] == nil {
		g.innerGrafo[from] = make([]Arco, 0, 1)
	}
	g.innerGrafo[from] = append(g.innerGrafo[from], Arco{to, peso})
}
