package main

import (
	"math"
	"strconv"
)

func DigitCountInt(x int) int {
	return len(strconv.Itoa(x))
}

func MaxInt(ints ...int) int {
	var max int64 = math.MinInt64
	for _, x := range ints {
		if int64(x) > max {
			max = int64(x)
		}
	}
	return int(max)
}
