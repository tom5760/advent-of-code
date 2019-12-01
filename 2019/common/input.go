package common

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

// ReadUint64Slice reads newline separated input as a slice of uint64s.
func ReadUint64Slice(r io.Reader) ([]uint64, error) {
	var inputs []uint64

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		value, err := strconv.ParseUint(scanner.Text(), 10, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to convert value to int: %w", err)
		}

		inputs = append(inputs, value)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to scan input: %w", err)
	}

	return inputs, nil
}
