package main

import (
	"log"
	"strconv"
	"strings"
)

const (
	AfterCount = 10
)

var TotalRecipes = []int{5, 9, 18, 2018, 409551}

type ScoreBoard struct {
	Board []int
	Elves [2]int
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

func (b *ScoreBoard) get(i int) int {
	return b.Board[b.Elves[i]]
}

func (b *ScoreBoard) makeRecipes() {
	sum := b.get(0) + b.get(1)

	tensDigit := sum / 10
	onesDigit := sum % 10

	if tensDigit != 0 {
		b.Board = append(b.Board, tensDigit)
	}
	b.Board = append(b.Board, onesDigit)
}

func (b *ScoreBoard) selectRecipe() {
	b.Elves[0] = (b.Elves[0] + b.get(0) + 1) % len(b.Board)
	b.Elves[1] = (b.Elves[1] + b.get(1) + 1) % len(b.Board)
}

func main() {
	for _, total := range TotalRecipes {
		board := ScoreBoard{
			Board: []int{3, 7},
			Elves: [...]int{0, 1},
		}

		// log.Println(board.Board)

		for len(board.Board) < total+AfterCount {
			board.Tick()
			// log.Println(board.Board)
		}

		log.Printf("(part 1) ten recipes after %d recipes: %s", total, board.Digits(total, AfterCount))
	}
}
