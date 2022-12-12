package day01

import (
	"testing"

	"github.com/tom5760/advent-of-code/2022/testutils"
)

func TestDay01(t *testing.T) {
	t.Parallel()

	testutils.Run(t, Parse, Part1, Part2, []testutils.Test[int, int]{
		{
			Name:  "example",
			Part1: 24000,
			Part2: 45000,
		},
		{
			Name:  "input",
			Part1: 69289,
			Part2: 205615,
		},
	})
}
