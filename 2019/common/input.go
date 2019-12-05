package common

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strconv"
)

// ReadStringSlice scans r and parses it into a slice of strings.  If split is
// nil, each number is separated by a newline.  Otherwise, the input is split
// with that function.
func ReadStringSlice(r io.Reader, split bufio.SplitFunc) ([]string, error) {
	var inputs []string

	scanner := bufio.NewScanner(r)

	if split != nil {
		scanner.Split(split)
	}

	for scanner.Scan() {
		inputs = append(inputs, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to scan input: %w", err)
	}

	return inputs, nil
}

// ReadIntSlice scans r and parses it into a slice of Ints.  If split is
// nil, each number is separated by a newline.  Otherwise, the input is split
// with that function.
func ReadIntSlice(r io.Reader, split bufio.SplitFunc) ([]int, error) {
	var inputs []int

	scanner := bufio.NewScanner(r)

	if split != nil {
		scanner.Split(split)
	}

	for scanner.Scan() {
		value, err := strconv.ParseInt(scanner.Text(), 10, 0)
		if err != nil {
			return nil, fmt.Errorf("failed to convert value to int: %w", err)
		}

		inputs = append(inputs, int(value))
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to scan input: %w", err)
	}

	return inputs, nil
}

// ReadUint64Slice scans r and parses it into a slice of uint64s.  If split is
// nil, each number is separated by a newline.  Otherwise, the input is split
// with that function.
func ReadUint64Slice(r io.Reader, split bufio.SplitFunc) ([]uint64, error) {
	var inputs []uint64

	scanner := bufio.NewScanner(r)

	if split != nil {
		scanner.Split(split)
	}

	for scanner.Scan() {
		value, err := strconv.ParseUint(scanner.Text(), 10, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to convert value to int: %w", err)
		}

		inputs = append(inputs, value)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to scan input: %w", err)
	}

	return inputs, nil
}

// ScanCommas splits input on commas.  Based on ScanLines from the stdlib:
// https://golang.org/src/bufio/scan.go?s=11802:11880#L335
func ScanCommas(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	if i := bytes.IndexByte(data, ','); i >= 0 {
		return i + 1, bytes.TrimSpace(data[0:i]), nil
	}

	if atEOF {
		return len(data), bytes.TrimSpace(data), nil
	}

	// Request more data.
	return 0, nil, nil
}
