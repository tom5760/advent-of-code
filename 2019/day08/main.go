package main

import (
	"fmt"
	"image/png"
	"log"
	"math"
	"os"

	"github.com/tom5760/advent-of-code/2019/common"
	"github.com/tom5760/advent-of-code/2019/sif"
)

const (
	width  = 25
	height = 6

	outputFile = "day08-part2.png"
)

func main() {
	os.Exit(run())
}

func run() int {
	input, err := common.ReadIntSlice(os.Stdin, common.ScanDigits)
	if err != nil {
		log.Println("failed to read input:", err)
		return 1
	}

	img, err := sif.Decode(width, height, input)
	if err != nil {
		log.Println("failed to parse input:", err)
		return 1
	}

	log.Println("(part 1):", Part1(img))

	if err := Part2(img); err != nil {
		log.Println("failed to do part 2:", err)
		return 1
	}

	log.Println("(part 2): written to", outputFile)

	return 0
}

// Part1 - Find the layer that contains the fewest 0 digits. On that layer,
// what is the number of 1 digits multiplied by the number of 2 digits?
func Part1(img *sif.Image) int {
	minZeroCount := math.MaxInt64
	minZeroLayer := 0

	for i, layer := range img.Layers {
		var zeroCount int
		for _, px := range layer {
			if px == 0 {
				zeroCount++
			}
		}
		if zeroCount < minZeroCount {
			minZeroCount = zeroCount
			minZeroLayer = i
		}
	}

	var oneCount, twoCount int
	for _, px := range img.Layers[minZeroLayer] {
		if px == 1 {
			oneCount++
		}
		if px == 2 {
			twoCount++
		}
	}

	return oneCount * twoCount
}

// Part2 - What message is produced after decoding your image?
func Part2(img *sif.Image) error {
	f, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer f.Close()

	if err := png.Encode(f, img); err != nil {
		return fmt.Errorf("failed to encode PNG: %w", err)
	}

	return nil
}
