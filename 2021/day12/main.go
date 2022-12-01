package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"golang.org/x/exp/slices"
)

func main() {
	if err := run(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func run() error {
	graph, err := ParseInput(os.Stdin)
	if err != nil {
		return fmt.Errorf("failed to parse input: %w", err)
	}

	fmt.Println("Part 1:", Part1(graph))
	// fmt.Println("Part 2:", Part2(grid))

	return nil
}

func ParseInput(r io.Reader) (map[string][]string, error) {
	graph := make(map[string][]string)
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		left, right, found := strings.Cut(scanner.Text(), "-")
		if !found {
			return nil, fmt.Errorf("malformed string")
		}

		graph[left] = append(graph[left], right)
		graph[right] = append(graph[right], left)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to scan input: %w", err)
	}

	return graph, nil
}

func Part1(graph map[string][]string) uint64 {
	var total uint64

	cur := "start"

	visited := make(map[string]uint8, len(graph))
	visited[cur]++

	frontier := slices.Clone(graph[cur])

	for len(frontier) > 0 {
		cur = frontier[0]
		frontier = slices.Delete(frontier, 0, 1)
		visited[cur]++

		if cur == "end" {
			total++
			continue
		}
		if strings.ToUpper(cur) == cur && visited[cur] >= 2 {
			continue
		}
		if strings.ToUpper(cur) != cur && visited[cur] >= 1 {
			continue
		}

		frontier = append(frontier, graph[cur]...)
	}

	log.Println("VISITED:", visited)
	return total
}
