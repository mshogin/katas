package main

import "testing"
import "github.com/stretchr/testify/assert"

func TestSortV3(t *testing.T) {
	a := assert.New(t)

	a.Equal([]float32{1,2,3,4,5,6,7,8,9},
		SortV3([]float32{9,8,4,3,2,5,6,1,7}),
	)
}


func TestSortV2(t *testing.T) {
	a := assert.New(t)

	a.Equal([]float32{1, 2, 3},
		SortV2([]float32{3, 2, 1}),
	)

	a.Equal([]float32{1, 2, 3, 4, 5, 6, 7, 8},
		SortV2([]float32{8, 7, 4, 5, 6, 3, 2, 1}),
	)
}

func TestSortV1(t *testing.T) {
	a := assert.New(t)

	expectedList := []float64{1,2,3}

	sortedList := SortV1([]float64{1,3,2})

	a.Equal( expectedList, sortedList)
}
