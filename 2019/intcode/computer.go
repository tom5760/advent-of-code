package intcode

import (
	"fmt"
)

// Computer implements an Intcode computer.
type Computer struct {
	Halted bool
	PC     int
	Memory []int
}

// NewComputer initializes a computer with a copy of the given memory.
func NewComputer(mem []int) *Computer {
	computer := &Computer{
		Memory: make([]int, len(mem)),
	}

	copy(computer.Memory, mem)

	return computer
}

// Run executes instructions until a halt instruction is executed.
func (c *Computer) Run() error {
	for !c.Halted {
		if err := c.Step(); err != nil {
			return err
		}
	}

	return nil
}

// Step executes a single instruction at the current program counter.
func (c *Computer) Step() error {
	op := c.opcode()

	inst, ok := isa[op]
	if !ok {
		return fmt.Errorf("unexepcted opcode: %d", op)
	}

	inst(c)

	return nil
}

// opcode returns the opcode at PC
func (c *Computer) opcode() opcode {
	return opcode(c.Memory[c.PC])
}

// arg dereferences a pointer at position PC+i.
func (c *Computer) arg(i int) int {
	addr := c.Memory[c.PC+i]
	return c.Memory[addr]
}

// ret stores a value v at pointer position PC+i.
func (c *Computer) ret(i, v int) {
	retAddr := c.Memory[c.PC+i]
	c.Memory[retAddr] = v
}

// halt sets the Halted flag to true, which will cause the Run function to
// return.
func (c *Computer) halt() {
	c.Halted = true
}
