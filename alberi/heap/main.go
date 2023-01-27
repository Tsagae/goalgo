package main

import "fmt"

func main() {
	arr := []int{7, 9, 12, 2, 15, 25, 3, 14, 10}
	//arr2 := []int{7, 9, 12, 6, 5, 3, 1}
	creaHeap(arr)
	fmt.Println(arr)
	inserisci(&arr, 16)
	fmt.Println(arr)
	arr = rimuovi(arr, 16)
	fmt.Println(arr)
}

func inserisci(arr *[]int, chiave int) {
	//aggiunge l'emento come ultima foglia e riordina dal basso
	*arr = append(*arr, chiave)
	riordinaDalBasso(*arr, len(*arr)-1)
}

func rimuovi(arr []int, chiave int) []int {
	//sostituisce l'elemento con l'ultima foglia, scarta l'ultimo e risistema
	index := indiceDaChiave(arr, chiave)
	arr[index], arr[len(arr)-1] = arr[len(arr)-1], arr[index]
	fmt.Println("scambio: ", arr)
	arr = arr[0 : len(arr)-1]
	if chiave > arr[index] {
		riordinaRadiceIterativo(arr, index)
	} else if chiave < arr[index] {
		riordinaDalBasso(arr, index)
	}
	return arr
}

func riordinaRadiceRicorsivo(heap []int, index int) {
	//guardo se la radice è più piccola del maggiore dei due figli
	sx, dx := getIndexFigli(heap, index)
	var maxIndex int
	if sx == -1 && dx == -1 {
		return
	} else if sx == -1 {
		maxIndex = dx
	} else if dx == -1 {
		maxIndex = sx
	} else {
		if heap[sx] > heap[dx] {
			maxIndex = sx
		} else {
			maxIndex = dx
		}
	}
	//se è vero scambio la radice con il figlio maggiore
	if heap[index] < heap[maxIndex] {
		heap[index], heap[maxIndex] = heap[maxIndex], heap[index]
		riordinaRadiceRicorsivo(heap, maxIndex)
	}
}

func riordinaRadiceIterativo(arr []int, index int) {
	for {
		//guardo se la radice è più piccola del maggiore dei due figli
		sx, dx := getIndexFigli(arr, index)
		var maxIndex int
		if sx == -1 && dx == -1 {
			return
		} else if sx == -1 {
			maxIndex = dx
		} else if dx == -1 {
			maxIndex = sx
		} else {
			if arr[sx] > arr[dx] {
				maxIndex = sx
			} else {
				maxIndex = dx
			}
		}
		//se è vero scambio la radice con il figlio maggiore
		if arr[index] < arr[maxIndex] {
			arr[index], arr[maxIndex] = arr[maxIndex], arr[index]
			index = maxIndex
		} else {
			return
		}
	}
}

func getIndexFigli(heap []int, index int) (int, int) {
	sx := index*2 + 1
	dx := index*2 + 2
	if sx >= len(heap) {
		sx = -1
	}
	if dx >= len(heap) {
		dx = -1
	}
	return sx, dx
}

// trasforma un array in un heap
func creaHeap(arr []int) {
	for i := len(arr) / 2; i >= 0; i-- {
		riordinaRadiceIterativo(arr, i)
	}
}

func heapSort(arr []int) {
	creaHeap(arr)
	for i := len(arr) - 1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		riordinaRadiceIterativo(arr[0:i], 0)
	}
}

func getIndexPadre(index int) int {
	if index%2 == 0 {
		index /= 2
		index--
	} else {
		index /= 2
	}
	return index
}

func riordinaDalBasso(arr []int, index int) {
	for {
		var lowerIndex int = index
		lowerIndex = getIndexPadre(index)
		if lowerIndex < 0 {
			return
		}
		if arr[lowerIndex] < arr[index] {
			arr[lowerIndex], arr[index] = arr[index], arr[lowerIndex]
			index = lowerIndex
		} else {
			return
		}
	}
}

func indiceDaChiave(arr []int, chiave int) int {
	var i, v int
	for i, v = range arr {
		if v == chiave {
			break
		}
	}
	return i
}
