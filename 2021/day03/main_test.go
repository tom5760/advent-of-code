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
			name: "example",
			expected: []uint64{
				0b00100,
				0b11110,
				0b10110,
				0b10111,
				0b10101,
				0b01111,
				0b00111,
				0b11100,
				0b10000,
				0b11001,
				0b00010,
				0b01010,
			},
			input: `00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`,
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
			expected: 198,
			input: []uint64{
				0b00100,
				0b11110,
				0b10110,
				0b10111,
				0b10101,
				0b01111,
				0b00111,
				0b11100,
				0b10000,
				0b11001,
				0b00010,
				0b01010,
			},
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
		gamma    uint64
		epsilon  uint64
		input    []uint64
	}{
		{
			name:     "example",
			expected: 230,
			input: []uint64{
				0b00100,
				0b11110,
				0b10110,
				0b10111,
				0b10101,
				0b01111,
				0b00111,
				0b11100,
				0b10000,
				0b11001,
				0b00010,
				0b01010,
			},
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
