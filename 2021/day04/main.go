package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/tom5760/advent-of-code/2021/input"
)

func main() {
	if err := run(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func run() error {
	game, err := ParseInput(os.Stdin)
	if err != nil {
		return fmt.Errorf("failed to parse input: %w", err)
	}

	wins := game.Play()

	fmt.Println("Part 1:", Part1(wins))
	fmt.Println("Part 2:", Part2(wins))

	return nil
}

// It automatically generates a random order in which to draw numbers and a
// random set of boards (your puzzle input).
//
// First line is numbers, then following inputs are bingo boards, separated by
// two newlines.
func ParseInput(r io.Reader) (*Game, error) {
	var (
		game Game
		err  error
	)

	scanner := bufio.NewScanner(r)

	scanner.Split(scanTwoNewlines)

	// Grab numbers
	scanner.Scan()

	game.Numbers, err = input.
		Parser[int64]{
		SplitFunc: scanComma,
		ParseFunc: input.Int(10, 64),
	}.Slice(bytes.NewReader(scanner.Bytes()))

	if err != nil {
		return nil, fmt.Errorf("failed to parse numbers: %w", err)
	}

	// Start grabbing boards
	parser := input.Parser[int64]{
		SplitFunc: bufio.ScanWords,
		ParseFunc: input.Int(10, 64),
	}

	for scanner.Scan() {
		numbers, err := parser.Slice(bytes.NewReader(scanner.Bytes()))
		if err != nil {
			return nil, fmt.Errorf("failed to parse boards: %w", err)
		}

		game.Boards = append(game.Boards, Board{
			Numbers: numbers,
			Marked:  make([]bool, len(numbers)),
		})
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to scan boards: %w", err)
	}

	return &game, nil
}

func scanComma(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	if i := bytes.IndexByte(data, ','); i >= 0 {
		return i + 1, dropCR(data[0:i]), nil
	}

	// If we're at EOF, we have a final, non-terminated line. Return it.
	if atEOF {
		return len(data), dropCR(data), nil
	}

	// Request more data.
	return 0, nil, nil
}

// Split on two consecutive newlines.
// Modified from bufio.ScanLines:
// https://cs.opensource.google/go/go/+/refs/tags/go1.17.3:src/bufio/scan.go;drc=93200b98c75500b80a2bf7cc31c2a72deff2741c;l=345-365
func scanTwoNewlines(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	if i := bytes.Index(data, []byte{'\n', '\n'}); i >= 0 {
		// We have a block terminated by two newlines.
		return i + 2, dropCR(data[0:i]), nil
	}

	// If we're at EOF, we have a final, non-terminated line. Return it.
	if atEOF {
		return len(data), dropCR(data), nil
	}

	// Request more data.
	return 0, nil, nil
}

// dropCR drops a terminal \r from the data.
func dropCR(data []byte) []byte {
	if len(data) > 0 && data[len(data)-1] == '\r' {
		return data[0 : len(data)-1]
	}
	return data
}

// The score of the winning board can now be calculated. Start by finding the
// sum of all unmarked numbers on that board; in this case, the sum is 188.
// Then, multiply that sum by the number that was just called when the board
// won, 24, to get the final score, 188 * 24 = 4512.
//
// To guarantee victory against the giant squid, figure out which board will
// win first. What will your final score be if you choose that board?
func Part1(wins []Win) int64 {
	if len(wins) == 0 {
		panic("nobody won!")
	}

	win := wins[0]

	return win.Board.UnmarkedSum() * win.Number
}

// Figure out which board will win last. Once it wins, what would its final
// score be?
func Part2(wins []Win) int64 {
	if len(wins) == 0 {
		panic("nobody won!")
	}

	win := wins[len(wins)-1]

	return win.Board.UnmarkedSum() * win.Number
}
