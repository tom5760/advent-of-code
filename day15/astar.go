package main

import (
	"math"
)

func AStar(arena *Arena, start, goal *Cell) (uint64, *Cell) {
	closedSet := make(map[*Cell]bool)
	openSet := make(map[*Cell]bool)

	openSet[start] = true

	cameFrom := make(map[*Cell]*Cell)

	gScore := make(map[*Cell]float64)
	gScore[start] = 0

	fScore := make(map[*Cell]float64)
	fScore[start] = heuristic(start, goal)

	getGScore := func(c *Cell) float64 {
		if score, ok := gScore[c]; ok {
			return score
		}
		return math.Inf(0)
	}

	getFScore := func(c *Cell) float64 {
		if score, ok := fScore[c]; ok {
			return score
		}
		return math.Inf(0)
	}

	for len(openSet) > 0 {
		var current *Cell

		minFScore := math.Inf(0)
		for cell := range openSet {
			f := getFScore(cell)
			if f < minFScore || (f == minFScore && Less(cell.Point, current.Point)) {
				current = cell
				minFScore = f
			}
		}

		if current == goal {
			return reconstructPath(cameFrom, current)
		}

		delete(openSet, current)
		closedSet[current] = true

		for _, neighbor := range current.EmptyNeighbors(arena) {
			if closedSet[neighbor] {
				continue
			}

			tentativeGScore := getGScore(current) + 1

			if !openSet[neighbor] {
				openSet[neighbor] = true
			} else if tentativeGScore > getGScore(neighbor) {
				continue
			} else if tentativeGScore == getGScore(neighbor) &&
				Less(neighbor.Point, current.Point) {
				continue
			}

			cameFrom[neighbor] = current
			gScore[neighbor] = tentativeGScore
			fScore[neighbor] = tentativeGScore + heuristic(neighbor, goal)
		}
	}

	return 0, nil
}

func reconstructPath(cameFrom map[*Cell]*Cell, current *Cell) (uint64, *Cell) {
	var distance uint64
	lastCell := current

	for cameFrom[current] != nil {
		lastCell = current
		current = cameFrom[current]
		distance++
	}

	return distance, lastCell
}

func heuristic(a, b *Cell) float64 {
	return float64(manhattanDistance(a.Point, b.Point))
}

func manhattanDistance(a, b Point) int64 {
	return abs(int64(a.X)-int64(b.X)) + abs(int64(a.Y)-int64(b.Y))
}

func abs(x int64) int64 {
	if x < 0 {
		return x * -1
	}
	return x
}
