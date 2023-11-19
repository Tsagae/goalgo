package sorting

import (
	"testing"
)

func TestHeapsort(t *testing.T) {
	testArr := []int{9, 3, 8, 0, 1, 6, 7, 2, 5, 4}
	HeapSort(testArr)
	for i, v := range testArr {
		if !(v == i) {
			t.Errorf("Incorrect sorting")
		}
	}
}
