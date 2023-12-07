package sorting

import (
	"cmp"
)

func QuickSort[T cmp.Ordered](arr []T) {
	// base case <= 1 elements: trivial solution
	if len(arr) <= 1 {
		return
	} else {
		//else
		// 1) partition the array in 2
		var pivot int = partition(arr, 0, len(arr)-1)
		// 2) recursively sort the two parts indipendently

		QuickSort(arr[:pivot])
		QuickSort(arr[pivot+1:])
	}
}

func partition[T cmp.Ordered](arr []T, i int, j int) int { //in loco
	//define a pivot (first element)
	var pivot int = i
	//until the indexes don't cross each other
	for i <= j {
		for ; arr[j] > arr[pivot]; j-- {
			//search the first element that is smaller or equal than the pivot from the right
		}
		for ; i <= j && arr[i] <= arr[pivot]; i++ {
			//search the first element that is greater than the pivot from the left
		}
		if i < j {
			//swap of the two elements
			arr[j], arr[i] = arr[i], arr[j]
			i++
			j--
		}
	}
	//when the two indexes cross i swap the element of index i-1 with the pivot
	arr[j], arr[pivot] = arr[pivot], arr[j]
	//return the index of the pivot
	return j
}
