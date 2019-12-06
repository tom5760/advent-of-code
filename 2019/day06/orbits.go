package main

import (
	"strings"
)

const (
	centerID = "COM"
)

// A Planet orbits a Parent, and has Children orbiting it.  Forms a tree
// structure.
type Planet struct {
	ID       string
	Parent   *Planet
	Children []*Planet
}

// The Universe has planets and a Center of mass.
type Universe struct {
	Center  *Planet
	Planets map[string]*Planet
}

// ParseUniverse parses lines of the form "ParentID)ChildID" into a universe.
func ParseUniverse(lines []string) *Universe {
	planets := map[string]*Planet{}

	getPlanet := func(id string) *Planet {
		p, ok := planets[id]
		if !ok {
			p = &Planet{ID: id}
			planets[id] = p
		}
		return p
	}

	for _, line := range lines {
		parts := strings.Split(line, ")")
		if len(parts) != 2 {
			panic("unexpected number of parts in orbit line")
		}

		center := getPlanet(parts[0])
		child := getPlanet(parts[1])

		center.Children = append(center.Children, child)
		child.Parent = center
	}

	center, ok := planets[centerID]
	if !ok {
		panic("no center of mass")
	}

	if center.Parent != nil {
		panic("center of mass has parent")
	}

	return &Universe{
		Center:  center,
		Planets: planets,
	}
}

// TotalOrbits returns the total number of direct and indirect orbits in the
// universe.  This is counting the total distances from every node in the tree
// to the root.
func TotalOrbits(u *Universe) int {
	var orbits int

	for _, planet := range u.Planets {
		for planet.Parent != nil {
			orbits++
			planet = planet.Parent
		}
	}

	return orbits
}

// TotalTransfers uses a breadth-first search to find the distance between two
// nodes.
func TotalTransfers(u *Universe, start, end *Planet) int {
	q := []*Planet{}
	parents := map[*Planet]*Planet{}
	disc := map[*Planet]bool{start: true}

	q = append(q, start)

	for len(q) > 0 {
		cur := q[0]
		q = q[1:]

		if cur == end {
			dist := 0
			x := cur

			for {
				if x == nil {
					// Don't count transfer for start -> A, A -> B, and X -> end.
					return dist - 3
				}
				x = parents[x]
				dist++
			}
		}

		for _, child := range cur.Children {
			if !disc[child] {
				disc[child] = true
				parents[child] = cur
				q = append(q, child)
			}
		}

		if cur.Parent != nil && !disc[cur.Parent] {
			disc[cur.Parent] = true
			parents[cur.Parent] = cur
			q = append(q, cur.Parent)
		}
	}

	panic("couldn't find path")
}
