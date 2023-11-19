package sorting

import "cmp"

func mergeSort[T cmp.Ordered](arr []T) []T {
	//TODO: change to in place implementation
	n := len(arr)
	if n > 1 {
		m := n / 2
		arrB := arr[:m]
		arrC := arr[m:]
		arrB = mergeSort(arrB)
		arrC = mergeSort(arrC)
		arr = merge(arrB, arrC)
	}
	return arr
}

func pop[T cmp.Ordered](arr []T) (T, []T) {
	return arr[0], arr[1:]
}

func merge[T cmp.Ordered](arr1 []T, arr2 []T) []T {
	sumDim := len(arr1) + len(arr2)
	res := make([]T, sumDim)
	i := 0
	for ; len(arr1) != 0 && len(arr2) != 0; i++ {
		if arr1[0] < arr2[0] {
			res[i], arr1 = pop(arr1)
		} else {
			res[i], arr2 = pop(arr2)
		}
	}

	if len(arr1) == 0 {
		arr1, arr2 = arr2, arr1
	}
	for _, v := range arr1 {
		res[i] = v
		i++
	}
	return res
}
