package main

import (
	"fmt"
	"image/png"
	"log"
	"math"
	"os"

	"github.com/tom5760/advent-of-code/2019/common"
	"github.com/tom5760/advent-of-code/2019/intcode"
	"github.com/tom5760/advent-of-code/2019/sif"
)

const (
	outputFile = "day11-part2.png"
)

func main() {
	os.Exit(run())
}

func run() int {
	program, err := common.ReadIntSlice(os.Stdin, common.ScanCommas)
	if err != nil {
		log.Println("failed to read input:", err)
		return 1
	}

	log.Println("(part 1) panels painted:", Part1(program))

	if err := Part2(program); err != nil {
		log.Println("failed to do part 2:", err)
		return 1
	}

	log.Println("(part 2): written to", outputFile)

	return 0
}

// Part1 - Build a new emergency hull painting robot and run the Intcode
// program on it. How many panels does it paint at least once?
func Part1(program []int) int {
	colorChan := make(chan int, 1)
	cmdChan := make(chan int, 2)

	computer := intcode.NewComputer(program)

	computer.InputChan = colorChan
	computer.OutputChan = cmdChan

	go computer.Run()

	robot := NewRobot()
	robot.Run(colorChan, cmdChan)

	return len(robot.Hull)
}

// Part2 - After starting the robot on a single white panel instead, what
// registration identifier does it paint on your hull?
func Part2(program []int) error {
	colorChan := make(chan int, 1)
	cmdChan := make(chan int, 2)

	computer := intcode.NewComputer(program)

	computer.InputChan = colorChan
	computer.OutputChan = cmdChan

	go computer.Run()

	robot := NewRobot()
	robot.Hull[Point{0, 0}] = sif.ColorWhite
	robot.Run(colorChan, cmdChan)

	minX := math.MaxInt64
	minY := math.MaxInt64
	maxX := math.MinInt64
	maxY := math.MinInt64

	for p := range robot.Hull {
		if p.X < minX {
			minX = p.X
		}
		if p.X > maxX {
			maxX = p.X
		}
		if p.Y < minY {
			minY = p.Y
		}
		if p.Y > maxY {
			maxY = p.Y
		}
	}

	img := sif.Image{
		Width:  (maxX - minX) + 1,
		Height: (maxY - minY) + 1,
	}

	for p, c := range robot.Hull {
		img.Set(p.X-minX, p.Y-minY, c)
	}

	f, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer f.Close()

	if err := png.Encode(f, &img); err != nil {
		return fmt.Errorf("failed to encode PNG: %w", err)
	}

	return nil
}
