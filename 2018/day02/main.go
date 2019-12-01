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

	log.Println("(part 1) checksum:", checksum(ids))
	log.Println("(part 2) common letters:", commonLetters(ids))
}

func readInput(r io.Reader) []string {
	var ids []string

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		ids = append(ids, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln("error reading input:", err)
	}

	return ids
}

func checksum(ids []string) uint64 {
	var twos, threes uint64

	for _, id := range ids {
		hasTwo, hasThree := countTwosThrees(id)
		if hasTwo {
			twos++
		}
		if hasThree {
			threes++
		}
	}

	return twos * threes
}

func countTwosThrees(id string) (hasTwo, hasThree bool) {
	letters := make(map[rune]uint64, len(id))
	for _, r := range id {
		letters[r]++
	}

	for _, counts := range letters {
		switch counts {
		case 2:
			hasTwo = true
		case 3:
			hasThree = true
		}
	}

	return hasTwo, hasThree
}

func commonLetters(ids []string) string {
	for _, a := range ids {
		for _, b := range ids {
			if len(a) != len(b) {
				continue
			}

			var diffCount uint64

			for i := 0; i < len(a); i++ {
				if a[i] != b[i] {
					diffCount++
				}
			}

			if diffCount == 1 {
				var sb strings.Builder

				for i := 0; i < len(a); i++ {
					if a[i] == b[i] {
						sb.WriteByte(a[i])
					}
				}

				return sb.String()
			}
		}
	}
	panic("couldn't find two correct box IDs")
}
