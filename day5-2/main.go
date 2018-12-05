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
	var polymer []byte

	polymer = readInput(os.Stdin)

	shortestLength := math.MaxInt64

	for char := 'A'; char <= 'Z'; char++ {
		filtered := filterPolymer(polymer, char)
		filtered = fullReact(filtered)
		if len(filtered) < shortestLength {
			shortestLength = len(filtered)
		}
	}

	log.Println("length of shortest polymer:", shortestLength)
}

func readInput(r io.Reader) []byte {
	polymer, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatalln("failed to read input:", err)
		return nil
	}

	return bytes.TrimSpace(polymer)
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

func fullReact(polymer []byte) []byte {
	var lastLen int

	for lastLen != len(polymer) {
		lastLen = len(polymer)
		polymer = react(polymer)
	}

	return polymer
}

func react(polymer []byte) []byte {
	for i := 0; i < len(polymer)-1; i++ {
		a := polymer[i]
		b := polymer[i+1]

		if a-b == 32 || b-a == 32 {
			return append(polymer[:i], polymer[i+2:]...)
		}
	}

	return polymer
}
