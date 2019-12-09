package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/tom5760/advent-of-code/2019/common"
)

const (
	expectedPart1 = 2500

//  expectedPart2 = 69113332
)

func TestPart1(t *testing.T) {
	img, err := readInput()
	if err != nil {
		t.Fatalf("failed to read input: %v", err)
		return
	}

	actual := Part1(img)

	if expectedPart1 != actual {
		t.Errorf("part 1 expected %v != %v", expectedPart1, actual)
	}
}

//func TestPart2(t *testing.T) {
//  memory, err := readInput()
//  if err != nil {
//    t.Fatalf("failed to read input: %v", err)
//    return
//  }

//  actual, err := Part2(memory)
//  if err != nil {
//    t.Fatalf("failed to run part 2: %v", err)
//    return
//  }

//  if expectedPart2 != actual {
//    t.Errorf("part 2 expected %v != %v", expectedPart2, actual)
//  }
//}

func readInput() (*SIF, error) {
	f, err := os.Open("input")
	if err != nil {
		return nil, fmt.Errorf("failed to open input file: %w", err)
	}
	defer f.Close()

	input, err := common.ReadIntSlice(f, common.ScanDigits)
	if err != nil {
		return nil, fmt.Errorf("failed to read input: %w", err)
	}

	img, err := ParseSIF(width, height, input)
	if err != nil {
		return nil, fmt.Errorf("failed to parse input: %w", err)
	}

	return img, nil
}
