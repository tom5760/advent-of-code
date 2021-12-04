package main

import (
	"fmt"
	"math"
	"strings"
	"text/tabwriter"

	"golang.org/x/exp/slices"
)

type (
	Game struct {
		Numbers []int64

		Boards []Board
	}

	Board struct {
		Numbers []int64
		Marked  []bool
	}

	Win struct {
		Index  int
		Board  *Board
		Number int64
	}
)

func (g *Game) Play() []Win {
	var wins []Win

	for _, n := range g.Numbers {
		for i := range g.Boards {
			board := &g.Boards[i]

			if board.Won() {
				continue
			}

			if board.Mark(n) {
				wins = append(wins, Win{
					Index:  i,
					Board:  board,
					Number: n,
				})
			}
		}
	}

	return wins
}

func (g *Game) String() string {
	var sb strings.Builder

	sb.WriteString("Numbers: ")

	for i, n := range g.Numbers {
		fmt.Fprintf(&sb, "%v", n)
		if i < len(g.Numbers)-1 {
			sb.Write([]byte{',', ' '})
		}
	}

	sb.WriteString("\nBoards:\n")

	for i, board := range g.Boards {
		sb.WriteString(board.String())
		if i < len(g.Numbers)-1 {
			sb.WriteByte('\n')
		}
	}

	return sb.String()
}

// Mark a number, return true if game is won.
func (b *Board) Mark(n int64) bool {
	i := slices.Index(b.Numbers, n)
	if i == -1 {
		return false
	}

	b.Marked[i] = true

	return b.Won()
}

func (b *Board) N() int {
	return int(math.Sqrt(float64(len(b.Numbers))))
}

const (
	tBold  = "\033[1m"
	tClear = "\033[0m"
)

func (b *Board) String() string {
	var sb strings.Builder

	w := tabwriter.NewWriter(&sb, 0, 0, 1, ' ', 0)
	n := b.N()

	for y := 0; y < n; y++ {
		for x := 0; x < n-1; x++ {
			fmt.Fprintf(w, "%v\t", b.xy(x, y, n))
		}

		fmt.Fprintf(w, "%v\n", b.xy((n-1), y, n))
	}

	w.Flush()

	return sb.String()
}

func (b *Board) xy(x, y, n int) string {
	xy := y*n + x
	mark := tClear
	unmark := tClear

	if b.Marked[xy] {
		mark = tBold
	}

	return fmt.Sprintf("%v%v%v\t", mark, b.Numbers[xy], unmark)
}

// Check if this board has won.
func (b *Board) Won() bool {
	n := b.N()

	return b.checkRows(n) || b.checkCols(n)
}

func (b *Board) UnmarkedSum() int64 {
	var sum int64
	for i, n := range b.Numbers {
		if !b.Marked[i] {
			sum += n
		}
	}
	return sum
}

func (b *Board) checkRows(n int) bool {
	for y := 0; y < n; y++ {
		if b.checkRow(n, y) {
			return true
		}
	}

	return false
}

func (b *Board) checkRow(n, y int) bool {
	for x := 0; x < n; x++ {
		if !b.Marked[y*n+x] {
			return false
		}
	}

	return true
}

func (b *Board) checkCols(n int) bool {
	for x := 0; x < n; x++ {
		if b.checkCol(n, x) {
			return true
		}
	}

	return false
}

func (b *Board) checkCol(n, x int) bool {
	for y := 0; y < n; y++ {
		if !b.Marked[y*n+x] {
			return false
		}
	}

	return true
}

//func (b *Board) checkDiag1(n int) bool {
//	for i := 0; i < n; i++ {
//		if !b.Marked[i*n+i] {
//			return false
//		}
//	}
//
//	return true
//}
//
//func (b *Board) checkDiag2(n int) bool {
//	for i := 0; i < n; i++ {
//		if !b.Marked[(n-i-1)*n+i] {
//			return false
//		}
//	}
//
//	return true
//}
