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
	Kind UnitKind
	Cell *Cell

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
			if cell.Unit == nil || (cell.Unit != nil && cell.Unit == u) {
				cells = append(cells, cell)
			}
		}
	}
	return cells
}

// Copied mostly from: https://www.reddit.com/r/adventofcode/comments/a6chwa/2018_day_15_solutions/ebtwcqr/
func (u *Unit) move(arena *Arena, targetCells []*Cell) {
	type data struct {
		cell *Cell
		dist uint64
	}

	visiting := []data{{u.Cell, 0}}
	meta := map[*Cell]data{u.Cell: data{nil, 0}}
	seen := make(map[*Cell]bool)

	for len(visiting) > 0 {
		cur := visiting[0]
		visiting = visiting[1:]

	neighborLoop:
		for _, neighbor := range cur.cell.EmptyNeighbors(arena) {
			neighborData, ok := meta[neighbor]
			if !ok {
				meta[neighbor] = data{cur.cell, cur.dist + 1}
			} else if cur.dist+1 < neighborData.dist {
				meta[neighbor] = data{cur.cell, cur.dist + 1}
			} else if cur.dist+1 == neighborData.dist {
				if Less(cur.cell.Point, neighborData.cell.Point) {
					meta[neighbor] = data{cur.cell, cur.dist + 1}
				}
			}

			if !seen[neighbor] {
				for _, vd := range visiting {
					if vd.cell == neighbor {
						continue neighborLoop
					}
				}
				visiting = append(visiting, data{neighbor, cur.dist + 1})
			}
		}

		seen[cur.cell] = true
	}

	var minDist uint64 = math.MaxUint64
	for _, targetCell := range targetCells {
		m, ok := meta[targetCell]
		if ok && m.dist < minDist {
			minDist = m.dist
		}
	}

	var minTargets []*Cell
	for _, targetCell := range targetCells {
		m, ok := meta[targetCell]
		if ok && m.dist == minDist {
			minTargets = append(minTargets, targetCell)
		}
	}

	if len(minTargets) == 0 {
		return
	}

	sort.Slice(minTargets, func(i, j int) bool {
		return Less(minTargets[i].Point, minTargets[j].Point)
	})

	closest := minTargets[0]

	for meta[closest].dist > 1 {
		closest = meta[closest].cell
	}

	closest.Enter(u)
}

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
			return Less(a.Cell.Point, b.Cell.Point)
		}
		return false
	})

	adjacentEnemies[0].damage(u.AP)
}

func (u *Unit) adjacentEnemies(enemies []*Unit) []*Unit {
	var adjacent []*Unit

	for _, enemy := range enemies {
		if Adjacent(u.Cell.Point, enemy.Cell.Point) {
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
