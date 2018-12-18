package main

import (
	"bufio"
	"io"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var veinRegex = regexp.MustCompile(`([xy])=(\d+), ([xy])=(\d+)\.\.(\d+)`)

type Point struct {
	X, Y int64
}

type CellKind int

const (
	SandCell CellKind = iota
	ClayCell
	SpringCell
	WaterCell
)

type Cell struct {
	Kind      CellKind
	Point     Point
	Reachable bool
}

func TrickleIterative(w *World, start *Cell) bool {
	openSet := map[*Cell]bool{start: true}
	closedSet := make(map[*Cell]bool)
	terminalSet := make(map[*Cell]bool)

	pop := func() *Cell {
		for cell := range openSet {
			delete(openSet, cell)
			return cell
		}
		panic("openSet is empty")
	}

	var current *Cell
	for len(openSet) > 0 {
		current = pop()

		current.Reachable = true
		closedSet[current] = true

		if current.Point.Y > w.MaxY {
			continue
		}

		down := current.Down(w)
		if down.Kind == SandCell {
			if !closedSet[down] {
				openSet[down] = true
			}
			continue
		}

		var added bool

		left := current.Left(w)
		if !closedSet[left] && left.Kind == SandCell {
			added = true
			openSet[left] = true
		}

		right := current.Right(w)
		if !closedSet[right] && right.Kind == SandCell {
			added = true
			openSet[right] = true
		}

		if !added {
			terminalSet[current] = true
		}
	}

	for cell := range terminalSet {
		if SealedLeft(w, cell) && SealedRight(w, cell) {
			cell.Kind = WaterCell
			return true
		}
	}
	return false
}

func SealedLeft(w *World, c *Cell) bool {
	p := Point{Y: c.Point.Y}
	for p.X = c.Point.X; p.X >= w.MinX; p.X-- {
		cell := w.Get(p)
		if cell.Kind == ClayCell || cell.Kind == WaterCell {
			return true
		}
		if cell.Down(w).Kind == SandCell {
			return false
		}
	}
	return false
}

func SealedRight(w *World, c *Cell) bool {
	p := Point{Y: c.Point.Y}
	for p.X = c.Point.X; p.X <= w.MaxX; p.X++ {
		cell := w.Get(p)
		if cell.Kind == ClayCell || cell.Kind == WaterCell {
			return true
		}
		if cell.Down(w).Kind == SandCell {
			return false
		}
	}
	return false
}

func (c *Cell) Down(w *World) *Cell {
	return w.Get(Point{c.Point.X, c.Point.Y + 1})
}

func (c *Cell) Left(w *World) *Cell {
	return w.Get(Point{c.Point.X - 1, c.Point.Y})
}

func (c *Cell) Right(w *World) *Cell {
	return w.Get(Point{c.Point.X + 1, c.Point.Y})
}

func (c *Cell) String() string {
	switch c.Kind {
	case SandCell:
		if c.Reachable {
			return "|"
		}
		return "."
	case ClayCell:
		return "#"
	case SpringCell:
		return "+"
	case WaterCell:
		return "~"
	default:
		panic("unexpected cell kind")
	}
}

type World struct {
	MinX, MinY, MaxX, MaxY int64

	Ground map[Point]*Cell
}

func (w *World) Flow() {
	i := 1
	for w.Tick() {
		// log.Println("tick", i, w)
		i++
	}
}

func (w *World) Tick() bool {
	cell := w.Get(Point{500, 1})
	return TrickleIterative(w, cell)
}

func (w *World) CountReachable() uint64 {
	var count uint64
	var p Point

	for p.Y = w.MinY; p.Y < w.MaxY; p.Y++ {
		for p.X = w.MinX; p.X <= w.MaxX; p.X++ {
			cell := w.Get(p)
			if cell.Reachable || cell.Kind == WaterCell {
				count++
			}
		}
	}

	return count
}

func (w *World) Get(p Point) *Cell {
	c, ok := w.Ground[p]
	if !ok {
		return &Cell{
			Kind:  SandCell,
			Point: p,
		}
	}
	return c
}

func (w *World) String() string {
	var sb strings.Builder

	sb.WriteByte('\n')

	// Figure out how many rows for x-axis labels
	xLen := MaxInt(DigitCountInt64(w.MinX), DigitCountInt64(w.MaxX))

	// Figure out how many columns for y-axis labels
	yLen := MaxInt(DigitCountInt64(w.MinY), DigitCountInt64(w.MaxY))

	// Draw x-axis labels
	for i := 0; i < xLen; i++ {
		// Add spacing for y-axis
		for j := 0; j < yLen+1; j++ {
			sb.WriteByte(' ')
		}

		for x := w.MinX - 1; x <= w.MaxX+1; x++ {
			xStr := strconv.FormatInt(x, 10)
			// if len(xStr)-1 > i {
			// 	sb.WriteByte(' ')
			// } else {
			sb.WriteByte(xStr[i])
			// }
		}

		sb.WriteByte('\n')
	}

	var p Point
	for p.Y = w.MinY; p.Y <= w.MaxY; p.Y++ {
		// Draw y-axis label
		yStr := strconv.FormatInt(p.Y, 10)
		for i := yLen - len(yStr); i > 0; i-- {
			sb.WriteByte(' ')
		}
		sb.WriteString(yStr)
		sb.WriteByte(' ')

		for p.X = w.MinX - 1; p.X <= w.MaxX+1; p.X++ {
			cell := w.Get(p)
			sb.WriteString(cell.String())
		}

		sb.WriteByte('\n')
	}

	return sb.String()
}

func readInput(r io.Reader) *World {
	scanner := bufio.NewScanner(r)

	world := &World{
		MinX:   math.MaxInt64,
		MinY:   math.MaxInt64,
		MaxX:   math.MinInt64,
		MaxY:   math.MinInt64,
		Ground: make(map[Point]*Cell),
	}

	for scanner.Scan() {
		matches := veinRegex.FindStringSubmatch(scanner.Text())
		if len(matches) != 6 {
			log.Fatalln("failed to parse input")
			return nil
		}

		axis1 := matches[1]
		axis2 := matches[3]

		value1, err := strconv.ParseInt(matches[2], 10, 64)
		if err != nil {
			log.Fatalln("failed to parse value:", err)
			return nil
		}
		rangeStart, err := strconv.ParseInt(matches[4], 10, 64)
		if err != nil {
			log.Fatalln("failed to parse range start:", err)
			return nil
		}
		rangeEnd, err := strconv.ParseInt(matches[5], 10, 64)
		if err != nil {
			log.Fatalln("failed to parse range end:", err)
			return nil
		}

		var p Point
		var axis *int64

		switch {
		case axis1 == "x" && axis2 == "y":
			p.X = value1
			axis = &p.Y

		case axis1 == "y" && axis2 == "x":
			p.Y = value1
			axis = &p.X

		default:
			panic("invalid axis")
		}

		for *axis = rangeStart; *axis <= rangeEnd; *axis++ {
			world.Ground[p] = &Cell{
				Kind:  ClayCell,
				Point: p,
			}
		}

		if p.X < world.MinX {
			world.MinX = p.X
		}
		if p.Y < world.MinY {
			world.MinY = p.Y
		}
		if p.X > world.MaxX {
			world.MaxX = p.X
		}
		if p.Y > world.MaxY {
			world.MaxY = p.Y
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln("failed to read input:", err)
		return nil
	}

	if world.MinY > 0 {
		world.MinY = 0
	}

	for y := world.MinY; y <= world.MaxY; y++ {
		for x := world.MinX; x <= world.MaxX; x++ {
			p := Point{x, y}
			if _, ok := world.Ground[p]; ok {
				continue
			}
			world.Ground[p] = &Cell{
				Kind:  SandCell,
				Point: p,
			}
		}
	}

	world.Ground[Point{500, 0}] = &Cell{
		Kind:  SpringCell,
		Point: Point{500, 0},
	}

	return world
}

func main() {
	world := readInput(os.Stdin)
	// log.Println("world:", world)

	world.Flow()
	log.Println("world:", world)

	log.Println("(part 1) water reachable tile count:", world.CountReachable())
}
