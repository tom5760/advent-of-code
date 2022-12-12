// Package day10 implements the solution for Advent of Code 2022 day 10.
//
// See: https://adventofcode.com/2022/day/10
package day10

import (
	"bufio"
	"bytes"
	"fmt"
	"strconv"
	"strings"

	"github.com/tom5760/advent-of-code/2022/inpututils"
	"github.com/tom5760/advent-of-code/2022/structs"
)

func Parse(name string) (Program, error) {
	var program Program

	err := inpututils.Scan(name, func(scanner *bufio.Scanner) error {
		for scanner.Scan() {
			fields := bytes.Fields(scanner.Bytes())
			op := string(fields[0])

			var (
				instruction Instruction
				err         error
			)

			switch op {
			case "addx":
				var arg int
				arg, err = strconv.Atoi(string(fields[1]))
				instruction = &OpAddX{Arg: arg}

			case "noop":
				instruction = &OpNoop{}
			default:
				return fmt.Errorf("unknown instruction %q", op)
			}

			if err != nil {
				return fmt.Errorf("failed to parse instruction %q: %w", op, err)
			}

			program = append(program, instruction)
		}

		return nil
	})

	return program, err
}

func Part1(program Program) int {
	computer := Computer{X: 1}

	var signal int

	for _, instruction := range program {
		computer.Execute(instruction, func(cycle int) {
			if (computer.Cycle-20)%40 == 0 {
				strength := computer.Cycle * computer.X
				signal += strength
			}
		})
	}

	return signal
}

func Part2(program Program) string {
	computer := Computer{X: 1}

	const (
		width  = 40
		height = 6
	)

	crt := structs.Grid[byte]{
		Width:  width,
		Height: height,
		Values: make([]byte, width*height),
	}

	for _, instruction := range program {
		computer.Execute(instruction, func(cycle int) {
			cycle -= 1

			n := (cycle % width) - computer.X
			var b byte = '.'
			if n >= -1 && n <= 1 {
				b = '#'
			}

			crt.Values[cycle%(crt.Height*crt.Width)] = b
		})
	}

	var sb strings.Builder

	sb.WriteByte('\n')

	for y := 0; y < crt.Height; y++ {
		i := y * crt.Width
		sb.Write(crt.Values[i : i+crt.Width])
		sb.WriteByte('\n')
	}

	return sb.String()
}

type (
	Program []Instruction

	Instruction interface {
		Execute(*Computer)
	}

	Computer struct {
		Cycle int
		clock func(int)

		X int
	}
)

func (c *Computer) Execute(instruction Instruction, clock func(int)) {
	c.clock = clock
	instruction.Execute(c)
	c.clock = nil
}

func (c *Computer) Tick(n int) {
	for i := 0; i < n; i++ {
		c.Cycle++
		if c.clock != nil {
			c.clock(c.Cycle)
		}
	}
}

type OpAddX struct{ Arg int }

func (op *OpAddX) Execute(c *Computer) {
	c.Tick(2)
	c.X += op.Arg
}

type OpNoop struct{}

func (op *OpNoop) Execute(c *Computer) {
	c.Tick(1)
}
