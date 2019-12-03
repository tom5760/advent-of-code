package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWireToLines(t *testing.T) {
	tests := []struct {
		wire  Wire
		lines Lines
	}{
		{
			wire: Wire{
				{Direction: DirectionRight, Length: 8},
				{Direction: DirectionUp, Length: 5},
				{Direction: DirectionLeft, Length: 5},
				{Direction: DirectionDown, Length: 3},
			},
			lines: Lines{
				{A: Point{0, 0}, B: Point{8, 0}},
				{A: Point{8, 0}, B: Point{8, 5}},
				{A: Point{8, 5}, B: Point{3, 5}},
				{A: Point{3, 5}, B: Point{3, 2}},
			},
		},
		{
			wire: Wire{
				{Direction: DirectionUp, Length: 7},
				{Direction: DirectionRight, Length: 6},
				{Direction: DirectionDown, Length: 4},
				{Direction: DirectionLeft, Length: 4},
			},
			lines: Lines{
				{A: Point{0, 0}, B: Point{0, 7}},
				{A: Point{0, 7}, B: Point{6, 7}},
				{A: Point{6, 7}, B: Point{6, 3}},
				{A: Point{6, 3}, B: Point{2, 3}},
			},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			actual := WireToLines(tt.wire)
			assert.Equal(t, tt.lines, actual)
		})
	}
}

func TestLineSegmentIntersection(t *testing.T) {
	tests := []struct {
		i, j LineSegment
		p    Point
		ok   bool
	}{
		{
			i:  LineSegment{A: Point{2, 3}, B: Point{6, 3}},
			j:  LineSegment{A: Point{3, 5}, B: Point{3, 2}},
			p:  Point{3, 3},
			ok: true,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			actualPoint, actualOK := LineSegmentIntersection(tt.i, tt.j)
			assert.Equal(t, tt.p, actualPoint)
			assert.Equal(t, tt.ok, actualOK)
		})
	}
}

func TestLinesIntersections(t *testing.T) {
	tests := []struct {
		a, b   Lines
		points []Point
	}{
		{
			a: Lines{
				{A: Point{0, 0}, B: Point{8, 0}},
				{A: Point{8, 0}, B: Point{8, 5}},
				{A: Point{8, 5}, B: Point{3, 5}},
				{A: Point{3, 5}, B: Point{3, 2}},
			},
			b: Lines{
				{A: Point{0, 0}, B: Point{0, 7}},
				{A: Point{0, 7}, B: Point{6, 7}},
				{A: Point{6, 7}, B: Point{6, 3}},
				{A: Point{6, 3}, B: Point{2, 3}},
			},
			points: []Point{
				{3, 3},
				{6, 5},
			},
		},
	}

	for i, tt := range tests {
		// Put points into a set so we can compare unordered.
		expectedPoints := map[Point]bool{}
		for _, p := range tt.points {
			expectedPoints[p] = true
		}

		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			actual := LinesIntersections(tt.a, tt.b)

			actualPoints := map[Point]bool{}
			for _, p := range actual {
				actualPoints[p] = true
			}

			assert.Equal(t, expectedPoints, actualPoints)
		})
	}
}

func TestStepsToPoint(t *testing.T) {
	tests := []struct {
		l Lines
		p Point
		d int
	}{
		{
			l: Lines{
				{A: Point{0, 0}, B: Point{8, 0}},
				{A: Point{8, 0}, B: Point{8, 5}},
				{A: Point{8, 5}, B: Point{3, 5}},
				{A: Point{3, 5}, B: Point{3, 2}},
			},
			p: Point{3, 3},
			d: 20,
		},
		{
			l: Lines{
				{A: Point{0, 0}, B: Point{8, 0}},
				{A: Point{8, 0}, B: Point{8, 5}},
				{A: Point{8, 5}, B: Point{3, 5}},
				{A: Point{3, 5}, B: Point{3, 2}},
			},
			p: Point{6, 5},
			d: 15,
		},
		{
			l: Lines{
				{A: Point{0, 0}, B: Point{0, 7}},
				{A: Point{0, 7}, B: Point{6, 7}},
				{A: Point{6, 7}, B: Point{6, 3}},
				{A: Point{6, 3}, B: Point{2, 3}},
			},
			p: Point{3, 3},
			d: 20,
		},
		{
			l: Lines{
				{A: Point{0, 0}, B: Point{0, 7}},
				{A: Point{0, 7}, B: Point{6, 7}},
				{A: Point{6, 7}, B: Point{6, 3}},
				{A: Point{6, 3}, B: Point{2, 3}},
			},
			p: Point{6, 5},
			d: 15,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			actual := StepsToPoint(tt.l, tt.p)
			assert.Equal(t, tt.d, actual)
		})
	}
}
