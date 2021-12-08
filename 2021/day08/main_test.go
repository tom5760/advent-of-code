package main

import (
	"reflect"
	"strings"
	"testing"
)

const exampleInput = `be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe
edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc
fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg
fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb
aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga | gecf egdcabf bgf bfgea
fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf | gebdcfa ecba ca fadegcb
dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf | cefg dcbef fcge gbcadfe
bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef
egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg | gbdfcae bgc cg cgb
gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce`

var exampleInputParsed = []Signals{
	{
		Samples: [10]Signal{
			0b00010010, 0b01111111, 0b01111110, 0b01111101, 0b01010110,
			0b01111100, 0b01111011, 0b00111110, 0b00101111, 0b00011010,
		},
		Output: [4]Signal{0b01111111, 0b00111110, 0b01111110, 0b01010110},
	},
	{
		Samples: [10]Signal{
			0b01111011, 0b01011110, 0b01000110, 0b01000100, 0b01111111,
			0b01111010, 0b01101111, 0b00011111, 0b01111110, 0b01110100,
		},
		Output: [4]Signal{0b01111110, 0b01000110, 0b01111111, 0b01000100},
	},
	{
		Samples: [10]Signal{
			0b01111011, 0b01000100, 0b00011111, 0b01101011, 0b01101111,
			0b01111110, 0b01001111, 0b01100101, 0b01000110, 0b01111111,
		},
		Output: [4]Signal{0b01000100, 0b01000100, 0b01101111, 0b1000110},
	},
	{
		Samples: [10]Signal{
			0b01111110, 0b00001110, 0b00111111, 0b01011011, 0b00100111,
			0b00000110, 0b00111101, 0b00011111, 0b01111101, 0b01111111,
		},
		Output: [4]Signal{0b00111111, 0b00011111, 0b01111101, 0b110},
	},
	{
		Samples: [10]Signal{
			0b01111111, 0b01100010, 0b01100000, 0b01110011, 0b00111011,
			0b01110100, 0b01010111, 0b01110111, 0b01011111, 0b01101111,
		},
		Output: [4]Signal{0b01110100, 0b01111111, 0b01100010, 0b1110011},
	},
	{
		Samples: [10]Signal{
			0b01110011, 0b00000101, 0b01110111, 0b01111111, 0b01111101,
			0b01101110, 0b00010111, 0b01111011, 0b01100111, 0b00100101,
		},
		Output: [4]Signal{0b01111111, 0b00010111, 0b00000101, 0b1111111},
	},
	{
		Samples: [10]Signal{
			0b01101110, 0b01101000, 0b01111111, 0b01110100, 0b01111011,
			0b00111111, 0b00111110, 0b01001111, 0b01111110, 0b01100000,
		},
		Output: [4]Signal{0b01110100, 0b00111110, 0b01110100, 0b1111111},
	},
	{
		Samples: [10]Signal{
			0b01111110, 0b01110111, 0b01110110, 0b01111101, 0b01001111,
			0b00011000, 0b00111010, 0b00011100, 0b01111111, 0b01011110,
		},
		Output: [4]Signal{0b00011000, 0b01110111, 0b01001111, 0b1110110},
	},
	{
		Samples: [10]Signal{
			0b01111011, 0b01111110, 0b01011100, 0b00110111, 0b01000110,
			0b01111111, 0b01000100, 0b01101111, 0b01111010, 0b01110110,
		},
		Output: [4]Signal{0b01111111, 0b01000110, 0b01000100, 0b1000110},
	},
	{
		Samples: [10]Signal{
			0b01100111, 0b01100100, 0b01111111, 0b01010111, 0b01100000,
			0b01011111, 0b01110001, 0b01110111, 0b00101111, 0b01111110,
		},
		Output: [4]Signal{0b01110001, 0b01100111, 0b01100000, 0b1010111},
	},
}

func TestParseInput(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		expected []Signals
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
		input    []Signals
	}{
		{
			name:     "example",
			expected: 26,
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
		input    []Signals
	}{
		{
			name:     "example",
			expected: 61229,
			input:    exampleInputParsed,
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
