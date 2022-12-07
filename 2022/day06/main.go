package day06

import (
	"bytes"
	"os"

	"github.com/tom5760/advent-of-code/2022/testutils"
)

func Parse(name string) ([]byte, error) {
	buf, err := os.ReadFile(name)

	return bytes.TrimSpace(buf), err
}

func Part1(buf []byte) int {
	return FindMarker(buf, 4)
}

func Part2(buf []byte) int {
	return FindMarker(buf, 14)
}

func FindMarker(buf []byte, markerLen int) int {
	for i := range buf {
		marker := buf[i : i+markerLen]

		testutils.GT.Log(string(marker))

		if unique(marker) {
			return i + markerLen
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
