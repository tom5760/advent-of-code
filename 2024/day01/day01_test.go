package day01_test

import (
	"testing"

	"github.com/tom5760/advent-of-code/aoc2024/aoc"
	"github.com/tom5760/advent-of-code/aoc2024/day01"
)

func TestDay01(t *testing.T) {
	aoc.Run(t, day01.Run, []aoc.Test{
		{
			Name:  "example",
			Part1: 11,
			Part2: 31,
		},
		{
			Name:  "input",
			Part1: 1941353,
			Part2: 22539317,
		},
	})
}
