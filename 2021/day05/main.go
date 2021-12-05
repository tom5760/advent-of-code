package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/tom5760/advent-of-code/2021/input"
)

func main() {
	if err := run(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func run() error {
	lines, err := ParseInput(os.Stdin)
	if err != nil {
		return fmt.Errorf("failed to parse input: %w", err)
	}

	fmt.Println("Part 1:", Part1(lines))
	fmt.Println("Part 2:", Part2(lines))

	return nil
}

func ParseInput(r io.Reader) ([]Line, error) {
	return input.Parser[Line]{ParseFunc: func(scanner *bufio.Scanner) (Line, error) {
		coords := bytes.Split(scanner.Bytes(), []byte{' ', '-', '>', ' '})
		if len(coords) != 2 {
			return Line{}, fmt.Errorf("unexpected line format: %v", coords)
		}

		var line Line

		for i := range line {
			var err error
			line[i], err = ParsePoint(coords[i])
			if err != nil {
				return line, fmt.Errorf("failed to parse point: %w", err)
			}
		}

		return line, nil
	}}.Slice(r)
}

func ParsePoint(b []byte) (Point, error) {
	parts := bytes.Split(b, []byte{','})
	if len(parts) != 2 {
		return Point{}, fmt.Errorf("unexpected point format: %v", parts)
	}

	var (
		p   Point
		err error
	)

	p.X, err = strconv.ParseInt(string(parts[0]), 10, 64)
	if err != nil {
		return p, fmt.Errorf("failed to parse X coordinate: %w", err)
	}

	p.Y, err = strconv.ParseInt(string(parts[1]), 10, 64)
	if err != nil {
		return p, fmt.Errorf("failed to parse Y coordinate: %w", err)
	}

	return p, nil
}

// To avoid the most dangerous areas, you need to determine the number of
// points where at least two lines overlap.
func Part1(lines []Line) int {
	var maxX, maxY int64

	for _, l := range lines {
		for _, p := range l {
			if maxX < p.X {
				maxX = p.X
			}

			if maxY < p.Y {
				maxY = p.Y
			}
		}
	}

	g := &Grid{
		Width:  maxX + 1,
		Height: maxY + 1,
	}

	g.Data = make([]int64, g.Width*g.Height)

	for _, line := range lines {
		g.DrawOrthogonal(line)
	}

	var count int

	for _, p := range g.Data {
		if p > 1 {
			count++
		}
	}

	return count
}

// Count diagonal lines
func Part2(lines []Line) int {
	var maxX, maxY int64

	for _, l := range lines {
		for _, p := range l {
			if maxX < p.X {
				maxX = p.X
			}

			if maxY < p.Y {
				maxY = p.Y
			}
		}
	}

	g := &Grid{
		Width:  maxX + 1,
		Height: maxY + 1,
	}

	g.Data = make([]int64, g.Width*g.Height)

	for _, line := range lines {
		g.Draw(line)
	}

	var count int

	for _, p := range g.Data {
		if p > 1 {
			count++
		}
	}

	return count
}
