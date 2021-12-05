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
		expected []Line
		input    string
	}{
		{
			name: "example",
			expected: []Line{
				{{0, 9}, {5, 9}},
				{{8, 0}, {0, 8}},
				{{9, 4}, {3, 4}},
				{{2, 2}, {2, 1}},
				{{7, 0}, {7, 4}},
				{{6, 4}, {2, 0}},
				{{0, 9}, {2, 9}},
				{{3, 4}, {1, 4}},
				{{0, 0}, {8, 8}},
				{{5, 5}, {8, 2}},
			},
			input: `0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`,
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
		input    []Line
	}{
		{
			name:     "example",
			expected: 5,
			input: []Line{
				{{0, 9}, {5, 9}},
				{{8, 0}, {0, 8}},
				{{9, 4}, {3, 4}},
				{{2, 2}, {2, 1}},
				{{7, 0}, {7, 4}},
				{{6, 4}, {2, 0}},
				{{0, 9}, {2, 9}},
				{{3, 4}, {1, 4}},
				{{0, 0}, {8, 8}},
				{{5, 5}, {8, 2}},
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
		input    []Line
	}{
		{
			name:     "example",
			expected: 12,
			input: []Line{
				{{0, 9}, {5, 9}},
				{{8, 0}, {0, 8}},
				{{9, 4}, {3, 4}},
				{{2, 2}, {2, 1}},
				{{7, 0}, {7, 4}},
				{{6, 4}, {2, 0}},
				{{0, 9}, {2, 9}},
				{{3, 4}, {1, 4}},
				{{0, 0}, {8, 8}},
				{{5, 5}, {8, 2}},
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
