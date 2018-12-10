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
	X, Y       int64
	velX, velY int64
}

func (p *Point) Tick() {
	p.X += p.velX
	p.Y += p.velY
}

func (p *Point) Untick() {
	p.X -= p.velX
	p.Y -= p.velY
}

type Space struct {
	points []*Point
	grid   map[int64]map[int64]*Point
}

func (s *Space) Update() {
	s.grid = make(map[int64]map[int64]*Point)

	for _, p := range s.points {
		col := s.grid[p.X]
		if col == nil {
			col = make(map[int64]*Point)
			s.grid[p.X] = col
		}

		col[p.Y] = p
	}
}

func (s *Space) Tick() {
	for _, p := range s.points {
		p.Tick()
	}
}

func (s *Space) Untick() {
	for _, p := range s.points {
		p.Untick()
	}
}

func (s *Space) Draw() {
	minX, minY, maxX, maxY := s.Bounds()

	s.Update()

	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			col := s.grid[x]
			if _, ok := col[y]; ok {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
}

func (s *Space) Bounds() (minX, minY, maxX, maxY int64) {
	minX = math.MaxInt64
	minY = math.MaxInt64
	maxX = math.MinInt64
	maxY = math.MinInt64

	for _, p := range s.points {
		if p.X < minX {
			minX = p.X
		}
		if p.X > maxX {
			maxX = p.X
		}
		if p.Y < minY {
			minY = p.Y
		}
		if p.Y > maxY {
			maxY = p.Y
		}
	}

	return minX, minY, maxX, maxY
}

func main() {
	space := Space{points: readInput(os.Stdin)}

	minBounds := int64(math.MaxInt64)

	var tick int

	for {
		space.Tick()
		tick++

		minX, minY, maxX, maxY := space.Bounds()
		var max, min int64
		if maxX > maxY {
			max = maxX - maxY
		} else {
			max = maxY - maxX
		}

		if minX > minY {
			min = minX - minY
		} else {
			min = minY - minX
		}

		bounds := max * min

		if bounds < minBounds {
			minBounds = bounds
		}

		if bounds > minBounds {
			log.Println("need to wait:", tick-1)
			space.Untick()
			space.Draw()
			return
		}
	}
}

const inputFormat = "position=<%d, %d> velocity=<%d, %d>"

func readInput(r io.Reader) []*Point {
	var points []*Point

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		var p Point

		if _, err := fmt.Sscanf(scanner.Text(), inputFormat, &p.X, &p.Y, &p.velX, &p.velY); err != nil {
			log.Fatalln("failed to parse input:", err)
			return nil
		}

		points = append(points, &p)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln("failed to read input:", err)
		return nil
	}

	return points
}
