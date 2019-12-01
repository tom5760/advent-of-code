package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"math"
	"os"
)

func main() {
	polymer := readInput(os.Stdin)

	log.Println("(part 1) length of reacted polymer:", len(copyReact(polymer)))
	log.Println("(part 2) length of shortest polymer:", shortestReact(polymer))
}

func readInput(r io.Reader) []byte {
	polymer, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatalln("failed to read input:", err)
		return nil
	}

	return bytes.TrimSpace(polymer)
}

// copyReact makes a copy, doesn't modify
func copyReact(polymer []byte) []byte {
	c := make([]byte, len(polymer))
	copy(c, polymer)
	return react(c)
}

func shortestReact(polymer []byte) int {
	shortestLength := math.MaxInt64

	for char := 'A'; char <= 'Z'; char++ {
		filtered := react(filterPolymer(polymer, char))
		if len(filtered) < shortestLength {
			shortestLength = len(filtered)
		}
	}

	return shortestLength
}

func filterPolymer(polymer []byte, char rune) []byte {
	return bytes.Map(func(r rune) rune {
		switch r {
		case char, char + 32:
			return -1
		default:
			return r
		}
	}, polymer)
}

func react(polymer []byte) []byte {
	i := 0
	for i < len(polymer)-1 {
		a := polymer[i]
		b := polymer[i+1]
		d := int(a) - int(b)

		if d == 32 || d == -32 {
			polymer = append(polymer[:i], polymer[i+2:]...)
			if i > 0 {
				i--
			}
		} else {
			i++
		}
	}

	return polymer
}
