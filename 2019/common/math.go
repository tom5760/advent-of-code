package common

// IntMin returns the smaller of x or y.
//
// math.Min works on float64s, not ints.
func IntMin(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// IntMax returns the larger of x or y.
//
// math.Min works on float64s, not ints.
func IntMax(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// IntAbs returns the absolute value of x.
//
// math.Min works on float64s, not ints.
func IntAbs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
