package main

import "fmt"

type Point struct {
	X, Y uint64
}

func (p Point) String() string {
	return fmt.Sprintf("%d,%d", p.X, p.Y)
}

func Adjacent(a, b Point) bool {
	if a.X == b.X {
		return a.Y == b.Y-1 || a.Y == b.Y+1
	}
	if a.Y == b.Y {
		return a.X == b.X-1 || a.X == b.X+1
	}
	return false
}

func Less(a, b Point) bool {
	if a.Y < b.Y {
		return true
	} else if a.Y == b.Y {
		return a.X < b.X
	}
	return false
}
