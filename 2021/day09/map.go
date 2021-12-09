package main

type FloorMap struct {
	Width  uint64
	Height uint64
	Data   []uint8
}

func (m *FloorMap) XY(x, y uint64) uint8 {
	return m.Data[y*m.Width+x]
}

func (m *FloorMap) LowPoints() []int {
	var points []int

	for i, d := range m.Data {
		x := uint64(i) % m.Width
		y := uint64(i) / m.Width

		ok := true

		// Left
		if x > 0 {
			ok = ok && d < m.Data[i-1]
		}

		// Right
		if x < (m.Width - 1) {
			ok = ok && d < m.Data[i+1]
		}

		// Up
		if y > 0 {
			ok = ok && d < m.Data[uint64(i)-m.Width]
		}

		// Down
		if y < (m.Height - 1) {
			ok = ok && d < m.Data[uint64(i)+m.Width]
		}

		if ok {
			// This point is the lowest around.
			points = append(points, i)
		}
	}

	return points
}

func (m *FloorMap) FindBasin(i int) uint64 {
	seen := make(map[int]bool, len(m.Data))

	frontier := make(map[int]bool, len(m.Data))
	frontier[i] = true

	var size uint64

	for len(frontier) > 0 {
		i := mapPop(frontier)
		if seen[i] {
			continue
		}

		seen[i] = true
		d := m.Data[i]

		// Locations of height 9 do not count as being in any basin.
		if d == 9 {
			continue
		}

		size++

		x := uint64(i) % m.Width
		y := uint64(i) / m.Width

		// Left
		if x > 0 && m.Data[i-1] >= d {
			frontier[i-1] = true
		}

		// Right
		if x < (m.Width-1) && m.Data[i+1] >= d {
			frontier[i+1] = true
		}

		// Up
		if y > 0 && m.Data[uint64(i)-m.Width] >= d {
			frontier[i-int(m.Width)] = true
		}

		// Down
		if y < (m.Height-1) && m.Data[uint64(i)+m.Width] >= d {
			frontier[i+int(m.Width)] = true
		}
	}

	return size
}

func mapPop(m map[int]bool) int {
	for v := range m {
		delete(m, v)
		return v
	}

	panic("empty map")
}
