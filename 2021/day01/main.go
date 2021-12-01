package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
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
func ParseInput(r io.Reader) ([]int, error) {
	var depths []int

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		n, err := strconv.ParseInt(scanner.Text(), 10, 0)
		if err != nil {
			return nil, fmt.Errorf("failed to parse '%v': %w", scanner.Text(), err)
		}

		depths = append(depths, int(n))
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to scan input: %w", err)
	}

	return depths, nil
}

// How many measurements are larger than the previous measurement?
func Part1(depths []int) int {
	if len(depths) < 2 {
		return 0
	}

	var count int

	for i := range depths[1:] {
		if depths[i+1] > depths[i] {
			count++
		}
	}

	return count
}

// Consider sums of a three-measurement sliding window. How many sums are
// larger than the previous sum?
func Part2(depths []int) int {
	if len(depths) < 4 {
		return 0
	}

	var count int

	for i := range depths[3:] {
		a := depths[i] + depths[i+1] + depths[i+2]
		b := depths[i+1] + depths[i+2] + depths[i+3]

		if b > a {
			count++
		}
	}

	return count
}
