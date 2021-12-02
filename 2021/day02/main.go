package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	if err := run(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func run() error {
	cmds, err := ParseInput(os.Stdin)
	if err != nil {
		return fmt.Errorf("failed to parse input: %w", err)
	}

	fmt.Println("Part 1:", Part1(cmds))
	fmt.Println("Part 2:", Part2(cmds))

	return nil
}

// It seems like the submarine can take a series of commands like forward 1,
// down 2, or up 3:
//
//
// The submarine seems to already have a planned course (your puzzle input).
// You should probably figure out where it's going.
func ParseInput(r io.Reader) ([]Command, error) {
	var cmds []Command

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		cmd, err := ParseCommand(scanner.Text())
		if err != nil {
			return nil, fmt.Errorf("failed to parse command: %w", err)
		}

		cmds = append(cmds, cmd)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to scan input: %w", err)
	}

	return cmds, nil
}

// Calculate the horizontal position and depth you would have after following
// the planned course. What do you get if you multiply your final horizontal
// position by your final depth?
func Part1(cmds []Command) int {
	var sub Submarine

	for _, cmd := range cmds {
		cmd.ExecV1(&sub)
	}

	return sub.X * sub.Y
}

// Using this new interpretation of the commands, calculate the horizontal
// position and depth you would have after following the planned course. What
// do you get if you multiply your final horizontal position by your final
// depth?
func Part2(cmds []Command) int {
	var sub Submarine

	for _, cmd := range cmds {
		cmd.ExecV2(&sub)
	}

	return sub.X * sub.Y
}
