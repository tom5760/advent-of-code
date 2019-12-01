package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/tom5760/advent-of-code/2019/common"
)

const (
	expectedPart1 = 3299598
	expectedPart2 = 4946546
)

func TestPart1(t *testing.T) {
	modules, err := readInput()
	if err != nil {
		t.Fatalf("failed to read input: %v", err)
		return
	}

	actual := Part1(modules)
	if expectedPart1 != actual {
		t.Errorf("part 1 expected %v != %v", expectedPart1, actual)
	}
}

func TestPart2(t *testing.T) {
	modules, err := readInput()
	if err != nil {
		t.Fatalf("failed to read input: %v", err)
		return
	}

	actual := Part2(modules)
	if expectedPart2 != actual {
		t.Errorf("part 2 expected %v != %v", expectedPart2, actual)
	}
}

func readInput() ([]uint64, error) {
	f, err := os.Open("input")
	if err != nil {
		return nil, fmt.Errorf("failed to open input file: %w", err)
	}
	defer f.Close()

	modules, err := common.ReadUint64Slice(f)
	if err != nil {
		return nil, fmt.Errorf("failed to read input: %w", err)
	}

	return modules, nil
}
