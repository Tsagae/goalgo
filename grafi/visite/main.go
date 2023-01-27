package main

import "fmt"

type Vertice struct {
	key string
	val int
}

type Queue struct {
	innerQueue []*Vertice
}

type Grafo struct {
	innerGrafo map[*Vertice][]*Vertice
}

type 

func main() {
	fmt.Println("prova")
	innerGrafo := make(map[*Vertice][]*Vertice)
	a := &Vertice{key: "a"}
	b := &Vertice{key: "b"}
	c := &Vertice{key: "c"}
	d := &Vertice{key: "d"}
	e := &Vertice{key: "e"}
	f := &Vertice{key: "f"}
	innerGrafo[a] = []*Vertice{b, c}
	innerGrafo[b] = []*Vertice{d, f, c, a}
	innerGrafo[c] = []*Vertice{a, b, f}
	innerGrafo[d] = []*Vertice{b, e}
	innerGrafo[e] = []*Vertice{d, f}
	innerGrafo[f] = []*Vertice{c, b, e}
	g := new(Grafo)
	g.innerGrafo = innerGrafo
	//g.visitaInAmpiezza(e)
	visitati := make(map[*Vertice]bool)
	g.visitaInProfondita(d, visitati)
}

func 

func (g *Grafo) visitaInProfondita(ver *Vertice, visitati map[*Vertice]bool) {
	visitati[ver] = true
	fmt.Println(ver)
	for _, v := range g.innerGrafo[ver] {
		if !visitati[v] {
			g.visitaInProfondita(v, visitati)
		}
	}
}

func (g *Grafo) visitaInAmpiezza(ver *Vertice) {
	queue := newQueue()
	visitati := make(map[*Vertice]bool)
	queue.add(ver)
	visitati[ver] = true
	for queue.size() != 0 {
		currentVer := queue.pop()
		fmt.Println(currentVer)
		for _, v := range g.innerGrafo[currentVer] {
			if !visitati[v] {
				queue.add(v)
				visitati[v] = true
			}
		}
	}
}

func newQueue() *Queue {
	s := new(Queue)
	s.innerQueue = make([]*Vertice, 0)
	return s
}

func (s *Queue) add(v *Vertice) {
	s.innerQueue = append(s.innerQueue, v)
}

func (s *Queue) peek() *Vertice {
	return s.innerQueue[0]
}

func (s *Queue) size() int {
	return len(s.innerQueue)
}

func (s *Queue) pop() *Vertice {
	var v *Vertice = s.peek()
	s.innerQueue = s.innerQueue[1:]
	return v
}

func (s *Queue) toString() string {
	var outStr string
	for _, v := range s.innerQueue {
		outStr += fmt.Sprintf("k: %s | ", v.key)
	}
	return outStr
}
