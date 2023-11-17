package merge

/*
	func main() {
		var a []int = []int{7, 10, 4, 2, 5, 11}
		a = mergeSort(a)
		fmt.Println(a)
	}
*/

func mergeSort(arr []int) []int {
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

func pop(arr *[]int) int {
	var toRet int = (*arr)[0]
	*arr = (*arr)[1:]
	return toRet
}

func merge(arr1 []int, arr2 []int) []int {
	sumDim := len(arr1) + len(arr2)
	res := make([]int, sumDim)
	i := 0
	for ; len(arr1) != 0 && len(arr2) != 0; i++ {
		if arr1[0] < arr2[0] {
			res[i] = pop(&arr1)
		} else {
			res[i] = pop(&arr2)
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
