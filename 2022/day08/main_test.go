package day08

import (
	"testing"

	"github.com/tom5760/advent-of-code/2022/testutils"
)

func TestDay08(t *testing.T) {
	t.Parallel()

	testutils.Run(t, Parse, Part1, Part2, []testutils.Test[int, int]{
		{
			Name:  "example",
			Part1: 21,
			Part2: 8,
		},
		{
			Name:  "input",
			Part1: 1690,
			Part2: 535680,
		},
	})
}
