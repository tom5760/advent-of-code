package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strings"
)

const (
	initialStateHeader = "initial state: "

	ruleWidth = 5
)

type Garden map[int64]bool

func (g Garden) HasPlant(i int64) bool {
	_, ok := g[i]
	return ok
}

func (g Garden) Tick(rules []Rule) Garden {
	next := make(Garden)

	minI, maxI := g.bounds()

	for i := minI; i <= maxI; i++ {
		if g.ApplyRules(i, rules) {
			next[i] = true
		}
	}

	return next
}

func (g Garden) ApplyRules(i int64, rules []Rule) bool {
	left2 := g.HasPlant(i - 2)
	left1 := g.HasPlant(i - 1)
	center := g.HasPlant(i)
	right1 := g.HasPlant(i + 1)
	right2 := g.HasPlant(i + 2)

	for _, rule := range rules {
		matches, hasPlant := rule.Apply(left2, left1, center, right1, right2)
		if matches {
			return hasPlant
		}
	}
	return false
}

func (g Garden) Count() int64 {
	var c int64

	for i := range g {
		c += i
	}

	return c
}

func (g Garden) String() string {
	var sb strings.Builder

	minI, maxI := g.bounds()
	for i := minI; i <= maxI; i++ {
		if _, ok := g[i]; ok {
			sb.WriteByte('#')
		} else {
			sb.WriteByte('.')
		}
	}

	return sb.String()
}

func (g Garden) bounds() (minI, maxI int64) {
	minI = math.MaxInt64
	maxI = math.MinInt64

	for i := range g {
		if i < minI {
			minI = i
		}
		if i > maxI {
			maxI = i
		}
	}

	// Need to make sure to cover the whole rule.
	return minI - ruleWidth, maxI + ruleWidth
}

type Rule struct {
	Left2, Left1, Center, Right1, Right2, Result bool
}

func (r Rule) String() string {
	left := plantsToString(r.Left2, r.Left1, r.Center, r.Right1, r.Right2)
	var result string
	if r.Result {
		result = "#"
	} else {
		result = "."
	}

	return fmt.Sprint(left, " => ", result)
}

func (r Rule) Apply(left2, left1, center, right1, right2 bool) (matches, hasPlant bool) {
	if (r.Left2 == left2) && (r.Left1 == left1) && (r.Center == center) &&
		(r.Right1 == right1) && (r.Right2 == right2) {
		return true, r.Result
	}

	return false, false
}

func main() {
	garden, rules := readInput(os.Stdin)

	log.Println("(part 1) sum:", simulate(garden, rules, 20))
	log.Println("(part 2) sum:", simulate(garden, rules, 50000000000))
}

func readInput(r io.Reader) (Garden, []Rule) {
	scanner := bufio.NewScanner(r)

	// read initial state line
	scanner.Scan()

	garden := make(Garden)
	for i, c := range scanner.Text()[len(initialStateHeader):] {
		if c == '#' {
			garden[int64(i)] = true
		}
	}

	// skip the blank line
	scanner.Scan()

	var rules []Rule

	// read rules
	for scanner.Scan() {
		bytes := scanner.Bytes()

		rules = append(rules, Rule{
			Left2:  bytes[0] == '#',
			Left1:  bytes[1] == '#',
			Center: bytes[2] == '#',
			Right1: bytes[3] == '#',
			Right2: bytes[4] == '#',
			Result: bytes[9] == '#',
		})
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln("failed to read input:", err)
		return nil, nil
	}

	return garden, rules
}

func simulate(garden Garden, rules []Rule, generations uint64) int64 {
	step := generations / 100
	for step*100 < generations {
		step++
	}

	for i := uint64(0); i < generations; i++ {
		garden = garden.Tick(rules)
		if i%step == 0 {
			log.Printf("%.0f%% complete (generation %d of %d)",
				float64(i)/float64(generations)*100, i, generations)
		}
	}

	return garden.Count()
}

func plantsToString(left2, left1, center, right1, right2 bool) string {
	var sb strings.Builder

	if left2 {
		sb.WriteByte('#')
	} else {
		sb.WriteByte('.')
	}
	if left1 {
		sb.WriteByte('#')
	} else {
		sb.WriteByte('.')
	}
	if center {
		sb.WriteByte('#')
	} else {
		sb.WriteByte('.')
	}
	if right1 {
		sb.WriteByte('#')
	} else {
		sb.WriteByte('.')
	}
	if right2 {
		sb.WriteByte('#')
	} else {
		sb.WriteByte('.')
	}

	return sb.String()
}
