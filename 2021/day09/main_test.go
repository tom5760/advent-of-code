package main

import (
	"reflect"
	"strings"
	"testing"
)

const exampleInput = `2199943210
3987894921
9856789892
8767896789
9899965678
`

var exampleInputParsed = &FloorMap{
	Width:  10,
	Height: 5,
	Data: []uint8{
		2, 1, 9, 9, 9, 4, 3, 2, 1, 0,
		3, 9, 8, 7, 8, 9, 4, 9, 2, 1,
		9, 8, 5, 6, 7, 8, 9, 8, 9, 2,
		8, 7, 6, 7, 8, 9, 6, 7, 8, 9,
		9, 8, 9, 9, 9, 6, 5, 6, 7, 8,
	},
}

func TestParseInput(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		expected *FloorMap
		input    string
	}{
		{
			name:     "example",
			expected: exampleInputParsed,
			input:    exampleInput,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			actual, err := ParseInput(strings.NewReader(test.input))
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if !reflect.DeepEqual(actual, test.expected) {
				t.Errorf("got %v; want %v", actual, test.expected)
			}
		})
	}
}

func TestPart1(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		expected uint64
		input    *FloorMap
	}{
		{
			name:     "example",
			expected: 15,
			input:    exampleInputParsed,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			actual := Part1(test.input)

			if actual != test.expected {
				t.Errorf("got %v; want %v", actual, test.expected)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		expected uint64
		input    *FloorMap
	}{
		{
			name:     "example",
			expected: 1134,
			input:    exampleInputParsed,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			_ = Part1(test.input)
			actual := Part2(test.input)

			if actual != test.expected {
				t.Errorf("got %v; want %v", actual, test.expected)
			}
		})
	}
}
