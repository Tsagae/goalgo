package sorting

import "cmp"

func QuickSort[T cmp.Ordered](arr []T) {
	// base case <= 1 elements: trivial solution
	if len(arr) <= 1 {
		return
	} else {
		//else
		// 1) partition the array in 2
		var pivot int = partition(arr, 0, len(arr)-1)
		// 2) recursively sort the two parts indipendently
		QuickSort(arr[0:pivot])
		QuickSort(arr[pivot:])
	}
}

func partition[T cmp.Ordered](arr []T, i int, f int) int { //in loco
	//define a pivot (first element)
	var pivot int = i
	//until the indexes don't cross each other
	for i != f {
		for ; arr[i] <= arr[pivot] && i != f; i++ {
			//i search the first element at the left of the pivot that is greater than the pivot
		}
		for ; arr[f] > arr[pivot] && i != f; f-- {
			//i search the first element at the right of the pivot that is smaller or equal than the pivot
		}
		//swap of the two elements
		arr[f], arr[i] = arr[i], arr[f]
	}
	//when the two indexes cross i swap the element of index i-1 with the pivot
	arr[i-1], arr[pivot] = arr[pivot], arr[i-1]
	//return the pivot
	return i
}
