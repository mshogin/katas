package sort

// partition1 ...
func partition1(arr []int) int {
	pivot_idx := len(arr) - 1

	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[pivot_idx] {

		}
	}
}

// QuickSort1 ...
func QuickSort1(arr []int) {
	pi := partition1(arr)
	// QuickSort1(arr[:pi])
	// QuickSort1(arr[pi+1:])
}
