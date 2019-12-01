package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

type Point struct {
	X, Y uint64
}

type Claim struct {
	ID    uint64
	Point Point

	Width, Height uint64

	overlapCount int
}

type Fabric map[Point][]*Claim

func (f Fabric) claim(c *Claim, x, y uint64) {
	point := Point{x, y}

	claims := append(f[point], c)
	f[point] = claims

	if len(claims) > 1 {
		for _, other := range claims {
			other.overlapCount++
		}
	}
}

func (f Fabric) claimArea(claim *Claim) {
	topLeft := claim.Point

	for i := topLeft.X; i < topLeft.X+claim.Width; i++ {
		for j := topLeft.Y; j < topLeft.Y+claim.Height; j++ {
			f.claim(claim, i, j)
		}
	}
}

func main() {
	claims := readInput(os.Stdin)
	fabric := make(Fabric)

	for _, claim := range claims {
		fabric.claimArea(claim)
	}

	log.Println("(part 1) area with multiple claims:", findOverlappedArea(fabric))
	log.Println("(part 2) ID of single claim:", findSingleClaim(fabric).ID)
}

const claimFormat = "#%d @ %d,%d: %dx%d"

func readInput(r io.Reader) []*Claim {
	var claims []*Claim

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		var claim Claim

		_, err := fmt.Sscanf(scanner.Text(), claimFormat,
			&claim.ID, &claim.Point.X, &claim.Point.Y, &claim.Width, &claim.Height)
		if err != nil {
			log.Fatalln("err parsing input:", err)
			return nil
		}

		claims = append(claims, &claim)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln("error reading input:", err)
		return nil
	}

	return claims
}

func findOverlappedArea(fabric Fabric) uint64 {
	var area uint64

	for _, claims := range fabric {
		if len(claims) > 1 {
			area++
		}
	}

	return area
}

func findSingleClaim(fabric Fabric) *Claim {
	for _, claims := range fabric {
		if len(claims) > 1 {
			continue
		}
		for _, claim := range claims {
			if claim.overlapCount == 0 {
				return claim
			}
		}
	}
	panic("no single claim")
}
