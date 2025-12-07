package day01

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strconv"
)

func Run(l *log.Logger, r io.Reader) (int, int, error) {
	var part1, part2 int

	dial := 50

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Bytes()
		if len(line) < 2 {
			return 0, 0, fmt.Errorf("unexpected line: %q", string(line))
		}

		value, err := strconv.Atoi(string(line[1:]))
		if err != nil {
			return 0, 0, fmt.Errorf("failed to parse line value: %q", string(line))
		}

		dir := line[0]

		var p2 int

		switch dir {
		case 'L':
			part2 += (((100 - dial) % 100) + value) / 100
			dial -= value

		case 'R':
			dial += value
			part2 += dial / 100

		default:
			return 0, 0, fmt.Errorf("unexpected direction value: %q", string(dir))
		}

		dial = (dial%100 + 100) % 100
		part2 += p2

		if dial == 0 {
			part1++
		}
	}
	if err := scanner.Err(); err != nil {
		return 0, 0, fmt.Errorf("failed to scan input: %w", err)
	}

	return part1, part2, nil
}
