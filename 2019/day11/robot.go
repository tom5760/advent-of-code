package main

import "github.com/tom5760/advent-of-code/2019/sif"

// Direction is a direction the robot can fase.
type Direction byte

// The directions a robot can fase.
const (
	DirectionUp    Direction = '^'
	DirectionDown  Direction = 'v'
	DirectionLeft  Direction = '<'
	DirectionRight Direction = '>'
)

// Clockwise returns the next clockwise direction.
func (d Direction) Clockwise() Direction {
	switch d {
	case DirectionUp:
		return DirectionRight
	case DirectionRight:
		return DirectionDown
	case DirectionDown:
		return DirectionLeft
	case DirectionLeft:
		return DirectionUp
	}
	panic("invalid direction")
}

// CounterClockwise returns the next counter-clockwise direction.
func (d Direction) CounterClockwise() Direction {
	switch d {
	case DirectionUp:
		return DirectionLeft
	case DirectionLeft:
		return DirectionDown
	case DirectionDown:
		return DirectionRight
	case DirectionRight:
		return DirectionUp
	}
	panic("invalid direction")
}

// Point is a point on the hull
type Point struct {
	X, Y int
}

// Robot is a robot that can paint the hull.
type Robot struct {
	Facing   Direction
	Location Point
	Hull     map[Point]sif.Color
}

// NewRobot initializes a new robot.
func NewRobot() *Robot {
	return &Robot{
		Facing:   DirectionUp,
		Location: Point{0, 0},
		Hull:     map[Point]sif.Color{},
	}
}

// Clockwise rotates the robot clockwise
func (r *Robot) Clockwise() {
	r.Facing = r.Facing.Clockwise()
}

// CouterClockwise rotates the robot counter-clockwise
func (r *Robot) CounterClockwise() {
	r.Facing = r.Facing.CounterClockwise()
}

// Forward moves the robot one space in the direction it is facing.
func (r *Robot) Forward() {
	switch r.Facing {
	case DirectionUp:
		r.Location.Y--
	case DirectionRight:
		r.Location.X++
	case DirectionDown:
		r.Location.Y++
	case DirectionLeft:
		r.Location.X--
	default:
		panic("invalid direction")
	}
}

// Scan returns the color of the tile under the robot.
func (r *Robot) Scan() sif.Color {
	color, ok := r.Hull[r.Location]
	if !ok {
		return sif.ColorBlack
	}
	return color
}

// Paint paints the color under the robot.
func (r *Robot) Paint(c sif.Color) {
	r.Hull[r.Location] = c
}

// Run runs the robot with the given channels.
func (r *Robot) Run(colorChan chan<- int, cmdChan <-chan int) {
	for {
		colorChan <- colorToInt(r.Scan())

		paintCmd, ok := <-cmdChan
		if !ok {
			return
		}

		r.Paint(intToColor(paintCmd))

		moveCmd, ok := <-cmdChan
		if !ok {
			return
		}

		switch moveCmd {
		case 0:
			r.CounterClockwise()
		case 1:
			r.Clockwise()
		default:
			panic("invalid move command")
		}

		r.Forward()
	}
}

func colorToInt(c sif.Color) int {
	switch c {
	case sif.ColorBlack:
		return 0
	case sif.ColorWhite:
		return 1
	}
	panic("invalid color")
}

func intToColor(c int) sif.Color {
	switch c {
	case 0:
		return sif.ColorBlack
	case 1:
		return sif.ColorWhite
	}
	panic("invalid color")
}
