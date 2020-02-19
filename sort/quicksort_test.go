package sort

import (
	"math/rand"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func systemSort(source []int) []int {
	result := make([]int, len(source))
	copy(result[:], source)
	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})
	return result
}

func unique(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func genereateArray() []int {
	N := rand.Intn(100)
	arr := make([]int, N)
	for j := range arr {
		arr[j] = rand.Intn(100)
	}
	return unique(arr)
}

func TestQuickSort(t *testing.T) {
	a := assert.New(t)

	for i := 0; i < 100; i++ {
		arr := genereateArray()
		QuickSort(arr)
		a.Equal(systemSort(arr), arr)
	}
}
