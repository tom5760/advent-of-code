package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
)

type Point struct {
	X, Y int64
}

type Entity struct {
	Point      Point
	velX, velY int64
}

func (p *Entity) Tick() {
	p.Point.X += p.velX
	p.Point.Y += p.velY
}

func (p *Entity) Untick() {
	p.Point.X -= p.velX
	p.Point.Y -= p.velY
}

type Space []Entity

func (s Space) Tick() {
	for i := range s {
		(&s[i]).Tick()
	}
}

func (s Space) Untick() {
	for i := range s {
		(&s[i]).Untick()
	}
}

func (s Space) Draw() {
	grid := make(map[Point]bool)
	for _, p := range s {
		grid[p.Point] = true
	}

	minX, minY, maxX, maxY := s.Bounds()

	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if _, ok := grid[Point{x, y}]; ok {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
}

func (s Space) Bounds() (minX, minY, maxX, maxY int64) {
	minX = math.MaxInt64
	minY = math.MaxInt64
	maxX = math.MinInt64
	maxY = math.MinInt64

	for _, p := range s {
		if p.Point.X < minX {
			minX = p.Point.X
		}
		if p.Point.X > maxX {
			maxX = p.Point.X
		}
		if p.Point.Y < minY {
			minY = p.Point.Y
		}
		if p.Point.Y > maxY {
			maxY = p.Point.Y
		}
	}

	return minX, minY, maxX, maxY
}

func main() {
	space := readInput(os.Stdin)

	minBounds := int64(math.MaxInt64)

	var tick int

	for {
		space.Tick()
		tick++

		minX, minY, maxX, maxY := space.Bounds()
		bounds := (maxX - minX) * (maxY - minY)

		if bounds < minBounds {
			minBounds = bounds
		}

		if bounds > minBounds {
			log.Println("seconds to wait:", tick-1)
			space.Untick()
			space.Draw()
			return
		}
	}
}

const inputFormat = "position=<%d, %d> velocity=<%d, %d>"

func readInput(r io.Reader) Space {
	var entities []Entity

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		var p Entity

		if _, err := fmt.Sscanf(scanner.Text(), inputFormat, &p.Point.X, &p.Point.Y, &p.velX, &p.velY); err != nil {
			log.Fatalln("failed to parse input:", err)
			return nil
		}

		entities = append(entities, p)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln("failed to read input:", err)
		return nil
	}

	return Space(entities)
}
