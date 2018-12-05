package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	polymer := react(readInput(os.Stdin))
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
