package sorting

import "cmp"

func SelectionSort[T cmp.Ordered](arr []T) {
	//Take the smallest and swap with[i]
	var minI int
	for i := 0; i < len(arr); i++ {
		minI = i
		for j := i; j < len(arr); j++ {
			if arr[j] < arr[minI] {
				minI = j
			}
		}
		arr[minI], arr[i] = arr[i], arr[minI]
	}
}

func InsertionSort[T cmp.Ordered](arr []T) {
	//From arr[0] to arr[i-1] is the sorted partition
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			} else {
				break
			}
		}
	}
}

func BubbleSort[T cmp.Ordered](arr []T) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
