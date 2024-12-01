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

// each column is split by three spaces
var columnDelimiter = []byte{' ', ' ', ' '}

func Parse(r io.Reader) (*Input, error) {
	var input Input

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		parts := bytes.Split(scanner.Bytes(), columnDelimiter)
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

func (d *Input) Part1() int {
	slices.Sort(d.Left)
	slices.Sort(d.Right)

	var distance int

	for i := range d.Left {
		left := d.Left[i]
		right := d.Right[i]

		distance += int(math.Abs(float64(left - right)))
	}

	return distance
}

func (d *Input) Part2() int {
	var similarity int

	for _, left := range d.Left {
		var count int
		for _, right := range d.Right {
			if left == right {
				count++
			}
		}

		similarity += left * count
	}

	return similarity
}
