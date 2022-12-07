package day01

import (
	"bufio"
	"fmt"
	"sort"
	"strconv"

	"github.com/tom5760/advent-of-code/2022/inpututils"
)

func Parse(name string) ([]int, error) {
	var (
		elves []int
		elf   int
	)

	err := inpututils.Scan(name, func(scanner *bufio.Scanner) error {
		for scanner.Scan() {
			line := scanner.Bytes()
			if len(line) == 0 {
				elves = append(elves, elf)
				elf = 0

				continue
			}

			calories, err := strconv.Atoi(string(line))
			if err != nil {
				return fmt.Errorf("failed to parse line: %w", err)
			}

			elf += calories
		}

		// Make sure to record the last elf.
		elves = append(elves, elf)

		// Sort the elves from highest to lowest.
		sort.Slice(elves, func(i, j int) bool { return elves[j] < elves[i] })

		return nil
	})

	return elves, err
}

func Part1(elves []int) int {
	return elves[0]
}

func Part2(elves []int) int {
	return elves[0] + elves[1] + elves[2]
}
