package aoc2023_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/tom5760/advent-of-code/aoc2023"
)

func TestDay1(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		part1 int
		part2 int
	}{
		{name: "example1", part1: 142, part2: 142},
		{name: "example2", part1: 209, part2: 281},
		{name: "input", part1: 53386, part2: 53312},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			inputPath := filepath.Join("day01", test.name)
			f, err := os.Open(inputPath)
			if err != nil {
				t.Errorf("failed to open test file %q: %v", inputPath, err)
				return
			}

			defer f.Close()

			part1, part2, err := aoc2023.Day01(f)
			if err != nil {
				t.Error(err)
				return
			}
			if test.part1 != part1 {
				t.Errorf("part 1; want: %v; got: %v", test.part1, part1)
			}
			if test.part2 != part2 {
				t.Errorf("part 2; want: %v; got: %v", test.part2, part2)
			}
		})
	}
}
