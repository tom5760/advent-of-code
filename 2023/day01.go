package aoc2023

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
)

func Day01(input io.Reader) (int, int, error) {
	scanner := bufio.NewScanner(input)

	var part1, part2 int

	for scanner.Scan() {
		part1 += day01part1(scanner.Bytes())
		part2 += day01part2(scanner.Bytes())
	}

	if err := scanner.Err(); err != nil {
		return 0, 0, fmt.Errorf("failed to read input: %w", err)
	}

	return part1, part2, nil
}

func day01part1(line []byte) int {
	var (
		haveFirst   bool
		first, last byte
	)

	for _, b := range line {
		if b < '0' || b > '9' {
			continue
		}

		last = b - '0'
		if !haveFirst {
			first = last
			haveFirst = true
		}
	}

	return int((first * 10) + last)
}

func day01part2(line []byte) int {
	var (
		haveFirst   bool
		first, last byte
	)

	for i := 0; i < len(line); i++ {
		var (
			v     byte
			found bool
		)

		cur := line[i:]

		if cur[0] >= '0' && cur[0] <= '9' {
			v = line[i] - '0'
			found = true
		} else {
			for j, buf := range numBufs {
				if !bytes.HasPrefix(cur, buf) {
					continue
				}

				v = byte(j + 1)
				found = true
				break
			}
		}

		if found {
			last = v
			if !haveFirst {
				first = last
				haveFirst = true
			}
		}
	}

	return int((first * 10) + last)
}

var numBufs = [][]byte{
	{'o', 'n', 'e'},
	{'t', 'w', 'o'},
	{'t', 'h', 'r', 'e', 'e'},
	{'f', 'o', 'u', 'r'},
	{'f', 'i', 'v', 'e'},
	{'s', 'i', 'x'},
	{'s', 'e', 'v', 'e', 'n'},
	{'e', 'i', 'g', 'h', 't'},
	{'n', 'i', 'n', 'e'},
}
