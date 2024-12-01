package aoc

import (
	"io"
	"os"
	"testing"
)

type Day interface {
	Part1() int
	Part2() int
}

type Test struct {
	Name  string
	Part1 int
	Part2 int
}

type InputParser[D Day] func(io.Reader) (D, error)

func Run[D Day](t *testing.T, parserFn InputParser[D], tests []Test) {
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

			input, err := parserFn(f)
			if err != nil {
				t.Fatalf("failed to parse test input: %v", err)
			}

			if v := input.Part1(); v != test.Part1 {
				t.Errorf("failed part 1; got %v; want %v", v, test.Part1)
			} else {
				t.Logf("Part 1: %v", v)
			}

			if v := input.Part2(); v != test.Part2 {
				t.Errorf("failed part 2; got %v; want %v", v, test.Part2)
			} else {
				t.Logf("Part 2: %v", v)
			}
		})
	}
}
