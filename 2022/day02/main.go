package day02

import (
	"bufio"
	"fmt"

	"github.com/tom5760/advent-of-code/2022/inpututils"
)

func Parse(name string) ([]Instruction, error) {
	var instructions []Instruction

	err := inpututils.Scan(name, func(scanner *bufio.Scanner) error {
		scanner.Split(bufio.ScanWords)

		for scanner.Scan() {
			opponent := scanner.Bytes()

			if !scanner.Scan() {
				return fmt.Errorf("missing response move")
			}

			response := scanner.Bytes()

			if len(opponent) != 1 {
				return fmt.Errorf("invalid opponent move")
			}
			if len(response) != 1 {
				return fmt.Errorf("invalid response move")
			}

			instructions = append(instructions, Instruction{
				Opponent: opponent[0],
				Response: response[0],
			})
		}

		return nil
	})

	return instructions, err
}

type Move int

const (
	Rock     Move = 1
	Paper    Move = 2
	Scissors Move = 3
)

type Instruction struct {
	Opponent byte
	Response byte
}

type Round struct {
	Opponent Move
	Response Move
}

func (r *Round) Score() int {
	const (
		Win  = 6
		Loss = 0
		Draw = 3
	)

	score := int(r.Response)

	switch r.Response {
	case Rock:
		switch r.Opponent {
		case Scissors:
			score += Win
		case Paper:
			score += Loss
		case Rock:
			score += Draw
		}

	case Paper:
		switch r.Opponent {
		case Rock:
			score += Win
		case Scissors:
			score += Loss
		case Paper:
			score += Draw
		}

	case Scissors:
		switch r.Opponent {
		case Paper:
			score += Win
		case Rock:
			score += Loss
		case Scissors:
			score += Draw
		}
	}

	return score
}

func ToMove(in byte) Move {
	switch in {
	case 'A', 'X':
		return Rock
	case 'B', 'Y':
		return Paper
	case 'C', 'Z':
		return Scissors
	default:
		panic("invalid move")
	}
}

func Part1(instructions []Instruction) int {
	var total int

	for _, instruction := range instructions {
		round := Round{
			Opponent: ToMove(instruction.Opponent),
			Response: ToMove(instruction.Response),
		}

		total += round.Score()
	}

	return total
}

func Part2(instructions []Instruction) int {
	var total int

	for _, instruction := range instructions {
		round := Round{
			Opponent: ToMove(instruction.Opponent),
		}

		switch instruction.Response {
		case 'X': // lose
			switch round.Opponent {
			case Rock:
				round.Response = Scissors
			case Paper:
				round.Response = Rock
			case Scissors:
				round.Response = Paper
			}

		case 'Y': // draw
			round.Response = round.Opponent

		case 'Z': // win
			switch round.Opponent {
			case Rock:
				round.Response = Paper
			case Paper:
				round.Response = Scissors
			case Scissors:
				round.Response = Rock
			}
		}

		total += round.Score()
	}

	return total
}
