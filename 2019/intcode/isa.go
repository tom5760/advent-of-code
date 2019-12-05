package intcode

//go:generate stringer -type=opcode -trimprefix=op
type opcode uint64

const (
	opAdd opcode = iota + 1
	opMul
	opInput
	opOutput
	opJmpT
	opJmpF
	opLT
	opEQ

	opHalt opcode = 99
)

type inst func(c *Computer)

var isa = map[opcode]inst{
	opAdd:    instAdd,
	opMul:    instMul,
	opInput:  instInput,
	opOutput: instOutput,
	opJmpT:   instJmpT,
	opJmpF:   instJmpF,
	opLT:     instLT,
	opEQ:     instEQ,
	opHalt:   instHalt,
}

func instAdd(c *Computer) {
	a := c.arg(1)
	b := c.arg(2)

	c.ret(3, a+b)

	c.PC += 4
}

func instMul(c *Computer) {
	a := c.arg(1)
	b := c.arg(2)

	c.ret(3, a*b)

	c.PC += 4
}

func instInput(c *Computer) {
	i := c.input()

	c.ret(1, i)

	c.PC += 2
}

func instOutput(c *Computer) {
	a := c.arg(1)

	c.output(a)

	c.PC += 2
}

// instJmpT - Opcode 5 is jump-if-true: if the first parameter is non-zero, it
// sets the instruction pointer to the value from the second parameter.
// Otherwise, it does nothing.
func instJmpT(c *Computer) {
	a := c.arg(1)
	b := c.arg(2)

	if a != 0 {
		c.PC = b
	} else {
		c.PC += 3
	}
}

// instJmpF - Opcode 6 is jump-if-false: if the first parameter is zero, it
// sets the instruction pointer to the value from the second parameter.
// Otherwise, it does nothing.
func instJmpF(c *Computer) {
	a := c.arg(1)
	b := c.arg(2)

	if a == 0 {
		c.PC = b
	} else {
		c.PC += 3
	}
}

// instLT - Opcode 7 is less than: if the first parameter is less than the
// second parameter, it stores 1 in the position given by the third parameter.
// Otherwise, it stores 0.
func instLT(c *Computer) {
	a := c.arg(1)
	b := c.arg(2)

	var rv int

	if a < b {
		rv = 1
	} else {
		rv = 0
	}

	c.ret(3, rv)

	c.PC += 4
}

// instEQ - Opcode 8 is equals: if the first parameter is equal to the second
// parameter, it stores 1 in the position given by the third parameter.
// Otherwise, it stores 0.
func instEQ(c *Computer) {
	a := c.arg(1)
	b := c.arg(2)

	var rv int

	if a == b {
		rv = 1
	} else {
		rv = 0
	}

	c.ret(3, rv)

	c.PC += 4
}

func instHalt(c *Computer) {
	c.halt()
}
