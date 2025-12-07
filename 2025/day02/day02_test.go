package day02_test

import (
	"testing"

	"github.com/tom5760/advent-of-code/aoc2025/aoc"
	"github.com/tom5760/advent-of-code/aoc2025/day02"
)

func TestDay02(t *testing.T) {
	aoc.Run(t, day02.Run, []aoc.Test{
		{
			Name:  "example",
			Part1: 1227775554,
			Part2: 4174379265,
		},
		{
			Name:  "input",
			Part1: 30599400849,
			Part2: 46270373595,
		},
	})
}
