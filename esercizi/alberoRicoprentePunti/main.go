package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

// grafo non orientato
type Grafo struct {
	innerGrafo map[string]map[string]int
}

type Arco struct {
	x    string
	y    string
	peso int
	next *Arco
}

type Punto struct {
	x int
	y int
}

func main() {
	/*
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
	*/
	g := parseInput()
	fmt.Println("numero vertici: ", len(g.innerGrafo))
	_, pesoTot := g.getAlberoRicoprente()
	fmt.Println("peso totale: ", pesoTot)
}

func parseInput() *Grafo {
	g := newGrafo()
	scanner := bufio.NewScanner(os.Stdin)
	tempInput := make([]*Punto, 0, 5)
	for scanner.Scan() {
		splittedString := strings.Split(scanner.Text(), " ")
		x, _ := strconv.Atoi(splittedString[0])
		y, _ := strconv.Atoi(splittedString[1])
		tempInput = append(tempInput, &Punto{x, y})
	}

	for i := 0; i < len(tempInput); i++ {
		for j := 0; j < len(tempInput); j++ {
			if i != j {
				fromName := fmt.Sprintf("%d,%d", tempInput[i].x, tempInput[i].y)
				toName := fmt.Sprintf("%d,%d", tempInput[j].x, tempInput[j].y)
				g.aggiungi(fromName, toName, int(math.Abs(float64(tempInput[i].x-tempInput[j].x))+math.Abs(float64(tempInput[i].y-tempInput[j].y))))
			}
		}
	}
	return g
}

func printList(archiOrdinatiHead *Arco) {
	for archiOrdinatiHead != nil {
		fmt.Println(archiOrdinatiHead)
		archiOrdinatiHead = archiOrdinatiHead.next
	}
}

func countListSize(archiOrdinatiHead *Arco) int {
	i := 0
	for archiOrdinatiHead != nil {
		archiOrdinatiHead = archiOrdinatiHead.next
		i++
	}
	return i
}

func (g *Grafo) getAlberoRicoprente() (*Grafo, int) {
	albero := newGrafo()
	trovati := make(map[string]bool)
	archiOrdinatiHead := g.creaListaDiArchi()
	var pesoTotale int = 0

	trovati[archiOrdinatiHead.x] = true

	fmt.Println("archiOrdinati size: ", countListSize(archiOrdinatiHead))
	var azione bool = true
	//porcheria ma almeno è fatta con una linkedlist
	var count int = 0
	for azione {
		azione = false
		curArco := archiOrdinatiHead
		var prevArco *Arco = nil
		i := 0
		for curArco != nil {
			count++
			//fmt.Println("archiOrdinati size: ", countListSize(archiOrdinatiHead))
			//fmt.Println("trovati: ", trovati)
			//fmt.Println("albero: ", albero.innerGrafo)
			if trovati[curArco.x] && trovati[curArco.y] {
				azione = true
				if prevArco == nil {
					archiOrdinatiHead = archiOrdinatiHead.next
					break
				} else {
					prevArco.next = curArco.next
				}
			} else if (trovati[curArco.x] || trovati[curArco.y]) || !(trovati[curArco.x] || trovati[curArco.y]) {
				azione = true
				trovati[curArco.x] = true
				trovati[curArco.y] = true
				albero.aggiungi(curArco.x, curArco.y, curArco.peso)
				pesoTotale += curArco.peso
				if prevArco == nil {
					archiOrdinatiHead = archiOrdinatiHead.next
				} else {
					prevArco.next = curArco.next
				}
				break
			}
			if prevArco == nil {
				prevArco = curArco
			} else {
				prevArco = prevArco.next
			}
			curArco = curArco.next
			i++
		}
	}
	fmt.Println("count: ", count)
	return albero, pesoTotale
}

func (g *Grafo) creaListaDiArchi() *Arco {
	outList := make([]*Arco, 0, 2*len(g.innerGrafo))
	for from, v := range g.innerGrafo {
		for to, peso := range v {
			outList = append(outList, &Arco{from, to, peso, nil})
		}
	}
	sort.Slice(outList, func(i, j int) bool {
		return outList[i].peso < outList[j].peso
	})
	curArco := outList[0]
	for i := 1; i < len(outList); i++ {
		curArco.next = outList[i]
		curArco = curArco.next
	}
	return outList[0]
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
