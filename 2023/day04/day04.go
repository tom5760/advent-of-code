package day04

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"strconv"
)

func Run(lg *log.Logger, input io.Reader) (int, int, error) {
	var part1, part2 int

	stack := []int{0}
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		line := scanner.Bytes()
		parts := bytes.Fields(line)

		var points int

		winning := make(map[int]bool)
		var sep int
		for i, part := range parts[2:] {
			p := string(part)
			if p == "|" {
				sep = 2 + i
				break
			}

			n, err := strconv.Atoi(p)
			if err != nil {
				return 0, 0, fmt.Errorf("failed to parse winning number %v: %w", i, err)
			}

			winning[n] = true
		}

		var wins int

		for i, part := range parts[sep+1:] {
			p := string(part)

			n, err := strconv.Atoi(p)
			if err != nil {
				return 0, 0, fmt.Errorf("failed to parse own number %v: %w", i, err)
			}

			if winning[n] {
				wins++
				if points == 0 {
					points = 1
				} else {
					points *= 2
				}
			}
		}

		part1 += points

		var repeat int

		if len(stack) > 0 {
			repeat, stack = stack[0], stack[1:]
		}

		repeat++
		part2 += repeat

		if len(stack) < wins {
			stack = append(stack, make([]int, wins-len(stack))...)
		}

		for i := 0; i < wins; i++ {
			stack[i] += repeat
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, 0, fmt.Errorf("failed to scan input: %w", err)
	}

	return part1, part2, nil
}
