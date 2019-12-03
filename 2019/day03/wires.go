package main

import (
	"math"

	"github.com/tom5760/advent-of-code/2019/common"
)

// Direction is the direction of a wire instruction.
type Direction byte

// The possible directions for wires.
const (
	DirectionUp    Direction = 'U'
	DirectionDown  Direction = 'D'
	DirectionLeft  Direction = 'L'
	DirectionRight Direction = 'R'
)

// WireSegment is a direction and a length of wire in that direction.
type WireSegment struct {
	Direction Direction
	Length    int
}

// Wire is a sequence of WireSegments.
type Wire []WireSegment

// Point is a coordinate in an XY plane.
type Point struct {
	X, Y int
}

// LineSegment is a line segment defined by two points.
type LineSegment struct {
	A, B Point
}

// Lines is a sequence of line segments.
type Lines []LineSegment

// WireSetToLineSet converts a slice of wires to a slice of lines.
func WireSetToLineSet(w []Wire) []Lines {
	lineSet := make([]Lines, len(w))

	for i, wire := range w {
		lineSet[i] = WireToLines(wire)
	}

	return lineSet
}

// WireToLines transforms a sequence of wire instructions to a sequence of line
// segments.
func WireToLines(w Wire) Lines {
	c := Point{}
	lines := make(Lines, len(w))

	for i, wire := range w {
		segment := LineSegment{A: c}

		switch wire.Direction {
		case DirectionUp:
			c.Y += wire.Length
		case DirectionDown:
			c.Y -= wire.Length
		case DirectionLeft:
			c.X -= wire.Length
		case DirectionRight:
			c.X += wire.Length
		}

		segment.B = c
		lines[i] = segment
	}

	return lines
}

// LineSetIntersections returns intersection points for a set of lines.
func LineSetIntersections(lines []Lines) []Point {
	points := map[Point]bool{}

	for i, a := range lines {
		for j, b := range lines {
			// Dont compare lines against themselves.
			if i == j {
				continue
			}

			for _, p := range LinesIntersections(a, b) {
				points[p] = true
			}
		}
	}

	filtered := make([]Point, 0, len(points))

	for point := range points {
		filtered = append(filtered, point)
	}

	return filtered
}

// LinesIntersections finds the intersecting points between two sets of line
// segments.  As a special case, intersections at the origin (0, 0) are
// filtered out (due to instructions).
func LinesIntersections(a, b Lines) []Point {
	var points []Point

	for _, i := range a {
		for _, j := range b {
			if point, ok := LineSegmentIntersection(i, j); ok {
				if point == (Point{0, 0}) {
					continue
				}

				points = append(points, point)
			}
		}
	}

	return points
}

// LineSegmentIntersection returns the intersection point of two line segments,
// or false of the segments do not intersect.
//
// Assumes the segments are axis-aligned.  Overlapping lines do not intersect.
func LineSegmentIntersection(i, j LineSegment) (Point, bool) {
	ixMin, ixMax := common.IntOrder(i.A.X, i.B.X)
	jxMin, jxMax := common.IntOrder(j.A.X, j.B.X)

	xMin := common.IntMax(ixMin, jxMin)
	xMax := common.IntMin(ixMax, jxMax)

	// Lines are horizontal, not intersecting.
	if xMin > xMax {
		return Point{}, false
	}

	iyMin, iyMax := common.IntOrder(i.A.Y, i.B.Y)
	jyMin, jyMax := common.IntOrder(j.A.Y, j.B.Y)

	yMin := common.IntMax(iyMin, jyMin)
	yMax := common.IntMin(iyMax, jyMax)

	if yMin > yMax {
		return Point{}, false
	}

	return Point{X: xMin, Y: yMin}, true
}

// ManhattanDistance computes the manhattan distance between two points.
func ManhattanDistance(a, b Point) int {
	return common.IntAbs(a.X-b.X) + common.IntAbs(a.Y-b.Y)
}

// ClosestPoint returns the point in set that is the closest to p by Manhattan
// distance.
func ClosestPoint(p Point, set []Point) (Point, int) {
	d := math.MaxInt64

	var closest Point

	for _, q := range set {
		nd := ManhattanDistance(p, q)
		if nd < d {
			d = nd
			closest = q
		}
	}

	return closest, d
}

// StepsToPoint returns the number of steps (integer coordinates) required to
// reach the given point p on line segments l.  Assumes p is on one of the line
// segments in l.
func StepsToPoint(l Lines, p Point) int {
	var steps int

	for _, seg := range l {
		horiz := seg.A.Y == seg.B.Y

		if horiz {
			if seg.A.Y == p.Y {
				if p.X >= common.IntMin(seg.A.X, seg.B.X) || p.X <= common.IntMax(seg.A.X, seg.B.X) {
					steps += common.IntAbs(seg.A.X - p.X)
					break
				}
			}
			steps += common.IntAbs(seg.A.X - seg.B.X)
		} else {
			if seg.A.X == p.X {
				if p.Y >= common.IntMin(seg.A.Y, seg.B.Y) || p.Y <= common.IntMax(seg.A.Y, seg.B.Y) {
					steps += common.IntAbs(seg.A.Y - p.Y)
					break
				}
			}
			steps += common.IntAbs(seg.A.Y - seg.B.Y)
		}
	}

	return steps
}
