package day04

import (
	"testing"

	"github.com/tom5760/advent-of-code/2022/testutils"
)

func TestDay04(t *testing.T) {
	t.Parallel()

	testutils.Run(t, Parse, Part1, Part2, []testutils.Test[int, int]{
		{
			Name:  "example",
			Part1: 2,
			Part2: 4,
		},
		{
			Name:  "input",
			Part1: 483,
			Part2: 874,
		},
	})
}
