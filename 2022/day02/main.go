package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/tom5760/advent-of-code/2022/input"
)

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

func main() {
	if err := run(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func run() error {
	parser := input.Parser[Instruction]{
		TokenFunc: bufio.ScanWords,
		ParseFunc: Parse(),
	}

	instructions, err := parser.ReadFileSlice("./day02/input")
	if err != nil {
		return fmt.Errorf("failed to parse input: %w", err)
	}

	Part1(instructions)
	Part2(instructions)

	return nil
}

func Parse() input.ParseFunc[Instruction] {
	var opponentIn, responseIn []byte

	return func(input []byte, outChan chan<- Instruction) error {
		if opponentIn == nil {
			opponentIn = input
			return nil
		}

		responseIn = input

		if len(opponentIn) != 1 {
			return fmt.Errorf("invalid opponent move")
		}

		if len(responseIn) != 1 {
			return fmt.Errorf("invalid response move")
		}

		outChan <- Instruction{
			Opponent: opponentIn[0],
			Response: responseIn[0],
		}

		opponentIn = nil
		responseIn = nil

		return nil
	}
}

func Part1(instructions []Instruction) {
	var total int

	for _, instruction := range instructions {
		round := Round{
			Opponent: ToMove(instruction.Opponent),
			Response: ToMove(instruction.Response),
		}

		total += round.Score()
	}

	log.Println("PART 1:", total)
}

func Part2(instructions []Instruction) {
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

	log.Println("PART 2:", total)
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
