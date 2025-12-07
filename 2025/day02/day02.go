package day02

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"slices"
	"strconv"
	"unicode"
	"unicode/utf8"
)

func Run(l *log.Logger, r io.Reader) (int, int, error) {
	var part1, part2 int

	scanner := bufio.NewScanner(r)
	scanner.Split(ScanFields(','))

	for scanner.Scan() {
		line := bytes.TrimSpace(scanner.Bytes())
		parts := bytes.Split(line, []byte{'-'})
		if len(parts) != 2 {
			return 0, 0, fmt.Errorf("unexpected line format: %q", string(line))
		}

		start, err := strconv.Atoi(string(parts[0]))
		if err != nil {
			return 0, 0, fmt.Errorf("failed to parse start: %w", err)
		}

		end, err := strconv.Atoi(string(parts[1]))
		if err != nil {
			return 0, 0, fmt.Errorf("failed to parse end: %w", err)
		}

		for i := start; i <= end; i++ {
			if isInvalid, isPart1 := IsInvalid(i); isInvalid {
				l.Printf("%v-%v invalid: %v", start, end, i)
				if isPart1 {
					part1 += i
				}

				part2 += i
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return 0, 0, fmt.Errorf("failed to scan input: %w", err)
	}

	return part1, part2, nil
}

// ScanFields is a split function for a [bufio.Scanner] that returns each
// delimiter-separated word of text, with surrounding spaces deleted.  It will
// never return an empty string.
//
// Modified from [bufio.ScanWords].
func ScanFields(d rune) func([]byte, bool) (int, []byte, error) {
	return func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		// Skip leading spaces.
		start := 0
		for width := 0; start < len(data); start += width {
			var r rune
			r, width = utf8.DecodeRune(data[start:])
			if !unicode.IsSpace(r) {
				break
			}
		}

		// Scan until delimiter, marking end of word.
		for width, i := 0, start; i < len(data); i += width {
			var r rune
			r, width = utf8.DecodeRune(data[i:])
			if r == d {
				return i + width, data[start:i], nil
			}
		}

		// If we're at EOF, we have a final, non-empty, non-terminated word.
		// Return it.
		if atEOF && len(data) > start {
			return len(data), data[start:], nil
		}

		// Request more data.
		return start, nil, nil
	}
}

func IsInvalid(v int) (bool, bool) {
	str := []byte(strconv.Itoa(v))

	half := len(str) / 2
	even := len(str)%2 == 0

outer:
	for i := half; i > 0; i-- {
		p := str[:i]

		for chunk := range slices.Chunk(str, i) {
			if !bytes.Equal(p, chunk) {
				continue outer
			}
		}

		return true, i == half && even
	}

	return false, false
}
