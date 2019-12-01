package main

import (
	"fmt"
	"sort"
	"strings"
)

type Results struct {
	Rounds  int
	TotalHP int
	Outcome int
}

type Arena struct {
	Width, Height uint64

	LastRound int

	Cells []*Cell
	Units []*Unit
}

func (a *Arena) Battle() {
	a.LastRound = 1
	for !a.Round() {
		// log.Println(a.LastRound, a)
	}
	a.LastRound--
}

// Round returns whether the battle is over.
func (a *Arena) Round() bool {
	if len(a.Units) == 0 {
		return true
	}

	sort.Slice(a.Units, func(i, j int) bool {
		return Less(a.Units[i].Cell.Point, a.Units[j].Cell.Point)
	})

	isOver := false
	for _, unit := range a.Units {
		if unit.Turn(a) {
			isOver = true
			break
		}
	}

	// Clean up dead units.
	for i := 0; i < len(a.Units); i++ {
		if a.Units[i].HP <= 0 {
			a.Units = append(a.Units[:i], a.Units[i+1:]...)
			i--
		}
	}

	if !isOver {
		a.LastRound++
	}
	return isOver
}

func (a *Arena) TotalHP() int {
	var total int
	for _, unit := range a.Units {
		total += int(unit.HP)
	}
	return total
}

func (a *Arena) Outcome() int {
	return a.LastRound * a.TotalHP()
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
			if unit.Cell.Point.Y == y {
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
