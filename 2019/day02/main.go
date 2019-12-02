package main

import (
	"fmt"
	"log"
	"os"

	"github.com/tom5760/advent-of-code/2019/common"
)

const (
	maxValue     = 99
	targetOutput = 19690720
)

func main() {
	os.Exit(run())
}

func run() int {
	memory, err := common.ReadUint64Slice(os.Stdin, common.ScanCommas)
	if err != nil {
		log.Println("failed to read input:", err)
		return 1
	}

	part1, err := Part1(memory)
	if err != nil {
		log.Println("failed to run part 1:", err)
		return 1
	}

	log.Println("part 1:", part1)

	part2, err := Part2(memory)
	if err != nil {
		log.Println("failed to run part 2:", err)
		return 1
	}

	log.Println("part 2:", part2)

	return 0
}

func RunComputer(memory []uint64, noun, verb uint64) (uint64, error) {
	computer := NewComputer(memory)

	computer.Memory[1] = noun
	computer.Memory[2] = verb

	if err := computer.Run(); err != nil {
		return 0, fmt.Errorf("failed to run program: %w", err)
	}

	return computer.Memory[0], nil
}

// For part 1: To do this, before running the program, replace position 1
// with the value 12 and replace position 2 with the value 2. What value is
// left at position 0 after the program halts?
func Part1(memory []uint64) (uint64, error) {
	return RunComputer(memory, 12, 2)
}

// For part 2: Find the input noun and verb that cause the program to produce
// the output 19690720. What is 100 * noun + verb?
func Part2(memory []uint64) (uint64, error) {
	for noun := uint64(0); noun < 100; noun++ {
		for verb := uint64(0); verb < 100; verb++ {
			ret, err := RunComputer(memory, noun, verb)
			if err != nil {
				return 0, err
			}
			if ret == targetOutput {
				return 100*noun + verb, nil
			}
		}
	}

	return 0, fmt.Errorf("solution not found")
}
