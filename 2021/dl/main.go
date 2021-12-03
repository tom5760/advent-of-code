package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/peterbourgon/ff/v3"
)

const (
	pathInput = "https://adventofcode.com/2021/day/%d/input"
)

func main() {
	if err := run(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func run() error {
	fs := flag.NewFlagSet("input", flag.ExitOnError)

	token := fs.String("token", "", "Session token from adventofcode.com.  Get from cookies.")
	day := fs.Uint("day", 0, "Day to fetch input from.")

	if err := ff.Parse(fs, os.Args[1:], ff.WithEnvVarPrefix("AOC")); err != nil {
		return fmt.Errorf("failed to parse token: %w", err)
	}

	if *token == "" {
		return fmt.Errorf("session token required")
	}

	if *day == 0 {
		return fmt.Errorf("puzzle day required")
	}

	outPath := fmt.Sprintf("day%.2d/input", *day)
	outf, err := os.Create(outPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}

	defer outf.Close()

	inputURL := fmt.Sprintf(pathInput, *day)
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, inputURL, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: *token,
	})

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to do request: %w", err)
	}

	if _, err := io.Copy(outf, resp.Body); err != nil {
		return fmt.Errorf("failed to copy response: %w", err)
	}

	return nil
}
