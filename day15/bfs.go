package main

import (
	"math"
	"sort"
)

// NextCell returns the next node in the shortest path from start to end.
func NextCell(arena *Arena, start, end *Cell) (*Cell, uint64) {
	paths := ShortestPaths(arena, start, end)

	if len(paths) == 0 {
		return nil, 0
	}

	type nextCell struct {
		next     *Cell
		distance uint64
	}

	var nextCells []nextCell
	for _, path := range paths {
		var next *Cell
		if len(path) == 0 {
			continue
		}
		if len(path) == 1 {
			next = path[0]
		} else {
			next = path[1]
		}

		nextCells = append(nextCells, nextCell{
			next:     next,
			distance: uint64(len(path)),
		})
	}

	if len(nextCells) == 0 {
		return nil, 0
	}

	sort.Slice(nextCells, func(i, j int) bool {
		return Less(nextCells[i].next.Point, nextCells[j].next.Point)
	})

	return nextCells[0].next, nextCells[0].distance
}

func ShortestPaths(arena *Arena, start, end *Cell) [][]*Cell {
	paths := BreadthFirstSearch(arena, start, end)

	var minDist uint64 = math.MaxUint64
	for _, path := range paths {
		l := uint64(len(path))
		if l < minDist {
			minDist = l
		}
	}

	var shortestPaths [][]*Cell
	for _, path := range paths {
		if uint64(len(path)) == minDist {
			shortestPaths = append(shortestPaths, path)
		}
	}

	return shortestPaths
}

func BreadthFirstSearch(arena *Arena, start, end *Cell) [][]*Cell {
	openSet := make(map[*Cell]bool)
	closedSet := make(map[*Cell]bool)
	routes := make(map[*Cell]*Cell)

	pop := func() *Cell {
		var c *Cell
		for cell := range openSet {
			c = cell
			break
		}
		delete(openSet, c)
		return c
	}

	var paths [][]*Cell

	openSet[start] = true
	routes[start] = nil

	for len(openSet) > 0 {
		current := pop()

		if current == end {
			paths = append(paths, constructPath(current, routes))
			continue
		}

		closedSet[current] = true

		for _, neighbor := range current.EmptyNeighbors(arena) {
			if closedSet[neighbor] {
				continue
			}

			if !openSet[neighbor] {
				openSet[neighbor] = true
				routes[neighbor] = current
			}
		}
	}

	return paths
}

func constructPath(current *Cell, routes map[*Cell]*Cell) []*Cell {
	backPath := []*Cell{current}

	for routes[current] != nil {
		current = routes[current]
		backPath = append(backPath, current)
	}

	var path []*Cell
	for i := len(backPath) - 1; i >= 0; i-- {
		path = append(path, backPath[i])
	}

	return path
}
