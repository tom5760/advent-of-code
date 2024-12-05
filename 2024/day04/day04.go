package day04

import (
	"fmt"
	"io"

	"github.com/tom5760/advent-of-code/aoc2024/grid"
)

const needle = "MAS"

func Run(r io.Reader) (int, int, error) {
	g, err := grid.Parse(r)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to parse grid: %w", err)
	}

	return Part1(g), Part2(g), nil
}

func Part1(g *grid.Grid) int {
	var count int

	check := func(points []grid.Point) {
		v := string(g.Values(points))
		if v == needle {
			count++
		}
	}

	for c := range g.Cells() {
		if c.B != 'X' {
			continue
		}

		// search all eight directions for the remaining 'MAS'

		// up
		check([]grid.Point{
			{X: c.X, Y: c.Y - 1},
			{X: c.X, Y: c.Y - 2},
			{X: c.X, Y: c.Y - 3},
		})

		// down
		check([]grid.Point{
			{X: c.X, Y: c.Y + 1},
			{X: c.X, Y: c.Y + 2},
			{X: c.X, Y: c.Y + 3},
		})

		// left
		check([]grid.Point{
			{X: c.X - 1, Y: c.Y},
			{X: c.X - 2, Y: c.Y},
			{X: c.X - 3, Y: c.Y},
		})

		// right
		check([]grid.Point{
			{X: c.X + 1, Y: c.Y},
			{X: c.X + 2, Y: c.Y},
			{X: c.X + 3, Y: c.Y},
		})

		// up left
		check([]grid.Point{
			{X: c.X - 1, Y: c.Y - 1},
			{X: c.X - 2, Y: c.Y - 2},
			{X: c.X - 3, Y: c.Y - 3},
		})

		// up right
		check([]grid.Point{
			{X: c.X + 1, Y: c.Y - 1},
			{X: c.X + 2, Y: c.Y - 2},
			{X: c.X + 3, Y: c.Y - 3},
		})

		// down right
		check([]grid.Point{
			{X: c.X + 1, Y: c.Y + 1},
			{X: c.X + 2, Y: c.Y + 2},
			{X: c.X + 3, Y: c.Y + 3},
		})

		// down left
		check([]grid.Point{
			{X: c.X - 1, Y: c.Y + 1},
			{X: c.X - 2, Y: c.Y + 2},
			{X: c.X - 3, Y: c.Y + 3},
		})
	}

	return count
}

func Part2(g *grid.Grid) int {
	var count int

	check := func(a, b []grid.Point) {
		x := string(g.Values(a))
		y := string(g.Values(b))
		if x == needle && y == needle {
			count++
		}
	}

	for c := range g.Cells() {
		if c.B != 'A' {
			continue
		}

		// search all four diagonals for the remaining 'M' and 'S'

		check(
			// up left - down right
			[]grid.Point{
				{X: c.X - 1, Y: c.Y - 1},
				{X: c.X, Y: c.Y},
				{X: c.X + 1, Y: c.Y + 1},
			},
			// up right - down left
			[]grid.Point{
				{X: c.X + 1, Y: c.Y - 1},
				{X: c.X, Y: c.Y},
				{X: c.X - 1, Y: c.Y + 1},
			},
		)

		check(
			// up left - down right
			[]grid.Point{
				{X: c.X - 1, Y: c.Y - 1},
				{X: c.X, Y: c.Y},
				{X: c.X + 1, Y: c.Y + 1},
			},
			// down left - up right
			[]grid.Point{
				{X: c.X - 1, Y: c.Y + 1},
				{X: c.X, Y: c.Y},
				{X: c.X + 1, Y: c.Y - 1},
			},
		)

		check(
			// down right - up left
			[]grid.Point{
				{X: c.X + 1, Y: c.Y + 1},
				{X: c.X, Y: c.Y},
				{X: c.X - 1, Y: c.Y - 1},
			},
			// up right - down left
			[]grid.Point{
				{X: c.X + 1, Y: c.Y - 1},
				{X: c.X, Y: c.Y},
				{X: c.X - 1, Y: c.Y + 1},
			},
		)

		check(
			// down right - up left
			[]grid.Point{
				{X: c.X + 1, Y: c.Y + 1},
				{X: c.X, Y: c.Y},
				{X: c.X - 1, Y: c.Y - 1},
			},
			// down left - up right
			[]grid.Point{
				{X: c.X - 1, Y: c.Y + 1},
				{X: c.X, Y: c.Y},
				{X: c.X + 1, Y: c.Y - 1},
			},
		)
	}

	return count
}
