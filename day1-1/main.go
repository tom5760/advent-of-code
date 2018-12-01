package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var total int64

	for scanner.Scan() {
		var n int64
		if _, err := fmt.Sscanf(scanner.Text(), "%d", &n); err != nil {
			log.Fatalln("error reading input:", err)
			return
		}

		total += n
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln("err scanning input:", err)
	}

	log.Println("Total:", total)
}
