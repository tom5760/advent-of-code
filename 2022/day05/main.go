package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	if err := run(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func run() error {
	stacks, instructions, err := ParseInput()
	if err != nil {
		return fmt.Errorf("failed to parse input: %w", err)
	}

	part2Stacks := stacks.Clone()

	Part1(stacks, instructions)
	Part2(part2Stacks, instructions)

	return nil
}

func ParseInput() (Stacks, []Instruction, error) {
	f, err := os.Open("./day05/input")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open input file: %w", err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var stacks [][]Crate
	parseStacks := func(line []byte) error {
		n := len(line)
		if n == 0 {
			return nil
		}

		for i, col := 0, 0; i < n; {
			if len(stacks) <= col {
				stacks = append(stacks, nil)
			}

			if line[i] == '[' {
				stacks[col] = append([]Crate{Crate(line[i+1])}, stacks[col]...)
			}

			col++
			i += 4
		}

		return nil
	}

	var instructions []Instruction
	parseInstructions := func(line []byte) error {
		fields := bytes.Fields(line)
		if len(fields) != 6 {
			return fmt.Errorf("invalid instruction line format")
		}

		count, err := strconv.Atoi(string(fields[1]))
		if err != nil {
			return fmt.Errorf("failed to parse count: %w", err)
		}

		source, err := strconv.Atoi(string(fields[3]))
		if err != nil {
			return fmt.Errorf("failed to parse source: %w", err)
		}

		destination, err := strconv.Atoi(string(fields[5]))
		if err != nil {
			return fmt.Errorf("failed to parse destination: %w", err)
		}

		instructions = append(instructions, Instruction{
			Count:       count,
			Source:      source,
			Destination: destination,
		})

		return nil
	}

	scanFunc := parseStacks
	for scanner.Scan() {
		line := scanner.Bytes()
		if len(line) == 0 {
			scanFunc = parseInstructions
			continue
		}

		if err := scanFunc(line); err != nil {
			return nil, nil, err
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return stacks, instructions, nil
}

func Part1(stacks Stacks, instructions []Instruction) {
	for _, instruction := range instructions {
		stacks.ExecuteV1(instruction)
	}

	fmt.Print("PART 1: ")

	for _, stack := range stacks {
		fmt.Print(stack[len(stack)-1])
	}

	fmt.Print("\n")
}

func Part2(stacks Stacks, instructions []Instruction) {
	for _, instruction := range instructions {
		stacks.ExecuteV2(instruction)
	}

	fmt.Print("PART 2: ")

	for _, stack := range stacks {
		fmt.Print(stack[len(stack)-1])
	}

	fmt.Print("\n")
}

type (
	Stacks [][]Crate

	Crate byte

	Instruction struct {
		Count       int
		Source      int
		Destination int
	}
)

func (c Crate) String() string {
	return string(c)
}

func (i Instruction) String() string {
	return fmt.Sprintf("move %v from %v to %v", i.Count, i.Source, i.Destination)
}

func (s Stacks) Clone() Stacks {
	ns := make(Stacks, len(s))
	for i, stack := range s {
		ns[i] = make([]Crate, len(stack))
		copy(ns[i], stack)
	}

	return ns
}

func (s Stacks) ExecuteV1(inst Instruction) {
	source := s[inst.Source-1]
	dest := s[inst.Destination-1]

	for i := 0; i < inst.Count; i++ {
		n := len(source) - 1
		crate := source[n]
		source = source[:n]
		dest = append(dest, crate)
	}

	s[inst.Source-1] = source
	s[inst.Destination-1] = dest
}

func (s Stacks) ExecuteV2(inst Instruction) {
	source := s[inst.Source-1]
	dest := s[inst.Destination-1]

	n := len(source) - inst.Count
	moved := source[n:]

	source = source[:n]
	dest = append(dest, moved...)

	s[inst.Source-1] = source
	s[inst.Destination-1] = dest
}
