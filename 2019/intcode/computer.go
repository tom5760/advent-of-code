package intcode

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/tom5760/advent-of-code/2019/common"
)

type paramMode int

const (
	paramModePosition  paramMode = 0
	paramModeImmediate paramMode = 1
)

// Computer implements an Intcode computer.
type Computer struct {
	Log bool

	Halted bool
	PC     int

	Memory, Inputs, Outputs []int

	tw *tabwriter.Writer
}

// NewComputer initializes a computer with a copy of the given memory.
func NewComputer(mem []int) *Computer {
	computer := &Computer{
		Memory: make([]int, len(mem)),

		tw: tabwriter.NewWriter(os.Stderr, 0, 2, 1, ' ', 0),
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
func (c *Computer) Step() (err error) {
	defer func() {
		if r := recover(); r != nil {
			c.log("\n")
			c.tw.Flush()
			panic(r)
		}
	}()

	op := c.opcode()

	c.log("PC: %d\t%s\t", c.PC, op)

	inst, ok := isa[op]
	if !ok {
		return fmt.Errorf("unexepcted opcode: %d", op)
	}

	inst(c)

	c.log("\n")

	return nil
}

// opcode returns the opcode at PC
func (c *Computer) opcode() opcode {
	// Only last two decimal digits are the opcode.
	return opcode(c.Memory[c.PC] % 100)
}

// mode returns the parameter mode for the argument at PC+i.
func (c *Computer) mode(i int) paramMode {
	return paramMode((c.Memory[c.PC] / common.IntPow10(i+1)) % 10)
}

// deref dereferences the argument value at PC+i.
func (c *Computer) deref(i int) int {
	addr := c.Memory[c.PC+i]

	c.log("%#.2x ", addr)
	c.log("(%.3d)\t", c.Memory[addr])
	return c.Memory[addr]
}

// immediate returns the argument value at PC+i.
func (c *Computer) immediate(i int) int {
	c.log("(%.3d)\t", c.Memory[c.PC+i])
	return c.Memory[c.PC+i]
}

// arg returns the argument value at PC+i.  Takes care to check the opcode for
// the parameter mode for whether to dereference or not.
func (c *Computer) arg(i int) int {
	switch c.mode(i) {
	case paramModePosition:
		return c.deref(i)

	case paramModeImmediate:
		return c.immediate(i)

	default:
		panic("unexpected param mode")
	}
}

// ret stores a value v at pointer position PC+i.  Assumes that return values
// are always in position parameter mode.
func (c *Computer) ret(i, v int) {
	retAddr := c.Memory[c.PC+i]
	c.log("= %#.2x (%.3d)\t", retAddr, v)
	c.Memory[retAddr] = v
}

// input pops the first value off of the input slice.
func (c *Computer) input() int {
	if len(c.Inputs) == 0 {
		panic("input empty")
	}

	i := c.Inputs[0]
	c.Inputs = c.Inputs[1:]

	return i
}

// output pushes a value onto the output slice.
func (c *Computer) output(i int) {
	c.Outputs = append(c.Outputs, i)
}

// halt sets the Halted flag to true, which will cause the Run function to
// return.
func (c *Computer) halt() {
	c.Halted = true
	c.tw.Flush()
}

func (c *Computer) log(format string, a ...interface{}) {
	if c.Log {
		fmt.Fprintf(c.tw, format, a...)
	}
}
