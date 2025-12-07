package day03

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
)

func Run(l *log.Logger, r io.Reader) (int, int, error) {
	var part1, part2 int

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Bytes()

		p1 := Joltage(line, 2)
		p2 := Joltage(line, 12)

		part1 += p1
		part2 += p2
	}
	if err := scanner.Err(); err != nil {
		return 0, 0, fmt.Errorf("failed to scan input: %w", err)
	}

	return part1, part2, nil
}

func Joltage(bank []byte, digitN int) int {
	digits := make([]byte, digitN)
	bankN := len(bank)

	i := 0
	j := bankN - digitN + 1

	for x := range digitN {
		k, digit := IndexMax(bank[i:j])

		digits[x] = digit
		i += k + 1
		j++
	}

	var v int

	for i, x := range digits {
		v += int(x-'0') * int(math.Pow(10, float64(digitN-(i+1))))
	}

	return v
}

func IndexMax(b []byte) (int, byte) {
	var (
		ri int
		rb byte
	)

	for i, v := range b {
		if v > rb {
			ri = i
			rb = v
		}
	}

	return ri, rb
}
