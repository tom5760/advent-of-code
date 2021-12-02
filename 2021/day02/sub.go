package main

import (
	"fmt"
	"strconv"
	"strings"
)

type (
	Submarine struct {
		Aim  int
		X, Y int
	}

	Command interface {
		// - `forward X` increases the horizontal position by X units.
		// - `down X` increases the depth by X units.
		// - `up X` decreases the depth by X units.
		ExecV1(*Submarine)

		// - `down X` increases your aim by X units.
		// - `up X` decreases your aim by X units.
		// - `forward X` does two things:
		//   - It increases your horizontal position by X units.
		//   - It increases your depth by your aim multiplied by X.
		ExecV2(*Submarine)
	}
)

func ParseCommand(cmd string) (Command, error) {
	fields := strings.Fields(cmd)

	if len(fields) != 2 {
		return nil, fmt.Errorf("unexpected number of fields")
	}

	name := fields[0]
	arg, err := strconv.ParseInt(fields[1], 10, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to parse argument '%v': %w", fields[1], err)
	}

	switch name {
	case "forward":
		return ForwardCommand(arg), nil
	case "down":
		return DownCommand(arg), nil
	case "up":
		return UpCommand(arg), nil
	default:
		return nil, fmt.Errorf("unknown command '%v'", name)
	}
}

type ForwardCommand int

func (c ForwardCommand) ExecV1(sub *Submarine) {
	sub.X += int(c)
}

func (c ForwardCommand) ExecV2(sub *Submarine) {
	sub.X += int(c)
	sub.Y += sub.Aim * int(c)
}

type DownCommand int

func (c DownCommand) ExecV1(sub *Submarine) {
	sub.Y += int(c)
}

func (c DownCommand) ExecV2(sub *Submarine) {
	sub.Aim += int(c)
}

type UpCommand int

func (c UpCommand) ExecV1(sub *Submarine) {
	sub.Y -= int(c)
}

func (c UpCommand) ExecV2(sub *Submarine) {
	sub.Aim -= int(c)
}
