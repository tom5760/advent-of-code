package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	ids := readInput(os.Stdin)

	for _, a := range ids {
		for _, b := range ids {
			if len(a) != len(b) {
				continue
			}

			var diffCount uint

			for i := 0; i < len(a); i++ {
				if a[i] != b[i] {
					diffCount++
				}
			}

			if diffCount == 1 {
				log.Println("found ID with only one difference")
				log.Println(a)
				log.Println(b)

				var sb strings.Builder

				for i := 0; i < len(a); i++ {
					if a[i] == b[i] {
						sb.WriteByte(a[i])
					}
				}

				log.Println("common letters:", sb.String())

				return
			}
		}
	}
}

func readInput(r io.Reader) []string {
	scanner := bufio.NewScanner(r)

	var ids []string

	for scanner.Scan() {
		ids = append(ids, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln("err scanning input:", err)
	}

	return ids
}
