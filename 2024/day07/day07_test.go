package day07_test

import (
	"testing"

	"github.com/tom5760/advent-of-code/aoc2024/aoc"
	"github.com/tom5760/advent-of-code/aoc2024/day07"
)

func TestDay07(t *testing.T) {
	aoc.Run(t, day07.Run, []aoc.Test{
		{
			Name:  "example",
			Part1: 3749,
			Part2: 11387,
		},
		{
			Name:  "input",
			Part1: 2501605301465,
			Part2: 44841372855953,
		},
	})
}
