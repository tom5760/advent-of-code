package util

import (
	"io"
	"log"
	"os"
	"testing"
)

type (
	DayFunc func(*log.Logger, io.Reader) (int, int, error)

	DayTest struct {
		Name  string
		Part1 int
		Part2 int
	}
)

func RunTests(t *testing.T, fn DayFunc, tests []DayTest) {
	t.Helper()
	t.Parallel()

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Helper()
			t.Parallel()

			f, err := os.Open(test.Name)
			if err != nil {
				t.Errorf("failed to open test file %q: %v", test.Name, err)
				return
			}

			defer f.Close()

			lg := NewTestLogger(t)

			part1, part2, err := fn(lg, f)
			if err != nil {
				t.Error(err)
				return
			}
			if test.Part1 != part1 {
				t.Errorf("part 1; want: %v; got: %v", test.Part1, part1)
			}
			if test.Part2 != part2 {
				t.Errorf("part 2; want: %v; got: %v", test.Part2, part2)
			}
		})
	}
}
