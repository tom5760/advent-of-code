package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
)

const SafeDistance = 10000

type Coordinate struct {
	X, Y int
}

func main() {
	coords := readInput(os.Stdin)

	minX, minY, maxX, maxY := findBoundaries(coords)
	var safeAreaSize int

	for x := minX; x < maxX; x++ {
		for y := minY; y < maxY; y++ {

			var totalDist int
			xy := Coordinate{x, y}

			for _, c := range coords {
				totalDist += manhattanDistance(xy, c)
			}
			if totalDist < SafeDistance {
				safeAreaSize++
			}
		}
	}

	log.Println("safe area size:", safeAreaSize)
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
