package day06

import (
	"testing"

	"github.com/tom5760/advent-of-code/2022/testutils"
)

func TestDay06(t *testing.T) {
	t.Parallel()

	testutils.Run(t, Parse, Part1, Part2, []testutils.Test[int, int]{
		{
			Name:  "example1",
			Part1: 7,
			Part2: 19,
		},
		{
			Name:  "example2",
			Part1: 5,
			Part2: 23,
		},
		{
			Name:  "example3",
			Part1: 6,
			Part2: 23,
		},
		{
			Name:  "example4",
			Part1: 10,
			Part2: 29,
		},
		{
			Name:  "example5",
			Part1: 11,
			Part2: 26,
		},
		{
			Name:  "input",
			Part1: 1647,
			Part2: 2447,
		},
	})
}
