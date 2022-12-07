package day07

import (
	"testing"

	"github.com/tom5760/advent-of-code/2022/testutils"
)

func TestDay07(t *testing.T) {
	t.Parallel()

	testutils.Run(t, Parse, Part1, Part2, []testutils.Test[int]{
		{
			Name:  "example",
			Part1: 95437,
			Part2: 24933642,
		},
		{
			Name:  "input",
			Part1: 919137,
			Part2: 2877389,
		},
	})
}
