package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseAsteroids(t *testing.T) {
	tests := []struct {
		input    string
		expected *Field
	}{
		{
			input: ".#..#\n.....\n#####\n....#\n...##\n",
			expected: &Field{
				Width:  5,
				Height: 5,
				Asteroids: map[Point]Object{
					Point{1, 0}: objectAsteroid,
					Point{4, 0}: objectAsteroid,
					Point{0, 2}: objectAsteroid,
					Point{1, 2}: objectAsteroid,
					Point{2, 2}: objectAsteroid,
					Point{3, 2}: objectAsteroid,
					Point{4, 2}: objectAsteroid,
					Point{4, 3}: objectAsteroid,
					Point{3, 4}: objectAsteroid,
					Point{4, 4}: objectAsteroid,
				},
			},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			actual, err := ParseAsteroids(strings.NewReader(tt.input))
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestLineDistance(t *testing.T) {
	tests := []struct {
		p1, p2 Point
		d      float64
	}{
		{
			p1: Point{0, 0},
			p2: Point{1, 0},
			d:  1,
		},
		{
			p1: Point{1, 0},
			p2: Point{0, 0},
			d:  1,
		},
		{
			p1: Point{0, 0},
			p2: Point{0, 1},
			d:  1,
		},
		{
			p1: Point{0, 1},
			p2: Point{0, 0},
			d:  1,
		},
		{
			p1: Point{0, 0},
			p2: Point{2, 0},
			d:  2,
		},
		{
			p1: Point{2, 0},
			p2: Point{0, 0},
			d:  2,
		},
		{
			p1: Point{0, 0},
			p2: Point{0, 2},
			d:  2,
		},
		{
			p1: Point{0, 2},
			p2: Point{0, 0},
			d:  2,
		},
		{
			p1: Point{0, 0},
			p2: Point{3, 4},
			d:  5,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			assert.Equal(t, tt.d, lineDistance(tt.p1, tt.p2))
		})
	}
}

func TestFindMonitoringStation(t *testing.T) {
	tests := []struct {
		field string
		point Point
		count int
	}{
		{
			field: `.#..#
.....
#####
....#
...##
`,
			point: Point{3, 4},
			count: 8,
		},
		{
			field: `......#.#.
#..#.#....
..#######.
.#.#.###..
.#..#.....
..#....#.#
#..#....#.
.##.#..###
##...#..#.
.#....####
`,
			point: Point{5, 8},
			count: 33,
		},
		{
			field: `#.#...#.#.
.###....#.
.#....#...
##.#.#.#.#
....#.#.#.
.##..###.#
..#...##..
..##....##
......#...
.####.###.
`,
			point: Point{1, 2},
			count: 35,
		},
		{
			field: `.#..#..###
####.###.#
....###.#.
..###.##.#
##.##.#.#.
....###..#
..#.#..#.#
#..#.#.###
.##...##.#
.....#.#..
`,
			point: Point{6, 3},
			count: 41,
		},
		{
			field: `.#..##.###...#######
##.############..##.
.#.######.########.#
.###.#######.####.#.
#####.##.#.##.###.##
..#####..#.#########
####################
#.####....###.#.#.##
##.#################
#####.##.###..####..
..######..##.#######
####.##.####...##..#
.#####..#.######.###
##...#.##########...
#.##########.#######
.####.#.###.###.#.##
....##.##.###..#####
.#.#.###########.###
#.#.#.#####.####.###
###.##.####.##.#..##
`,
			point: Point{11, 13},
			count: 210,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			field, err := ParseAsteroids(strings.NewReader(tt.field))
			if !assert.NoError(t, err) {
				return
			}

			actualPoint, actualCount := FindMonitoringStation(field)
			assert.Equal(t, tt.point, actualPoint)
			assert.Equal(t, tt.count, actualCount)
		})
	}
}
