package day06

import (
	"errors"
	"fmt"
	"io"

	"github.com/tom5760/advent-of-code/aoc2024/grid"
)

const (
	tileFloor byte = '.'
	tileWall  byte = '#'
	tileGuard byte = '^'
)

func Run(r io.Reader) (int, int, error) {
	g, err := grid.Parse(r)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to parse input: %w", err)
	}

	start, err := findStart(g)
	if err != nil {
		return 0, 0, err
	}

	part1, _ := Simulate(g, start)

	var part2 int
	for point, b := range g.Cells() {
		if b != tileFloor {
			continue
		}

		g.Set(point.X, point.Y, tileWall)
		if _, looped := Simulate(g, start); looped {
			part2++
		}
		g.Set(point.X, point.Y, tileFloor)
	}

	return part1, part2, nil
}

func findStart(g *grid.Grid) (*grid.Point, error) {
	for point, b := range g.Cells() {
		if b == tileGuard {
			return &point, nil
		}
	}

	return nil, errors.New("guard not found")
}

type Direction byte

const (
	Up Direction = 1 << iota
	Down
	Left
	Right
)

type Sim struct {
	grid *grid.Grid
	pos  *grid.Point
	dir  Direction

	track map[grid.Point]Direction
}

// Returns number of steps, and false if left the grid, or true if looped.
func Simulate(g *grid.Grid, start *grid.Point) (int, bool) {
	s := Sim{
		grid:  g,
		pos:   start,
		dir:   Up,
		track: make(map[grid.Point]Direction, g.W*g.H),
	}

	for {
		s.track[*s.pos] |= s.dir

		nextPos := s.nextPos()
		if nextPos == nil {
			// went off of the grid
			return len(s.track), false
		}

		if s.track[*nextPos]&s.dir != 0 {
			// loop
			return len(s.track), true
		}

		nextTile := s.grid.Get(nextPos.X, nextPos.Y)
		if nextTile == tileWall {
			s.rotate()
			continue
		}

		s.pos = nextPos
	}
}

func (s *Sim) nextPos() *grid.Point {
	p := *s.pos

	switch s.dir {
	case Up:
		p.Y--
	case Down:
		p.Y++
	case Left:
		p.X--
	case Right:
		p.X++
	default:
		panic("invalid direction")
	}

	if p.X < 0 || p.X >= s.grid.W ||
		p.Y < 0 || p.Y >= s.grid.H {
		return nil
	}

	return &p
}

func (s *Sim) rotate() {
	switch s.dir {
	case Up:
		s.dir = Right
	case Down:
		s.dir = Left
	case Left:
		s.dir = Up
	case Right:
		s.dir = Down
	default:
		panic("invalid direction")
	}
}
