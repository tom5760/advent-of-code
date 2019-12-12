package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	expectedPart1 = 334
	expectedPart2 = 1119
)

func TestPart1(t *testing.T) {
	field, err := readInput()
	if err != nil {
		t.Fatalf("failed to read input: %v", err)
		return
	}

	assert.Equal(t, expectedPart1, Part1(field))
}

func TestPart2(t *testing.T) {
	field, err := readInput()
	if err != nil {
		t.Fatalf("failed to read input: %v", err)
		return
	}

	station, _ := FindMonitoringStation(field)

	assert.Equal(t, expectedPart2, Part2(field, station, vaporizeGoal))
}

func readInput() (*Field, error) {
	f, err := os.Open("input")
	if err != nil {
		return nil, fmt.Errorf("failed to open input file: %w", err)
	}
	defer f.Close()

	field, err := ParseAsteroids(f)
	if err != nil {
		return nil, fmt.Errorf("failed to read input: %w", err)
	}

	return field, nil
}
