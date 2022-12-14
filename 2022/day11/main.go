// Package day11 implements the solution for Advent of Code 2022 day 11.
//
// See: https://adventofcode.com/2022/day/11
package day11

import (
	"bufio"
	"bytes"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/tom5760/advent-of-code/2022/inpututils"
	"golang.org/x/exp/slices"
)

func Parse(name string) ([]Monkey, error) {
	var (
		monkeys []Monkey
		cur     Monkey
	)

	err := inpututils.Scan(name, func(scanner *bufio.Scanner) error {
		for scanner.Scan() {
			lineBuf := scanner.Bytes()

			if len(lineBuf) == 0 {
				monkeys = append(monkeys, cur)
				cur = Monkey{}
				continue
			}

			line := string(bytes.TrimSpace(lineBuf))

			if args, ok := inpututils.TrimPrefix(line, "Monkey "); ok {
				// last character is a ':'
				id, err := strconv.Atoi(args[:len(args)-1])
				if err != nil {
					return fmt.Errorf("failed to parse monkey ID: %w", err)
				}

				cur.ID = id

				continue
			}

			if args, ok := inpututils.TrimPrefix(line, "Starting items: "); ok {
				itemStrs := strings.Split(args, ", ")
				cur.Items = make([]int, len(itemStrs))
				for i, itemStr := range itemStrs {
					item, err := strconv.Atoi(itemStr)
					if err != nil {
						return fmt.Errorf("failed to parse item string: %w", err)
					}

					cur.Items[i] = item
				}

				continue
			}

			if args, ok := inpututils.TrimPrefix(line, "Operation: new = old "); ok {
				fields := strings.Fields(args)
				if len(fields) != 2 {
					return fmt.Errorf("unexpected operation format")
				}
				operator := fields[0]

				if fields[1] == "old" {
					switch operator {
					case "*":
						cur.Operation = OperationSquare()
					default:
						return fmt.Errorf("unexpected operation operator %q", operator)
					}
				} else {
					arg, err := strconv.Atoi(fields[1])
					if err != nil {
						return fmt.Errorf("failed to parse operation operand: %w", err)
					}

					switch operator {
					case "*":
						cur.Operation = OperationMultiply(arg)
					case "+":
						cur.Operation = OperationAdd(arg)
					default:
						return fmt.Errorf("unexpected operation operator %q", operator)
					}
				}

				continue
			}

			if args, ok := inpututils.TrimPrefix(line, "Test: divisible by "); ok {
				factor, err := strconv.Atoi(args)
				if err != nil {
					return fmt.Errorf("failed to parse test factor: %w", err)
				}

				cur.Factor = factor

				continue
			}

			if args, ok := inpututils.TrimPrefix(line, "If true: throw to monkey "); ok {
				id, err := strconv.Atoi(args)
				if err != nil {
					return fmt.Errorf("failed to parse true ID: %w", err)
				}

				cur.TrueID = id

				continue
			}

			if args, ok := inpututils.TrimPrefix(line, "If false: throw to monkey "); ok {
				id, err := strconv.Atoi(args)
				if err != nil {
					return fmt.Errorf("failed to parse true ID: %w", err)
				}

				cur.FalseID = id

				continue
			}

			return fmt.Errorf("unexpected line %q", line)
		}

		// Include the last monkey
		monkeys = append(monkeys, cur)

		return nil
	})

	return monkeys, err
}

func Part1(monkeys []Monkey) int {
	test := CloneMonkeys(monkeys)
	return Simulate(test, 20, 3)
}

func Part2(monkeys []Monkey) int {
	test := CloneMonkeys(monkeys)
	return Simulate(test, 10000, 0)
}

func Simulate(monkeys []Monkey, rounds, decay int) int {
	gcm := 1

	for _, monkey := range monkeys {
		gcm *= monkey.Factor
	}

	for round := 0; round < rounds; round++ {
		for i := range monkeys {
			monkey := &monkeys[i]
			monkey.Inspect(monkeys, decay, gcm)
		}
	}

	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].InspectCount > monkeys[j].InspectCount
	})

	return monkeys[0].InspectCount * monkeys[1].InspectCount
}

type Monkey struct {
	ID        int
	Items     []int
	Operation func(int) int

	Factor  int
	TrueID  int
	FalseID int

	InspectCount int
}

func CloneMonkeys(monkeys []Monkey) []Monkey {
	rv := slices.Clone(monkeys)

	// slices.Clone is a shallow clone, so if we want a clone of Items, we have
	// to do it ourselves.
	for i, monkey := range monkeys {
		rv[i].Items = slices.Clone(monkey.Items)
	}

	return rv
}

func (m *Monkey) Inspect(monkeys []Monkey, decay, gcm int) {
	for _, worry := range m.Items {
		worry = m.Operation(worry)
		if decay > 0 {
			worry /= decay
		}

		worry = worry % gcm

		var next *Monkey
		if worry%m.Factor == 0 {
			next = &monkeys[m.TrueID]
		} else {
			next = &monkeys[m.FalseID]
		}

		next.Items = append(next.Items, worry)
	}

	m.InspectCount += len(m.Items)
	m.Items = m.Items[:0]
}

func OperationMultiply(factor int) func(int) int {
	return func(old int) int {
		return old * factor
	}
}

func OperationAdd(increment int) func(int) int {
	return func(old int) int {
		return old + increment
	}
}

func OperationSquare() func(int) int {
	return func(old int) int {
		return old * old
	}
}
