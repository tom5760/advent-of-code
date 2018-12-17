package main

import (
	"math"
	"strconv"
)

func DigitCountInt64(x int64) int {
	return len(strconv.FormatInt(x, 10))
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
