package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestQuickSort1 ...
func TestQuickSort1(t *testing.T) {
	a := assert.New(t)

	arr := []int{3, 2, 1}
	QuickSort1(arr)

	a.Equal([]int{1, 2, 3}, arr)
}
