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

// System is a collection of bodies.
type System struct {
	Bodies []Body
	pairs  []pair
}

type pair struct {
	a, da, b, db *int
}

// ParseSystem parses line input into a System.
func ParseSystem(lines []string) (*System, error) {
	n := len(lines)
	s := &System{
		Bodies: make([]Body, n),
		pairs:  make([]pair, 0, (n*(n-1))/2),
	}

	for i, line := range lines {
		p := &s.Bodies[i]
		_, err := fmt.Sscanf(line, "<x=%d, y=%d, z=%d>", &p.X, &p.Y, &p.Z)
		if err != nil {
			return nil, fmt.Errorf("failed to scan line: %w", err)
		}
	}

	for i := 0; i < len(s.Bodies); i++ {
		a := &s.Bodies[i]
		for j := i + 1; j < len(s.Bodies); j++ {
			b := &s.Bodies[j]
			s.pairs = append(s.pairs,
				pair{&a.X, &a.DX, &b.X, &b.DX},
				pair{&a.Y, &a.DY, &b.Y, &b.DY},
				pair{&a.Z, &a.DZ, &b.Z, &b.DZ},
			)
		}
	}

	return s, nil
}

// Step simulates one step of a System.
func (s System) Step() {
	s.velocity()
	s.position()
}

// Energy returns the total energy of the system.
func (s System) Energy() int {
	var sum int
	for _, body := range s.Bodies {
		sum += body.Energy()
	}
	return sum
}

func (s System) velocity() {
	for _, p := range s.pairs {
		if *p.a > *p.b {
			*p.da--
			*p.db++
		} else if *p.a < *p.b {
			*p.da++
			*p.db--
		}
	}
}

func (s System) position() {
	for i := range s.Bodies {
		body := &s.Bodies[i]
		body.X += body.DX
		body.Y += body.DY
		body.Z += body.DZ
	}
}
