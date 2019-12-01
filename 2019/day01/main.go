package main

import (
	"log"
	"os"

	"github.com/tom5760/advent-of-code/2019/common"
)

func main() {
	os.Exit(run())
}

func run() int {
	modules, err := common.ReadUint64Slice(os.Stdin)
	if err != nil {
		log.Println("failed to read input:", err)
		return 1
	}

	log.Println("(part 1) total fuel for modules only:", Part1(modules))
	log.Println("(part 2) total fuel including fuel weight:", Part2(modules))

	return 0
}

// What is the sum of the fuel requirements for all of the modules on your
// spacecraft?
func Part1(modules []uint64) uint64 {
	return common.MapSumUint64(modules, FuelForModule)
}

// What is the sum of the fuel requirements for all of the modules on your
// spacecraft when also taking into account the mass of the added fuel?
func Part2(modules []uint64) uint64 {
	return common.MapSumUint64(modules, TotalFuelForModule)
}
