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

	for _, row := range g.Rows() {
		sb.Write(row)
		sb.WriteByte('\n')
	}

	return sb.String()
}

func (g *Grid) addr(x, y int) int {
	return g.W*y + x
}

func (g *Grid) Get(x, y int) byte {
	return g.b[g.addr(x, y)]
}

func (g *Grid) Set(x, y int, b byte) {
	g.b[g.addr(x, y)] = b
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

		v = append(v, g.Get(p.X, p.Y))
	}

	return v
}

type Cell struct {
	Point
	B *byte
}

func (g *Grid) Cells() iter.Seq2[Point, byte] {
	return func(yield func(Point, byte) bool) {
		var p Point

		for p.Y = range g.H {
			for p.X = range g.W {
				b := g.Get(p.X, p.Y)
				if !yield(p, b) {
					return
				}
			}
		}
	}
}

func (g *Grid) Rows() iter.Seq2[int, []byte] {
	return func(yield func(int, []byte) bool) {
		for y := range g.H {
			i := g.addr(0, y)
			row := g.b[i : i+g.W]
			if !yield(y, row) {
				return
			}
		}
	}
}
