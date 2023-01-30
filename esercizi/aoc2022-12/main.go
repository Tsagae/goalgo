package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Punto struct {
	x int
	y int
}

func main() {
	g, s := getInput()
	printGrid(g)
	points := trovaStart(grid)
	var min int = math.MaxInt
	for _, v := range points {
		curVal = bfsCostoCammino(g, v)
		if(curVal < min){
			min = curVal
		}
	}
}

func trovaStart(grid [][]int) []*Punto {
	punti := make([]*Punto, 0)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				punti = append(punti, &{i, j})
			}
		}
	}
}

func bfsCostoCammino(grid [][]int, start *Punto) int {
	distanze := make([][]int, len(grid))
	for i, _ := range distanze {
		distanze[i] = make([]int, len(grid[i]))
	}

	visitati := make([][]bool, len(grid))
	for i, _ := range visitati {
		visitati[i] = make([]bool, len(grid[i]))
	}

	queue := make([]*Punto, 1)
	queue[0] = start

	for len(queue) != 0 {
		//printQueue(queue)
		curPunto := queue[0]
		x := curPunto.x
		y := curPunto.y

		if grid[x][y] == 27 {
			fmt.Println("trovato")
			return distanze[x][y]
		}

		queue = queue[1:]
		if visitati[x][y] {
			continue
		}
		visitati[x][y] = true
		//sinistra
		updateDistanzaAndCoda(x, y, x-1, y, grid, distanze, visitati, &queue)
		// destra
		updateDistanzaAndCoda(x, y, x+1, y, grid, distanze, visitati, &queue)
		//sopra
		updateDistanzaAndCoda(x, y, x, y-1, grid, distanze, visitati, &queue)
		// sotto
		updateDistanzaAndCoda(x, y, x, y+1, grid, distanze, visitati, &queue)
	}
	return -1
}

func updateDistanzaAndCoda(x int, y int, x2 int, y2 int, grid [][]int, distanze [][]int, visitati [][]bool, queue *[]*Punto) {
	if checkBounds(x2, y2, grid) {
		if grid[x2][y2] <= grid[x][y]+1 {
			curDist := distanze[x][y]
			distanze[x2][y2] = curDist + 1
			if !visitati[x2][y2] {
				*queue = append(*queue, &Punto{x2, y2})
			}
		}
	}
}

func checkBounds(x int, y int, grid [][]int) bool {
	if x < 0 || x >= len(grid) {
		return false
	}
	if y < 0 || y >= len(grid[x]) {
		return false
	}
	return true
}

func printQueue(queue []*Punto) {
	for _, v := range queue {
		fmt.Println(*v, " ")
	}
}

func printGrid(grid [][]int) {
	size := len(grid)
	for i := 0; i < size; i++ {
		for j := 0; j < len(grid[i]); j++ {
			fmt.Print(grid[i][j], " ")
		}
		fmt.Print("\n")
	}
}

func getInput() ([][]int, *Punto) {
	/*

		Sabqponm
		abcryxxl
		accszExk
		acctuvwj
		abdefghi
	*/
	var start *Punto = new(Punto)
	grid := make([][]int, 0)
	var lineSize int = 0
	var lineNumber int = 0

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		lineSize = len(line)
		grid = append(grid, make([]int, lineSize))
		for i, v := range line {
			if v >= 'a' && v <= 'z' {
				grid[lineNumber][i] = int(v) - 96
			} else if v == 'S' {
				grid[lineNumber][i] = 1
				start.x = lineNumber
				start.y = i
			} else if v == 'E' {
				grid[lineNumber][i] = 27
			}
		}
		lineNumber++
	}
	return grid, start
}
