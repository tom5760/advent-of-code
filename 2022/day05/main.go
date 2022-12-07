package day05

import (
	"bufio"
	"bytes"
	"fmt"
	"strconv"
	"strings"

	"github.com/tom5760/advent-of-code/2022/inpututils"
)

func Parse(name string) (Problem, error) {
	var problem Problem

	err := inpututils.Scan(name, func(scanner *bufio.Scanner) error {
		// Scan initial crate stacks
		for scanner.Scan() {
			line := scanner.Bytes()

			n := len(line)
			if n == 0 {
				break
			}

			for i, col := 0, 0; i < n; {
				if len(problem.stacks) <= col {
					problem.stacks = append(problem.stacks, nil)
				}

				if line[i] == '[' {
					problem.stacks[col] = append([]Crate{Crate(line[i+1])}, problem.stacks[col]...)
				}

				col++
				i += 4
			}
		}

		// Scan instructions
		for scanner.Scan() {
			line := scanner.Bytes()

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

			problem.instructions = append(problem.instructions, Instruction{
				Count:       count,
				Source:      source,
				Destination: destination,
			})
		}

		return nil
	})

	return problem, err
}

func Part1(problem Problem) string {
	stacks := problem.stacks.Clone()

	for _, instruction := range problem.instructions {
		stacks.ExecuteV1(instruction)
	}

	var sb strings.Builder

	for _, stack := range stacks {
		sb.WriteByte(byte(stack[len(stack)-1]))
	}

	return sb.String()
}

func Part2(problem Problem) string {
	stacks := problem.stacks.Clone()

	for _, instruction := range problem.instructions {
		stacks.ExecuteV2(instruction)
	}

	var sb strings.Builder

	for _, stack := range stacks {
		sb.WriteByte(byte(stack[len(stack)-1]))
	}

	return sb.String()
}

type (
	Problem struct {
		stacks       Stacks
		instructions []Instruction
	}

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
