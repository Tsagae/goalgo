package floydwarshall

import (
	"fmt"
	"math"
)

type Grafo struct {
	innerGrafo [][]int
}

type Arco struct {
	from int
	to   int
	peso int
}

/*
func main() {
	fmt.Println("prova")
	g := newGrafo(4, []Arco{Arco{1, 0, 4}, Arco{0, 2, -2}, Arco{1, 2, 3}, Arco{2, 3, 2}, Arco{3, 1, -1}})
	stampaMatriceQuadrata(g.innerGrafo)
	cammini := g.floydWarshall()
	stampaMatriceQuadrata(cammini)

}
*/

func (g *Grafo) floydWarshall() [][]int {
	mat := g.inizializzaMatrice()
	n := len(g.innerGrafo)
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if !(mat[i][k] == math.MaxInt || mat[k][j] == math.MaxInt) {
					if mat[i][j] > mat[i][k]+mat[k][j] {
						mat[i][j] = mat[i][k] + mat[k][j]
					}
				}
			}
		}
	}
	return mat
}

func (g *Grafo) inizializzaMatrice() [][]int {
	var n int = len(g.innerGrafo)
	mat := make([][]int, n)
	for i := 0; i < n; i++ {
		mat[i] = make([]int, n)
		for j := 0; j < n; j++ {
			mat[i][j] = g.innerGrafo[i][j]
		}
	}
	return mat
}

func newGrafo(n int, archi []Arco) *Grafo {
	innerGrafo := make([][]int, n)
	for i := 0; i < n; i++ {
		innerGrafo[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == j {
				innerGrafo[i][j] = 0
			} else {
				innerGrafo[i][j] = math.MaxInt
			}
		}
	}
	for len(archi) != 0 {
		v := archi[0]
		innerGrafo[v.from][v.to] = v.peso
		archi = archi[1:]
	}
	fmt.Println(archi)
	return &Grafo{innerGrafo}
}
func stampaMatriceQuadrata(mat [][]int) {
	n := len(mat)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%d ", mat[i][j])
		}
		fmt.Print("\n")
	}
}
