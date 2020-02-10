package sort

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func partition(arr []int) int {
	if len(arr) < 2 {
		return 0
	}

	highIdx := 0
	pivotIdx := len(arr) - 1

	for i := 0; i < pivotIdx; i++ {
		if arr[i] < arr[pivotIdx] {
			swap(arr, highIdx, i)
			highIdx++
		}
	}
	swap(arr, pivotIdx, highIdx)
	return highIdx
}

// QuickSort - quick sort implementation
func QuickSort(arr []int) {
	pivotIdx := partition(arr)
	if pivotIdx > 1 {
		QuickSort(arr[:pivotIdx])
	}
	if len(arr) > 2 {
		QuickSort(arr[pivotIdx+1:])
	}
}
