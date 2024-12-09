package day06_test

import (
	"testing"

	"github.com/tom5760/advent-of-code/aoc2024/aoc"
	"github.com/tom5760/advent-of-code/aoc2024/day06"
)

func TestDay06(t *testing.T) {
	aoc.Run(t, day06.Run, []aoc.Test{
		{
			Name:  "example",
			Part1: 41,
			Part2: 6,
		},
		{
			Name:  "input",
			Part1: 5177,
			Part2: 1686,
		},
	})
}
