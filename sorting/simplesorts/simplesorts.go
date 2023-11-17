package simplesorts

/*
func main() {
	var a []int = []int{11, 6, 10, 1, 4}
	fmt.Println(a)
	bubbleSort(a)
	fmt.Println(a)

}
*/

func selectionSort(arr []int) []int {
	//Prendo il minore e lo scambio con arr[i]
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
	return arr
}

func insertionSort(arr []int) []int {
	//da arr[0] a arr[i-1] è la parte ordinata
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			} else {
				break
			}
		}
	}
	return arr
}

func bubbleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}
