package sort

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// systemSort ...
func systemSort(source []int) []int {
	result := make([]int, len(source))
	copy(result[:], source)
	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})
	return result
}

// TestQuickSort ...
func TestQuickSort(t *testing.T) {
	a := assert.New(t)

	arr := []int{6, 7, 3, 2, 1}

	QuickSort(arr)

	a.Equal(systemSort(arr), arr)
}
