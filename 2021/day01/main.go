package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/tom5760/advent-of-code/2021/input"
)

func main() {
	if err := run(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func run() error {
	depths, err := ParseInput(os.Stdin)
	if err != nil {
		return fmt.Errorf("failed to parse input: %w", err)
	}

	fmt.Println("Part 1:", Part1(depths))
	fmt.Println("Part 2:", Part2(depths))

	return nil
}

// As the submarine drops below the surface of the ocean, it automatically
// performs a sonar sweep of the nearby sea floor. On a small screen, the sonar
// sweep report (your puzzle input) appears: each line is a measurement of the
// sea floor depth as the sweep looks further and further away from the
// submarine.
func ParseInput(r io.Reader) ([]int64, error) {
	return input.Parser[int64]{
		ParseFunc: input.Int(10, 64),
	}.Slice(r)
}

// How many measurements are larger than the previous measurement?
func Part1(depths []int64) int64 {
	if len(depths) < 2 {
		return 0
	}

	var count int64

	for i := range depths[1:] {
		if depths[i+1] > depths[i] {
			count++
		}
	}

	return count
}

// Consider sums of a three-measurement sliding window. How many sums are
// larger than the previous sum?
func Part2(depths []int64) int64 {
	if len(depths) < 4 {
		return 0
	}

	var count int64

	for i := range depths[3:] {
		a := depths[i] + depths[i+1] + depths[i+2]
		b := depths[i+1] + depths[i+2] + depths[i+3]

		if b > a {
			count++
		}
	}

	return count
}
