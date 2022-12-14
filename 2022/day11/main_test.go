package day11

import (
	"testing"

	"github.com/tom5760/advent-of-code/2022/testutils"
)

func TestDay11(t *testing.T) {
	t.Parallel()

	testutils.Run(t, Parse, Part1, Part2, []testutils.Test[int, int]{
		{
			Name:  "example",
			Part1: 10605,
			Part2: 2713310158,
		},
		{
			Name:  "input",
			Part1: 111210,
			Part2: 15447387620,
		},
	})
}
