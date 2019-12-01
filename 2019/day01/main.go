package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	os.Exit(run())
}

func run() int {
	modules, err := readInput(os.Stdin)
	if err != nil {
		log.Println("failed to read input:", err)
		return 1
	}

	moduleTotal := sumModules(modules, FuelForModule)
	log.Println("(part 1) total fuel for modules only:", moduleTotal)

	fuelTotal := sumModules(modules, TotalFuelForModule)
	log.Println("(part 2) total fuel including fuel weight:", fuelTotal)

	return 0
}

func readInput(r io.Reader) ([]uint64, error) {
	var modules []uint64

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		mass, err := strconv.ParseUint(scanner.Text(), 10, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to convert mass to int: %w", err)
		}

		modules = append(modules, mass)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to scan input: %w", err)
	}

	return modules, nil
}

func sumModules(modules []uint64, predicate func(uint64) uint64) uint64 {
	var total uint64

	for _, mass := range modules {
		total += predicate(mass)
	}

	return total
}
