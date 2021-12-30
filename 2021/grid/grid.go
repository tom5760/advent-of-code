package grid

import (
	"bufio"
	"fmt"
	"io"
)

type Grid struct {
	Width  uint64
	Height uint64
	Data   []uint8
}

func Parse(r io.Reader) (*Grid, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanBytes)

	var (
		g        Grid
		widthSet bool
	)

	for scanner.Scan() {
		b := scanner.Bytes()[0]

		if b == '\n' {
			g.Height++
			widthSet = true

			continue
		}

		// Assume ASCII digit.
		if b < 48 || b > 57 {
			panic("unexpected map byte")
		}

		if !widthSet {
			g.Width++
		}

		g.Data = append(g.Data, b-48)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to scan input: %w", err)
	}

	return &g, nil
}

func Copy(a *Grid) *Grid {
	b := &Grid{
		Width:  a.Width,
		Height: a.Height,
		Data:   make([]uint8, len(a.Data)),
	}

	copy(b.Data, a.Data)

	return b
}

func (g *Grid) Index(x, y uint64) uint64 {
	return y*g.Width + x
}

func (g *Grid) Coord(i uint64) (x, y uint64) {
	x = uint64(i) % g.Width
	y = uint64(i) / g.Width

	return x, y
}

func (g *Grid) Get(x, y uint64) uint8 {
	return g.Data[g.Index(x, y)]
}

func (g *Grid) Ptr(x, y uint64) *uint8 {
	return &g.Data[g.Index(x, y)]
}

func (g *Grid) Set(x, y uint64, v uint8) {
	g.Data[g.Index(x, y)] = v
}
