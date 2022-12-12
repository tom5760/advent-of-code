// Package day09 implements the solution for Advent of Code 2022 day 9.
//
// See: https://adventofcode.com/2022/day/9
package day09

import (
	"bufio"
	"bytes"
	"fmt"
	"strconv"

	"github.com/tom5760/advent-of-code/2022/inpututils"
)

func Parse(name string) ([]Motion, error) {
	var motions []Motion

	err := inpututils.Scan(name, func(scanner *bufio.Scanner) error {
		for scanner.Scan() {
			fields := bytes.Fields(scanner.Bytes())
			if len(fields) != 2 {
				return fmt.Errorf("unexpected field format")
			}

			amount, err := strconv.Atoi(string(fields[1]))
			if err != nil {
				return fmt.Errorf("failed to parse amount: %w", err)
			}

			motions = append(motions, Motion{
				Direction: Direction(fields[0][0]),
				Amount:    amount,
			})
		}

		return nil
	})

	return motions, err
}

func Part1(motions []Motion) int {
	return Simulate(motions, 2)
}

func Part2(motions []Motion) int {
	return Simulate(motions, 10)
}

func Simulate(motions []Motion, length int) int {
	rope := make(Rope, length)
	tail := &rope[len(rope)-1]
	visited := make(map[Coordinate]bool)

	for _, motion := range motions {
		for i := 0; i < motion.Amount; i++ {
			rope.Move(motion.Direction)
			visited[*tail] = true
		}
	}

	var total int
	for _, ok := range visited {
		if ok {
			total++
		}
	}

	return total
}

type (
	Direction byte

	Motion struct {
		Direction Direction
		Amount    int
	}

	Coordinate struct {
		X int
		Y int
	}

	Rope []Coordinate
)

const (
	Up    Direction = 'U'
	Down  Direction = 'D'
	Left  Direction = 'L'
	Right Direction = 'R'
)

func (d Direction) String() string {
	return string(d)
}

func (c *Coordinate) Move(d Direction) {
	switch d {
	case Up:
		c.Y--
	case Down:
		c.Y++
	case Left:
		c.X--
	case Right:
		c.X++
	}
}

func (r Rope) Move(d Direction) {
	head := &r[0]
	head.Move(d)

	for i := 1; i < len(r); i++ {
		tail := &r[i]

		// If the head is ever two steps directly up, down, left, or right from the
		// tail, the tail must also move one step in that direction so it remains
		// close enough.
		if head.X == tail.X {
			if head.Y-tail.Y == 2 {
				tail.Move(Down)
			} else if head.Y-tail.Y == -2 {
				tail.Move(Up)
			}
		} else if head.Y == tail.Y {
			if head.X-tail.X == 2 {
				tail.Move(Right)
			} else if head.X-tail.X == -2 {
				tail.Move(Left)
			}
		} else {
			// Otherwise, if the head and tail aren't touching and aren't in the same row
			// or column, the tail always moves one step diagonally to keep up.
			if head.X-tail.X == 2 {
				if head.Y > tail.Y {
					tail.Move(Down)
				} else {
					tail.Move(Up)
				}
				tail.Move(Right)
			} else if head.X-tail.X == -2 {
				if head.Y > tail.Y {
					tail.Move(Down)
				} else {
					tail.Move(Up)
				}
				tail.Move(Left)
			} else if head.Y-tail.Y == 2 {
				if head.X > tail.X {
					tail.Move(Right)
				} else {
					tail.Move(Left)
				}
				tail.Move(Down)
			} else if head.Y-tail.Y == -2 {
				if head.X > tail.X {
					tail.Move(Right)
				} else {
					tail.Move(Left)
				}
				tail.Move(Up)
			}
		}

		head = tail
	}
}
