package day05

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"slices"
	"strconv"
)

func Run(r io.Reader) (int, int, error) {
	rules, updates, err := Parse(r)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to parse input: %w", err)
	}

	var part1, part2 int

	for _, update := range updates {
		i := update.IsCorrect(rules)

		if i == -1 {
			part1 += update[(len(update)-1)/2]
		} else {
			for i != -1 {
				update.Fix(rules[i])
				i = update.IsCorrect(rules)
			}

			part2 += update[(len(update)-1)/2]
		}
	}

	return part1, part2, nil
}

type Rule struct {
	A, B int
}

type Update []int

func Parse(r io.Reader) ([]Rule, []Update, error) {
	var (
		rules   []Rule
		updates []Update
	)

	scanner := bufio.NewScanner(r)
	onSection1 := true

	for scanner.Scan() {
		line := scanner.Bytes()
		if len(line) == 0 {
			onSection1 = false
			continue
		}

		if onSection1 {
			rule, err := parseRule(line)
			if err != nil {
				return nil, nil, fmt.Errorf("failed to parse rule: %w", err)
			}

			rules = append(rules, rule)

		} else {
			update, err := parseUpdate(line)
			if err != nil {
				return nil, nil, fmt.Errorf("failed to parse update: %w", err)
			}

			updates = append(updates, update)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, fmt.Errorf("failed to scan input: %w", err)
	}

	return rules, updates, nil
}

func parseRule(line []byte) (Rule, error) {
	parts := bytes.Split(line, []byte{'|'})
	if len(parts) != 2 {
		return Rule{}, errors.New("unexpected rule format")
	}

	a, err := strconv.Atoi(string(parts[0]))
	if err != nil {
		return Rule{}, fmt.Errorf("failed to parse left rule: %w", err)
	}

	b, err := strconv.Atoi(string(parts[1]))
	if err != nil {
		return Rule{}, fmt.Errorf("failed to parse right rule: %w", err)
	}

	r := Rule{A: a, B: b}

	return r, nil
}

func parseUpdate(line []byte) (Update, error) {
	parts := bytes.Split(line, []byte{','})
	if len(parts) < 2 {
		return nil, fmt.Errorf("unexpected update format: %q", string(line))
	}

	update := make(Update, len(parts))

	for i, part := range parts {
		v, err := strconv.Atoi(string(part))
		if err != nil {
			return nil, fmt.Errorf("failed to parse update value: %w", err)
		}

		update[i] = v
	}

	return update, nil
}

func (u Update) IsCorrect(rules []Rule) int {
	for i, rule := range rules {
		ai := slices.Index(u, rule.A)
		bi := slices.Index(u, rule.B)

		if ai == -1 || bi == -1 {
			continue
		}

		if ai > bi {
			return i
		}
	}

	return -1
}

func (u Update) Fix(rule Rule) {
	ai := slices.Index(u, rule.A)
	bi := slices.Index(u, rule.B)

	u[ai], u[bi] = u[bi], u[ai]
}
