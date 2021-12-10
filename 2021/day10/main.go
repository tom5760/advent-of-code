package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"sort"

	"github.com/tom5760/advent-of-code/2021/input"
)

func main() {
	if err := run(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func run() error {
	lines, err := ParseInput(os.Stdin)
	if err != nil {
		return fmt.Errorf("failed to parse input: %w", err)
	}

	fmt.Println("Part 1:", Part1(lines))
	fmt.Println("Part 2:", Part2(lines))

	return nil
}

// The navigation subsystem syntax is made of several lines containing chunks.
// There are one or more chunks on each line, and chunks contain zero or more
// other chunks. Adjacent chunks are not separated by any delimiter; if one
// chunk stops, the next chunk (if any) can immediately start. Every chunk must
// open and close with one of four legal pairs of matching characters:
//
// - If a chunk opens with (, it must close with ).
// - If a chunk opens with [, it must close with ].
// - If a chunk opens with {, it must close with }.
// - If a chunk opens with <, it must close with >.
func ParseInput(r io.Reader) ([]string, error) {
	return input.Parser[string]{ParseFunc: input.String}.Slice(r)
}

// Did you know that syntax checkers actually have contests to see who can get
// the high score for syntax errors in a file? It's true! To calculate the
// syntax error score for a line, take the first illegal character on the line
// and look it up in the following table:
//
//  ): 3 points.
//  ]: 57 points.
//  }: 1197 points.
//  >: 25137 points.
//
// Find the first illegal character in each corrupted line of the navigation
// subsystem. What is the total syntax error score for those errors?
func Part1(lines []string) uint64 {
	var total uint64

	for _, line := range lines {
		total += LineError(line)
	}

	return total
}

func LineError(line string) uint64 {
	var stack []byte

	for _, b := range []byte(line) {
		switch b {
		case '(':
			stack = append(stack, ')')
		case '[':
			stack = append(stack, ']')
		case '{':
			stack = append(stack, '}')
		case '<':
			stack = append(stack, '>')

		default:
			i := len(stack) - 1
			expect := stack[i]
			stack = stack[:i]

			// corrupted
			if b != expect {
				switch b {
				case ')':
					return 3
				case ']':
					return 57
				case '}':
					return 1197
				case '>':
					return 25137
				}
			}
		}
	}

	// Line is valid
	return 0
}

// Now, discard the corrupted lines. The remaining lines are incomplete.
//
// Incomplete lines don't have any incorrect characters - instead, they're
// missing some closing characters at the end of the line. To repair the
// navigation subsystem, you just need to figure out the sequence of closing
// characters that complete all open chunks in the line.
//
// You can only use closing characters (), ], }, or >), and you must add them
// in the correct order so that only legal pairs are formed and all chunks end
// up closed.
//
// Start with a total score of 0. Then, for each character, multiply the total
// score by 5 and then increase the total score by the point value given for
// the character in the following table:
//
//  ): 1 point.
//  ]: 2 points.
//  }: 3 points.
//  >: 4 points.
//
// Find the completion string for each incomplete line, score the completion
// strings, and sort the scores. What is the middle score?
func Part2(lines []string) uint64 {
	var scores []uint64

	for _, line := range lines {
		score := LineScore(line)
		if score == 0 {
			// line was corrupted, or complete, ignore
			continue
		}

		scores = append(scores, score)
	}

	sort.Slice(scores, func(i, j int) bool { return scores[i] < scores[j] })

	return scores[len(scores)/2]
}

func LineScore(line string) uint64 {
	var stack []byte

	for _, b := range []byte(line) {
		switch b {
		case '(':
			stack = append(stack, ')')
		case '[':
			stack = append(stack, ']')
		case '{':
			stack = append(stack, '}')
		case '<':
			stack = append(stack, '>')

		default:
			i := len(stack) - 1
			expect := stack[i]
			stack = stack[:i]

			// corrupted
			if b != expect {
				return 0
			}
		}
	}

	// Line is valid

	if len(stack) == 0 {
		// line is complete
		return 0
	}

	var score uint64

	// Reverse stack for completion order
	for i := len(stack) - 1; i >= 0; i-- {
		b := stack[i]
		score *= 5

		switch b {
		case ')':
			score += 1
		case ']':
			score += 2
		case '}':
			score += 3
		case '>':
			score += 4
		default:
			panic("unexpected byte")
		}
	}

	return score
}
