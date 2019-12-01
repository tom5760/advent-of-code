package common

// MapSumUint64 applies the predicate to each value in data, returning the sum
// of each mapped value.
func MapSumUint64(data []uint64, predicate func(uint64) uint64) uint64 {
	var total uint64

	for _, value := range data {
		total += predicate(value)
	}

	return total
}
