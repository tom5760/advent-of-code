package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	if err := run(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func run() error {
	f, err := os.Open("./day01/input")
	if err != nil {
		return fmt.Errorf("failed to open input: %w", err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var (
		elves []uint64
		elf   uint64
	)

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			elves = append(elves, elf)
			elf = 0

			continue
		}

		calories, err := strconv.ParseUint(line, 10, 0)
		if err != nil {
			return fmt.Errorf("failed to parse input: %w", err)
		}

		elf += calories
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("failed to scan input: %w", err)
	}

	sort.Slice(elves, func(i, j int) bool { return elves[j] < elves[i] })

	log.Println("PART 1:", elves[0])
	log.Println("PART 2:", elves[0]+elves[1]+elves[2])

	return nil
}
