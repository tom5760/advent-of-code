// Package day14 implements the solution for Advent of Code 2022 day 14.
//
// See: https://adventofcode.com/2022/day/14
package day14

import (
	"bufio"
	"bytes"
	"fmt"
	"strconv"

	"github.com/tom5760/advent-of-code/2022/inpututils"
	"github.com/tom5760/advent-of-code/2022/structs"
	"golang.org/x/exp/maps"
)

const (
	startX = 500
	startY = 0
)

func Parse(name string) (Cave, error) {
	var rocks []Rock

	err := inpututils.Scan(name, func(scanner *bufio.Scanner) error {
		for scanner.Scan() {
			coords := bytes.Split(scanner.Bytes(), []byte{' ', '-', '>', ' '})
			if len(coords) < 2 {
				return fmt.Errorf("unexpected rock length")
			}

			rock := make(Rock, len(coords))
			for i := range rock {
				xy := bytes.Split(coords[i], []byte{','})
				if len(xy) != 2 {
					return fmt.Errorf("unexpected coordinate length")
				}

				x, err := strconv.Atoi(string(xy[0]))
				if err != nil {
					return fmt.Errorf("failed to parse X coordinate: %w", err)
				}

				y, err := strconv.Atoi(string(xy[1]))
				if err != nil {
					return fmt.Errorf("failed to parse Y coordinate: %w", err)
				}

				rock[i] = structs.Coordinate{X: x, Y: y}
			}

			rocks = append(rocks, rock)
		}

		return nil
	})

	var cave Cave
	cave.Init()

	for _, rock := range rocks {
		for i, to := range rock[1:] {
			from := rock[i]
			structs.ForLine(from, to, func(c structs.Coordinate) {
				cave.Set(c.X, c.Y, TileRock)
			})
		}
	}

	cave.Set(startX, startY, TileStart)

	return cave, err
}

func Part1(cave Cave) int {
	return CountSand(cave, false)
}

func Part2(cave Cave) int {
	// +1 to count the sand that blocked the entrance.
	return CountSand(cave, true) + 1
}

func CountSand(incave Cave, hasFloor bool) int {
	cave := incave
	cave.Values = maps.Clone(cave.Values)

	fmt.Println("START:\n", cave.String())

	var total int
	for DropSand(cave, startX, startY, hasFloor) {
		total++
	}

	fmt.Println("END:\n", cave.String())

	return total
}

type (
	Tile byte

	Cave = structs.SparseGrid[Tile]

	Rock []structs.Coordinate
)

const (
	TileEmpty Tile = '.'
	TileRock  Tile = '#'
	TileSand  Tile = 'o'
	TileStart Tile = '+'
)

func (t Tile) String() string {
	return string(t)
}

func DropSand(cave Cave, startX, startY int, hasFloor bool) bool {
	x := startX
	y := startY

	nextX := x
	nextY := y

	for {
		x = nextX
		y = nextY

		if !hasFloor && y > cave.MaxY {
			// We have fallen into the void
			return false
		}

		// Nothing here, try to fall down.
		nextY++

		if hasFloor && nextY == cave.MaxY+2 {
			// Hit the floor, Settle where we are
			cave.Set(x, y, TileSand)
			return true
		}

		if v, ok := cave.Get(nextX, nextY); !ok || v == TileEmpty {
			// Can fall straight down
			continue
		}

		// Try down-left
		nextX = x - 1
		if v, ok := cave.Get(nextX, nextY); !ok || v == TileEmpty {
			// Can fall down-left
			continue
		}

		// Try down-right
		nextX = x + 1
		if v, ok := cave.Get(nextX, nextY); !ok || v == TileEmpty {
			// Can fall down-right
			continue
		}

		if x == startX && y == startY {
			// We blocked the start
			return false
		}

		// Settle where we are
		cave.Set(x, y, TileSand)
		return true
	}
}
