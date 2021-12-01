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
		expected []int
		input    string
	}{
		{
			name: "example",
			expected: []int{
				199,
				200,
				208,
				210,
				200,
				207,
				240,
				269,
				260,
				263,
			},
			input: `199
200
208
210
200
207
240
269
260
263`,
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
		expected int
		input    []int
	}{
		{
			name:     "example",
			expected: 7,
			input: []int{
				199,
				200,
				208,
				210,
				200,
				207,
				240,
				269,
				260,
				263,
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
		expected int
		input    []int
	}{
		{
			name:     "example",
			expected: 5,
			input: []int{
				199,
				200,
				208,
				210,
				200,
				207,
				240,
				269,
				260,
				263,
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
