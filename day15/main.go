package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

func readInput(r io.Reader) *Arena {
	scanner := bufio.NewScanner(r)

	arena := new(Arena)
	var point Point

	for scanner.Scan() {
		if point.Y+1 > arena.Height {
			arena.Height = point.Y + 1
		}

		for _, b := range scanner.Bytes() {
			if point.X+1 > arena.Width {
				arena.Width = point.X + 1
			}

			cell := &Cell{
				Kind:  CellKind(b),
				Point: point,
			}

			if b == 'G' || b == 'E' {
				unit := &Unit{
					Kind: UnitKind(b),
					AP:   DefaultAP,
					HP:   DefaultHP,
				}
				cell.Kind = Floor
				cell.Enter(unit)

				arena.Units = append(arena.Units, unit)
			}

			arena.Cells = append(arena.Cells, cell)

			point.X++
		}

		point.X = 0
		point.Y++
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln("failed to read input:", err)
		return nil
	}

	return arena
}

func main() {
	arena := readInput(os.Stdin)
	arena.Battle()

	log.Printf("(part 1) battle outcome after %d rounds, %d total HP: %d",
		arena.LastRound, arena.TotalHP(), arena.Outcome())
}
