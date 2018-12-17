package main

import (
	"bufio"
	"bytes"
	"io"
	"io/ioutil"
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
	initialInput, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalln("failed to read input:", err)
		return
	}

	arena := readInput(bytes.NewReader(initialInput))
	arena.Battle()

	log.Printf("(part 1) battle outcome after %d rounds, %d total HP: %d",
		arena.LastRound, arena.TotalHP(), arena.Outcome())

	elfAP := int64(4)
	for {
		arena = readInput(bytes.NewReader(initialInput))

		var initialElves int
		for _, unit := range arena.Units {
			if unit.Kind == Elf {
				initialElves++
				unit.AP = elfAP
			}
		}

		arena.Battle()

		var finalElves int
		for _, unit := range arena.Units {
			if unit.Kind == Elf {
				finalElves++
			}
		}

		if initialElves == finalElves {
			log.Printf("(part 2) battle outcome with %d AP after %d rounds, %d total HP: %d",
				elfAP, arena.LastRound, arena.TotalHP(), arena.Outcome())
			return
		}

		elfAP++
	}
}
