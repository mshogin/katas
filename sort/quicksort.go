package sort

func partition(arr []int) int {
	pivotIdx := len(arr) - 1
	highIdx := 0
	for i := 0; i < pivotIdx; i++ {
		if arr[i] < arr[pivotIdx] {
			arr[i], arr[highIdx] = arr[highIdx], arr[i]
			highIdx++
		}
	}
	arr[highIdx], arr[pivotIdx] = arr[pivotIdx], arr[highIdx]
	return highIdx
}

// QuickSort ...
func QuickSort(arr []int) {
	pi := partition(arr)
	if pi > 0 && len(arr[:pi]) > 1 {
		QuickSort(arr[:pi])
	}
	if len(arr[pi+1:]) > 1 {
		QuickSort(arr[pi+1:])
	}
}
