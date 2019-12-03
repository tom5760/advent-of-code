package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/tom5760/advent-of-code/2019/common"
)

var (
	origin = Point{0, 0}
)

func main() {
	os.Exit(run())
}

func run() int {
	lineSet, err := readInput(os.Stdin)
	if err != nil {
		log.Println("failed to read input:", err)
		return 1
	}

	log.Println("(part 1):", Part1(lineSet))
	log.Println("(part 2):", Part2(lineSet))

	return 0
}

func readInput(r io.Reader) ([]Lines, error) {
	wireStrs, err := common.ReadStringSlice(r, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to read input: %w", err)
	}

	wireSet := make([]Wire, len(wireStrs))

	for i, str := range wireStrs {
		instructions, err := common.ReadStringSlice(
			strings.NewReader(str), common.ScanCommas)
		if err != nil {
			return nil, fmt.Errorf("failed to parse wire line: %w", err)
		}

		wire := make(Wire, len(instructions))

		for j, inst := range instructions {
			lenS := inst[1:]

			length, err := strconv.ParseInt(lenS, 10, 0)
			if err != nil {
				return nil, fmt.Errorf("failed to parse wire length: %w", err)
			}

			wire[j] = WireSegment{
				Direction: Direction(inst[0]),
				Length:    int(length),
			}
		}

		wireSet[i] = wire
	}

	return WireSetToLineSet(wireSet), nil
}

// Part1 - What is the Manhattan distance from the central port to the closest
// intersection?
func Part1(lineSet []Lines) int {
	points := LineSetIntersections(lineSet)
	_, distance := ClosestPoint(origin, points)

	return distance
}

// Part2 - What is the fewest combined steps the wires must take to reach an
// intersection?
func Part2(lineSet []Lines) int {
	points := LineSetIntersections(lineSet)

	steps := math.MaxInt64

	for _, point := range points {
		var sum int

		for _, line := range lineSet {
			sum += StepsToPoint(line, point)
		}

		if sum < steps {
			steps = sum
		}
	}

	return steps
}
