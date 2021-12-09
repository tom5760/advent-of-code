package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
)

func main() {
	if err := run(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func run() error {
	floor, err := ParseInput(os.Stdin)
	if err != nil {
		return fmt.Errorf("failed to parse input: %w", err)
	}

	fmt.Println("Part 1:", Part1(floor))
	fmt.Println("Part 2:", Part2(floor))

	return nil
}

// If you can model how the smoke flows through the caves, you might be able to
// avoid it and be that much safer. The submarine generates a heightmap of the
// floor of the nearby caves for you (your puzzle input).
//
// Each number corresponds to the height of a particular location, where 9 is
// the highest and 0 is the lowest a location can be.
func ParseInput(r io.Reader) (*FloorMap, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanBytes)

	var (
		m        FloorMap
		widthSet bool
	)

	for scanner.Scan() {
		b := scanner.Bytes()[0]

		if b == '\n' {
			m.Height++
			widthSet = true

			continue
		}

		if b < 48 || b > 57 {
			panic("unexpected map byte")
		}

		if !widthSet {
			m.Width++
		}

		// Assume ASCII digit.
		m.Data = append(m.Data, b-48)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to scan input: %w", err)
	}

	return &m, nil
}

// Your first goal is to find the low points - the locations that are lower
// than any of its adjacent locations. Most locations have four adjacent
// locations (up, down, left, and right); locations on the edge or corner of
// the map have three or two adjacent locations, respectively. (Diagonal
// locations do not count as adjacent.)
//
// The risk level of a low point is 1 plus its height.
//
// Find all of the low points on your heightmap. What is the sum of the risk
// levels of all low points on your heightmap?
func Part1(m *FloorMap) uint64 {
	points := m.LowPoints()

	var risk uint64

	for _, i := range points {
		risk += (uint64(m.Data[i]) + 1)
	}

	return risk
}

// Next, you need to find the largest basins so you know what areas are most
// important to avoid.
//
// A basin is all locations that eventually flow downward to a single low
// point. Therefore, every low point has a basin, although some basins are very
// small. Locations of height 9 do not count as being in any basin, and all
// other locations will always be part of exactly one basin.
//
// The size of a basin is the number of locations within the basin, including
// the low point.
//
// What do you get if you multiply together the sizes of the three largest
// basins?
func Part2(m *FloorMap) uint64 {
	lows := m.LowPoints()
	basins := make([]uint64, len(lows))

	for i, j := range lows {
		basins[i] = m.FindBasin(j)
	}

	sort.Slice(basins, func(i, j int) bool { return basins[i] > basins[j] })

	return basins[0] * basins[1] * basins[2]
}
