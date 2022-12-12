// Package day08 implements the solution for Advent of Code 2022 day 8.
//
// See: https://adventofcode.com/2022/day/8
package day08

import (
	"bufio"
	"fmt"
	"math"
	"strconv"

	"github.com/tom5760/advent-of-code/2022/inpututils"
	"github.com/tom5760/advent-of-code/2022/sliceutils"
	"github.com/tom5760/advent-of-code/2022/structs"
)

func Parse(name string) (structs.Grid[int], error) {
	var grid structs.Grid[int]

	err := inpututils.Scan(name, func(scanner *bufio.Scanner) error {
		scanner.Split(bufio.ScanBytes)

		widthKnown := false

		for scanner.Scan() {
			input := scanner.Bytes()
			if len(input) != 1 {
				return fmt.Errorf("failed to read byte?")
			}

			b := input[0]
			switch b {
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				if !widthKnown {
					grid.Width++
				}

			case '\n':
				widthKnown = true
				grid.Height++

				continue
			default:
				return fmt.Errorf("unexpected height %q", b)
			}

			height, err := strconv.Atoi(string(b))
			if err != nil {
				return fmt.Errorf("failed to parse height: %w", err)
			}

			grid.Values = append(grid.Values, height)
		}

		return nil
	})

	return grid, err
}

func Part1(grid structs.Grid[int]) int {
	visible := structs.Grid[bool]{
		Height: grid.Height,
		Width:  grid.Width,
		Values: make([]bool, grid.Height*grid.Width),
	}

	setVisible := func(x, y int) { visible.Set(x, y, true) }

	// Scan rows...
	for y := 0; y < grid.Height; y++ {
		WalkVisible(grid, 0, y, MoveRight, setVisible)
		WalkVisible(grid, grid.Width-1, y, MoveLeft, setVisible)
	}

	// Scan columns...
	for x := 0; x < grid.Width; x++ {
		WalkVisible(grid, x, 0, MoveDown, setVisible)
		WalkVisible(grid, x, grid.Height-1, MoveUp, setVisible)
	}

	return sliceutils.Count(visible.Values, func(isVisible bool) bool { return isVisible })
}

func Part2(grid structs.Grid[int]) int {
	scores := structs.Grid[int]{
		Height: grid.Height,
		Width:  grid.Width,
		Values: make([]int, grid.Height*grid.Width),
	}

	for y := 0; y < grid.Height; y++ {
		for x := 0; x < grid.Width; x++ {
			up := CountDistance(grid, x, y, MoveUp)
			down := CountDistance(grid, x, y, MoveDown)
			left := CountDistance(grid, x, y, MoveLeft)
			right := CountDistance(grid, x, y, MoveRight)

			scores.Set(x, y, up*down*left*right)
		}
	}

	return sliceutils.Max(scores.Values)
}

func Move(g *structs.Grid[int], x, y int, moveFn MoveFunc, fn func(int, int, int) bool) {
	fn(x, y, g.Get(x, y))

	for moveFn(g, &x, &y) {
		if !fn(x, y, g.Get(x, y)) {
			return
		}
	}
}

type MoveFunc func(grid *structs.Grid[int], x, y *int) bool

func WalkVisible(grid structs.Grid[int], startX, startY int, moveFn MoveFunc, fn func(int, int)) {
	maxHeight := math.MinInt

	Move(&grid, startX, startY, moveFn, func(x, y, height int) bool {
		if height > maxHeight {
			fn(x, y)
			maxHeight = height
		}

		return true
	})
}

func CountDistance(grid structs.Grid[int], startX, startY int, moveFn MoveFunc) int {
	var total int
	maxHeight := grid.Get(startX, startY)

	Move(&grid, startX, startY, moveFn, func(x, y, height int) bool {
		total++
		return height < maxHeight
	})

	// minus one to subtract counting the start point.
	return total - 1
}

func MoveLeft(grid *structs.Grid[int], x, y *int) bool {
	if *x > 0 {
		*x--
		return true
	}

	return false
}

func MoveRight(grid *structs.Grid[int], x, y *int) bool {
	if *x < grid.Width-1 {
		*x++
		return true
	}

	return false
}

func MoveUp(grid *structs.Grid[int], x, y *int) bool {
	if *y > 0 {
		*y--
		return true
	}

	return false
}

func MoveDown(grid *structs.Grid[int], x, y *int) bool {
	if *y < grid.Height-1 {
		*y++
		return true
	}

	return false
}
