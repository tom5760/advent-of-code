package main

import (
	"math"
	"sort"
)

func AStar(arena *Arena, start, goal *Cell) (uint64, *Cell) {
	closedSet := make(map[*Cell]bool)
	openSet := make(map[*Cell]bool)

	openSet[start] = true

	cameFrom := make(map[*Cell][]*Cell)

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
			if f < minFScore {
				current = cell
				minFScore = f
			}
		}

		if current == goal {
			// log.Println("reconstructing path", start.Point, goal.Point)
			return uint64(getGScore(goal)), reconstructPath(cameFrom, start, current)
		}

		delete(openSet, current)
		closedSet[current] = true

		currentGScore := getGScore(current) + 1

		for _, neighbor := range current.EmptyNeighbors(arena) {
			if closedSet[neighbor] {
				continue
			}

			if !openSet[neighbor] {
				openSet[neighbor] = true
			}

			if currentGScore > getGScore(neighbor) {
				continue
			}

			if currentGScore < getGScore(neighbor) {
				// cameFrom[neighbor] = append(cameFrom[neighbor], current)
				cameFrom[neighbor] = []*Cell{current}
				gScore[neighbor] = currentGScore
				fScore[neighbor] = currentGScore + heuristic(neighbor, goal)

			} else {
				cameFrom[neighbor] = append(cameFrom[neighbor], current)
				gScore[neighbor] = currentGScore
				fScore[neighbor] = currentGScore + heuristic(neighbor, goal)
			}
		}
	}

	return 0, nil
}

func reconstructPath(cameFrom map[*Cell][]*Cell, start, current *Cell) *Cell {
	if start == current {
		return start
	}

	nextSet := make(map[*Cell]bool)
	recursePath(cameFrom, nextSet, start, current)

	var nextCells []*Cell
	for cell := range nextSet {
		if len(nextSet) > 1 {
			// log.Println("cell:", cell.Point)
		}
		nextCells = append(nextCells, cell)
	}

	sort.Slice(nextCells, func(i, j int) bool {
		return Less(nextCells[i].Point, nextCells[j].Point)
	})

	if len(nextSet) > 1 {
		// log.Println("next cell:", nextCells[0].Point)
	}

	return nextCells[0]

	// var distance uint64
	// lastCell := current

	// for cameFrom[current] != nil {
	// 	lastCell = current
	// 	current = cameFrom[current]
	// 	distance++
	// }

	// return distance, lastCell
}

func recursePath(cameFrom map[*Cell][]*Cell, nextSet map[*Cell]bool, start, current *Cell) {
	ancestors := cameFrom[current]

	if len(ancestors) == 0 {
		return
	}

	if ancestors[0] == start {
		nextSet[current] = true
		return
	}

	for _, ancestor := range ancestors {
		recursePath(cameFrom, nextSet, start, ancestor)
	}
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
