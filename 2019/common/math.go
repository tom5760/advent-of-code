package common

import (
	"math"
)

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

// IntPow10 returns 10**n, the base-10 exponential of n.
func IntPow10(n int) int {
	return int(math.Pow10(n))
}

// PermutationsInt calls f with every permutation of the input slice x.
//
// Implements non-recursive Heap's Algorithm:
//   https://en.wikipedia.org/wiki/Heap%27s_algorithm
func PermutationsInt(x []int, f func([]int)) {
	c := make([]int, len(x))

	f(x)

	for i := 0; i < len(x); {
		if c[i] < i {
			if i%2 == 0 {
				x[0], x[i] = x[i], x[0]
			} else {
				x[c[i]], x[i] = x[i], x[c[i]]
			}

			f(x)
			c[i]++
			i = 0
		} else {
			c[i] = 0
			i++
		}
	}
}

// From https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/

// GCD greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// LCM find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
