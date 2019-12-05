package intcode

//go:generate stringer -type=opcode -trimprefix=op
type opcode uint64

const (
	opAdd  opcode = 1
	opMul  opcode = 2
	opHalt opcode = 99
)

type inst func(c *Computer)

var isa = map[opcode]inst{
	opAdd:  instAdd,
	opMul:  instMul,
	opHalt: instHalt,
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

func instHalt(c *Computer) {
	c.halt()
}
