package day15

import (
	"testing"

	"github.com/tom5760/advent-of-code/2022/testutils"
)

func TestDay15(t *testing.T) {
	t.Parallel()

	testutils.Run(t, Parse, Part1, Part2, []testutils.Test[int, int]{
		{
			Name:  "example",
			Part1: 26,
			Part2: 0,
		},
		{
			Name:  "input",
			Part1: 0,
			Part2: 0,
		},
	})
}
