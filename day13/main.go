package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type Point struct {
	X, Y uint64
}

func (p Point) String() string {
	return fmt.Sprintf("%d,%d", p.X, p.Y)
}

type RailKind byte

const (
	Vertical     RailKind = '|'
	Horizontal   RailKind = '-'
	Curve        RailKind = '/'
	BackCurve    RailKind = '\\'
	Intersection RailKind = '+'
)

func (k RailKind) String() string {
	return string(k)
}

type Rail struct {
	Point Point
	Kind  RailKind

	Carts map[*Cart]bool
}

// Enter returns true if there was a collision.
func (r *Rail) Enter(c *Cart) bool {
	r.Carts[c] = true
	return len(r.Carts) > 1
}

func (r *Rail) Leave(c *Cart) {
	delete(r.Carts, c)
}

func (r *Rail) String() string {
	if len(r.Carts) > 1 {
		return "X"
	}
	if len(r.Carts) > 0 {
		for cart := range r.Carts {
			return cart.Direction.String()
		}
	}
	return r.Kind.String()
}

type CartDirection byte

const (
	Up    CartDirection = '^'
	Right CartDirection = '>'
	Down  CartDirection = 'v'
	Left  CartDirection = '<'
)

func (d CartDirection) TurnRight() CartDirection {
	switch d {
	case Up:
		return Right
	case Right:
		return Down
	case Down:
		return Left
	case Left:
		return Up
	}
	panic("invalid direction")
}

func (d CartDirection) TurnLeft() CartDirection {
	switch d {
	case Up:
		return Left
	case Right:
		return Up
	case Down:
		return Right
	case Left:
		return Down
	}
	panic("invalid direction")
}

func (d CartDirection) String() string {
	return string(d)
}

type CartTurn int

const (
	TurnLeft CartTurn = iota
	GoStraight
	TurnRight
)

type Cart struct {
	Point     Point
	Direction CartDirection
	NextTurn  CartTurn
}

// Move returns true if there was a collision.
func (c *Cart) Move(rails map[Point]*Rail) bool {
	rails[c.Point].Leave(c)

	switch c.Direction {
	case Up:
		c.Point.Y--
	case Right:
		c.Point.X++
	case Down:
		c.Point.Y++
	case Left:
		c.Point.X--
	}

	return rails[c.Point].Enter(c)
}

// Tick returns true if there was a collision.
func (c *Cart) Tick(rails map[Point]*Rail) bool {
	if c.Move(rails) {
		return true
	}

	rail := rails[c.Point]
	switch rail.Kind {
	case Curve:
		switch c.Direction {
		case Up, Down:
			c.Direction = c.Direction.TurnRight()
		case Left, Right:
			c.Direction = c.Direction.TurnLeft()
		}

	case BackCurve:
		switch c.Direction {
		case Up, Down:
			c.Direction = c.Direction.TurnLeft()
		case Left, Right:
			c.Direction = c.Direction.TurnRight()
		}

	case Intersection:
		switch c.NextTurn {
		case TurnLeft:
			c.Direction = c.Direction.TurnLeft()

		case TurnRight:
			c.Direction = c.Direction.TurnRight()
		}
		c.NextTurn = (c.NextTurn + 1) % 3
	}

	return false
}

type Mine struct {
	Width, Height uint64

	Rails map[Point]*Rail
	Carts map[*Cart]bool
}

// Tick returns a slice of collision points.
func (m *Mine) Tick() []*Point {
	var crashes []*Point

	for cart := range m.Carts {
		if cart.Tick(m.Rails) {
			rail := m.Rails[cart.Point]

			for crashed := range rail.Carts {
				delete(m.Carts, crashed)
				delete(rail.Carts, crashed)
			}

			crashes = append(crashes, &cart.Point)
		}
	}

	return crashes
}

func (m *Mine) String() string {
	var sb strings.Builder

	sb.WriteByte('\n')

	for y := uint64(0); y < m.Height; y++ {
		for x := uint64(0); x < m.Width; x++ {
			point := Point{x, y}
			if rail, ok := m.Rails[point]; ok {
				sb.WriteString(rail.String())
			} else {
				sb.WriteString(" ")
			}
		}
		sb.WriteByte('\n')
	}

	return sb.String()
}

func readInput(r io.Reader) ([]*Rail, []*Cart) {
	scanner := bufio.NewScanner(r)

	var rails []*Rail
	var carts []*Cart
	var x, y uint64

	for scanner.Scan() {
		for _, b := range scanner.Bytes() {
			var cart *Cart
			var rail *Rail

			switch b {
			case '^', 'v', '>', '<':
				cart = &Cart{Direction: CartDirection(b)}

			case '|', '-', '/', '\\', '+':
				rail = &Rail{Kind: RailKind(b)}
			}

			point := Point{x, y}

			if cart != nil {
				cart.Point = point
				carts = append(carts, cart)

				rail = new(Rail)

				switch cart.Direction {
				case Up, Down:
					rail.Kind = Vertical
				case Left, Right:
					rail.Kind = Horizontal
				}
			}

			if rail != nil {
				rail.Point = point
				rail.Carts = make(map[*Cart]bool)
				rails = append(rails, rail)
			}

			x++
		}

		x = 0
		y++
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln("failed to read input:", err)
		return nil, nil
	}

	return rails, carts
}

func buildMine(rails []*Rail, carts []*Cart) *Mine {
	mine := Mine{
		Rails: make(map[Point]*Rail),
		Carts: make(map[*Cart]bool),
	}

	for _, rail := range rails {
		mine.Rails[rail.Point] = rail

		if rail.Point.X+1 > mine.Width {
			mine.Width = rail.Point.X + 1
		}
		if rail.Point.Y+1 > mine.Height {
			mine.Height = rail.Point.Y + 1
		}
	}

	for _, cart := range carts {
		mine.Carts[cart] = true
		mine.Rails[cart.Point].Enter(cart)
	}

	return &mine
}

func main() {
	mine := buildMine(readInput(os.Stdin))
	// log.Print(mine)

	var firstCrash bool

	for len(mine.Carts) > 1 {
		crashes := mine.Tick()
		// log.Print(mine)

		if len(crashes) > 0 && !firstCrash {
			log.Println("(part 1) location of first crash:", crashes[0])
			firstCrash = true
		}
	}

	for cart := range mine.Carts {
		log.Println("(part 2) location of last cart:", cart.Point)
		return
	}
}
