package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"math/bits"
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
	signals, err := ParseInput(os.Stdin)
	if err != nil {
		return fmt.Errorf("failed to parse input: %w", err)
	}

	fmt.Println("Part 1:", Part1(signals))
	fmt.Println("Part 2:", Part2(signals))

	return nil
}

// For each display, you watch the changing signals for a while, make a note of
// all ten unique signal patterns you see, and then write down a single four
// digit output value (your puzzle input). Using the signal patterns, you
// should be able to work out which pattern corresponds to which digit.
func ParseInput(r io.Reader) ([]Signals, error) {
	return input.Parser[Signals]{
		ParseFunc: func(scanner *bufio.Scanner) (Signals, error) {
			parts := bytes.Split(scanner.Bytes(), []byte{' ', '|', ' '})
			if len(parts) != 2 {
				return Signals{}, fmt.Errorf("failed to split signal")
			}

			samples := bytes.Fields(parts[0])
			outputs := bytes.Fields(parts[1])

			if len(samples) != 10 {
				return Signals{}, fmt.Errorf("unexpected number of samples")
			}

			if len(outputs) != 4 {
				return Signals{}, fmt.Errorf("unexpected number of outputs")
			}

			var (
				signals Signals
				err     error
			)

			for i, digit := range samples {
				if signals.Samples[i], err = ParseSignal(digit); err != nil {
					return Signals{}, fmt.Errorf("failed to parse sample digit: %w", err)
				}
			}

			for i, digit := range outputs {
				if signals.Output[i], err = ParseSignal(digit); err != nil {
					return Signals{}, fmt.Errorf("failed to parse output digit: %w", err)
				}
			}

			return signals, nil
		},
	}.Slice(r)
}

// In the output values, how many times do digits 1, 4, 7, or 8 appear?
func Part1(signals []Signals) uint64 {
	var count uint64

	for _, signal := range signals {
		for _, output := range signal.Output {
			switch bits.OnesCount8(uint8(output)) {
			// 1 = 2 segments, 4 = 4 segments, 7 = 3 segments, 8 = 7 segments
			case 2, 4, 3, 7:
				count++
			}
		}
	}

	return count
}

// For each entry, determine all of the wire/segment connections and decode the
// four-digit output values. What do you get if you add up all of the output
// values?
func Part2(signals []Signals) uint64 {
	var sum uint64

	for _, signal := range signals {
		sum += signal.Decode()
	}

	return sum
}
