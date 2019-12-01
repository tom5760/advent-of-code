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

type Point struct {
	X, Y int
}

func (p Point) Up() Point {
	return Point{p.X, p.Y - 1}
}

func (p Point) Down() Point {
	return Point{p.X, p.Y + 1}
}

func (p Point) Left() Point {
	return Point{p.X - 1, p.Y}
}

func (p Point) Right() Point {
	return Point{p.X + 1, p.Y}
}
