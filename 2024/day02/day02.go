package day02

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"slices"
	"strconv"
)

type (
	Input  []Report
	Report []int
)

// Each column is split by one space.
var delimiter = []byte{' '}

func Parse(r io.Reader) (Input, error) {
	var input Input

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		parts := bytes.Split(scanner.Bytes(), delimiter)
		report := make(Report, len(parts))

		for i, part := range parts {
			level, err := strconv.Atoi(string(part))
			if err != nil {
				return nil, fmt.Errorf("failed to parse level: %w", err)
			}

			report[i] = level
		}

		if len(report) == 0 {
			return nil, errors.New("empty report")
		}

		input = append(input, report)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to scan input: %w", err)
	}

	return input, nil
}

func (r Report) IsSafe() bool {
	if len(r) == 1 {
		return true
	}

	var isSet, isIncreasing bool

	for i, cur := range r[1:] {
		prev := r[i]

		switch cur - prev {
		case 1, 2, 3:
			if isSet {
				if !isIncreasing {
					return false
				}
			} else {
				isSet = true
				isIncreasing = true
			}

		case -1, -2, -3:
			if isSet {
				if isIncreasing {
					return false
				}
			} else {
				isSet = true
				isIncreasing = false
			}

		default:
			return false
		}
	}

	return true
}

func (r Report) IsSafeProblemDampener() bool {
	if r.IsSafe() {
		return true
	}

	for i := range r {
		if r.removeLevel(i).IsSafe() {
			return true
		}
	}

	return false
}

func (r Report) removeLevel(i int) Report {
	return slices.Concat(r[:i], r[i+1:])
}

func (v Input) Part1() int {
	var count int

	for _, report := range v {
		if report.IsSafe() {
			count++
		}
	}

	return count
}

func (v Input) Part2() int {
	var count int

	for _, report := range v {
		if report.IsSafeProblemDampener() {
			count++
		}
	}

	return count
}
