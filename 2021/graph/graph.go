package graph

type Graph[N any] struct {
	Nodes []N
}

func Parse(r io.Reader) (*Graph, error) {
}
