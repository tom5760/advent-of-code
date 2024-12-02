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

// Each column is split by three spaces.
var delimiter = []byte{' ', ' ', ' '}

func Parse(r io.Reader) (*Input, error) {
	var input Input

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		parts := bytes.Split(scanner.Bytes(), delimiter)
		if len(parts) != 2 {
			return nil, errors.New("unexpected line format")
		}

		left, err := strconv.Atoi(string(parts[0]))
		if err != nil {
			return nil, fmt.Errorf("failed to parse left column: %w", err)
		}

		right, err := strconv.Atoi(string(parts[1]))
		if err != nil {
			return nil, fmt.Errorf("failed to parse right column: %w", err)
		}

		input.Left = append(input.Left, left)
		input.Right = append(input.Right, right)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to scan input: %w", err)
	}

	if len(input.Left) != len(input.Right) {
		return nil, errors.New("columns are not equal length")
	}

	return &input, nil
}

func (v *Input) Part1() int {
	slices.Sort(v.Left)
	slices.Sort(v.Right)

	var distance int

	for i := range v.Left {
		left := v.Left[i]
		right := v.Right[i]

		distance += int(math.Abs(float64(left - right)))
	}

	return distance
}

func (v *Input) Part2() int {
	var similarity int

	for _, left := range v.Left {
		var count int
		for _, right := range v.Right {
			if left == right {
				count++
			}
		}

		similarity += left * count
	}

	return similarity
}
