package main

import (
	"log"
	"os"
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

	log.Println("(part 1): asteroids detected:", Part1(field))

	return 0
}

// Part1 - Find the best location for a new monitoring station. How many other
// asteroids can be detected from that location?
func Part1(field *Field) int {
	_, count := FindMonitoringStation(field)
	return count
}
