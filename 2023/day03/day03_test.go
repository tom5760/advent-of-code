package day03_test

import (
	"testing"

	"github.com/tom5760/advent-of-code/aoc2023/day03"
	"github.com/tom5760/advent-of-code/aoc2023/util"
)

func TestDay03(t *testing.T) {
	util.RunTests(t, day03.Run, []util.DayTest{
		{Name: "example1", Part1: 4361, Part2: 467835},
		{Name: "example2", Part1: 11157, Part2: 4531843},
		{Name: "input", Part1: 557705, Part2: 84266818},
	})
}
