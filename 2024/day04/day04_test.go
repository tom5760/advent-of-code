package day04_test

import (
	"testing"

	"github.com/tom5760/advent-of-code/aoc2024/aoc"
	"github.com/tom5760/advent-of-code/aoc2024/day04"
)

func TestDay04(t *testing.T) {
	aoc.Run(t, day04.Run, []aoc.Test{
		{
			Name:  "example",
			Part1: 18,
			Part2: 9,
		},
		{
			Name:  "input",
			Part1: 2567,
			Part2: 2029,
		},
	})
}
