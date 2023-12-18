package day04_test

import (
	"testing"

	"github.com/tom5760/advent-of-code/aoc2023/day04"
	"github.com/tom5760/advent-of-code/aoc2023/util"
)

func TestDay04(t *testing.T) {
	util.RunTests(t, day04.Run, []util.DayTest{
		{Name: "example1", Part1: 13, Part2: 30},
		{Name: "input", Part1: 19135, Part2: 5704953},
	})
}
