package day02_test

import (
	"testing"

	"github.com/tom5760/advent-of-code/aoc2024/aoc"
	"github.com/tom5760/advent-of-code/aoc2024/day02"
)

func TestDay02(t *testing.T) {
	aoc.Run(t, day02.Parse, []aoc.Test{
		{
			Name:  "example",
			Part1: 2,
			Part2: 4,
		},
		{
			Name:  "input",
			Part1: 585,
			Part2: 626,
		},
	})
}
