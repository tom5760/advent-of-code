package day12

import (
	"testing"

	"github.com/tom5760/advent-of-code/2022/testutils"
)

func TestDay12(t *testing.T) {
	t.Parallel()

	testutils.Run(t, Parse, Part1, Part2, []testutils.Test[int, int]{
		{
			Name:  "example",
			Part1: 31,
			Part2: 29,
		},
		{
			Name:  "input",
			Part1: 481,
			Part2: 480,
		},
	})
}
