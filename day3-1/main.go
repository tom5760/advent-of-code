package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type fabric map[uint]map[uint]uint

func (f fabric) claim(x, y uint) {
	col := f[x]
	if col == nil {
		col = make(map[uint]uint)
		f[x] = col
	}

	col[y]++
}

func (f fabric) claimArea(x, y, width, height uint) {
	for i := x; i < x+width; i++ {
		for j := y; j < y+height; j++ {
			f.claim(i, j)
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

		f.claimArea(x, y, width, height)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln("err scanning input:", err)
	}

	var multipleClaims uint

	for _, col := range f {
		for _, numClaims := range col {
			if numClaims >= 2 {
				multipleClaims++
			}
		}
	}

	log.Println(multipleClaims, "square inches are within two or more claims")
}
