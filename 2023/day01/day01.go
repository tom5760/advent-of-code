package day01

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
)

func Run(lg *log.Logger, input io.Reader) (int, int, error) {
	scanner := bufio.NewScanner(input)

	var part1, part2 int

	for scanner.Scan() {
		line := scanner.Bytes()
		part1 += doPart1(line)
		part2 += doPart2(line)
	}

	if err := scanner.Err(); err != nil {
		return 0, 0, fmt.Errorf("failed to read input: %w", err)
	}

	return part1, part2, nil
}

func doPart1(line []byte) int {
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

var numBufs = [][]byte{
	[]byte("one"),
	[]byte("two"),
	[]byte("three"),
	[]byte("four"),
	[]byte("five"),
	[]byte("six"),
	[]byte("seven"),
	[]byte("eight"),
	[]byte("nine"),
}

func doPart2(line []byte) int {
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
