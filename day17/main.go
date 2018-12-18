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

type Vein struct {
	Axis1, Axis2         string
	Value1               int
	RangeStart, RangeEnd int
}

type Tile byte

const (
	SpringTile    Tile = '+'
	SandTile      Tile = '.'
	ClayTile      Tile = '#'
	WaterTile     Tile = '~'
	ReachableTile Tile = '|'
)

type Point struct {
	X, Y int
}

func (p Point) Up() Point {
	return Point{p.X, p.Y - 1}
}

func (p Point) Down() Point {
	return Point{p.X, p.Y + 1}
}

func (p Point) Left() Point {
	return Point{p.X - 1, p.Y}
}

func (p Point) Right() Point {
	return Point{p.X + 1, p.Y}
}

type World struct {
	MinX, MinY, MaxX, MaxY int
	Width, Height          int
	Tiles                  [][]Tile
}

func (w *World) CountReachable() int {
	var count int
	for y := w.MinY; y <= w.MaxY; y++ {
		for x := w.MinX - 1; x <= w.MaxX+1; x++ {
			tile := w.Tile(x, y)
			if tile != nil && (*tile == ReachableTile || *tile == WaterTile) {
				count++
			}
		}
	}

	return count
}

func (w *World) CountStable() int {
	var count int
	for y := w.MinY; y <= w.MaxY; y++ {
		for x := w.MinX - 1; x <= w.MaxX+1; x++ {
			tile := w.Tile(x, y)
			if tile != nil && *tile == WaterTile {
				count++
			}
		}
	}

	return count
}

func (w *World) Tile(x, y int) *Tile {
	// Readjust to 0, 0
	x = x - w.MinX + 1

	if x < 0 || x >= w.Width || y < 0 || y >= w.Height {
		return nil
	}

	return &w.Tiles[y][x]
}

func (w *World) Point(p Point) *Tile {
	return w.Tile(p.X, p.Y)
}

func (w *World) TileIs(p Point, tile Tile) bool {
	t := w.Point(p)
	return t != nil && *t == tile
}

func (w *World) String() string {
	var sb strings.Builder

	sb.WriteByte('\n')

	// Figure out how many rows for x-axis labels
	xLen := MaxInt(DigitCountInt(w.MinX), DigitCountInt(w.MaxX))

	// Figure out how many columns for y-axis labels
	yLen := MaxInt(DigitCountInt(w.MinY), DigitCountInt(w.MaxY))

	// Draw x-axis labels
	for i := 0; i < xLen; i++ {
		// Add spacing for y-axis
		for j := 0; j < yLen+1; j++ {
			sb.WriteByte(' ')
		}

		for x := w.MinX - 1; x <= w.MaxX+1; x++ {
			xStr := strconv.Itoa(x)
			sb.WriteByte(xStr[i])
		}

		sb.WriteByte('\n')
	}

	for y := 0; y <= w.MaxY; y++ {
		// Draw y-axis label
		yStr := strconv.Itoa(y)
		for i := yLen - len(yStr); i > 0; i-- {
			sb.WriteByte(' ')
		}
		sb.WriteString(yStr)
		sb.WriteByte(' ')

		for x := w.MinX - 1; x <= w.MaxX+1; x++ {
			sb.WriteString(string(*w.Tile(x, y)))
		}

		sb.WriteByte('\n')
	}

	return sb.String()
}

func worldSize(veins []Vein) (minX, minY, maxX, maxY int) {
	minX = math.MaxInt64
	minY = math.MaxInt64
	maxX = math.MinInt64
	maxY = math.MinInt64

	for _, vein := range veins {
		var minAxis1, maxAxis1, minAxis2, maxAxis2 *int

		if vein.Axis1 == "x" {
			minAxis1 = &minX
			maxAxis1 = &maxX
			minAxis2 = &minY
			maxAxis2 = &maxY
		} else {
			minAxis1 = &minY
			maxAxis1 = &maxY
			minAxis2 = &minX
			maxAxis2 = &maxX
		}

		if vein.Value1 < *minAxis1 {
			*minAxis1 = vein.Value1
		}
		if vein.Value1 > *maxAxis1 {
			*maxAxis1 = vein.Value1
		}
		if vein.RangeStart < *minAxis2 {
			*minAxis2 = vein.RangeStart
		}
		if vein.RangeEnd > *maxAxis2 {
			*maxAxis2 = vein.RangeEnd
		}
	}

	return minX, minY, maxX, maxY
}

func makeTileMatrix(width, height int) [][]Tile {
	tileStorage := make([]Tile, width*height)
	tiles := make([][]Tile, height)

	for y := 0; y < height; y++ {
		x := y * width
		tiles[y] = tileStorage[x : x+width]
	}

	return tiles
}

func buildWorld(veins []Vein) *World {
	var world World

	world.MinX, world.MinY, world.MaxX, world.MaxY = worldSize(veins)

	// Add three to include MaxX, and one space to either side for water to flow.
	world.Width = (world.MaxX - world.MinX) + 3
	// Add one to include MaxY, and everything up to y=0
	world.Height = world.MaxY + 1

	world.Tiles = makeTileMatrix(world.Width, world.Height)

	for y := 0; y < world.Height; y++ {
		for x := 0; x < world.Width; x++ {
			world.Tiles[y][x] = SandTile
		}
	}

	for _, vein := range veins {
		var x, y int
		var axis *int

		if vein.Axis1 == "x" {
			x = vein.Value1
			axis = &y
		} else {
			y = vein.Value1
			axis = &x
		}

		for *axis = vein.RangeStart; *axis <= vein.RangeEnd; *axis++ {
			*world.Tile(x, y) = ClayTile
		}
	}

	*world.Tile(500, 0) = SpringTile

	return &world
}

func readInput(r io.Reader) []Vein {
	scanner := bufio.NewScanner(r)

	var veins []Vein

	for scanner.Scan() {
		matches := veinRegex.FindStringSubmatch(scanner.Text())
		if len(matches) != 6 {
			log.Fatalln("failed to parse input")
			return nil
		}

		vein := Vein{
			Axis1: matches[1],
			Axis2: matches[3],
		}

		var err error
		vein.Value1, err = strconv.Atoi(matches[2])
		if err != nil {
			log.Fatalln("failed to parse value:", err)
			return nil
		}
		vein.RangeStart, err = strconv.Atoi(matches[4])
		if err != nil {
			log.Fatalln("failed to parse range start:", err)
			return nil
		}
		vein.RangeEnd, err = strconv.Atoi(matches[5])
		if err != nil {
			log.Fatalln("failed to parse range end:", err)
			return nil
		}

		veins = append(veins, vein)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln("failed to read input:", err)
		return nil
	}

	return veins
}

func main() {
	world := buildWorld(readInput(os.Stdin))
	// log.Println("world:", world)

	Flow(world)
	// log.Println("world:", world)

	log.Println("(part 1) water reachable tile count:", world.CountReachable())
	log.Println("(part 2) water stable tile count:", world.CountStable())
}
