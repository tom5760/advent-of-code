package inpututils

import (
	"bufio"
	"fmt"
	"os"
)

func Scan(name string, scanFunc func(*bufio.Scanner) error) error {
	f, err := os.Open(name)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	if err := scanFunc(scanner); err != nil {
		return err
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("failed to scan input: %w", err)
	}

	return nil
}
