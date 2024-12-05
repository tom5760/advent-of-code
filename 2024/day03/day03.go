package day03

import (
	"fmt"
	"io"
)

func Run(r io.Reader) (int, int, error) {
	parser := Parse(r)
	do := true

	var part1, part2 int

	for op := range parser.Ops() {
		switch op := op.(type) {
		case OpMul:
			v := op.A * op.B

			part1 += v
			if do {
				part2 += v
			}

		case OpDo:
			do = true
		case OpDont:
			do = false

		default:
			return 0, 0, fmt.Errorf("unexpected op type: %T", op)
		}
	}

	if err := parser.Err(); err != nil {
		return 0, 0, fmt.Errorf("failed to run parser: %w", err)
	}

	return part1, part2, nil
}
