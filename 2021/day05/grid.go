package main

import (
	"fmt"
	"strings"
	"text/tabwriter"
)

type (
	Point struct {
		X, Y int64
	}

	Line [2]Point

	Grid struct {
		Width  int64
		Height int64

		Data []int64
	}
)

func (g *Grid) String() string {
	var sb strings.Builder

	w := tabwriter.NewWriter(&sb, 0, 0, 1, ' ', 0)

	for y := int64(0); y < g.Height; y++ {
		for x := int64(0); x < g.Width; x++ {
			fmt.Fprintf(w, "%v\t", g.Data[y*g.Width+x])
		}

		w.Write([]byte{'\n'})
	}

	w.Flush()

	return sb.String()
}

func (g *Grid) DrawOrthogonal(l Line) {
	// Assume lines are orthogonal for now.
	switch {
	case l[0].X == l[1].X:
		// Vertical
		min, max := minmax(l[0].Y, l[1].Y)
		for y := min; y <= max; y++ {
			g.Mark(l[0].X, y)
		}

	case l[0].Y == l[1].Y:
		// Horizontal
		min, max := minmax(l[0].X, l[1].X)
		for x := min; x <= max; x++ {
			g.Mark(x, l[0].Y)
		}

	default:
		// skip non-orthogonal lines
	}
}

func (g *Grid) Draw(l Line) {
	var xdiff, ydiff int64
	x1, y1 := l[0].X, l[0].Y
	x2, y2 := l[1].X, l[1].Y

	if x1 > x2 {
		xdiff = -1
	} else if x1 < x2 {
		xdiff = 1
	}

	if y1 > y2 {
		ydiff = -1
	} else if y1 < y2 {
		ydiff = 1
	}

	x, y := x1, y1

	for x != x2 || y != y2 {
		g.Mark(x, y)
		x += xdiff
		y += ydiff
	}

	g.Mark(x, y)
}

func (g *Grid) Mark(x, y int64) {
	g.Data[y*g.Width+x]++
}

func minmax(x, y int64) (min, max int64) {
	if y < x {
		return y, x
	}
	return x, y
}
