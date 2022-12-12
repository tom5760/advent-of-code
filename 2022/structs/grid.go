package structs

type Grid[T any] struct {
	Height int
	Width  int
	Values []T
}

func (g *Grid[T]) Get(x, y int) T {
	return g.Values[y*g.Width+x]
}

func (g *Grid[T]) Set(x, y int, v T) {
	g.Values[y*g.Width+x] = v
}
