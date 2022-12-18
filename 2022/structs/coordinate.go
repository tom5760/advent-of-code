package structs

type Coordinate struct {
	X, Y int
}

func ForLine(from, to Coordinate, fn func(Coordinate)) {
	// Assume line is orthogonal
	var inc func(*Coordinate)

	switch {
	case from.X < to.X:
		inc = func(c *Coordinate) { c.X++ }
	case from.X > to.X:
		inc = func(c *Coordinate) { c.X-- }

	case from.Y < to.Y:
		inc = func(c *Coordinate) { c.Y++ }
	case from.Y > to.Y:
		inc = func(c *Coordinate) { c.Y-- }

	default:
		panic("no line?")
	}

	for from.X != to.X || from.Y != to.Y {
		fn(from)
		inc(&from)
	}

	fn(from)
}

// Distance returns the Manhattan distance between a and b.
func Distance(a, b Coordinate) int {
	return abs(a.X-b.X) + abs(a.Y-b.Y)
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}
