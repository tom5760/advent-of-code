package intcode

import (
	"fmt"
	"strings"
)

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

			i += 4

		case opHalt:
			fmt.Fprintf(&sb, "%s\n", op)
			i++

		default:
			fmt.Fprintf(&sb, "%d ", op)
			i++
		}
	}

	return sb.String()
}
