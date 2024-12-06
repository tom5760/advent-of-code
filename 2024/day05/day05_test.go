package day05_test

import (
	"testing"

	"github.com/tom5760/advent-of-code/aoc2024/aoc"
	"github.com/tom5760/advent-of-code/aoc2024/day05"
)

func TestDay05(t *testing.T) {
	aoc.Run(t, day05.Run, []aoc.Test{
		{
			Name:  "example",
			Part1: 143,
			Part2: 123,
		},
		{
			Name:  "input",
			Part1: 6384,
			Part2: 5353,
		},
	})
}
