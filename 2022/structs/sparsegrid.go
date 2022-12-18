package structs

import (
	"fmt"
	"math"
	"strings"
)

type SparseGrid[T any] struct {
	MinX, MaxX int
	MinY, MaxY int

	Values map[Coordinate]T
}

func (g *SparseGrid[T]) Init() {
	g.MinX = math.MaxInt
	g.MinY = math.MaxInt

	g.MaxX = math.MinInt
	g.MaxY = math.MinInt

	g.Values = make(map[Coordinate]T)
}

func (g *SparseGrid[T]) Get(x, y int) (T, bool) {
	return g.GetC(Coordinate{X: x, Y: y})
}

func (g *SparseGrid[T]) GetC(c Coordinate) (T, bool) {
	v, ok := g.Values[c]
	return v, ok
}

func (g *SparseGrid[T]) Set(x, y int, v T) {
	g.SetC(Coordinate{X: x, Y: y}, v)
}

func (g *SparseGrid[T]) SetC(c Coordinate, v T) {
	if c.X < g.MinX {
		g.MinX = c.X
	}
	if c.X > g.MaxX {
		g.MaxX = c.X
	}

	if c.Y < g.MinY {
		g.MinY = c.Y
	}
	if c.Y > g.MaxY {
		g.MaxY = c.Y
	}

	g.Values[c] = v
}

func (g *SparseGrid[T]) String() string {
	var sb strings.Builder
	const border = 2

	for y := g.MinY - border; y <= g.MaxY+border; y++ {
		for x := g.MinX - border; x <= g.MaxX+border; x++ {
			v, ok := g.Get(x, y)
			if ok {
				fmt.Fprintf(&sb, "%v", v)
			} else {
				sb.WriteByte(' ')
			}
		}
		sb.WriteByte('\n')
	}

	return sb.String()
}
