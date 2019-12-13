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

	bodies, err := ParseBodies(lines)
	if err != nil {
		log.Println("failed to parse input:", err)
		return 1
	}

	log.Println("(part 1): total energy:", Part1(bodies))

	return 0
}

// Part1 - What is the total energy in the system after simulating the moons
// given in your scan for 1000 steps?
func Part1(bodies Bodies) int {
	bodies.Run(steps)
	return bodies.Energy()
}
