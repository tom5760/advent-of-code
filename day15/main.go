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
					Kind:  UnitKind(b),
					Point: point,
					Cell:  cell,
					AP:    DefaultAP,
					HP:    DefaultHP,
				}

				cell.Kind = Floor
				cell.Unit = unit

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
	log.Println("initial:", arena)

	round := 1
	for !arena.Tick() {
		// for _, unit := range arena.Units {
		// 	log.Println(unit, unit.Point, "HP", unit.HP)
		// }
		log.Println("round", round, arena)

		round++
		// time.Sleep(500 * time.Millisecond)
	}
	round--

	var totalHP int
	for _, unit := range arena.Units {
		totalHP += int(unit.HP)
	}

	log.Printf("(part 1) battle outcome after %d rounds, %d total HP: %d", round, totalHP, round*totalHP)
}
