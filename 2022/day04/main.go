package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/tom5760/advent-of-code/2022/input"
)

func main() {
	if err := run(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func run() error {
	parser := input.Parser[Pair]{
		ParseFunc: func(line []byte, outChan chan<- Pair) error {
			pair := bytes.SplitN(line, []byte{','}, 2)
			if len(pair) != 2 {
				return fmt.Errorf("failed to split pair")
			}

			firstSections := bytes.SplitN(pair[0], []byte{'-'}, 2)
			if len(firstSections) != 2 {
				return fmt.Errorf("failed to split first sections")
			}

			secondSections := bytes.SplitN(pair[1], []byte{'-'}, 2)
			if len(secondSections) != 2 {
				return fmt.Errorf("failed to split second sections")
			}

			firstStart, err := strconv.Atoi(string(firstSections[0]))
			if err != nil {
				return fmt.Errorf("failed to parse first pair section start: %w", err)
			}

			firstEnd, err := strconv.Atoi(string(firstSections[1]))
			if err != nil {
				return fmt.Errorf("failed to parse first pair section start: %w", err)
			}

			secondStart, err := strconv.Atoi(string(secondSections[0]))
			if err != nil {
				return fmt.Errorf("failed to parse second pair section start: %w", err)
			}

			secondEnd, err := strconv.Atoi(string(secondSections[1]))
			if err != nil {
				return fmt.Errorf("failed to parse second pair section start: %w", err)
			}

			outChan <- Pair{
				First: Sections{
					Start: firstStart,
					End:   firstEnd,
				},
				Second: Sections{
					Start: secondStart,
					End:   secondEnd,
				},
			}

			return nil
		},
	}

	pairs, err := parser.ReadFileSlice("./day04/input")
	if err != nil {
		return fmt.Errorf("failed to parse input: %w", err)
	}

	Part1(pairs)
	Part2(pairs)

	return nil
}

type (
	Pair struct {
		First  Sections
		Second Sections
	}

	Sections struct {
		Start int
		End   int
	}
)

func (s Sections) Contain(other Sections) bool {
	return s.Start <= other.Start && s.End >= other.End
}

func (s Sections) Overlap(other Sections) bool {
	return s.Start <= other.End && s.End >= other.Start
}

func Part1(pairs []Pair) {
	var contained int

	for _, pair := range pairs {
		if pair.First.Contain(pair.Second) || pair.Second.Contain(pair.First) {
			contained++
		}
	}

	fmt.Println("PART 1:", contained)
}

func Part2(pairs []Pair) {
	var overlaps int

	for _, pair := range pairs {
		if pair.First.Overlap(pair.Second) || pair.Second.Overlap(pair.First) {
			overlaps++
		}
	}

	fmt.Println("PART 2:", overlaps)
}
