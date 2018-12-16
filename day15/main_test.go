package main

import (
	"fmt"
	"os"
	"testing"
)

type example struct {
	Rounds  int
	TotalHP int
	Outcome int
}

func TestMain(t *testing.T) {
	tests := []example{
		{
			Rounds:  47,
			TotalHP: 590,
			Outcome: 27730,
		},
		{
			Rounds:  37,
			TotalHP: 982,
			Outcome: 36334,
		},
		{
			Rounds:  46,
			TotalHP: 859,
			Outcome: 39514,
		},
		{
			Rounds:  35,
			TotalHP: 793,
			Outcome: 27755,
		},
		{
			Rounds:  54,
			TotalHP: 536,
			Outcome: 28944,
		},
		{
			Rounds:  20,
			TotalHP: 937,
			Outcome: 18740,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("example%d", i+1), func(t *testing.T) {
			f, err := os.Open(fmt.Sprintf("input-example%d", i+1))
			if err != nil {
				t.Error("failed to open test file:", err)
				return
			}
			defer f.Close()

			arena := readInput(f)
			arena.Battle()

			if test.Rounds != arena.LastRound {
				t.Error("incorrect number of rounds", test.Rounds, arena.LastRound)
			}
			if test.TotalHP != arena.TotalHP() {
				t.Error("incorrect total HP", test.TotalHP, arena.TotalHP())
			}
			if test.Outcome != arena.Outcome() {
				t.Error("incorrect total HP", test.Outcome, arena.Outcome())
			}
		})
	}
}
