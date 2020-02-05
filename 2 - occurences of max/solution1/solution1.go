package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int64{18, 90, 90, 13, 90, 75, 90, 8, 90, 43}
	valueMap := make(map[int64]int)

	var max int64 = math.MinInt64
	for _, n := range arr {
		if n > max {
			max = n
		}
		valueMap[n] = valueMap[n] + 1
	}

	fmt.Println(valueMap[max])
}
