package main

import (
	"strings"
	"testing"

	"github.com/tom5760/advent-of-code/2021/grid"
)

const exampleInput = `5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526
`

func TestPart1(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		expected uint64
		input    string
	}{
		{
			name:     "example",
			expected: 1656,
			input:    exampleInput,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			grid, err := grid.Parse(strings.NewReader(test.input))
			if err != nil {
				t.Fatalf("unexpected input parse error: %v", err)
			}

			actual := Part1(grid)

			if actual != test.expected {
				t.Errorf("got %v; want %v", actual, test.expected)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		expected uint64
		input    string
	}{
		{
			name:     "example",
			expected: 195,
			input:    exampleInput,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			grid, err := grid.Parse(strings.NewReader(test.input))
			if err != nil {
				t.Fatalf("unexpected input parse error: %v", err)
			}

			actual := Part2(grid)

			if actual != test.expected {
				t.Errorf("got %v; want %v", actual, test.expected)
			}
		})
	}
}
