package main

import (
	"fmt"
	"log"
	"os"

	"github.com/tom5760/advent-of-code/2019/common"
	"github.com/tom5760/advent-of-code/2019/intcode"
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

	part1, err := Part1(memory)
	if err != nil {
		log.Println("failed to run part 1:", err)
		return 1
	}

	log.Println("(part 1) diagnostic code:", part1)

	part2, err := Part2(memory)
	if err != nil {
		log.Println("failed to run part 2:", err)
		return 1
	}

	log.Println("(part 2) diagnostic code:", part2)

	return 0
}

// RunComputer runs the given program with inputs.
func RunComputer(memory []int, input []int) []int {
	computer := intcode.NewComputer(memory)

	//computer.Log = true
	computer.Inputs(input...)

	go computer.Run()

	return computer.Outputs()
}

// Part1 - After providing 1 to the only input instruction and passing all the
// tests, what diagnostic code does the program produce?
func Part1(memory []int) (int, error) {
	outputs := RunComputer(memory, []int{1})

	if len(outputs) == 0 {
		return 0, fmt.Errorf("produced no outputs")
	}

	for i, output := range outputs[:len(outputs)-1] {
		if output != 0 {
			return 0, fmt.Errorf("output %d returned non-zero (%d)", i, output)
		}
	}

	return outputs[len(outputs)-1], nil
}

// Part2 - What is the diagnostic code for system ID 5?
func Part2(memory []int) (int, error) {
	outputs := RunComputer(memory, []int{5})

	if len(outputs) != 1 {
		return 0, fmt.Errorf("unexpected number of outputs")
	}

	return outputs[0], nil
}
