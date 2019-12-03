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

// IntOrder orders x, y from smaller to larger.
func IntOrder(x, y int) (smaller, larger int) {
	if x < y {
		return x, y
	}
	return y, x
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
