package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	inputs := scanInputs(os.Stdin)

	var total int64
	var doneTotal, doneDuplicate bool

	frequencies := make(map[int64]bool)
	frequencies[total] = true

	for !(doneTotal && doneDuplicate) {
		for _, n := range inputs {
			total += n

			if !doneDuplicate && frequencies[total] {
				doneDuplicate = true
				log.Println("(part 2) first duplicated frequency:", total)
			}

			frequencies[total] = true
		}

		if !doneTotal {
			doneTotal = true
			log.Println("(part 1) frequency total:", total)
		}
	}
}

func scanInputs(r io.Reader) []int64 {
	var inputs []int64

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var n int64

		n, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			log.Fatalln("error parsing input:", err)
			return inputs
		}

		inputs = append(inputs, n)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln("error reading input:", err)
	}

	return inputs
}
