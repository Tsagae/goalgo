package sorting

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testArr []int = []int{9, 3, 3, -1, 9, -2, 8, 0, -2, 2, 5, 4}

func sortTest(t *testing.T, sort func([]int)) {
	arr := make([]int, len(testArr))
	copy(arr, testArr)
	sort(arr)
	assert.IsNonDecreasing(t, arr)
}

func TestSelectionSort(t *testing.T) {
	sortTest(t, SelectionSort)
}

func TestInsertionSort(t *testing.T) {
	sortTest(t, InsertionSort)
}

func TestBubbleSort(t *testing.T) {
	sortTest(t, BubbleSort)
}

func TestHeapSort(t *testing.T) {
	sortTest(t, HeapSort)
}

func TestQuickSort(t *testing.T) {
	sortTest(t, QuickSort)
}

func TestMergeSort(t *testing.T) {
	arr := make([]int, len(testArr))
	copy(arr, testArr)
	arr = mergeSort(arr)
	assert.IsNonDecreasing(t, arr)
}
