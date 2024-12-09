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

	for p, b := range g.Cells() {
		if b != 'X' {
			continue
		}

		// search all eight directions for the remaining 'MAS'

		// up
		check([]grid.Point{
			{X: p.X, Y: p.Y - 1},
			{X: p.X, Y: p.Y - 2},
			{X: p.X, Y: p.Y - 3},
		})

		// down
		check([]grid.Point{
			{X: p.X, Y: p.Y + 1},
			{X: p.X, Y: p.Y + 2},
			{X: p.X, Y: p.Y + 3},
		})

		// left
		check([]grid.Point{
			{X: p.X - 1, Y: p.Y},
			{X: p.X - 2, Y: p.Y},
			{X: p.X - 3, Y: p.Y},
		})

		// right
		check([]grid.Point{
			{X: p.X + 1, Y: p.Y},
			{X: p.X + 2, Y: p.Y},
			{X: p.X + 3, Y: p.Y},
		})

		// up left
		check([]grid.Point{
			{X: p.X - 1, Y: p.Y - 1},
			{X: p.X - 2, Y: p.Y - 2},
			{X: p.X - 3, Y: p.Y - 3},
		})

		// up right
		check([]grid.Point{
			{X: p.X + 1, Y: p.Y - 1},
			{X: p.X + 2, Y: p.Y - 2},
			{X: p.X + 3, Y: p.Y - 3},
		})

		// down right
		check([]grid.Point{
			{X: p.X + 1, Y: p.Y + 1},
			{X: p.X + 2, Y: p.Y + 2},
			{X: p.X + 3, Y: p.Y + 3},
		})

		// down left
		check([]grid.Point{
			{X: p.X - 1, Y: p.Y + 1},
			{X: p.X - 2, Y: p.Y + 2},
			{X: p.X - 3, Y: p.Y + 3},
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

	for p, b := range g.Cells() {
		if b != 'A' {
			continue
		}

		// search all four diagonals for the remaining 'M' and 'S'

		check(
			// up left - down right
			[]grid.Point{
				{X: p.X - 1, Y: p.Y - 1},
				{X: p.X, Y: p.Y},
				{X: p.X + 1, Y: p.Y + 1},
			},
			// up right - down left
			[]grid.Point{
				{X: p.X + 1, Y: p.Y - 1},
				{X: p.X, Y: p.Y},
				{X: p.X - 1, Y: p.Y + 1},
			},
		)

		check(
			// up left - down right
			[]grid.Point{
				{X: p.X - 1, Y: p.Y - 1},
				{X: p.X, Y: p.Y},
				{X: p.X + 1, Y: p.Y + 1},
			},
			// down left - up right
			[]grid.Point{
				{X: p.X - 1, Y: p.Y + 1},
				{X: p.X, Y: p.Y},
				{X: p.X + 1, Y: p.Y - 1},
			},
		)

		check(
			// down right - up left
			[]grid.Point{
				{X: p.X + 1, Y: p.Y + 1},
				{X: p.X, Y: p.Y},
				{X: p.X - 1, Y: p.Y - 1},
			},
			// up right - down left
			[]grid.Point{
				{X: p.X + 1, Y: p.Y - 1},
				{X: p.X, Y: p.Y},
				{X: p.X - 1, Y: p.Y + 1},
			},
		)

		check(
			// down right - up left
			[]grid.Point{
				{X: p.X + 1, Y: p.Y + 1},
				{X: p.X, Y: p.Y},
				{X: p.X - 1, Y: p.Y - 1},
			},
			// down left - up right
			[]grid.Point{
				{X: p.X - 1, Y: p.Y + 1},
				{X: p.X, Y: p.Y},
				{X: p.X + 1, Y: p.Y - 1},
			},
		)
	}

	return count
}
