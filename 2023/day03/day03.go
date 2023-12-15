package day03

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"strings"
)

type (
	grid map[point]any

	point struct {
		x, y int
	}

	number struct {
		number  int
		symbols []*symbol
	}

	symbol struct {
		kind  byte
		parts []*number
	}
)

func Run(lg *log.Logger, input io.Reader) (int, int, error) {
	var part1, part2 int

	g, err := parseGrid(input)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to parse input: %w", err)
	}

	parts := map[*number]bool{}

	for _, v := range g {
		if num, ok := v.(*number); ok && len(num.symbols) > 0 {
			parts[num] = true
		}

		if sym, ok := v.(*symbol); ok && sym.kind == '*' && len(sym.parts) == 2 {
			part2 += sym.parts[0].number * sym.parts[1].number
		}
	}

	for prt := range parts {
		part1 += prt.number
	}

	return part1, part2, nil
}

func parseGrid(r io.Reader) (grid, error) {
	scanner := bufio.NewScanner(r)

	g := grid{}

	var p point

	for scanner.Scan() {
		line := scanner.Bytes()

		var curNum *number

		for i, b := range line {
			p.x = i

			if b >= '0' && b <= '9' {
				if curNum == nil {
					curNum = new(number)

					// check behind for symbols
					checkPart(g, curNum, p.x-1, p.y)   // left
					checkPart(g, curNum, p.x-1, p.y-1) // left-up
				}

				curNum.number *= 10
				curNum.number += int(b - '0')

				g[p] = curNum

				// check behind for symbols
				checkPart(g, curNum, p.x, p.y-1) // up

				continue
			}

			if curNum != nil {
				// part finished
				// check behind for symbols
				checkPart(g, curNum, p.x, p.y-1) // right-up
				curNum = nil
			}

			if b == '.' {
				continue
			}

			// found a symbol
			sym := &symbol{
				kind: b,
			}
			g[p] = sym

			// check behind for parts
			seen := make(map[*number]bool, 4)
			checkSymbol(g, sym, seen, p.x-1, p.y)   // left
			checkSymbol(g, sym, seen, p.x, p.y-1)   // up
			checkSymbol(g, sym, seen, p.x-1, p.y-1) // left-up
			checkSymbol(g, sym, seen, p.x+1, p.y-1) // right-up
		}

		p.y++
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to scan input: %w", err)
	}

	return g, nil
}

func checkPart(g grid, p *number, x, y int) {
	if v, ok := g[point{x, y}]; ok {
		if s, ok := v.(*symbol); ok {
			p.symbols = append(p.symbols, s)
			s.parts = append(s.parts, p)
		}
	}
}

func checkSymbol(g grid, s *symbol, seen map[*number]bool, x, y int) {
	if v, ok := g[point{x, y}]; ok {
		if p, ok := v.(*number); ok {
			if seen[p] {
				return
			}

			s.parts = append(s.parts, p)
			p.symbols = append(p.symbols, s)
			seen[p] = true
		}
	}
}

func (g grid) String() string {
	maxX := math.MinInt
	maxY := math.MinInt

	minX := math.MaxInt
	minY := math.MaxInt

	for p := range g {
		if p.x < minX {
			minX = p.x
		}
		if p.y < minY {
			minY = p.y
		}
		if p.x > maxX {
			maxX = p.x
		}
		if p.y > maxY {
			maxY = p.y
		}
	}

	var b strings.Builder

	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			v, ok := g[point{x, y}]
			if !ok {
				b.WriteByte('.')
				continue
			}

			switch v := v.(type) {
			case *number:
				n, _ := fmt.Fprint(&b, v.number)
				x += n - 1

			case *symbol:
				b.WriteByte(v.kind)
			}
		}
		b.WriteByte('\n')
	}

	return b.String()
}
