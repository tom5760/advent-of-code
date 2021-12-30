package main

import (
	"fmt"
	"log"
	"os"

	"github.com/tom5760/advent-of-code/2021/grid"
)

const (
	steps = 100
)

func main() {
	if err := run(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func run() error {
	grid, err := grid.Parse(os.Stdin)
	if err != nil {
		return fmt.Errorf("failed to parse input: %w", err)
	}

	fmt.Println("Part 1:", Part1(grid))
	fmt.Println("Part 2:", Part2(grid))

	return nil
}

func Part1(g *grid.Grid) uint64 {
	var total uint64

	sim := NewSim(g)

	for step := 0; step < steps; step++ {
		sim.Step()

		for _, v := range sim.Flashed {
			if v {
				total++
			}
		}
	}

	return total
}

func Part2(g *grid.Grid) uint64 {
	sim := NewSim(g)
	var step uint64

	for {
		step++
		sim.Step()

		allFlashed := true

		for _, v := range sim.Flashed {
			if !v {
				allFlashed = false
				break
			}
		}

		if allFlashed {
			return step
		}
	}
}

type Sim struct {
	Grid    *grid.Grid
	Flashed []bool
}

func NewSim(g *grid.Grid) *Sim {
	return &Sim{
		Grid:    grid.Copy(g),
		Flashed: make([]bool, len(g.Data)),
	}
}

func (s *Sim) Step() {
	for i := range s.Flashed {
		s.Flashed[i] = false
	}

	// First, the energy level of each octopus increases by 1.
	for i, v := range s.Grid.Data {
		s.Grid.Data[i] = v + 1
	}

	// Then, any octopus with an energy level greater than 9 flashes. This
	// increases the energy level of all adjacent octopuses by 1, including
	// octopuses that are diagonally adjacent. If this causes an octopus to
	// have an energy level greater than 9, it also flashes. This process
	// continues as long as new octopuses keep having their energy level
	// increased beyond 9. (An octopus can only flash at most once per step.)
	newFlash := true
	for newFlash {
		newFlash = false

		for i, v := range s.Grid.Data {
			if s.Flashed[i] {
				continue
			}

			if v > 9 {
				s.Flashed[i] = true
				newFlash = true

				x, y := s.Grid.Coord(uint64(i))

				if x > 0 {
					*s.Grid.Ptr(x-1, y)++
				}
				if x < s.Grid.Width-1 {
					*s.Grid.Ptr(x+1, y)++
				}

				if y > 0 {
					*s.Grid.Ptr(x, y-1)++
				}
				if y < s.Grid.Height-1 {
					*s.Grid.Ptr(x, y+1)++
				}

				if x > 0 && y > 0 {
					*s.Grid.Ptr(x-1, y-1)++
				}
				if x < s.Grid.Width-1 && y > 0 {
					*s.Grid.Ptr(x+1, y-1)++
				}
				if x > 0 && y < s.Grid.Height-1 {
					*s.Grid.Ptr(x-1, y+1)++
				}
				if x < s.Grid.Width-1 && y < s.Grid.Height-1 {
					*s.Grid.Ptr(x+1, y+1)++
				}
			}
		}
	}

	// Finally, any octopus that flashed during this step has its energy level
	// set to 0, as it used all of its energy to flash.
	for i := range s.Flashed {
		if s.Flashed[i] {
			s.Grid.Data[i] = 0
		}
	}
}
