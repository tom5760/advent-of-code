package main

import (
	"fmt"
	"log"
	"os"

	"github.com/tom5760/advent-of-code/2019/common"
	"github.com/tom5760/advent-of-code/2019/intcode"
)

const (
	maxValue     = 100
	targetOutput = 19690720
)

func main() {
	os.Exit(run())
}

func run() int {
	memory, err := common.ReadIntSlice(os.Stdin, common.ScanCommas)
	if err != nil {
		log.Println("failed to read input:", err)
		return 1
	}

	log.Println("part 1:", Part1(memory))

	part2, err := Part2(memory)
	if err != nil {
		log.Println("failed to run part 2:", err)
		return 1
	}

	log.Println("part 2:", part2)

	return 0
}

// RunComputer runs the given program with arguments.
func RunComputer(memory []int, noun, verb int) int {
	computer := intcode.NewComputer(memory)

	computer.Memory[1] = noun
	computer.Memory[2] = verb

	computer.Run()

	return computer.Memory[0]
}

// Part1 - To do this, before running the program, replace position 1
// with the value 12 and replace position 2 with the value 2. What value is
// left at position 0 after the program halts?
func Part1(memory []int) int {
	return RunComputer(memory, 12, 2)
}

// Part2 - Find the input noun and verb that cause the program to produce
// the output 19690720. What is 100 * noun + verb?
func Part2(memory []int) (int, error) {
	for noun := 0; noun < maxValue; noun++ {
		for verb := 0; verb < maxValue; verb++ {
			if RunComputer(memory, noun, verb) == targetOutput {
				return 100*noun + verb, nil
			}
		}
	}

	return 0, fmt.Errorf("solution not found")
}
