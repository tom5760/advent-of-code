package day14

import (
	"testing"

	"github.com/tom5760/advent-of-code/2022/testutils"
)

func TestDay14(t *testing.T) {
	t.Parallel()

	testutils.Run(t, Parse, Part1, Part2, []testutils.Test[int, int]{
		{
			Name:  "example",
			Part1: 24,
			Part2: 93,
		},
		{
			Name:  "input",
			Part1: 873,
			Part2: 24813,
		},
	})
}
