package main

type CellKind byte

const (
	Wall  CellKind = '#'
	Floor CellKind = '.'
)

type Cell struct {
	Kind  CellKind
	Point Point

	Unit *Unit
}

func (c *Cell) Enter(unit *Unit) {
	if unit.Cell != nil {
		unit.Cell.Leave(unit)
	}

	unit.Cell = c
	c.Unit = unit
}

func (c *Cell) Leave(unit *Unit) {
	unit.Cell.Unit = nil
	unit.Cell = nil
	c.Unit = nil
}

func (c *Cell) String() string {
	if c.Unit != nil {
		return c.Unit.String()
	}
	return string(c.Kind)
}

func (c *Cell) EmptyNeighbors(arena *Arena) []*Cell {
	var neighbors []*Cell

	points := []Point{
		{c.Point.X, c.Point.Y - 1},
		{c.Point.X - 1, c.Point.Y},
		{c.Point.X + 1, c.Point.Y},
		{c.Point.X, c.Point.Y + 1},
	}

	for _, point := range points {
		if point.X < 0 || point.X > arena.Width-1 ||
			point.Y < 0 || point.Y > arena.Height-1 {
			continue
		}

		neighbor := arena.cell(point.X, point.Y)
		if neighbor.Kind == Floor && neighbor.Unit == nil {
			neighbors = append(neighbors, neighbor)
		}
	}

	return neighbors
}

// FloorNeighbors returns neighboring cells that are floors, regardless of
// whether they are currently occupied or not.
func (c *Cell) FloorNeighbors(arena *Arena) []*Cell {
	var neighbors []*Cell

	points := []Point{
		{c.Point.X, c.Point.Y - 1},
		{c.Point.X - 1, c.Point.Y},
		{c.Point.X + 1, c.Point.Y},
		{c.Point.X, c.Point.Y + 1},
	}

	for _, point := range points {
		if point.X < 0 || point.X > arena.Width-1 ||
			point.Y < 0 || point.Y > arena.Height-1 {
			continue
		}

		neighbor := arena.cell(point.X, point.Y)
		if neighbor.Kind == Floor {
			neighbors = append(neighbors, neighbor)
		}
	}

	return neighbors
}
