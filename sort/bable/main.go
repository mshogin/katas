package main

import "fmt"
import "math"

// N=44
func SortV3(list []float32) []float32 {
	N := 0
	for i:=0; i<len(list)-1; i++ {
		N++
		for j:=0; j<len(list) - i-1; j++ {
			N++
			if list[j]>list[j+1] {
				list[j], list[j+1] = list[j+1],list[j]
			}
		}
	}
	fmt.Println(N)
	return list
}


// N=len(list)^2
func SortV2(list []float32) []float32{
	N := 0
	for _ = range(list) { // main loop
		N++
		for j := 0; j < len(list)-1; j++ {
			N++
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
			}

		}
	}
	fmt.Println(N)
	return list  
}

func SortV1(list []float64) []float64{
	var min float64 
	result := []float64{}
	for len(list) > 0 {
		if min, list = extractMin(list); list != nil {
			result = append(result, min)
		}
	}
	return result
}

func extractMin(list []float64) (float64, []float64) {
	if len(list) < 1 {
		return 0, nil
	}

	newlist := []float64{}

	min := list[0]
	for _, v := range list[1:] {
		newlist = append(newlist, math.Max( v, min))
		min = math.Min(v, min)
	}
	return min, newlist
}
