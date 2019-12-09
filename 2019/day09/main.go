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
	program, err := common.ReadIntSlice(os.Stdin, common.ScanCommas)
	if err != nil {
		log.Println("failed to read input:", err)
		return 1
	}

	part1, err := Part1(program)
	if err != nil {
		log.Println("failed to run part 1:", err)
		return 1
	}

	log.Println("(part 1) BOOST keycode:", part1)

	part2, err := Part2(program)
	if err != nil {
		log.Println("failed to run part 2:", err)
		return 1
	}

	log.Println("(part 2) coordinates:", part2)

	return 0
}

// Part1 - Once your Intcode computer is fully functional, the BOOST program
// should report no malfunctioning opcodes when run in test mode; it should
// only output a single value, the BOOST keycode. What BOOST keycode does it
// produce?
func Part1(program []int) (int, error) {
	computer := intcode.NewComputer(program)

	computer.Inputs(1)
	computer.Run()

	outputs := computer.Outputs()

	if len(outputs) != 1 {
		return 0, fmt.Errorf("unexpected number of outputs")
	}

	return outputs[0], nil
}

// Part2 - Run the BOOST program in sensor boost mode. What are the coordinates
// of the distress signal?
func Part2(program []int) (int, error) {
	computer := intcode.NewComputer(program)

	computer.Inputs(2)
	computer.Run()

	outputs := computer.Outputs()

	if len(outputs) != 1 {
		log.Println(outputs)
		return 0, fmt.Errorf("unexpected number of outputs")
	}

	return outputs[0], nil
}
