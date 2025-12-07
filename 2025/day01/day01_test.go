package day01_test

import (
	"testing"

	"github.com/tom5760/advent-of-code/aoc2025/aoc"
	"github.com/tom5760/advent-of-code/aoc2025/day01"
)

func TestDay01(t *testing.T) {
	aoc.Run(t, day01.Run, []aoc.Test{
		{
			Name:  "example",
			Part1: 3,
			Part2: 6,
		},
		{
			Name:  "input",
			Part1: 1150,
			Part2: 6738,
		},
	})
}
