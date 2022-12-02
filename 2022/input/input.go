package input

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"sync"
)

type Parser[T any] struct {
	ParseFunc func(inChan <-chan []byte, outChan chan<- T) error
	SplitFunc bufio.SplitFunc
}

func (p *Parser[T]) ReadSlice(r io.Reader) ([]T, error) {
	scanner := bufio.NewScanner(r)

	if p.SplitFunc != nil {
		scanner.Split(p.SplitFunc)
	}

	inChan := make(chan []byte, 2)
	outChan := make(chan T, 2)
	parseChan := make(chan struct{}) // signals parse failure

	var wg sync.WaitGroup
	wg.Add(3)

	var parseErr error
	go func() {
		defer wg.Done()
		defer close(parseChan)

		parseErr = p.ParseFunc(inChan, outChan)
	}()

	go func() {
		defer wg.Done()
		defer close(inChan)

		for scanner.Scan() {
			select {
			case inChan <- bytes.TrimSpace(scanner.Bytes()):
			case <-parseChan:
				return
			}
		}
	}()

	var result []T
	go func() {
		defer wg.Done()

		for r := range outChan {
			result = append(result, r)
		}
	}()

	wg.Wait()

	if err := scanner.Err(); err != nil {
		return result, fmt.Errorf("failed to scan: %w", err)
	}
	if parseErr != nil {
		return result, fmt.Errorf("failed to parse: %w", parseErr)
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
