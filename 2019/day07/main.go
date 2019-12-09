package main

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/tom5760/advent-of-code/2019/common"
	"github.com/tom5760/advent-of-code/2019/intcode"
)

const (
	numAmplifiers = 5
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

	log.Println("(part 1) highest thruster signal:", part1)

	part2, err := Part2(program)
	if err != nil {
		log.Println("failed to run part 2:", err)
		return 1
	}

	log.Println("(part 2) highest thruster signal:", part2)

	return 0
}

// Part1 - Try every combination of phase settings on the amplifiers. What is
// the highest signal that can be sent to the thrusters?
func Part1(program []int) (int, error) {
	phaseInput := []int{0, 1, 2, 3, 4}
	maxOutput := 0

	var err error

	common.PermutationsInt(phaseInput, func(phases []int) {
		lastOutput := 0

		for i := 0; i < numAmplifiers; i++ {
			computer := intcode.NewComputer(program)

			computer.Inputs(phases[i], lastOutput)
			computer.OutputChan = make(chan int, 1)

			go computer.Run()

			outputs := computer.Outputs()

			if err != nil {
				return
			}

			if len(outputs) != 1 {
				err = fmt.Errorf("unexpected number of outputs")
				return
			}

			lastOutput = outputs[0]
		}

		if lastOutput > maxOutput {
			maxOutput = lastOutput
		}
	})

	return maxOutput, err
}

// Part2 - Try every combination of the new phase settings on the amplifier
// feedback loop. What is the highest signal that can be sent to the thrusters?
func Part2(program []int) (int, error) {
	phaseInput := []int{5, 6, 7, 8, 9}
	maxOutput := 0

	var err error

	common.PermutationsInt(phaseInput, func(phases []int) {
		computers := make([]*intcode.Computer, numAmplifiers)

		// create computers
		for i := range computers {
			computers[i] = intcode.NewComputer(program)
		}

		// hook up inputs to previous computer's output
		for i, cmp := range computers {
			if i == 0 {
				i = len(computers)
			}
			cmp.InputChan = computers[i-1].OutputChan
		}

		// set the initial phase input
		for i, p := range phases {
			computers[i].InputChan <- p
		}

		// first computer gets initial 0 input
		computers[0].InputChan <- 0

		var wg sync.WaitGroup

		for i, cmp := range computers {
			wg.Add(1)
			go func(i int, c *intcode.Computer) {
				c.Run()
				wg.Done()
			}(i, cmp)
		}

		wg.Wait()

		lastOutputs := computers[len(computers)-1].Outputs()
		if len(lastOutputs) != 1 {
			err = fmt.Errorf("unexpected number of outputs")
			return
		}

		if lastOutputs[0] > maxOutput {
			maxOutput = lastOutputs[0]
		}
	})

	return maxOutput, err
}
