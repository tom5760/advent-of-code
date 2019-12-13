package main

import (
	"log"
	"os"

	"github.com/tom5760/advent-of-code/2019/common"
)

const (
	steps = 1000
)

func main() {
	os.Exit(run())
}

func run() int {
	lines, err := common.ReadStringSlice(os.Stdin, nil)
	if err != nil {
		log.Println("failed to read input:", err)
		return 1
	}

	system, err := ParseSystem(lines)
	if err != nil {
		log.Println("failed to parse input:", err)
		return 1
	}

	log.Println("(part 1): total energy:", Part1(system))

	// regenerate system to reset to initial state
	system, err = ParseSystem(lines)
	if err != nil {
		log.Println("failed to parse input:", err)
		return 1
	}

	log.Println("(part 2): steps:", Part2(system))

	return 0
}

// Part1 - What is the total energy in the system after simulating the moons
// given in your scan for 1000 steps?
func Part1(system *System) int {
	for i := 0; i < steps; i++ {
		system.Step()
	}
	return system.Energy()
}

// Part2 - How many steps does it take to reach the first state that exactly
// matches a previous state?
func Part2(system *System) int {
	original := make([]Body, len(system.Bodies))
	copy(original, system.Bodies)

	var xCycle, yCycle, zCycle int

	check := func(i int) {
		x := true
		y := true
		z := true

		for j := 0; j < len(original); j++ {
			og := original[j]
			now := system.Bodies[j]

			if x && (og.X != now.X || og.DX != now.DX) {
				x = false
			}
			if y && (og.Y != now.Y || og.DY != now.DY) {
				y = false
			}
			if z && (og.Z != now.Z || og.DZ != now.DZ) {
				z = false
			}
		}

		if xCycle == 0 && x {
			xCycle = i
		}
		if yCycle == 0 && y {
			yCycle = i
		}
		if zCycle == 0 && z {
			zCycle = i
		}
	}

	for i := 0; xCycle == 0 || yCycle == 0 || zCycle == 0; i++ {
		system.Step()
		check(i + 1)
	}

	return common.LCM(xCycle, yCycle, zCycle)
}
