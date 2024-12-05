package day03_test

import (
	"strings"
	"testing"

	"github.com/tom5760/advent-of-code/aoc2024/aoc"
	"github.com/tom5760/advent-of-code/aoc2024/day03"
)

func TestDay03(t *testing.T) {
	aoc.Run(t, day03.Run, []aoc.Test{
		{
			Name:  "example1",
			Part1: 161,
			Part2: 161,
		},
		{
			Name:  "example2",
			Part1: 161,
			Part2: 48,
		},
		{
			Name:  "input",
			Part1: 178886550,
			Part2: 87163705,
		},
	})
}

func TestLex(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		input  string
		tokens []day03.Token
	}{
		{
			name:  "empty",
			input: "",
			tokens: []day03.Token{
				{Kind: day03.TokenEOF},
			},
		},
		{
			name:  "left_paren",
			input: "(",
			tokens: []day03.Token{
				{Kind: day03.TokenParenLeft, Value: "("},
				{Kind: day03.TokenEOF},
			},
		},
		{
			name:  "right_paren",
			input: ")",
			tokens: []day03.Token{
				{Kind: day03.TokenParenRight, Value: ")"},
				{Kind: day03.TokenEOF},
			},
		},
		{
			name:  "comma",
			input: ",",
			tokens: []day03.Token{
				{Kind: day03.TokenComma, Value: ","},
				{Kind: day03.TokenEOF},
			},
		},
		{
			name:  "number",
			input: "1234567890",
			tokens: []day03.Token{
				{Kind: day03.TokenNumber, Value: "1234567890"},
				{Kind: day03.TokenEOF},
			},
		},
		{
			name:  "symbol_mul",
			input: "mul",
			tokens: []day03.Token{
				{Kind: day03.TokenSymbol, Value: "mul"},
				{Kind: day03.TokenEOF},
			},
		},
		{
			name:  "other",
			input: "&^%!",
			tokens: []day03.Token{
				{Kind: day03.TokenOther, Value: "&^%!"},
				{Kind: day03.TokenEOF},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			l := day03.Lex(strings.NewReader(test.input))

			var i int

			for actual := range l.Tokens() {
				expected := test.tokens[i]
				i++

				if actual.Kind != expected.Kind {
					t.Errorf("token %v kind mismatched; got: %v; want: %v", i, actual.Kind, expected.Kind)
				}
				if actual.Value != expected.Value {
					t.Errorf("token %v value mismatched; got: %q; want: %q", i, actual.Value, expected.Value)
				}
			}

			if n := len(test.tokens); n != i {
				t.Errorf("didn't get all tokens; got: %v; want: %v", i, n)
			}
		})
	}
}
