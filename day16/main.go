package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

type Registers [4]int

type Sample struct {
	Before      Registers
	Instruction Instruction
	After       Registers
	OpCode      OpCode
}

const (
	beforeFormat      = "Before: [%d, %d, %d, %d]"
	afterFormat       = "After: [%d, %d, %d, %d]"
	instructionFormat = "%d %d %d %d"
)

func readInput(r io.Reader) ([]Sample, []Instruction) {
	scanner := bufio.NewScanner(r)

	var samples []Sample
	var program []Instruction

	var sample Sample
	var instruction Instruction

	type stateFn func(x string) (stateFn, error)
	var scanBefore, scanInstruction, scanAfter stateFn

	scanBefore = func(x string) (stateFn, error) {
		_, err := fmt.Sscanf(x, beforeFormat,
			&sample.Before[0], &sample.Before[1], &sample.Before[2], &sample.Before[3])
		return scanInstruction, err
	}

	scanInstruction = func(x string) (stateFn, error) {
		_, err := fmt.Sscanf(x, instructionFormat,
			&instruction.OP, &instruction.A, &instruction.B, &instruction.C)
		return scanAfter, err
	}

	scanAfter = func(x string) (stateFn, error) {
		_, err := fmt.Sscanf(x, afterFormat,
			&sample.After[0], &sample.After[1], &sample.After[2], &sample.After[3])
		return scanBefore, err
	}

	state := scanBefore
	blankLines := 0

	for scanner.Scan() {
		var err error
		if blankLines < 3 {
			if len(scanner.Bytes()) == 0 {
				blankLines++
				if blankLines == 1 {
					sample.Instruction = instruction
					samples = append(samples, sample)
				}
			} else {
				blankLines = 0
				state, err = state(scanner.Text())
			}
		} else {
			_, err = scanInstruction(scanner.Text())
			program = append(program, instruction)
		}

		if err != nil {
			log.Fatalln("failed to parse input:", err, scanner.Text())
			return nil, nil
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln("failed to read input:", err)
		return nil, nil
	}

	return samples, program
}

func main() {
	samples, program := readInput(os.Stdin)

	log.Println("samples:", len(samples))

	threeCount := 0
	for _, sample := range samples {
		opcodes := Check(sample)
		if len(opcodes) >= 3 {
			threeCount++
		}
	}

	log.Println("(part 1) samples with 3 or more opcodes:", threeCount)

	isa := make(map[int]OpCode)

	for len(ISA) > 0 {
		for _, sample := range samples {
			opcodes := Check(sample)
			if len(opcodes) == 1 {
				isa[sample.Instruction.OP] = opcodes[0]
				ReduceISA(sample)
			}
		}
	}

	var registers Registers
	for _, instruction := range program {
		registers = isa[instruction.OP](instruction, registers)
	}

	log.Println("(part 2) register 0:", registers[0])
}
