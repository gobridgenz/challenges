package main

import (
	"fmt"
	"math"
)

type elt struct {
	value int32
	count int32
}

func findNumberOfMax(ar []int32) int32 {
	maxElt := elt{value: math.MinInt32, count: 0}
	for _, c := range ar {
		if maxElt.value == c {
			maxElt.count++
			// reset counter if we find another max
		} else if c > maxElt.value {
			maxElt.count = 1
			maxElt.value = c
		}
	}
	return maxElt.count
}

func main() {
	arr := []int32{18, 90, 90, 13, 90, 75, 90, 8, 90, 43}
	fmt.Println(findNumberOfMax(arr))
}
