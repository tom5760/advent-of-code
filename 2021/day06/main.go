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
	fish, err := ParseInput(os.Stdin)
	if err != nil {
		return fmt.Errorf("failed to parse input: %w", err)
	}

	fmt.Println("Part 1:", Part1(fish))
	fmt.Println("Part 2:", Part2(fish))

	return nil
}

// Realizing what you're trying to do, the submarine automatically produces a
// list of the ages of several hundred nearby lanternfish (your puzzle input).
func ParseInput(r io.Reader) ([]uint64, error) {
	return input.
		Parser[uint64]{
		SplitFunc: input.ScanIndexByte(','),
		ParseFunc: input.Uint(10, 64),
	}.Slice(r)
}

// Surely, each lanternfish creates a new lanternfish once every 7 days.
//
// Furthermore, you reason, a new lanternfish would surely need slightly longer
// before it's capable of producing more lanternfish: two more days for its
// first cycle.
//
// Find a way to simulate lanternfish. How many lanternfish would there be
// after 80 days?
func Part1(fish []uint64) uint64 {
	return Simulate(fish, 80)
}

// How many lanternfish would there be after 256 days?
func Part2(fish []uint64) uint64 {
	return Simulate(fish, 256)
}

func Simulate(fish []uint64, total int) uint64 {
	// Store number of fish at days 0-8.
	days := make([]uint64, 9)

	for _, age := range fish {
		days[age]++
	}

	for day := 0; day < total; day++ {
		// Shift day 0 off the front, slide everything else off.
		newfish := days[0]
		days = append(days[:0], days[1:]...)[:9]

		// Day 0 fish reset to 6...
		days[6] += newfish

		// And each generate a new fish at day 8
		days[8] = newfish
	}

	var count uint64
	for _, fish := range days {
		count += fish
	}

	return count
}
