package main

import "fmt"

type Vertice struct {
	key string
	val int
}

type Queue struct {
	innerQueue []*Vertice
}

func main() {
	fmt.Println("prova")
	s := newQueue()
	s.add(&Vertice{key: "a"})
	s.add(&Vertice{key: "b"})
	s.add(&Vertice{key: "c"})
	fmt.Println(s.toString())
	fmt.Println(s.pop())
	fmt.Println(s.toString())

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
