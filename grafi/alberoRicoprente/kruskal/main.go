	package main

	import (
		"fmt"
		"sort"
	)

	type Grafo struct {
		innerGrafo map[string]map[string]int
	}

	type Arco struct {
		x    string
		y    string
		peso int
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
		listaArchi := g.getListaArchi()
		albero := newGrafo()
		for _, v := range listaArchi {
			albero.aggiungi(v.x, v.y, v.peso)
			if albero.dfsTrovaCicli("", v.x, make(map[string]bool)) { // da fare con quickunion e find con compressione di cammino invece di usare dfs
				albero.rimuovi(v.x, v.y)
			}
		}
		return albero
	}

	// restituisce una lista di archi in ordine crescente di peso
	func (g *Grafo) getListaArchi() []Arco {
		archiSet := make(map[Arco]bool)
		for k, v := range g.innerGrafo {
			for k2, v2 := range v {
				arco := Arco{x: k, y: k2, peso: v2}
				if !(archiSet[arco] || archiSet[Arco{x: k2, y: k, peso: v2}]) {
					archiSet[arco] = true
				}
			}
		}
		archiList := make([]Arco, 0, len(archiSet))
		for k := range archiSet {
			archiList = append(archiList, k)
		}
		sort.SliceStable(archiList, func(i, j int) bool {
			return archiList[i].peso < archiList[j].peso
		})

		return archiList
	}

	// return true se ha trovato un ciclo
	func (g *Grafo) dfsTrovaCicli(from string, ver string, visitati map[string]bool) bool {
		visitati[ver] = true
		//fmt.Println("visitando: ", ver) //TODO debug print
		//fmt.Println("intorno: ", g.innerGrafo[ver])
		for k := range g.innerGrafo[ver] {
			if k != from {
				//fmt.Println("da: ", from, " k: ", k, " visitati: ", visitati)
				if visitati[k] {
					return true
				} else {
					visitati[k] = true
					if g.dfsTrovaCicli(ver, k, visitati) {
						return true
					}
				}
			}
		}
		return false
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
