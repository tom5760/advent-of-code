package main

import (
	"fmt"
	"log"
	"math"
)

const (
	SerialNumber = 1133

	GridSize   = 300
	SquareSize = 3
)

type Coordinate struct {
	X, Y int64
}

func (c Coordinate) String() string {
	// Origin is 1, 1 so add...
	return fmt.Sprintf("%d,%d", c.X+1, c.Y+1)
}

func main() {
	grid := powerGrid()
	summedTable := summedAreaTable(grid)

	coord, power := maxArea(summedTable, SquareSize)
	log.Printf("(part 1) fuel cell %v has max power of %d", coord, power)

	var maxCoord Coordinate
	var maxPower int64 = math.MinInt64
	var maxSize int64

	for i := int64(1); i <= 300; i++ {
		coord, power := maxArea(summedTable, i)
		// log.Printf("fuel cell %v,%d has max power of %d", coord, i, power)
		if power > maxPower {
			maxPower = power
			maxCoord = coord
			maxSize = i
		}
	}

	log.Printf("(part 2) best size %v,%d with power of %d",
		maxCoord, maxSize, maxPower)
}

func powerGrid() map[Coordinate]int64 {
	powers := make(map[Coordinate]int64, GridSize*GridSize)

	for x := int64(0); x < GridSize; x++ {
		for y := int64(0); y < GridSize; y++ {
			powers[Coordinate{x, y}] = cellPower(x, y)
		}
	}

	return powers
}

func cellPower(x, y int64) int64 {
	// Find the fuel cell's rack ID, which is its X coordinate plus 10.
	rackID := x + 10

	// Begin with a power level of the rack ID times the Y coordinate.
	powerLevel := rackID * y

	// Increase the power level by the value of the grid serial number (your puzzle input).
	powerLevel += SerialNumber

	// Set the power level to itself multiplied by the rack ID.
	powerLevel *= rackID

	// Keep only the hundreds digit of the power level (so 12345 becomes 3; numbers
	// with no hundreds digit become 0).
	powerLevel = (powerLevel / 100) % 10

	// Subtract 5 from the power level.
	powerLevel -= 5

	return powerLevel
}

// https://en.wikipedia.org/wiki/Summed-area_table
func summedAreaTable(grid map[Coordinate]int64) map[Coordinate]int64 {
	table := make(map[Coordinate]int64, GridSize*GridSize)

	// fill in top-left 2x2 area
	table[Coordinate{0, 0}] = grid[Coordinate{0, 0}]
	table[Coordinate{1, 0}] = grid[Coordinate{1, 0}] + table[Coordinate{0, 0}]
	table[Coordinate{0, 1}] = grid[Coordinate{0, 1}] + table[Coordinate{0, 0}]

	for x := int64(1); x < GridSize; x++ {
		for y := int64(1); y < GridSize; y++ {
			table[Coordinate{x, y}] = grid[Coordinate{x, y}] +
				table[Coordinate{x, y - 1}] +
				table[Coordinate{x - 1, y}] -
				table[Coordinate{x - 1, y - 1}]
		}
	}

	return table
}

func maxArea(summedArea map[Coordinate]int64, squareSize int64) (Coordinate, int64) {
	squares := make(map[Coordinate]int64)

	for x := int64(0); x < GridSize-squareSize; x++ {
		for y := int64(0); y < GridSize-squareSize; y++ {
			A := Coordinate{x, y}
			B := Coordinate{x, y + squareSize}
			C := Coordinate{x + squareSize, y}
			D := Coordinate{x + squareSize, y + squareSize}

			squares[Coordinate{x, y}] = summedArea[D] + summedArea[A] - summedArea[B] - summedArea[C]
		}
	}

	var maxPower int64 = math.MinInt64
	var maxCoord Coordinate
	for coord, power := range squares {
		if power > maxPower {
			maxPower = power
			maxCoord = coord
		}
	}

	return maxCoord, maxPower
}
