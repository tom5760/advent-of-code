package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/tom5760/advent-of-code/2019/common"
)

func main() {
	os.Exit(run())
}

func run() int {
	db, err := readInput(os.Stdin)
	if err != nil {
		log.Println(err)
		return 1
	}

	log.Println("(part 1) total ORE:", Part1(db))

	return 0
}

func readInput(r io.Reader) (*ChemDB, error) {
	lines, err := common.ReadStringSlice(r, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to read input: %w", err)
	}

	return Parse(lines)
}

// Part1 - Given the list of reactions in your puzzle input, what is the
// minimum amount of ORE required to produce exactly 1 FUEL?
func Part1(db *ChemDB) int {
	return db.Totals(chemStart)
}
