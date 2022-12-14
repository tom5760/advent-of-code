package structs

import (
	"bufio"
	"fmt"

	"github.com/tom5760/advent-of-code/2022/inpututils"
)

type (
	Coordinate struct {
		X, Y int
	}

	Grid[T any] struct {
		Height int
		Width  int
		Values []T
	}
)

func ScanGrid[T any](name string, parseFn func(byte) (T, error)) (Grid[T], error) {
	var grid Grid[T]

	err := inpututils.Scan(name, func(scanner *bufio.Scanner) error {
		scanner.Split(bufio.ScanBytes)

		var widthKnown bool

		for scanner.Scan() {
			input := scanner.Bytes()

			b := input[0]
			if b == '\n' {
				widthKnown = true
				grid.Height++

				continue
			}

			if !widthKnown {
				grid.Width++
			}

			v, err := parseFn(b)
			if err != nil {
				return fmt.Errorf("failed to parse grid cell: %w", err)
			}

			grid.Values = append(grid.Values, v)
		}

		return nil
	})

	return grid, err
}

func (g *Grid[T]) ToIndex(x, y int) int {
	return y*g.Width + x
}

func (g *Grid[T]) ToCoord(i int) Coordinate {
	return Coordinate{
		X: i % g.Width,
		Y: i / g.Width,
	}
}

func (g *Grid[T]) Get(x, y int) T {
	return g.Values[g.ToIndex(x, y)]
}

func (g *Grid[T]) Set(x, y int, v T) {
	g.Values[g.ToIndex(x, y)] = v
}

func (g *Grid[T]) Up(c Coordinate) (Coordinate, bool) {
	c.Y--
	return c, c.Y >= 0
}

func (g *Grid[T]) Down(c Coordinate) (Coordinate, bool) {
	c.Y++
	return c, c.Y < g.Height
}

func (g *Grid[T]) Left(c Coordinate) (Coordinate, bool) {
	c.X--
	return c, c.X >= 0
}

func (g *Grid[T]) Right(c Coordinate) (Coordinate, bool) {
	c.X++
	return c, c.X < g.Width
}
