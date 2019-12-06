package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/tom5760/advent-of-code/2019/common"
)

const (
	expectedPart1 = 119831
	expectedPart2 = 322
)

func TestPart1(t *testing.T) {
	universe, err := readInput()
	if err != nil {
		t.Fatalf("failed to read input: %v", err)
		return
	}

	actual := Part1(universe)

	if expectedPart1 != actual {
		t.Errorf("part 1 expected %v != %v", expectedPart1, actual)
	}
}

func TestPart2(t *testing.T) {
	universe, err := readInput()
	if err != nil {
		t.Fatalf("failed to read input: %v", err)
		return
	}

	actual := Part2(universe)

	if expectedPart2 != actual {
		t.Errorf("part 2 expected %v != %v", expectedPart2, actual)
	}
}

func readInput() (*Universe, error) {
	f, err := os.Open("input")
	if err != nil {
		return nil, fmt.Errorf("failed to open input file: %w", err)
	}
	defer f.Close()

	lines, err := common.ReadStringSlice(f, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to read input: %w", err)
	}

	return ParseUniverse(lines), nil
}
