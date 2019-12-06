package main

import (
	"log"
	"os"

	"github.com/tom5760/advent-of-code/2019/common"
)

const (
	sourceID = "YOU"
	targetID = "SAN"
)

func main() {
	os.Exit(run())
}

func run() int {
	lines, err := common.ReadStringSlice(os.Stdin, nil)
	if err != nil {
		log.Println("failed to read input:", err)
	}

	universe := ParseUniverse(lines)

	log.Println("(part 1) total orbits:", Part1(universe))
	log.Println("(part 2) total transfers:", Part2(universe))

	return 0
}

// Part1 - What is the total number of direct and indirect orbits in your map
// data?
func Part1(u *Universe) int {
	return TotalOrbits(u)
}

// Part2 - What is the minimum number of orbital transfers required to move
// from the object YOU are orbiting to the object SAN is orbiting? (Between the
// objects they are orbiting - not between YOU and SAN.)
func Part2(u *Universe) int {
	source := u.Planets[sourceID]
	target := u.Planets[targetID]
	return TotalTransfers(u, source, target)
}
