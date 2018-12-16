package main

import (
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

// Turn returns true if the battle is over.
func (u *Unit) Turn(arena *Arena) bool {
	if u.HP <= 0 {
		return false
	}

	enemies := u.findEnemies(arena.Units)
	if len(enemies) == 0 {
		return true
	}

	targetCells := u.findTargetCells(arena, enemies)
	if len(targetCells) == 0 {
		return false
	}

	u.move(arena, targetCells)
	u.attack(enemies)

	return false
}

func (u *Unit) String() string {
	return string(u.Kind)
}

func (u *Unit) findEnemies(units []*Unit) []*Unit {
	var enemies []*Unit

	for _, unit := range units {
		if unit.Kind != u.Kind && unit.HP > 0 {
			enemies = append(enemies, unit)
		}
	}

	return enemies
}

func (u *Unit) findTargetCells(arena *Arena, enemies []*Unit) []*Cell {
	var cells []*Cell
	for _, enemy := range enemies {
		for _, cell := range enemy.Cell.FloorNeighbors(arena) {
			if cell.Unit == nil || cell.Unit == u {
				cells = append(cells, cell)
			}
		}
	}
	return cells
}

func (u *Unit) move(arena *Arena, targetCells []*Cell) {
	type nextCell struct {
		next     *Cell
		distance uint64
	}

	var nextCells []nextCell
	for _, targetCell := range targetCells {
		cell, distance := NextCell(arena, u.Cell, targetCell)
		if cell == nil {
			continue
		}
		nextCells = append(nextCells, nextCell{
			next:     cell,
			distance: distance,
		})
	}

	if len(nextCells) == 0 {
		return
	}

	sort.Slice(nextCells, func(i, j int) bool {
		a := nextCells[i]
		b := nextCells[j]

		if a.distance < b.distance {
			return true
		}
		if a.distance > b.distance {
			return false
		}
		return Less(a.next.Point, b.next.Point)
	})

	// log.Println("moving from", u.Cell.Point, "->", nextCells[0].next.Point)
	nextCells[0].next.Enter(u)
}

// func (u *Unit) move(arena *Arena, targetCells []*Cell) {
// 	type reachableCell struct {
// 		dist   uint64
// 		next   *Cell
// 		target *Cell
// 	}

// 	var reachable []reachableCell
// 	var minDistance uint64 = math.MaxUint64

// 	for _, targetCell := range targetCells {
// 		r := reachableCell{target: targetCell}
// 		r.dist, r.next = AStar(arena, u.Cell, targetCell)
// 		if r.next != nil {
// 			reachable = append(reachable, r)
// 			if r.dist < minDistance {
// 				minDistance = r.dist
// 			}
// 		}
// 	}

// 	if len(reachable) == 0 {
// 		return
// 	}

// 	var minimum []reachableCell
// 	for _, r := range reachable {
// 		if r.dist == minDistance {
// 			// log.Println("minimum reachable", r.next.Point)
// 			minimum = append(minimum, r)
// 		}
// 	}

// 	sort.Slice(minimum, func(i, j int) bool {
// 		return Less(reachable[i].target.Point, reachable[j].target.Point)
// 	})

// 	nextCell := minimum[0].next
// 	// log.Println("moving to ", nextCell.Point)

// 	if u.Cell != nextCell {
// 		nextCell.Enter(u)
// 	}
// }

func (u *Unit) attack(enemies []*Unit) {
	adjacentEnemies := u.adjacentEnemies(enemies)

	if len(adjacentEnemies) == 0 {
		return
	}

	sort.Slice(adjacentEnemies, func(i, j int) bool {
		a := adjacentEnemies[i]
		b := adjacentEnemies[j]

		if a.HP < b.HP {
			return true
		}
		if a.HP == b.HP {
			return Less(a.Point, b.Point)
		}
		return false
	})

	adjacentEnemies[0].damage(u.AP)
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

func (u *Unit) damage(d int64) {
	u.HP -= d
	if u.HP <= 0 {
		u.HP = 0
		u.kill()
	}
}

func (u *Unit) kill() {
	if u.Cell != nil {
		u.Cell.Leave(u)
	}
}
