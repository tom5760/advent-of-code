package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {
	if err := run(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func run() error {
	buf, err := os.ReadFile("./day06/input")
	if err != nil {
		return fmt.Errorf("failed to open input file: %w", err)
	}

	buf = bytes.TrimSpace(buf)

	fmt.Println("PART 1:", FindMarker(buf, 4))
	fmt.Println("PART 2:", FindMarker(buf, 14))

	return nil
}

func FindMarker(buf []byte, markerLen int) int {
	marker := make([]byte, markerLen)

	for i, c := range buf {
		marker[i%markerLen] = c

		if i > markerLen && unique(marker) {
			return i + 1
		}
	}

	panic("no marker found")
}

func unique(buf []byte) bool {
	for i, x := range buf {
		for _, y := range buf[i+1:] {
			if x == y {
				return false
			}
		}
	}

	return true
}
