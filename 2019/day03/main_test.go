package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	expectedPart1 = 1431
	expectedPart2 = 48012
)

func TestPart1(t *testing.T) {
	tests := []struct {
		wires    []Wire
		distance int
	}{
		{
			wires: []Wire{
				{
					{Direction: DirectionRight, Length: 8},
					{Direction: DirectionUp, Length: 5},
					{Direction: DirectionLeft, Length: 5},
					{Direction: DirectionDown, Length: 3},
				},
				{
					{Direction: DirectionUp, Length: 7},
					{Direction: DirectionRight, Length: 6},
					{Direction: DirectionDown, Length: 4},
					{Direction: DirectionLeft, Length: 4},
				},
			},
			distance: 6,
		},
		{
			wires: []Wire{
				{
					{Direction: DirectionRight, Length: 75},
					{Direction: DirectionDown, Length: 30},
					{Direction: DirectionRight, Length: 83},
					{Direction: DirectionUp, Length: 83},
					{Direction: DirectionLeft, Length: 12},
					{Direction: DirectionDown, Length: 49},
					{Direction: DirectionRight, Length: 71},
					{Direction: DirectionUp, Length: 7},
					{Direction: DirectionLeft, Length: 72},
				},
				{
					{Direction: DirectionUp, Length: 62},
					{Direction: DirectionRight, Length: 66},
					{Direction: DirectionUp, Length: 55},
					{Direction: DirectionRight, Length: 34},
					{Direction: DirectionDown, Length: 71},
					{Direction: DirectionRight, Length: 55},
					{Direction: DirectionDown, Length: 58},
					{Direction: DirectionRight, Length: 83},
				},
			},
			distance: 159,
		},
		{
			wires: []Wire{
				{
					{Direction: DirectionRight, Length: 98},
					{Direction: DirectionUp, Length: 47},
					{Direction: DirectionRight, Length: 26},
					{Direction: DirectionDown, Length: 63},
					{Direction: DirectionRight, Length: 33},
					{Direction: DirectionUp, Length: 87},
					{Direction: DirectionLeft, Length: 62},
					{Direction: DirectionDown, Length: 20},
					{Direction: DirectionRight, Length: 33},
					{Direction: DirectionUp, Length: 53},
					{Direction: DirectionRight, Length: 51},
				},
				{
					{Direction: DirectionUp, Length: 98},
					{Direction: DirectionRight, Length: 91},
					{Direction: DirectionDown, Length: 20},
					{Direction: DirectionRight, Length: 16},
					{Direction: DirectionDown, Length: 67},
					{Direction: DirectionRight, Length: 40},
					{Direction: DirectionUp, Length: 7},
					{Direction: DirectionRight, Length: 15},
					{Direction: DirectionUp, Length: 6},
					{Direction: DirectionRight, Length: 7},
				},
			},
			distance: 135,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			lineSet := WireSetToLineSet(tt.wires)
			actual := Part1(lineSet)

			assert.Equal(t, tt.distance, actual)
		})
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		wires []Wire
		steps int
	}{
		{
			wires: []Wire{
				{
					{Direction: DirectionRight, Length: 8},
					{Direction: DirectionUp, Length: 5},
					{Direction: DirectionLeft, Length: 5},
					{Direction: DirectionDown, Length: 3},
				},
				{
					{Direction: DirectionUp, Length: 7},
					{Direction: DirectionRight, Length: 6},
					{Direction: DirectionDown, Length: 4},
					{Direction: DirectionLeft, Length: 4},
				},
			},
			steps: 30,
		},
		{
			wires: []Wire{
				{
					{Direction: DirectionRight, Length: 75},
					{Direction: DirectionDown, Length: 30},
					{Direction: DirectionRight, Length: 83},
					{Direction: DirectionUp, Length: 83},
					{Direction: DirectionLeft, Length: 12},
					{Direction: DirectionDown, Length: 49},
					{Direction: DirectionRight, Length: 71},
					{Direction: DirectionUp, Length: 7},
					{Direction: DirectionLeft, Length: 72},
				},
				{
					{Direction: DirectionUp, Length: 62},
					{Direction: DirectionRight, Length: 66},
					{Direction: DirectionUp, Length: 55},
					{Direction: DirectionRight, Length: 34},
					{Direction: DirectionDown, Length: 71},
					{Direction: DirectionRight, Length: 55},
					{Direction: DirectionDown, Length: 58},
					{Direction: DirectionRight, Length: 83},
				},
			},
			steps: 610,
		},
		{
			wires: []Wire{
				{
					{Direction: DirectionRight, Length: 98},
					{Direction: DirectionUp, Length: 47},
					{Direction: DirectionRight, Length: 26},
					{Direction: DirectionDown, Length: 63},
					{Direction: DirectionRight, Length: 33},
					{Direction: DirectionUp, Length: 87},
					{Direction: DirectionLeft, Length: 62},
					{Direction: DirectionDown, Length: 20},
					{Direction: DirectionRight, Length: 33},
					{Direction: DirectionUp, Length: 53},
					{Direction: DirectionRight, Length: 51},
				},
				{
					{Direction: DirectionUp, Length: 98},
					{Direction: DirectionRight, Length: 91},
					{Direction: DirectionDown, Length: 20},
					{Direction: DirectionRight, Length: 16},
					{Direction: DirectionDown, Length: 67},
					{Direction: DirectionRight, Length: 40},
					{Direction: DirectionUp, Length: 7},
					{Direction: DirectionRight, Length: 15},
					{Direction: DirectionUp, Length: 6},
					{Direction: DirectionRight, Length: 7},
				},
			},
			steps: 410,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			lineSet := WireSetToLineSet(tt.wires)
			actual := Part2(lineSet)

			assert.Equal(t, tt.steps, actual)
		})
	}
}

func TestPart1_input(t *testing.T) {
	lines, err := readTestInput()
	if err != nil {
		t.Fatalf("failed to read input: %v", err)
		return
	}

	actual := Part1(lines)

	assert.Equal(t, expectedPart1, actual)
}

func TestPart2_input(t *testing.T) {
	lines, err := readTestInput()
	if err != nil {
		t.Fatalf("failed to read input: %v", err)
		return
	}

	actual := Part2(lines)

	assert.Equal(t, expectedPart2, actual)
}

func readTestInput() ([]Lines, error) {
	f, err := os.Open("input")
	if err != nil {
		return nil, fmt.Errorf("failed to open input file: %w", err)
	}
	defer f.Close()

	return readInput(f)
}
