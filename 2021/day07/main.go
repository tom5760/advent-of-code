package main

import (
	"fmt"
	"io"
	"log"
	"math"
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
	crabs, err := ParseInput(os.Stdin)
	if err != nil {
		return fmt.Errorf("failed to parse input: %w", err)
	}

	fmt.Println("Part 1:", Part1(crabs))
	fmt.Println("Part 2:", Part2(crabs))

	return nil
}

// You quickly make a list of the horizontal position of each crab (your puzzle input).
func ParseInput(r io.Reader) ([]uint64, error) {
	return input.
		Parser[uint64]{
		SplitFunc: input.ScanIndexByte(','),
		ParseFunc: input.Uint(10, 64),
	}.Slice(r)
}

// Determine the horizontal position that the crabs can align to using the
// least fuel possible. How much fuel must they spend to align to that
// position?
func Part1(crabs []uint64) uint64 {
	min, max := minmax(crabs)

	var minDistance uint64 = math.MaxUint64

	for i := min; i <= max; i++ {
		d := distance(crabs, i)
		if d < minDistance {
			minDistance = d
		}
	}

	return minDistance
}

// As it turns out, crab submarine engines don't burn fuel at a constant rate.
// Instead, each change of 1 step in horizontal position costs 1 more unit of
// fuel than the last: the first step costs 1, the second step costs 2, the
// third step costs 3, and so on.
//
// Determine the horizontal position that the crabs can align to using the
// least fuel possible so they can make you an escape route! How much fuel must
// they spend to align to that position?
func Part2(crabs []uint64) uint64 {
	min, max := minmax(crabs)

	var minDistance uint64 = math.MaxUint64

	for i := min; i <= max; i++ {
		d := distance2(crabs, i)
		if d < minDistance {
			minDistance = d
		}
	}

	return minDistance
}

func minmax(xs []uint64) (uint64, uint64) {
	var (
		min uint64 = math.MaxUint64
		max uint64 = 0
	)

	for _, x := range xs {
		if x < min {
			min = x
		}
		if x > max {
			max = x
		}
	}

	return min, max
}

func distance(xs []uint64, n uint64) uint64 {
	var d uint64

	for _, x := range xs {
		d += uint64(abs(int64(x) - int64(n)))
	}

	return d
}

func distance2(xs []uint64, n uint64) uint64 {
	var d uint64

	for _, x := range xs {
		d += rangeSum(uint64(abs(int64(x) - int64(n))))
	}

	return d
}

func abs(x int64) int64 {
	if x < 0 {
		return x * -1
	}

	return x
}

func rangeSum(x uint64) uint64 {
	var sum uint64
	for i := uint64(1); i <= x; i++ {
		sum += i
	}
	return sum
}
