package day01

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"math"
	"slices"
	"strconv"
)

type Input struct {
	Left  []int
	Right []int
}

func Run(r io.Reader) (int, int, error) {
	var left, right []int

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		fields := bytes.Fields(scanner.Bytes())
		if len(fields) != 2 {
			return 0, 0, errors.New("unexpected line format")
		}

		l, err := strconv.Atoi(string(fields[0]))
		if err != nil {
			return 0, 0, fmt.Errorf("failed to parse left column: %w", err)
		}

		r, err := strconv.Atoi(string(fields[1]))
		if err != nil {
			return 0, 0, fmt.Errorf("failed to parse right column: %w", err)
		}

		left = append(left, l)
		right = append(right, r)
	}

	if err := scanner.Err(); err != nil {
		return 0, 0, fmt.Errorf("failed to scan input: %w", err)
	}

	if len(left) != len(right) {
		return 0, 0, errors.New("columns are not equal length")
	}

	slices.Sort(left)
	slices.Sort(right)

	p1 := part1(left, right)
	p2 := part2(left, right)

	return p1, p2, nil
}

func part1(left, right []int) int {
	var distance int

	for i := range left {
		l := left[i]
		r := right[i]

		distance += int(math.Abs(float64(l - r)))
	}

	return distance
}

func part2(left, right []int) int {
	var similarity int

	for _, l := range left {
		var count int
		for _, r := range right {
			if l == r {
				count++
			}
		}

		similarity += l * count
	}

	return similarity
}
