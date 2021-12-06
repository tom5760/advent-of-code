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

	scanner.Split(input.ScanIndex([]byte{'\n', '\n'}))

	// Grab numbers
	scanner.Scan()

	game.Numbers, err = input.
		Parser[int64]{
		SplitFunc: input.ScanIndexByte(','),
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
