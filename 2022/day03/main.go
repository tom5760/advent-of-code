package day03

import (
	"bufio"

	"github.com/tom5760/advent-of-code/2022/inpututils"
)

func Parse(name string) ([]Rucksack, error) {
	var sacks []Rucksack

	err := inpututils.Scan(name, func(scanner *bufio.Scanner) error {
		for scanner.Scan() {
			line := scanner.Bytes()
			items := make([]Item, len(line))

			for i := range items {
				items[i] = Item(line[i])
			}

			n := len(items) / 2

			sacks = append(sacks, Rucksack{
				Items: items,

				First:  items[:n],
				Second: items[n:],
			})
		}

		return nil
	})

	return sacks, err
}

type Item byte

func (i Item) Priority() int {
	// Assumes items are ASCII letters A-Z a-z.

	b := byte(i)
	if b >= 97 && b <= 122 {
		return int(b - 96)
	}

	if b >= 65 && b <= 90 {
		return int(b - 38)
	}

	panic("unexpected item value")
}

type Rucksack struct {
	Items []Item

	First  []Item
	Second []Item
}

func CommonItem(sacks ...[]Item) Item {
	counts := make(map[Item]int, len(sacks[0]))

	for _, sack := range sacks {
		// Keep track of the unique item we've seen this sack.
		set := make(map[Item]bool, len(sacks[0]))
		for _, item := range sack {
			set[item] = true
		}

		// Add a tally for each unique item.
		for item := range set {
			counts[item]++
		}
	}

	for item, count := range counts {
		// If we've seen an item in each sack, it's the common one.
		if count == len(sacks) {
			return item
		}
	}

	panic("no common item found")
}

func Part1(sacks []Rucksack) int {
	var total int

	for _, sack := range sacks {
		item := CommonItem(sack.First, sack.Second)
		total += item.Priority()
	}

	return total
}

func Part2(sacks []Rucksack) int {
	var total int

	for i := 0; i < len(sacks); i += 3 {
		item := CommonItem(sacks[i].Items, sacks[i+1].Items, sacks[i+2].Items)
		total += item.Priority()
	}

	return total
}
