// Package day13 implements the solution for Advent of Code 2022 day 13.
//
// See: https://adventofcode.com/2022/day/13
package day13

import (
	"bufio"
	"container/list"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/tom5760/advent-of-code/2022/inpututils"
)

func Parse(name string) ([]Pair, error) {
	var pairs []Pair

	err := inpututils.Scan(name, func(scanner *bufio.Scanner) error {
		var pair Pair

		packet := &pair.Left

		for scanner.Scan() {
			line := scanner.Bytes()

			if len(line) == 0 {
				pairs = append(pairs, pair)
				pair = Pair{}
				packet = &pair.Left

				continue
			}

			var (
				stack  list.List
				el     *list.Element
				numstr string
			)

			for _, b := range line {
				switch b {
				case '[':
					if el == nil {
						el = stack.PushBack(packet)
					} else {
						el = stack.PushBack(new(ListData))
					}

				case ']':
					if numstr != "" {
						num, err := strconv.Atoi(numstr)
						if err != nil {
							return fmt.Errorf("failed to parse number: %w", err)
						}

						v := stack.Back().Value.(*ListData)
						*v = append(*v, IntData(num))
						numstr = ""
					}

					last := stack.Remove(stack.Back()).(*ListData)
					el = stack.Back()
					if el != nil {
						v := el.Value.(*ListData)
						*v = append(*v, *last)
					}

				case '1', '2', '3', '4', '5', '6', '7', '8', '9', '0':
					numstr += string(b)

				case ',':
					if numstr != "" {
						num, err := strconv.Atoi(numstr)
						if err != nil {
							return fmt.Errorf("failed to parse number: %w", err)
						}

						v := stack.Back().Value.(*ListData)
						*v = append(*v, IntData(num))
						numstr = ""
					}

				default:
					return fmt.Errorf("unexpected byte %q", b)
				}
			}

			packet = &pair.Right
		}

		pairs = append(pairs, pair)

		return nil
	})

	return pairs, err
}

func Part1(pairs []Pair) int {
	var sum int

	for i, pair := range pairs {
		if inOrder := pair.InOrder(); inOrder {
			sum += i + 1
		}
	}

	return sum
}

func Part2(pairs []Pair) int {
	packets := make([]Data, 0, (len(pairs)*2)+2)

	for _, pair := range pairs {
		packets = append(packets, pair.Left, pair.Right)
	}

	packets = append(packets,
		ListData{ListData{IntData(2)}},
		ListData{ListData{IntData(6)}},
	)

	sort.Slice(packets, func(i, j int) bool {
		return Compare(packets[i], packets[j]) == -1
	})

	var i1, i2 int

	for i, packet := range packets {
		if v, ok := packet.(ListData); ok {
			if len(v) != 1 {
				continue
			}

			if v, ok := v[0].(ListData); ok {
				if len(v) != 1 {
					continue
				}

				if v, ok := v[0].(IntData); ok {
					switch v {
					case 2:
						i1 = i + 1
					case 6:
						i2 = i + 1
					}
				}
			}
		}
	}

	return i1 * i2
}

type (
	Pair struct {
		Left, Right ListData
	}

	Data interface {
		fmt.Stringer
	}

	ListData []Data

	IntData int
)

func (p *Pair) InOrder() bool {
	return Compare(p.Left, p.Right) == -1
}

func Compare(left, right Data) int {
	// If both values are integers, the lower integer should come first.
	leftInt, leftIntOK := left.(IntData)
	rightInt, rightIntOK := right.(IntData)

	if leftIntOK && rightIntOK {
		switch {
		case leftInt < rightInt:
			return -1
		case leftInt > rightInt:
			return 1
		default:
			return 0
		}
	}

	// If both values are lists, compare the first value of each list, then the
	// second value, and so on.
	leftList, leftListOK := left.(ListData)
	rightList, rightListOK := right.(ListData)

	if leftListOK && rightListOK {

		for i := 0; true; i++ {
			if i >= len(leftList) && i >= len(rightList) {
				return 0
			}
			if i >= len(leftList) {
				return -1
			}
			if i >= len(rightList) {
				return 1
			}

			if out := Compare(leftList[i], rightList[i]); out != 0 {
				return out
			}
		}

		return 0
	}

	// If exactly one value is an integer, convert the integer to a list which
	// contains that integer as its only value, then retry the comparison.
	if leftIntOK && !rightIntOK {
		return Compare(ListData{leftInt}, right)
	}

	if !leftIntOK && rightIntOK {
		return Compare(left, ListData{rightInt})
	}

	panic("unknown case")
}

func (p *Pair) String() string {
	var sb strings.Builder

	sb.WriteString(p.Left.String())
	sb.WriteByte('\n')
	sb.WriteString(p.Right.String())
	sb.WriteByte('\n')

	return sb.String()
}

func (p ListData) String() string {
	var sb strings.Builder

	sb.WriteByte('[')

	for i, d := range p {
		sb.WriteString(d.String())
		if i < len(p)-1 {
			sb.WriteByte(',')
		}
	}

	sb.WriteByte(']')

	return sb.String()
}

func (d IntData) String() string {
	return fmt.Sprintf("%v", int(d))
}
