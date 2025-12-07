package day04_test

import (
	"testing"

	"github.com/tom5760/advent-of-code/aoc2025/aoc"
	"github.com/tom5760/advent-of-code/aoc2025/day04"
)

func TestDay04(t *testing.T) {
	aoc.Run(t, day04.Run, []aoc.Test{
		{
			Name:  "example",
			Part1: 13,
			Part2: 43,
		},
		{
			Name:  "input",
			Part1: 1518,
			Part2: 8665,
		},
	})
}
