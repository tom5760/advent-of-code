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
	opRelBase

	opHalt opcode = 99
)

var isa = map[opcode]func(c *Computer){
	// Opcode 1 adds together numbers read from two positions and stores the
	// result in a third position.
	opAdd: func(c *Computer) {
		a := c.arg(1)
		b := c.arg(2)

		c.ret(3, a+b)

		c.PC += 4
	},

	// Opcode 2 works exactly like opcode 1, except it multiplies the two inputs
	// instead of adding them.
	opMul: func(c *Computer) {
		a := c.arg(1)
		b := c.arg(2)

		c.ret(3, a*b)

		c.PC += 4
	},

	// Opcode 3 takes a single integer as input and saves it to the position
	// given by its only parameter.
	opInput: func(c *Computer) {
		i := c.input()

		c.ret(1, i)

		c.PC += 2
	},

	// Opcode 4 outputs the value of its only parameter.
	opOutput: func(c *Computer) {
		a := c.arg(1)

		c.output(a)

		c.PC += 2
	},

	// Opcode 5 is jump-if-true: if the first parameter is non-zero, it sets the
	// instruction pointer to the value from the second parameter.  Otherwise, it
	// does nothing.
	opJmpT: func(c *Computer) {
		a := c.arg(1)
		b := c.arg(2)

		if a != 0 {
			c.PC = b
		} else {
			c.PC += 3
		}
	},

	// Opcode 6 is jump-if-false: if the first parameter is zero, it sets the
	// instruction pointer to the value from the second parameter.  Otherwise, it
	// does nothing.
	opJmpF: func(c *Computer) {
		a := c.arg(1)
		b := c.arg(2)

		if a == 0 {
			c.PC = b
		} else {
			c.PC += 3
		}
	},

	// Opcode 7 is less than: if the first parameter is less than the second
	// parameter, it stores 1 in the position given by the third parameter.
	// Otherwise, it stores 0.
	opLT: func(c *Computer) {
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
	},

	// Opcode 8 is equals: if the first parameter is equal to the second
	// parameter, it stores 1 in the position given by the third parameter.
	// Otherwise, it stores 0.
	opEQ: func(c *Computer) {
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
	},

	// Opcode 9 adjusts the relative base by the value of its only parameter. The
	// relative base increases (or decreases, if the value is negative) by the
	// value of the parameter.
	opRelBase: func(c *Computer) {
		c.RB += c.arg(1)
		c.PC += 2
	},

	// Opcode 99 means that the program is finished and should immediately halt.
	opHalt: func(c *Computer) {
		c.halt()
	},
}
