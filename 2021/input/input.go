// Package input has convenience functions for common input types.
package input

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

// Tinkering with Go 1.18's generics.  Good reference:
//   https://bitfieldconsulting.com/golang/generics

type Parser[T any] struct {
	ParseFunc func(*bufio.Scanner) (T, error)

	SplitFunc bufio.SplitFunc // if nil, defaults to SplitLines
}

func (p Parser[T]) Slice(r io.Reader) ([]T, error) {
	scanner := bufio.NewScanner(r)

	if p.SplitFunc != nil {
		scanner.Split(p.SplitFunc)
	}

	var slice []T

	for scanner.Scan() {
		v, err := p.ParseFunc(scanner)
		if err != nil {
			return nil, fmt.Errorf("failed to parse line: %w", err)
		}

		slice = append(slice, v)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to scan input: %w", err)
	}

	return slice, nil
}

func Int(base, bitSize int) func(*bufio.Scanner) (int64, error) {
	return func(scanner *bufio.Scanner) (int64, error) {
		return strconv.ParseInt(scanner.Text(), base, bitSize)
	}
}

func String(scanner *bufio.Scanner) (string, error) {
	return scanner.Text(), nil
}
