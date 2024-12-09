package day07

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"strconv"
)

func Run(r io.Reader) (int, int, error) {
	equations, err := ParseInput(r)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to parse input: %w", err)
	}

	var part1, part2 int
	for _, eq := range equations {
		if eq.Check(false) {
			part1 += eq.Result
		}
		if eq.Check(true) {
			part2 += eq.Result
		}
	}

	return part1, part2, nil
}

type Equation struct {
	Result int
	Terms  []int
}

func ParseInput(r io.Reader) ([]Equation, error) {
	var equations []Equation

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		equation, err := ParseEquation(scanner.Bytes())
		if err != nil {
			return nil, fmt.Errorf("failed to parse equation: %w", err)
		}

		equations = append(equations, *equation)
	}

	return equations, nil
}

func ParseEquation(line []byte) (*Equation, error) {
	fields := bytes.Fields(line)
	if len(fields) < 2 {
		return nil, errors.New("unexpected equation format: too few items")
	}

	resultBuf := fields[0]

	if resultBuf[len(resultBuf)-1] != ':' {
		return nil, errors.New("unexpected equation result format")
	}

	result, err := strconv.Atoi(string(resultBuf[:len(resultBuf)-1]))
	if err != nil {
		return nil, fmt.Errorf("failed to parse result: %w", err)
	}

	terms := make([]int, len(fields)-1)

	for i, buf := range fields[1:] {
		term, err := strconv.Atoi(string(buf))
		if err != nil {
			return nil, fmt.Errorf("failed to parse term: %w", err)
		}

		terms[i] = term
	}

	equation := &Equation{
		Result: result,
		Terms:  terms,
	}

	return equation, nil
}

type Operator byte

const (
	Add Operator = iota
	Mul
	Concat
)

func (e *Equation) Check(withConcat bool) bool {
	ops := make([]Operator, len(e.Terms)-1)
	for i := range ops {
		ops[i] = Add
	}

outer:
	for {
		res := e.Terms[0]

		for i, op := range ops {
			t := e.Terms[i+1]

			switch op {
			case Add:
				res += t

			case Mul:
				res *= t

			case Concat:
				v := fmt.Sprintf("%v%v", res, t)
				res, _ = strconv.Atoi(v)

			default:
				panic(fmt.Sprintf("unexpected operator: %#v", op))
			}
		}

		if res == e.Result {
			return true
		}

		for i, op := range ops {
			if withConcat {
				switch op {
				case Add:
					ops[i] = Mul
					continue outer
				case Mul:
					ops[i] = Concat
					continue outer
				case Concat:
					ops[i] = Add
				}
			} else {
				switch op {
				case Add:
					ops[i] = Mul
					continue outer
				case Mul:
					ops[i] = Add
				}
			}
		}

		return false
	}
}
