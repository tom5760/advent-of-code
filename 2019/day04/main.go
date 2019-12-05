package main

import (
	"log"
	"os"
)

const (
	inputMin = 165432
	inputMax = 707912
)

func main() {
	os.Exit(run())
}

func run() int {
	log.Println("(part 1): valid password count:", Part1())
	log.Println("(part 2): valid stronger password count:", Part2())

	return 0
}

// Part1 - How many different passwords within the range given in your puzzle
// input meet these criteria?
func Part1() int {
	var count int

	for i := inputMin; i < inputMax; i++ {
		if ValidatePasswordV1(i) == nil {
			count++
		}
	}

	return count
}

// Part2 - How many different passwords within the range given in your puzzle
// input meet all of the criteria?
func Part2() int {
	var count int

	for i := inputMin; i < inputMax; i++ {
		if ValidatePasswordV2(i) == nil {
			count++
		}
	}

	return count
}
