package main

import (
	"log"
	"strconv"
	"strings"
)

const (
	AfterCount   = 10
	TotalRecipes = 409551
)

var Search = []int{4, 0, 9, 5, 5, 1}

type ScoreBoard struct {
	Board []int
	Elves [2]int

	Search  []int
	EndI    int
	searchI int
}

func (b *ScoreBoard) Digits(total, after int) string {
	var sb strings.Builder
	for _, x := range b.Board[total : total+10] {
		sb.WriteString(strconv.Itoa(x))
	}
	return sb.String()
}

func (b *ScoreBoard) Tick() {
	b.makeRecipes()
	b.selectRecipe()
}

func (b *ScoreBoard) Check() bool {
	return b.searchI == len(b.Search)
}

func (b *ScoreBoard) get(i int) int {
	return b.Board[b.Elves[i]]
}

func (b *ScoreBoard) makeRecipes() {
	sum := b.get(0) + b.get(1)

	tensDigit := sum / 10
	onesDigit := sum % 10

	if tensDigit != 0 {
		b.Board = append(b.Board, tensDigit)
		b.search(tensDigit)
	}

	b.Board = append(b.Board, onesDigit)
	b.search(onesDigit)
}

func (b *ScoreBoard) search(digit int) {
	if b.searchI >= len(b.Search) {
		return
	}

	if b.Search[b.searchI] == digit {
		b.searchI++
	} else {
		if b.searchI != 0 {
			b.searchI = 0
			b.search(digit)
		}
	}

	if b.searchI >= len(b.Search) {
		b.EndI = len(b.Board) - len(b.Search)
	}
}

func (b *ScoreBoard) selectRecipe() {
	b.Elves[0] = (b.Elves[0] + b.get(0) + 1) % len(b.Board)
	b.Elves[1] = (b.Elves[1] + b.get(1) + 1) % len(b.Board)
}

func main() {
	board := ScoreBoard{
		Board: []int{3, 7},
		Elves: [...]int{0, 1},
	}

	for len(board.Board) < TotalRecipes+AfterCount {
		board.Tick()
	}

	log.Printf("(part 1) ten recipes after %d recipes: %s", TotalRecipes, board.Digits(TotalRecipes, AfterCount))

	board = ScoreBoard{
		Board:  []int{3, 7},
		Elves:  [...]int{0, 1},
		Search: Search,
	}

	for !board.Check() {
		board.Tick()
	}

	log.Println("(part 2) recipie count:", len(board.Board[:board.EndI]))
}
