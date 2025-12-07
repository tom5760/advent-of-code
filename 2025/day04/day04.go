package day04

import (
	"fmt"
	"io"
	"log"

	"github.com/tom5760/advent-of-code/aoc2025/grid"
)

const (
	CellEmpty   = '.'
	CellPaper   = '@'
	CellRemoved = 'x'
)

func Run(l *log.Logger, r io.Reader) (int, int, error) {
	var part1, part2 int

	g, err := grid.Parse(r)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to parse grid: %w", err)
	}

	for i := 0; ; i++ {
		for p, v := range g.Cells() {
			if v != CellPaper {
				continue
			}

			var total int

			for _, v2 := range g.Adjacent(p) {
				switch v2 {
				case CellPaper, CellRemoved:
					total++
				}
			}

			if total < 4 {
				if i == 0 {
					part1++
				}

				g.Set(p.X, p.Y, CellRemoved)
			}
		}

		var totalRemoved int

		for p, v := range g.Cells() {
			if v == CellRemoved {
				totalRemoved++
				g.Set(p.X, p.Y, CellEmpty)
			}
		}

		if totalRemoved == 0 {
			break
		}

		part2 += totalRemoved
	}

	return part1, part2, nil
}
