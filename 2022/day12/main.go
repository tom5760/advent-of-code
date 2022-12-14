// Package day12 implements the solution for Advent of Code 2022 day 12.
//
// See: https://adventofcode.com/2022/day/12
package day12

import (
	"container/list"
	"math"

	"github.com/tom5760/advent-of-code/2022/structs"
)

func Parse(name string) (Hills, error) {
	return structs.ScanGrid(name, func(b byte) (byte, error) {
		return b, nil
	})
}

func Part1(hills Hills) int {
	start := FindStart(hills)

	return ShortestPath(hills, start)
}

func Part2(hills Hills) int {
	min := math.MaxInt

	for i, v := range hills.Values {
		switch v {
		case 'S', 'a':
			path := ShortestPath(hills, hills.ToCoord(i))
			if path < min {
				min = path
			}
		}
	}

	return min
}

type Hills = structs.Grid[byte]

func FindStart(hills Hills) structs.Coordinate {
	for i, v := range hills.Values {
		if v == 'S' {
			return hills.ToCoord(i)
		}
	}

	panic("start not found")
}

func ShortestPath(hills Hills, start structs.Coordinate) int {
	visited := make(map[structs.Coordinate]bool, len(hills.Values))
	visited[start] = true

	parents := make(map[structs.Coordinate]structs.Coordinate, len(hills.Values))

	var queue list.List
	queue.PushBack(start)

	check := func(curHeight byte, cur, next structs.Coordinate) {
		if visited[next] {
			return
		}

		nextHeight := ToHeight(hills.Get(next.X, next.Y))
		if nextHeight == curHeight+1 || nextHeight <= curHeight {
			visited[next] = true
			parents[next] = cur
			queue.PushBack(next)
		}
	}

	for queue.Len() > 0 {
		coord := queue.Remove(queue.Front()).(structs.Coordinate)
		height := hills.Get(coord.X, coord.Y)

		if height == 'E' {
			var total int

			coord, ok := parents[coord]
			for ok {
				total++
				coord, ok = parents[coord]
			}

			return total
		}

		height = ToHeight(height)

		if up, ok := hills.Up(coord); ok {
			check(height, coord, up)
		}
		if down, ok := hills.Down(coord); ok {
			check(height, coord, down)
		}
		if left, ok := hills.Left(coord); ok {
			check(height, coord, left)
		}
		if right, ok := hills.Right(coord); ok {
			check(height, coord, right)
		}
	}

	return math.MaxInt
}

func ToHeight(val byte) byte {
	switch val {
	case 'S':
		val = 'a'
	case 'E':
		val = 'z'
	}
	return val - 97
}
