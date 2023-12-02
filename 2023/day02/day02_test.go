package day02_test

import (
	"testing"

	"github.com/tom5760/advent-of-code/aoc2023/day02"
	"github.com/tom5760/advent-of-code/aoc2023/util"
)

func TestDay02(t *testing.T) {
	util.RunTests(t, day02.Run, []util.DayTest{
		{Name: "example1", Part1: 8, Part2: 2286},
		{Name: "input", Part1: 1734, Part2: 70387},
	})
}
