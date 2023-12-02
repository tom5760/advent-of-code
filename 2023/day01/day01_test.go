package day01_test

import (
	"testing"

	"github.com/tom5760/advent-of-code/aoc2023/day01"
	"github.com/tom5760/advent-of-code/aoc2023/util"
)

func TestDay01(t *testing.T) {
	util.RunTests(t, day01.Run, []util.DayTest{
		{Name: "example1", Part1: 142, Part2: 142},
		{Name: "example2", Part1: 209, Part2: 281},
		{Name: "input", Part1: 53386, Part2: 53312},
	})
}
