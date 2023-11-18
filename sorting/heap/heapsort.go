package heap

func HeapSort(arr []int) {
	createHeap(arr)
	for i := len(arr) - 1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		reorderRootRecursive(arr[0:i], 0)
	}
}

func reorderRootRecursive(heap []int, index int) {
	//if thr root is smaller than the biggest of the two children
	l, r := getIndexChildren(heap, index)
	var maxIndex int
	if l == -1 && r == -1 {
		return
	} else if l == -1 {
		maxIndex = r
	} else if r == -1 {
		maxIndex = l
	} else {
		if heap[l] > heap[r] {
			maxIndex = l
		} else {
			maxIndex = r
		}
	}
	//if true swap the root with the biggest child
	if heap[index] < heap[maxIndex] {
		heap[index], heap[maxIndex] = heap[maxIndex], heap[index]
		reorderRootRecursive(heap, maxIndex)
	}
}

func createHeap(arr []int) {
	for i := len(arr) / 2; i >= 0; i-- {
		reorderRootRecursive(arr, i)
	}
}

func getIndexChildren(arr []int, index int) (int, int) {
	l := index*2 + 1
	r := index*2 + 2
	if l >= len(arr) {
		l = -1
	}
	if r >= len(arr) {
		r = -1
	}
	return l, r
}
