// Package input has convenience functions for common input types.
package input

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strconv"
	"strings"
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
		return strconv.ParseInt(strings.TrimSpace(scanner.Text()), base, bitSize)
	}
}

func Uint(base, bitSize int) func(*bufio.Scanner) (uint64, error) {
	return func(scanner *bufio.Scanner) (uint64, error) {
		return strconv.ParseUint(strings.TrimSpace(scanner.Text()), base, bitSize)
	}
}

func String(scanner *bufio.Scanner) (string, error) {
	return scanner.Text(), nil
}

func ScanIndexByte(sep byte) bufio.SplitFunc {
	return func (data []byte, atEOF bool) (advance int, token []byte, err error) {
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}

		if i := bytes.IndexByte(data, ','); i >= 0 {
			return i + 1, dropCR(data[0:i]), nil
		}

		// If we're at EOF, we have a final, non-terminated line. Return it.
		if atEOF {
			return len(data), dropCR(data), nil
		}

		// Request more data.
		return 0, nil, nil
	}
}

func ScanIndex(sep []byte) bufio.SplitFunc {
	return func (data []byte, atEOF bool) (advance int, token []byte, err error) {
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}

		if i := bytes.Index(data, sep); i >= 0 {
			return i + len(sep), dropCR(data[0:i]), nil
		}

		// If we're at EOF, we have a final, non-terminated line. Return it.
		if atEOF {
			return len(data), dropCR(data), nil
		}

		// Request more data.
		return 0, nil, nil
	}
}

// copied from bufio package.
// dropCR drops a terminal \r from the data.
func dropCR(data []byte) []byte {
	if len(data) > 0 && data[len(data)-1] == '\r' {
		return data[0 : len(data)-1]
	}
	return data
}

