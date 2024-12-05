package aoc

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"iter"
	"strconv"
)

type Scanner struct {
	s *bufio.Scanner
}

func NewScanner(r io.Reader) *Scanner {
	return &Scanner{
		s: bufio.NewScanner(r),
	}
}

func (s *Scanner) Split(split bufio.SplitFunc) {
	s.s.Split(split)
}

func (s *Scanner) Err() error {
	return s.s.Err()
}

func (s *Scanner) Scan() iter.Seq[[]byte] {
	return func(yield func([]byte) bool) {
		for s.s.Scan() {
			if !yield(s.s.Bytes()) {
				return
			}
		}
	}
}

func IterRowsInt(r io.Reader) iter.Seq2[[]int, error] {
	return func(yield func([]int, error) bool) {
		scanner := bufio.NewScanner(r)
		for scanner.Scan() {
			fields := bytes.Fields(scanner.Bytes())
			ints := make([]int, len(fields))

			for i, field := range fields {
				v, err := strconv.Atoi(string(field))
				if err != nil {
					yield(nil, fmt.Errorf("failed to parse field: %w", err))
					return
				}

				ints[i] = v
			}

			if !yield(ints, nil) {
				return
			}
		}
		if err := scanner.Err(); err != nil {
			yield(nil, fmt.Errorf("failed to scan: %w", err))
			return
		}
	}
}
