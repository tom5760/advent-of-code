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
		expected []Command
		input    string
	}{
		{
			name: "example",
			expected: []Command{
				ForwardCommand(5),
				DownCommand(5),
				ForwardCommand(8),
				UpCommand(3),
				DownCommand(8),
				ForwardCommand(2),
			},
			input: `forward 5
down 5
forward 8
up 3
down 8
forward 2`,
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
		input    []Command
	}{
		{
			name:     "example",
			expected: 150,
			input: []Command{
				ForwardCommand(5),
				DownCommand(5),
				ForwardCommand(8),
				UpCommand(3),
				DownCommand(8),
				ForwardCommand(2),
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
		input    []Command
	}{
		{
			name:     "example",
			expected: 900,
			input: []Command{
				ForwardCommand(5),
				DownCommand(5),
				ForwardCommand(8),
				UpCommand(3),
				DownCommand(8),
				ForwardCommand(2),
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
