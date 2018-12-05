package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	polymer := readInput(os.Stdin)

	var lastLen int

	for lastLen != len(polymer) {
		lastLen = len(polymer)
		polymer = react(polymer)
	}

	log.Println("length of polymer:", len(polymer))
}

func readInput(r io.Reader) []byte {
	polymer, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatalln("failed to read input:", err)
		return nil
	}

	return bytes.TrimSpace(polymer)
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
