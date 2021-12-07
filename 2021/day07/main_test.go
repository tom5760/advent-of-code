package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestParseInput(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		expected []uint64
		input    string
	}{
		{
			name:     "example",
			expected: []uint64{16, 1, 2, 0, 4, 2, 7, 1, 2, 14},
			input:    "16,1,2,0,4,2,7,1,2,14\n",
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
		input    []uint64
	}{
		{
			name:     "example",
			expected: 37,
			input:    []uint64{16, 1, 2, 0, 4, 2, 7, 1, 2, 14},
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
		input    []uint64
	}{
		{
			name:     "example",
			expected: 168,
			input:    []uint64{16, 1, 2, 0, 4, 2, 7, 1, 2, 14},
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			actual := Part2(test.input)

			if actual != test.expected {
				t.Errorf("got %v; want %v", actual, test.expected)
			}
		})
	}
}
