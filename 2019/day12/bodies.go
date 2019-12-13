package main

import (
	"fmt"

	"github.com/tom5760/advent-of-code/2019/common"
)

// Body is a body in space.
type Body struct {
	X, Y, Z    int
	DX, DY, DZ int
}

// Potential is the potential energy of the body.
func (b *Body) Potential() int {
	return common.IntAbs(b.X) + common.IntAbs(b.Y) + common.IntAbs(b.Z)
}

// Kinetic is the kinetic energy of the body.
func (b *Body) Kinetic() int {
	return common.IntAbs(b.DX) + common.IntAbs(b.DY) + common.IntAbs(b.DZ)
}

// Energy is the total energy of the body.
func (b *Body) Energy() int {
	return b.Potential() * b.Kinetic()
}

// Bodies is a collection of bodies.
type Bodies []Body

// ParseBodies parses line input into a set of bodies.
func ParseBodies(lines []string) (Bodies, error) {
	planets := make(Bodies, len(lines))

	for i, line := range lines {
		p := &planets[i]
		_, err := fmt.Sscanf(line, "<x=%d, y=%d, z=%d>", &p.X, &p.Y, &p.Z)
		if err != nil {
			return nil, fmt.Errorf("failed to scan line: %w", err)
		}
	}

	return planets, nil
}

// Run runs the simulation for step steps.
func (b Bodies) Run(steps int) {
	for i := 0; i < steps; i++ {
		b.Step()
	}
}

// Step simulates one step of a set of bodies.
func (b Bodies) Step() {
	b.velocity()
	b.position()
}

// Energy returns the total energy of the system.
func (b Bodies) Energy() int {
	var sum int
	for _, body := range b {
		sum += body.Energy()
	}
	return sum
}

func (b Bodies) velocity() {
	for _, pair := range b.pairwise() {
		a := pair[0]
		b := pair[1]
		if a.X > b.X {
			a.DX--
			b.DX++
		}
		if a.X < b.X {
			a.DX++
			b.DX--
		}
		if a.Y > b.Y {
			a.DY--
			b.DY++
		}
		if a.Y < b.Y {
			a.DY++
			b.DY--
		}
		if a.Z > b.Z {
			a.DZ--
			b.DZ++
		}
		if a.Z < b.Z {
			a.DZ++
			b.DZ--
		}
	}
}

func (b Bodies) position() {
	for i := range b {
		body := &b[i]
		body.X += body.DX
		body.Y += body.DY
		body.Z += body.DZ
	}
}

func (b Bodies) pairwise() [][2]*Body {
	var pairs [][2]*Body

	for i := 0; i < len(b); i++ {
		for j := i + 1; j < len(b); j++ {
			pair := [2]*Body{&b[i], &b[j]}
			pairs = append(pairs, pair)
		}
	}

	return pairs
}
