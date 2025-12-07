package day03_test

import (
	"testing"

	"github.com/tom5760/advent-of-code/aoc2025/aoc"
	"github.com/tom5760/advent-of-code/aoc2025/day03"
)

func TestDay03(t *testing.T) {
	aoc.Run(t, day03.Run, []aoc.Test{
		{
			Name:  "example",
			Part1: 357,
			Part2: 3121910778619,
		},
		{
			Name:  "input",
			Part1: 17427,
			Part2: 173161749617495,
		},
	})
}
