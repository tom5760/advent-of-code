package main

import (
	"fmt"
	"log"
	"os"
	"sort"
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
	parser := input.Parser[uint64]{
		ParseFunc: Parse(),
	}

	elves, err := parser.ReadFileSlice("./day01/input")
	if err != nil {
		return fmt.Errorf("failed to parse input: %w", err)
	}

	sort.Slice(elves, func(i, j int) bool { return elves[j] < elves[i] })

	log.Println("PART 1:", elves[0])
	log.Println("PART 2:", elves[0]+elves[1]+elves[2])

	return nil
}

func Parse() input.ParseFunc[uint64] {
	var elf uint64

	return func(input []byte, outChan chan<- uint64) error {
		if len(input) == 0 {
			outChan <- elf
			elf = 0

			return nil
		}

		calories, err := strconv.ParseUint(string(input), 10, 0)
		if err != nil {
			return fmt.Errorf("failed to parse input: %w", err)
		}

		elf += calories

		return nil
	}
}
