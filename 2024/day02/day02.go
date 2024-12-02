package day02

import (
	"errors"
	"io"
	"slices"

	"github.com/tom5760/advent-of-code/aoc2024/aoc"
)

type Report []int

func Run(r io.Reader) (int, int, error) {
	var part1, part2 int

	for fields, err := range aoc.IterRowsInt(r) {
		if err != nil {
			return part1, part2, err
		}
		if len(fields) == 0 {
			return part1, part2, errors.New("empty report")
		}

		report := Report(fields)

		if report.IsSafe() {
			part1++
			part2++
		} else if report.IsSafeProblemDampener() {
			part2++
		}
	}

	return part1, part2, nil
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
