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
		expected *Game
		input    string
	}{
		{
			name: "example",
			expected: &Game{
				Numbers: []int64{
					7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12,
					22, 18, 20, 8, 19, 3, 26, 1,
				},
				Boards: []Board{
					{
						Numbers: []int64{
							22, 13, 17, 11, 0,
							8, 2, 23, 4, 24,
							21, 9, 14, 16, 7,
							6, 10, 3, 18, 5,
							1, 12, 20, 15, 19,
						},
						Marked: make([]bool, 25),
					},
					{
						Numbers: []int64{
							3, 15, 0, 2, 22,
							9, 18, 13, 17, 5,
							19, 8, 7, 25, 23,
							20, 11, 10, 24, 4,
							14, 21, 16, 12, 6,
						},
						Marked: make([]bool, 25),
					},
					{
						Numbers: []int64{
							14, 21, 17, 24, 4,
							10, 16, 15, 9, 19,
							18, 8, 23, 26, 20,
							22, 11, 13, 6, 5,
							2, 0, 12, 3, 7,
						},
						Marked: make([]bool, 25),
					},
				},
			},
			input: `7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19

 3 15  0  2 22
 9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
 2  0 12  3  7`,
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
		expected int64
		input    *Game
	}{
		{
			name:     "example",
			expected: 4512,
			input: &Game{
				Numbers: []int64{
					7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12,
					22, 18, 20, 8, 19, 3, 26, 1,
				},
				Boards: []Board{
					{
						Numbers: []int64{
							22, 13, 17, 11, 0,
							8, 2, 23, 4, 24,
							21, 9, 14, 16, 7,
							6, 10, 3, 18, 5,
							1, 12, 20, 15, 19,
						},
						Marked: make([]bool, 25),
					},
					{
						Numbers: []int64{
							3, 15, 0, 2, 22,
							9, 18, 13, 17, 5,
							19, 8, 7, 25, 23,
							20, 11, 10, 24, 4,
							14, 21, 16, 12, 6,
						},
						Marked: make([]bool, 25),
					},
					{
						Numbers: []int64{
							14, 21, 17, 24, 4,
							10, 16, 15, 9, 19,
							18, 8, 23, 26, 20,
							22, 11, 13, 6, 5,
							2, 0, 12, 3, 7,
						},
						Marked: make([]bool, 25),
					},
				},
			},
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			wins := test.input.Play()
			actual := Part1(wins)

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
		expected int64
		input    *Game
	}{
		{
			name:     "example",
			expected: 1924,
			input: &Game{
				Numbers: []int64{
					7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12,
					22, 18, 20, 8, 19, 3, 26, 1,
				},
				Boards: []Board{
					{
						Numbers: []int64{
							22, 13, 17, 11, 0,
							8, 2, 23, 4, 24,
							21, 9, 14, 16, 7,
							6, 10, 3, 18, 5,
							1, 12, 20, 15, 19,
						},
						Marked: make([]bool, 25),
					},
					{
						Numbers: []int64{
							3, 15, 0, 2, 22,
							9, 18, 13, 17, 5,
							19, 8, 7, 25, 23,
							20, 11, 10, 24, 4,
							14, 21, 16, 12, 6,
						},
						Marked: make([]bool, 25),
					},
					{
						Numbers: []int64{
							14, 21, 17, 24, 4,
							10, 16, 15, 9, 19,
							18, 8, 23, 26, 20,
							22, 11, 13, 6, 5,
							2, 0, 12, 3, 7,
						},
						Marked: make([]bool, 25),
					},
				},
			},
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			wins := test.input.Play()
			actual := Part2(wins)

			if actual != test.expected {
				t.Errorf("got %v; want %v", actual, test.expected)
			}
		})
	}
}
