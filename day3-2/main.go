package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type claim struct {
	id, x, y, width, height uint

	overlapCount int
}

type fabric map[uint]map[uint][]*claim

func (f fabric) claim(c *claim, x, y uint) {
	col := f[x]
	if col == nil {
		col = make(map[uint][]*claim)
		f[x] = col
	}

	col[y] = append(col[y], c)

	if len(col[y]) > 1 {
		for _, claim := range col[y] {
			claim.overlapCount++
		}
	}
}

func (f fabric) claimArea(id, x, y, width, height uint) {
	claim := &claim{id: id, x: x, y: y, width: width, height: height}

	for i := x; i < x+width; i++ {
		for j := y; j < y+height; j++ {
			f.claim(claim, i, j)
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	f := make(fabric)

	for scanner.Scan() {
		var id, x, y, width, height uint
		if _, err := fmt.Sscanf(scanner.Text(), "#%d @ %d,%d: %dx%d", &id, &x, &y, &width, &height); err != nil {
			log.Fatalln("err reading input line:", err)
			return
		}

		f.claimArea(id, x, y, width, height)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln("err scanning input:", err)
	}

	for _, col := range f {
		for _, claims := range col {
			if len(claims) > 1 {
				continue
			}

			for _, claim := range claims {
				if claim.overlapCount == 0 {
					log.Println("id of single claim", claim.id)
					return
				}
			}
		}
	}

	log.Println("no single claims?")
}
