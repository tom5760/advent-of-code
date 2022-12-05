package input

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

type (
	ParseFunc[T any] func(input []byte, outChan chan<- T) error

	Parser[T any] struct {
		TokenFunc bufio.SplitFunc
		ParseFunc ParseFunc[T]
	}
)

func (p *Parser[T]) ReadSlice(r io.Reader) ([]T, error) {
	scanner := bufio.NewScanner(r)

	if p.TokenFunc != nil {
		scanner.Split(p.TokenFunc)
	}

	var result []T
	outChan := make(chan T, 1)

	for scanner.Scan() {
		input := bytes.TrimSpace(scanner.Bytes())
		p.ParseFunc(input, outChan)

		select {
		case v := <-outChan:
			result = append(result, v)
		default:
		}
	}

	if err := scanner.Err(); err != nil {
		return result, fmt.Errorf("failed to scan: %w", err)
	}

	return result, nil
}

func (p *Parser[T]) ReadFileSlice(filePath string) ([]T, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}

	defer f.Close()

	return p.ReadSlice(f)
}
