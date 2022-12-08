package sliceutils

import "golang.org/x/exp/constraints"

func Count[T any](s []T, fn func(T) bool) int {
	var total int

	for _, v := range s {
		if fn(v) {
			total++
		}
	}

	return total
}

func Max[T constraints.Ordered](s []T) T {
	var max T

	for _, v := range s {
		if v > max {
			max = v
		}
	}

	return max
}
