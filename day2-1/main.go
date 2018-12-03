package main

import (
	"bufio"
	"bytes"
	"log"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var twos, threes uint

	for scanner.Scan() {
		var idTwos, idThrees uint
		id := scanner.Bytes()

		for char := byte('a'); char <= 'z'; char++ {
			switch bytes.Count(id, []byte{char}) {
			case 2:
				idTwos++

			case 3:
				idThrees++
			}
		}

		if idTwos > 0 {
			twos++
		}

		if idThrees > 0 {
			threes++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln("err scanning input:", err)
	}

	log.Println("checksum:", twos*threes)
}
