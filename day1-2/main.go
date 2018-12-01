package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	inputs := scanInputs(os.Stdin)

	var total int64

	frequencies := make(map[int64]bool)
	frequencies[total] = true

	for {
		for _, n := range inputs {
			total += n

			if frequencies[total] {
				log.Println("Frequency:", total)
				return
			}
			frequencies[total] = true
		}
	}
}

func scanInputs(r io.Reader) []int64 {
	inputs := make([]int64, 0)

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		var n int64

		if _, err := fmt.Sscanf(scanner.Text(), "%d", &n); err != nil {
			log.Fatalln("error reading input:", err)
			return inputs
		}

		inputs = append(inputs, n)
	}

	return inputs
}
