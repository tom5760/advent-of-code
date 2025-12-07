package grid

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"iter"
	"strings"
)

type Grid struct {
	b    []byte
	W, H int
}

func Parse(r io.Reader) (*Grid, error) {
	scanner := bufio.NewScanner(r)

	var g Grid

	for scanner.Scan() {
		line := scanner.Bytes()

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

var adjacent = []Point{
	{-1, -1}, // up-left
	{0, -1},  // up
	{1, -1},  // up-right
	{1, 0},   // right
	{1, 1},   // down-right
	{0, 1},   // down
	{-1, 1},  // down-left
	{-1, 0},  // left
}

// Iterates over the eight adjacent points.
func (g *Grid) Adjacent(p Point) iter.Seq2[Point, byte] {
	return func(yield func(Point, byte) bool) {
		for _, d := range adjacent {
			np := Point{
				X: p.X + d.X,
				Y: p.Y + d.Y,
			}

			if np.X >= 0 && np.X < g.W && np.Y >= 0 && np.Y < g.H {
				b := g.Get(np.X, np.Y)
				if !yield(np, b) {
					return
				}
			}
		}
	}
}

type Point struct {
	X, Y int
}
