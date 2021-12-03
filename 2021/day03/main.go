package main

import (
	"fmt"
	"io"
	"log"
	"math/bits"
	"os"
	"strings"

	"github.com/tom5760/advent-of-code/2021/input"
)

func main() {
	if err := run(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func run() error {
	reports, err := ParseInput(os.Stdin)
	if err != nil {
		return fmt.Errorf("failed to parse input: %w", err)
	}

	fmt.Println("Part 1:", Part1(reports))
	fmt.Println("Part 2:", Part2(reports))

	return nil
}

// The diagnostic report (your puzzle input) consists of a list of binary
// numbers which, when decoded properly, can tell you many useful things about
// the conditions of the submarine. The first parameter to check is the power
// consumption.
func ParseInput(r io.Reader) ([]uint64, error) {
	return input.Parser[uint64]{ParseFunc: input.Uint(2, 64)}.Slice(r)
}

// You need to use the binary numbers in the diagnostic report to generate two
// new binary numbers (called the gamma rate and the epsilon rate). The power
// consumption can then be found by multiplying the gamma rate by the epsilon
// rate.
//
// Each bit in the gamma rate can be determined by finding the most common bit
// in the corresponding position of all numbers in the diagnostic report.
//
// The epsilon rate is calculated in a similar way; rather than use the most
// common bit, the least common bit from each position is used.
func Part1(reports []uint64) uint64 {
	gamma, epsilon := rates(reports)

	return gamma * epsilon
}

// Next, you should verify the life support rating, which can be determined by
// multiplying the oxygen generator rating by the CO2 scrubber rating.
//
// Both the oxygen generator rating and the CO2 scrubber rating are values that
// can be found in your diagnostic report - finding them is the tricky part.
// Both values are located using a similar process that involves filtering out
// values until only one remains. Before searching for either rating value,
// start with the full list of binary numbers from your diagnostic report and
// consider just the first bit of those numbers. Then:
//
// - Keep only numbers selected by the bit criteria for the type of rating
//   value for which you are searching. Discard numbers which do not match the
//   bit criteria.
//
// - If you only have one number left, stop; this is the rating value for which
//   you are searching.
//
// - Otherwise, repeat the process, considering the next bit to the right.
//
// The bit criteria depends on which type of rating value you want to find:
//
// - To find oxygen generator rating, determine the most common value (0 or 1)
//   in the current bit position, and keep only numbers with that bit in that
//   position. If 0 and 1 are equally common, keep values with a 1 in the
//   position being considered.
//
// - To find CO2 scrubber rating, determine the least common value (0 or 1) in
//   the current bit position, and keep only numbers with that bit in that
//   position. If 0 and 1 are equally common, keep values with a 0 in the
//   position being considered.
func Part2(reports []uint64) uint64 {
	oxygen := filterReports(reports, true)
	co2 := filterReports(reports, false)

	return oxygen * co2
}

func bitLength(reports []uint64) int {
	var x int

	for _, report := range reports {
		y := 64 - bits.LeadingZeros64(report)
		if y > x {
			x = y
		}
	}

	return x
}

func rates(reports []uint64) (gamma, epsilon uint64) {
	if len(reports) == 0 {
		return 0, 0
	}

	n := bitLength(reports)
	counts := make([]int, n)

	for _, report := range reports {
		for i := 0; i < n; i++ {
			switch (report >> i) & 1 {
			case 0:
				counts[i]--
			case 1:
				counts[i]++
			}
		}
	}

	for i, count := range counts {
		if count >= 0 {
			gamma |= (1 << i)
		} else {
			epsilon |= (1 << i)
		}
	}

	return gamma, epsilon
}

func filterReports(reports []uint64, useGamma bool) uint64 {
	cur := make([]uint64, len(reports))
	next := make([]uint64, 0, len(reports))

	copy(cur, reports)

	n := bitLength(reports)

	for i := n - 1; i >= 0; i-- {
		gamma, epsilon := rates(cur)

		filter := gamma
		if !useGamma {
			filter = epsilon
		}

		bit := (filter >> i & 1)

		for _, r := range cur {
			if (r>>i)&1 == bit {
				next = append(next, r)
			}
		}

		if len(next) == 1 {
			return next[0]
		}

		cur, next = next, cur
		next = next[:0]
	}

	panic("failed to filter to single value")
}

func printReports(reports []uint64) string {
	var sb strings.Builder

	n := bitLength(reports)

	sb.WriteByte('[')

	for _, r := range reports {
		fmt.Fprintf(&sb, "%.*b", n, r)
		sb.WriteByte(' ')
	}

	sb.WriteByte(']')

	return sb.String()
}
