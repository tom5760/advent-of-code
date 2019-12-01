package main

import (
	"strconv"
	"testing"
)

type InstructionTest struct {
	OpCode      OpCode
	Instruction Instruction
	Initial     Registers
	Expected    Registers
}

var tests = []InstructionTest{
	{
		OpCode:      addr,
		Instruction: Instruction{0, 0, 1, 0},
		Initial:     [4]int{1, 2, 0, 0},
		Expected:    [4]int{3, 2, 0, 0},
	},
	{
		OpCode:      addi,
		Instruction: Instruction{0, 0, 5, 0},
		Initial:     [4]int{1, 2, 0, 0},
		Expected:    [4]int{6, 2, 0, 0},
	},
	{
		OpCode:      mulr,
		Instruction: Instruction{0, 0, 1, 0},
		Initial:     [4]int{2, 2, 0, 0},
		Expected:    [4]int{4, 2, 0, 0},
	},
	{
		OpCode:      muli,
		Instruction: Instruction{0, 0, 5, 0},
		Initial:     [4]int{2, 2, 0, 0},
		Expected:    [4]int{10, 2, 0, 0},
	},
	{
		OpCode:      banr,
		Instruction: Instruction{0, 0, 1, 0},
		Initial:     [4]int{2, 2, 0, 0},
		Expected:    [4]int{2, 2, 0, 0},
	},
	{
		OpCode:      bani,
		Instruction: Instruction{0, 0, 5, 0},
		Initial:     [4]int{2, 2, 0, 0},
		Expected:    [4]int{0, 2, 0, 0},
	},
	{
		OpCode:      borr,
		Instruction: Instruction{0, 0, 1, 0},
		Initial:     [4]int{2, 2, 0, 0},
		Expected:    [4]int{2, 2, 0, 0},
	},
	{
		OpCode:      bori,
		Instruction: Instruction{0, 0, 5, 0},
		Initial:     [4]int{2, 2, 0, 0},
		Expected:    [4]int{7, 2, 0, 0},
	},
	{
		OpCode:      setr,
		Instruction: Instruction{0, 0, 0, 3},
		Initial:     [4]int{2, 2, 0, 0},
		Expected:    [4]int{2, 2, 0, 2},
	},
	{
		OpCode:      seti,
		Instruction: Instruction{0, 7, 0, 3},
		Initial:     [4]int{2, 2, 0, 0},
		Expected:    [4]int{2, 2, 0, 7},
	},
	{
		OpCode:      gtir,
		Instruction: Instruction{0, 7, 0, 3},
		Initial:     [4]int{2, 2, 0, 0},
		Expected:    [4]int{2, 2, 0, 1},
	},
	{
		OpCode:      gtri,
		Instruction: Instruction{0, 1, 3, 3},
		Initial:     [4]int{3, 2, 0, 1},
		Expected:    [4]int{3, 2, 0, 0},
	},
	{
		OpCode:      gtrr,
		Instruction: Instruction{0, 2, 1, 3},
		Initial:     [4]int{2, 2, 1, 1},
		Expected:    [4]int{2, 2, 1, 0},
	},
}

func TestInstruction(t *testing.T) {
	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			actual := test.OpCode(test.Instruction, test.Initial)
			if actual != test.Expected {
				t.Error("fail", test.Expected, actual)
			}
		})
	}
}
