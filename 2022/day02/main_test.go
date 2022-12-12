package day02

import (
	"testing"

	"github.com/tom5760/advent-of-code/2022/testutils"
)

func TestDay02(t *testing.T) {
	t.Parallel()

	testutils.Run(t, Parse, Part1, Part2, []testutils.Test[int, int]{
		{
			Name:  "example",
			Part1: 15,
			Part2: 12,
		},
		{
			Name:  "input",
			Part1: 11906,
			Part2: 11186,
		},
	})
}
