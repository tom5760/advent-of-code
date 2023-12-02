package day02

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"strconv"
)

type (
	game struct {
		ID           int
		Observations []observation
	}

	observation struct {
		Red   int
		Green int
		Blue  int
	}
)

func Run(lg *log.Logger, input io.Reader) (int, int, error) {
	var part1, part2 int

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		line := scanner.Bytes()
		game, err := parseGame(line)
		if err != nil {
			return 0, 0, err
		}

		part1 += doPart1(game)
		part2 += doPart2(game)
	}

	if err := scanner.Err(); err != nil {
		return 0, 0, fmt.Errorf("failed to scan input: %w", err)
	}

	return part1, part2, nil
}

func parseGame(line []byte) (game, error) {
	line = bytes.TrimPrefix(line, []byte("Game "))

	// Every line starts with "Game: "
	sep := bytes.IndexByte(line, ':')
	str := string(line[:sep])
	line = line[sep+2:]

	id, err := strconv.Atoi(str)
	if err != nil {
		return game{}, fmt.Errorf("failed to parse game ID: %w", err)
	}

	var (
		observations []observation
		cur          observation
	)
	for {
		sep := bytes.IndexByte(line, ' ')
		str := string(line[:sep])
		line = line[sep+1:]

		n, err := strconv.Atoi(str)
		if err != nil {
			return game{}, fmt.Errorf("failed to parse count: %w", err)
		}

		switch {
		case bytes.HasPrefix(line, []byte("red")):
			cur.Red = n
			line = line[3:]
		case bytes.HasPrefix(line, []byte("green")):
			cur.Green = n
			line = line[5:]
		case bytes.HasPrefix(line, []byte("blue")):
			cur.Blue = n
			line = line[4:]
		default:
			return game{}, fmt.Errorf("unexpected observation type")
		}

		if len(line) == 0 {
			break
		}

		switch line[0] {
		case ';':
			observations = append(observations, cur)
		case ',':
		}
		line = line[2:]
	}

	observations = append(observations, cur)

	game := game{
		ID:           id,
		Observations: observations,
	}

	return game, nil
}

func doPart1(game game) int {
	const (
		constraintRed   = 12
		constraintGreen = 13
		constraintBlue  = 14
	)

	for _, obs := range game.Observations {
		if obs.Red > constraintRed || obs.Green > constraintGreen || obs.Blue > constraintBlue {
			return 0
		}
	}

	return game.ID
}

func doPart2(game game) int {
	var min observation

	for _, obs := range game.Observations {
		if obs.Red > min.Red {
			min.Red = obs.Red
		}
		if obs.Green > min.Green {
			min.Green = obs.Green
		}
		if obs.Blue > min.Blue {
			min.Blue = obs.Blue
		}
	}

	return min.Red * min.Green * min.Blue
}
