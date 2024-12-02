package aoc

import (
	"io"
	"os"
	"testing"
)

type (
	Test struct {
		Name  string
		Part1 int
		Part2 int
	}

	DayFunc func(r io.Reader) (int, int, error)
)

func Run(t *testing.T, runFn DayFunc, tests []Test) {
	t.Helper()
	t.Parallel()

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()

			f, err := os.Open(test.Name)
			if err != nil {
				t.Fatalf("failed to open test input file: %v", err)
				return
			}

			defer f.Close()

			part1, part2, err := runFn(f)
			if err != nil {
				t.Fatalf("failed to parse test input: %v", err)
			}

			t.Logf("part 1: %v", part1)
			t.Logf("part 2: %v", part2)

			if part1 != test.Part1 {
				t.Errorf("failed part 1; got %v; want %v", part1, test.Part1)
			}

			if part2 != test.Part2 {
				t.Errorf("failed part 2; got %v; want %v", part2, test.Part2)
			}
		})
	}
}
