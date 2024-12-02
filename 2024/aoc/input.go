package aoc

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"iter"
	"strconv"
)

func IterRowsInt(r io.Reader) iter.Seq2[[]int, error] {
	return func(yield func([]int, error) bool) {
		scanner := bufio.NewScanner(r)
		for scanner.Scan() {
			fields := bytes.Fields(scanner.Bytes())
			ints := make([]int, len(fields))

			for i, field := range fields {
				v, err := strconv.Atoi(string(field))
				if err != nil {
					yield(nil, fmt.Errorf("failed to parse field: %w", err))
					return
				}

				ints[i] = v
			}

			if !yield(ints, nil) {
				return
			}
		}
		if err := scanner.Err(); err != nil {
			yield(nil, fmt.Errorf("failed to scan: %w", err))
			return
		}
	}
}
