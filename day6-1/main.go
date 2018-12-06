package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
)

const alphabet = "01234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

type Coordinate struct {
	X, Y int
}

func main() {
	coords := readInput(os.Stdin)
	areas := countAreas(coords)

	// drawAreas(coords)

	var maxArea int

	for _, area := range areas {
		if area > maxArea {
			maxArea = area
		}
	}

	log.Println("largest area size:", maxArea)
}

func readInput(r io.Reader) []Coordinate {
	var coords []Coordinate

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		var c Coordinate
		if _, err := fmt.Sscanf(scanner.Text(), "%d, %d", &c.X, &c.Y); err != nil {
			log.Fatalln("failed to scan line:", err)
			return nil
		}

		coords = append(coords, c)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln("failed to read input:", err)
		return nil
	}

	return coords
}

func findBoundaries(coords []Coordinate) (minX, minY, maxX, maxY int) {
	minX = math.MaxInt64
	minY = math.MaxInt64

	for _, c := range coords {
		if c.X < minX {
			minX = c.X
		} else if c.X > maxX {
			maxX = c.X
		}

		if c.Y < minY {
			minY = c.Y
		} else if c.Y > maxY {
			maxY = c.Y
		}
	}

	return minX, minY, maxX, maxY
}

func drawAreas(coords []Coordinate) {
	minX, minY, maxX, maxY := findBoundaries(coords)

	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			i := closest(Coordinate{x, y}, coords)
			if i == -1 {
				fmt.Print(".")
				continue
			}

			c := coords[i]
			if c.X == x && c.Y == y {
				fmt.Print("*")
			} else {
				fmt.Print(string(alphabet[i]))
			}
		}
		fmt.Printf("\n")
	}
}

func countAreas(coords []Coordinate) []int {
	minX, minY, maxX, maxY := findBoundaries(coords)
	areas := make([]int, len(coords))

	markInfinite := func(x, y int) {
		i := closest(Coordinate{x, y}, coords)
		// This point is closest to more than one point, don't count it.
		if i == -1 {
			return
		}

		c := coords[i]
		if c.X == x && c.Y == y {
			return
		}

		areas[i] = -1
	}

	// Mark the borders of infinite areas
	for x := minX; x <= maxX; x++ {
		for _, y := range []int{minY, maxY} {
			markInfinite(x, y)
		}
	}
	for y := minY; y <= maxY; y++ {
		for _, x := range []int{minX, maxX} {
			markInfinite(x, y)
		}
	}

	for x := minX + 1; x < maxX; x++ {
		for y := minY + 1; y < maxY; y++ {
			closestI := closest(Coordinate{x, y}, coords)
			// This point is closest to more than one point, don't count it.
			if closestI == -1 {
				continue
			}

			// This is an infinite area, don't count it.
			if areas[closestI] == -1 {
				continue
			}

			areas[closestI]++
		}
	}

	return areas
}

func closest(source Coordinate, coords []Coordinate) int {
	var minI, minDist int
	minDist = math.MaxInt64

	for i, c := range coords {
		dist := manhattanDistance(source, c)
		if dist < minDist {
			minI = i
			minDist = dist
		} else if dist == minDist {
			// Use this to mark that the source is equidistant from multiple coordinates.
			minI = -1
		}
	}

	return minI
}

func manhattanDistance(a, b Coordinate) int {
	x := a.X - b.X
	if x < 0 {
		x *= -1
	}

	y := a.Y - b.Y
	if y < 0 {
		y *= -1
	}

	return x + y
}
