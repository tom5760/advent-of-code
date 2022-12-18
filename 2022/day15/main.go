// Package day15 implements the solution for Advent of Code 2022 day 15.
//
// See: https://adventofcode.com/2022/day/15
package day15

import (
	"bufio"
	"bytes"
	"fmt"
	"regexp"
	"strconv"

	"github.com/tom5760/advent-of-code/2022/inpututils"
	"github.com/tom5760/advent-of-code/2022/structs"
)

var inputRegexp = regexp.MustCompile(`Sensor at x=(-?\d+), y=(-?\d+): closest beacon is at x=(-?\d+), y=(-?\d+)`)

func Parse(name string) (Sensors, error) {
	var sensors Sensors

	err := inpututils.Scan(name, func(scanner *bufio.Scanner) error {
		// First line is target
		scanner.Scan()
		targetY, err := strconv.Atoi(string(bytes.TrimPrefix(scanner.Bytes(), []byte{'y', '='})))
		if err != nil {
			return fmt.Errorf("failed to find target Y: %w", err)
		}

		sensors.TargetY = targetY

		for scanner.Scan() {
			matches := inputRegexp.FindSubmatch(scanner.Bytes())
			if len(matches) != 5 {
				return fmt.Errorf("unexpected sensor format")
			}

			x1, err := strconv.Atoi(string(matches[1]))
			if err != nil {
				return fmt.Errorf("failed to parse x1: %w", err)
			}
			y1, err := strconv.Atoi(string(matches[2]))
			if err != nil {
				return fmt.Errorf("failed to parse y1: %w", err)
			}

			x2, err := strconv.Atoi(string(matches[3]))
			if err != nil {
				return fmt.Errorf("failed to parse x2: %w", err)
			}
			y2, err := strconv.Atoi(string(matches[4]))
			if err != nil {
				return fmt.Errorf("failed to parse y2: %w", err)
			}

			sensors.Sensors = append(sensors.Sensors, Sensor{
				Self:   structs.Coordinate{X: x1, Y: y1},
				Beacon: structs.Coordinate{X: x2, Y: y2},
			})
		}

		return nil
	})

	return sensors, err
}

func Part1(sensors Sensors) int {
	var cave structs.SparseGrid[Tile]
	cave.Init()

	mark := func(c structs.Coordinate) {
		if _, ok := cave.GetC(c); !ok {
			cave.SetC(c, TileExcluded)
		}
	}

	for _, sensor := range sensors.Sensors {
		cave.SetC(sensor.Self, TileSensor)
		cave.SetC(sensor.Beacon, TileBeacon)

		distance := structs.Distance(sensor.Self, sensor.Beacon)

		// Exclude all cells in a diamond around the sensor.
		for d := distance; d > 0; d-- {
			// Top right
			for i := 0; i < d; i++ {
				from := structs.Coordinate{X: sensor.Self.X + i, Y: sensor.Self.Y}
				to := structs.Coordinate{X: sensor.Self.X + i, Y: sensor.Self.Y - d + i}
				structs.ForLine(from, to, mark)
			}

			// Bottom right
			for i := 0; i < d; i++ {
				from := structs.Coordinate{X: sensor.Self.X + i, Y: sensor.Self.Y}
				to := structs.Coordinate{X: sensor.Self.X + i, Y: sensor.Self.Y + d - i}
				structs.ForLine(from, to, mark)
			}

			// Top left
			for i := 0; i < d; i++ {
				from := structs.Coordinate{X: sensor.Self.X - i, Y: sensor.Self.Y}
				to := structs.Coordinate{X: sensor.Self.X - i, Y: sensor.Self.Y - d + i}
				structs.ForLine(from, to, mark)
			}

			// Bottom left
			for i := 0; i < d; i++ {
				from := structs.Coordinate{X: sensor.Self.X - i, Y: sensor.Self.Y}
				to := structs.Coordinate{X: sensor.Self.X - i, Y: sensor.Self.Y + d - i}
				structs.ForLine(from, to, mark)
			}
		}
	}

	var total int
	for x := cave.MinX; x <= cave.MaxX; x++ {
		if v, ok := cave.Get(x, sensors.TargetY); ok && v == TileExcluded {
			total++
		}
	}

	return total
}

func Part2(sensors Sensors) int {
	return 0
}

type (
	Sensors struct {
		Sensors []Sensor
		TargetY int
	}

	Sensor struct {
		Self   structs.Coordinate
		Beacon structs.Coordinate
	}

	Tile byte
)

const (
	TileSensor   Tile = 'S'
	TileBeacon   Tile = 'B'
	TileExcluded Tile = '#'
)

func (t Tile) String() string {
	return string(t)
}
