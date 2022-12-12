package day09

import (
	"testing"

	"github.com/tom5760/advent-of-code/2022/testutils"
)

func TestDay09(t *testing.T) {
	t.Parallel()

	testutils.Run(t, Parse, Part1, Part2, []testutils.Test[int]{
		{
			Name:  "example",
			Part1: 13,
			Part2: 1,
		},
		{
			Name:  "input",
			Part1: 6018,
			Part2: 2619,
		},
	})
}
