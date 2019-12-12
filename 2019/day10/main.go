package main

import (
	"log"
	"os"
)

const (
	// -1 for zero based
	vaporizeGoal = 200 - 1
)

func main() {
	os.Exit(run())
}

func run() int {
	field, err := ParseAsteroids(os.Stdin)
	if err != nil {
		log.Println("failed to parse asteroids:", err)
		return 1
	}

	station, count := FindMonitoringStation(field)

	log.Println("(part 1): asteroids detected:", count)
	log.Println("(part 2): vaporized asteroid:", Part2(field, station, vaporizeGoal))

	return 0
}

// Part1 - Find the best location for a new monitoring station. How many other
// asteroids can be detected from that location?
func Part1(field *Field) int {
	_, count := FindMonitoringStation(field)
	return count
}

// Part2 - The Elves are placing bets on which will be the 200th asteroid to be
// vaporized. Win the bet by determining which asteroid that will be; what do
// you get if you multiply its X coordinate by 100 and then add its Y
// coordinate? (For example, 8,2 becomes 802.)
func Part2(field *Field, station Point, goalIndex int) int {
	vaporized := Vaporize(field, station)

	if len(vaporized)-1 < goalIndex {
		panic("no asteroids vaporized?")
	}

	goal := vaporized[goalIndex]
	return goal.X*100 + goal.Y
}
