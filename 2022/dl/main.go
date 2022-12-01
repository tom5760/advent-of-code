package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

const (
	pathInput = "https://adventofcode.com/2022/day/%d/input"
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
	dayStr := fs.String("day", "", "Day to fetch input from.")

	if err := fs.Parse(os.Args[1:]); err != nil {
		return fmt.Errorf("failed to parse arguments: %w", err)
	}

	if *token == "" {
		return fmt.Errorf("session token required")
	}

	if *dayStr == "" {
		return fmt.Errorf("puzzle day required")
	}

	day, err := strconv.ParseInt(*dayStr, 10, 64)
	if err != nil {
		return fmt.Errorf("failed to parse day: %w", err)
	}

	outPath := fmt.Sprintf("day%.2d/input", day)
	outf, err := os.Create(outPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}

	defer outf.Close()

	inputURL := fmt.Sprintf(pathInput, day)
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
