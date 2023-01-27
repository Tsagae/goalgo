package main

import "fmt"

func main() {
	var a []int = []int{7, 10, 4, 2, 5, 11}
	a = quickSort(a)
	fmt.Println(a)
}

func quickSort(arr []int) []int {
	// caso base <= 1 elementi soluzione immediata
	if len(arr) <= 1 {
		return arr
	} else {
		//else
		// 1) dividi l'array in due parti (partizionamento)
		var perno int = partiziona(arr, 0, len(arr)-1)
		// 2) ordina separatamente con chiamata ricorsiva
		quickSort(arr[0:perno])
		quickSort(arr[perno:])
		// 3) combina le due soluzioni
	}
	return arr
}

func partiziona(arr []int, i int, f int) int { //in loco
	//prendo un perno (primo elemento)
	var perno int = i
	fmt.Printf("perno: %d ", arr[perno])
	fmt.Println("arr iniziale: ", arr)
	//finché gli indici non si incrociano
	for i != f {
		for ; arr[i] <= arr[perno] && i != f; i++ {
			//cerco il primo elemento a sinistra del perno maggiore del perno
		}
		fmt.Printf("i: %d ", i)
		for ; arr[f] > arr[perno] && i != f; f-- {
			//cerco il primo elemento a destra del perno minore o uguale del perno
		}
		fmt.Printf("f: %d ", f)
		fmt.Print("arr pre scambio: ", arr)
		//scambio i due elementi
		arr[f], arr[i] = arr[i], arr[f]
		fmt.Println(" arr post scambio: ", arr)
	}
	//quando i due indici si incontrano scambio l'elemento con il perno
	arr[i-1], arr[perno] = arr[perno], arr[i-1]
	//restiuisco l'indice del perno
	fmt.Printf("perno: %d ", arr[i])
	fmt.Println("arr finale: ", arr)
	fmt.Println("-----------------")
	return i
}
