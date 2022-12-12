package day03

import (
	"testing"

	"github.com/tom5760/advent-of-code/2022/testutils"
)

func TestDay03(t *testing.T) {
	t.Parallel()

	testutils.Run(t, Parse, Part1, Part2, []testutils.Test[int, int]{
		{
			Name:  "example",
			Part1: 157,
			Part2: 70,
		},
		{
			Name:  "input",
			Part1: 7581,
			Part2: 2525,
		},
	})
}
