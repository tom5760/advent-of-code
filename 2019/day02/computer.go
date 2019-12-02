package main

import (
	"fmt"
	"strings"
)

const (
	instLength = 4
)

//go:generate stringer -type=opcode -trimprefix=op
type opcode uint64

const (
	opAdd opcode = 1
	opMul opcode = 2
	opEnd opcode = 99
)

// Computer implements an Intcode computer.
type Computer struct {
	PC     uint64
	Memory []uint64
}

// NewComputer initializes a computer with a copy of the given memory.
func NewComputer(mem []uint64) *Computer {
	computer := &Computer{
		Memory: make([]uint64, len(mem)),
	}

	copy(computer.Memory, mem)

	return computer
}

// Run begins executing instructions at the current program counter.
func (c *Computer) Run() error {
	for {
		op := opcode(c.Memory[c.PC])

		switch op {
		case opAdd:
			c.add()

		case opMul:
			c.mul()

		case opEnd:
			return nil

		default:
			return fmt.Errorf("unexpected opcode: %v", op)
		}
	}
}

// arg dereferences a pointer at position PC+i.
func (c *Computer) arg(i uint64) uint64 {
	addr := c.Memory[c.PC+i]
	return c.Memory[addr]
}

// ret stores a value v at pointer position PC+i.
func (c *Computer) ret(i, v uint64) {
	retAddr := c.Memory[c.PC+i]
	c.Memory[retAddr] = v
}

func (c *Computer) add() {
	a := c.arg(1)
	b := c.arg(2)

	c.ret(3, a+b)

	c.PC += 4
}

func (c *Computer) mul() {
	a := c.arg(1)
	b := c.arg(2)

	c.ret(3, a*b)

	c.PC += 4
}

// String returns a string representation of the computer's memory
func (c *Computer) String() string {
	var sb strings.Builder

	fmt.Fprintf(&sb, "PC: %v\nMemory:\n", c.PC)

	for i := 0; i < len(c.Memory); {
		op := opcode(c.Memory[i])

		switch op {
		case opAdd:
			fallthrough
		case opMul:
			if len(c.Memory)-i <= 3 {
				fmt.Fprintf(&sb, "%d ", op)
				i++

				continue
			}

			aAddr := c.Memory[i+1]
			bAddr := c.Memory[i+2]
			cAddr := c.Memory[i+3]

			a := c.Memory[aAddr]
			b := c.Memory[bAddr]
			c := c.Memory[cAddr]

			fmt.Fprintf(&sb, "%s %#.2x (%.3d) %#.2x (%.3d) %#.2x (%.3d)\n",
				op, aAddr, a, bAddr, b, cAddr, c)

			i += instLength

		case opEnd:
			fmt.Fprintf(&sb, "%s\n", op)
			i++

		default:
			fmt.Fprintf(&sb, "%d ", op)
			i++
		}
	}

	return sb.String()
}
