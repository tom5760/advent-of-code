package main

import (
	"fmt"
	"sort"
	"strings"
)

type Arena struct {
	Width, Height uint64

	Cells []*Cell
	Units []*Unit
}

// Tick returns whether the battle is over.
func (a *Arena) Tick() bool {
	// log.Println("round start")
	if len(a.Units) == 0 {
		return true
	}

	a.sortUnits()

	for i := 0; i < len(a.Units); i++ {
		unit := a.Units[i]
		// log.Println("tick for", unit, unit.Point)

		battleOver, killedUnit := unit.Tick(a)

		if killedUnit != nil {
			// log.Println(unit, unit.Point, "killed", killedUnit, killedUnit.Point)
			killedUnit.Cell.Unit = nil
			killedUnit.Cell = nil

			for j, u := range a.Units {
				if u == killedUnit {
					a.Units = append(a.Units[:j], a.Units[j+1:]...)
					if i > 0 && i > j {
						i--
					}
					break
				}
			}
		}

		if battleOver {
			return true
		}
	}

	return false
}

func (a *Arena) String() string {
	var sb strings.Builder

	sb.WriteByte('\n')

	for y := uint64(0); y < a.Height; y++ {
		for x := uint64(0); x < a.Width; x++ {
			cell := a.cell(x, y)
			sb.WriteString(cell.String())
		}

		sb.WriteString("   ")

		for _, unit := range a.Units {
			if unit.Point.Y == y {
				sb.WriteString(fmt.Sprintf("%v(%d), ", unit, unit.HP))
			}
		}

		if y < a.Height-1 {
			sb.WriteByte('\n')
		}
	}

	return sb.String()
}

func (a *Arena) cell(x, y uint64) *Cell {
	return a.Cells[y*a.Width+x]
}

// Sort units into reading order.
func (a *Arena) sortUnits() {
	sort.Slice(a.Units, func(i, j int) bool {
		return Less(a.Units[i].Point, a.Units[j].Point)
	})
}
