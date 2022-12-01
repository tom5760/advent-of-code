package main

import (
	"bufio"
	"os"
	"strconv"
	"testing"
)

func TestPart1(t *testing.T) {
	f, err := os.Open("./input")
	if err != nil {
		t.Fatal("failed to open input", err)
		return
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var (
		elves [][]uint64
		elf   []uint64
	)

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			elves = append(elves, elf)
			elf = []uint64{}

			continue
		}

		calories, err := strconv.ParseUint(scanner.Text(), 10, 0)
		if err != nil {
			t.Fatal("failed to parse input", err)
			return
		}

		elf = append(elf, calories)
	}

	if err := scanner.Err(); err != nil {
		t.Fatal("failed to scan input", err)
		return
	}

	var maxelf uint64
	for _, elf := range elves {
		var total uint64

		for _, calories := range elf {
			total += calories
		}

		if total > maxelf {
			maxelf = total
		}
	}

	const want = 69289

	if maxelf != want {
		t.Fatalf("got: %v, want: %v", maxelf, want)
	}
}
