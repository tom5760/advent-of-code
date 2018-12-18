package main

type PointSet map[Point]bool

func (s PointSet) Pop() Point {
	for point := range s {
		delete(s, point)
		return point
	}
	panic("set is empty")
}

func Flow(w *World) {
	openSet := make(PointSet)
	openSet[Point{500, 1}] = true

	var tick int

	for len(openSet) > 0 {
		terminalSet := flowPhase1(w, openSet)
		flowPhase2(w, openSet, terminalSet)
		tick++
	}
}

func flowPhase1(w *World, openSet PointSet) PointSet {
	closedSet := make(PointSet)
	terminalSet := make(PointSet)

	checkAdd := func(p Point) bool {
		tile := w.Point(p)
		if tile != nil && (*tile == ReachableTile || *tile == SandTile) {
			if !closedSet[p] {
				openSet[p] = true
			}
			return true
		}
		return false
	}

	for len(openSet) > 0 {
		current := openSet.Pop()

		closedSet[current] = true
		*w.Point(current) = ReachableTile

		if w.Point(current.Down()) == nil {
			continue
		}

		if checkAdd(current.Down()) {
			continue
		}

		if !checkAdd(current.Left()) && w.TileIs(current.Left(), ClayTile) {
			terminalSet[current] = true
		}

		if !checkAdd(current.Right()) && w.TileIs(current.Right(), ClayTile) {
			terminalSet[current] = true
		}
	}

	return terminalSet
}

func flowPhase2(w *World, openSet, terminalSet PointSet) {
	for terminal := range terminalSet {
		delete(terminalSet, terminal)

		var walkFn func(p Point) Point

		// Figure out which way to walk
		if w.TileIs(terminal.Left(), ReachableTile) {
			walkFn = func(p Point) Point {
				return p.Left()
			}
		} else if w.TileIs(terminal.Right(), ReachableTile) {
			walkFn = func(p Point) Point {
				return p.Right()
			}
		} else {
			// We are in a one-tile pit, fill with water, and check above again.
			*w.Point(terminal) = WaterTile
			openSet[terminal.Up()] = true
			continue
		}

		// Walk left or right until we hit another tile in the terminal set.
		var next Point
		for next = walkFn(terminal); !terminalSet[next] && w.TileIs(next, ReachableTile) &&
			next.X >= w.MinX-1 && next.X <= w.MaxX+1; next = walkFn(next) {
		}

		// If we don't hit another terminal, then water isn't pooling here.
		if !terminalSet[next] {
			continue
		}

		delete(terminalSet, next)

		// Fill the space between terminals with water
		var stop bool
		for p := terminal; !stop; p = walkFn(p) {
			*w.Point(p) = WaterTile

			// See if water should flow up one level.
			if w.TileIs(p.Up(), ReachableTile) {
				openSet[p.Up()] = true
			}
			stop = p == next
		}
	}
}
