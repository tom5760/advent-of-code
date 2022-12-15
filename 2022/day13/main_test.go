package day13

import (
	"testing"

	"github.com/tom5760/advent-of-code/2022/testutils"
)

func TestDay13(t *testing.T) {
	t.Parallel()

	testutils.Run(t, Parse, Part1, Part2, []testutils.Test[int, int]{
		{
			Name:  "example",
			Part1: 13,
			Part2: 140,
		},
		{
			Name:  "input",
			Part1: 5506,
			Part2: 21756,
		},
	})
}
