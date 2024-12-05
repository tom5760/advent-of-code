package grid

import (
	"errors"
	"fmt"
	"io"
	"iter"
	"strings"

	"github.com/tom5760/advent-of-code/aoc2024/aoc"
)

type Grid struct {
	b    []byte
	W, H int
}

func Parse(r io.Reader) (*Grid, error) {
	scanner := aoc.NewScanner(r)

	var g Grid

	for line := range scanner.Scan() {
		switch n := len(line); {
		case g.W == 0:
			g.W = n
		case g.W != n:
			return nil, errors.New("unexpected line length")
		}

		g.H++
		g.b = append(g.b, line...)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to scan input: %w", err)
	}

	return &g, nil
}

func (g *Grid) String() string {
	var sb strings.Builder

	for row := range g.Rows() {
		sb.Write(row.B)
		sb.WriteByte('\n')
	}

	return sb.String()
}

type Cell struct {
	X, Y int
	B    byte
}

func (g *Grid) Value(x, y int) byte {
	return g.b[g.W*y+x]
}

type Point struct {
	X, Y int
}

func (g *Grid) Values(points []Point) []byte {
	v := make([]byte, 0, len(points))

	for _, p := range points {
		if p.X < 0 || p.X >= g.W ||
			p.Y < 0 || p.Y >= g.H {
			continue
		}

		v = append(v, g.Value(p.X, p.Y))
	}

	return v
}

func (g *Grid) Cells() iter.Seq[Cell] {
	return func(yield func(Cell) bool) {
		var c Cell

		for c.Y = range g.H {
			for c.X = range g.W {
				c.B = g.Value(c.X, c.Y)
				if !yield(c) {
					return
				}
			}
		}
	}
}

type Row struct {
	Y int
	B []byte
}

func (g *Grid) Rows() iter.Seq[Row] {
	return func(yield func(Row) bool) {
		var r Row

		for r.Y = range g.H {
			n := g.W * r.Y
			r.B = g.b[n : n+g.W]
			if !yield(r) {
				return
			}
		}
	}
}
