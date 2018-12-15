package main

import (
	"math"
	"sort"
)

const (
	DefaultAP = 3
	DefaultHP = 200
)

type UnitKind byte

const (
	Elf    UnitKind = 'E'
	Goblin UnitKind = 'G'
)

type Unit struct {
	Kind  UnitKind
	Point Point
	Cell  *Cell

	AP int64
	HP int64
}

// Tick returns true if the battle is over, and a unit that died, if any.
func (u *Unit) Tick(arena *Arena) (bool, *Unit) {
	enemies := u.findEnemies(arena.Units)
	if len(enemies) == 0 {
		return true, nil
	}

	if adjacent := u.adjacentEnemies(enemies); len(adjacent) > 0 {
		return false, u.attack(adjacent)
	}

	u.move(arena, enemies)

	if adjacent := u.adjacentEnemies(enemies); len(adjacent) > 0 {
		return false, u.attack(adjacent)
	}

	return false, nil
}

func (u *Unit) String() string {
	return string(u.Kind)
}

func (u *Unit) findEnemies(units []*Unit) []*Unit {
	var enemies []*Unit

	for _, unit := range units {
		if unit.Kind != u.Kind {
			enemies = append(enemies, unit)
		}
	}

	return enemies
}

func (u *Unit) adjacentEnemies(enemies []*Unit) []*Unit {
	var adjacent []*Unit

	for _, enemy := range enemies {
		if Adjacent(u.Point, enemy.Point) {
			adjacent = append(adjacent, enemy)
		}
	}

	return adjacent
}

// attach returns a unit that died, if any.
func (u *Unit) attack(targets []*Unit) *Unit {
	sort.Slice(targets, func(i, j int) bool {
		a := targets[i]
		b := targets[j]

		if a.HP < b.HP {
			return true
		}
		if a.HP == b.HP {
			return Less(a.Point, b.Point)
		}
		return false
	})

	target := targets[0]

	target.HP -= u.AP
	// log.Println(u, u.Point, "attacks", target, target.Point, "HP:", target.HP)

	if target.HP <= 0 {
		return target
	}

	return nil
}

func (u *Unit) targetCells(arena *Arena, targets []*Unit) []*Cell {
	var cells []*Cell

	for _, target := range targets {
		for _, cell := range target.Cell.EmptyNeighbors(arena) {
			cells = append(cells, cell)
		}
	}

	return cells
}

func (u *Unit) move(arena *Arena, enemies []*Unit) {
	targetCells := u.targetCells(arena, enemies)

	type cellDistance struct {
		cell     *Cell
		distance uint64
	}

	var reachable []cellDistance
	var minDistance uint64 = math.MaxUint64

	for _, cell := range targetCells {
		distance, next := AStar(arena, u.Cell, cell)
		if next != nil {
			reachable = append(reachable, cellDistance{next, distance})

			if distance < minDistance {
				minDistance = distance
			}
		}
	}

	if len(reachable) == 0 {
		return
	}

	var nextCells []*Cell
	for _, r := range reachable {
		if r.distance == minDistance {
			nextCells = append(nextCells, r.cell)
		}
	}

	sort.Slice(nextCells, func(i, j int) bool {
		return Less(nextCells[i].Point, nextCells[j].Point)
	})

	// log.Println(u, u.Point, "moves to", nextCells[0].Point)
	nextCells[0].Enter(u)
}
