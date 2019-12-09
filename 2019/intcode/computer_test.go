package intcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComputer(t *testing.T) {
	tests := []struct {
		Start, End, Inputs, Outputs []int
	}{
		{
			Start: []int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50},
			End:   []int{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50},
		},
		{
			Start: []int{1, 0, 0, 0, 99},
			End:   []int{2, 0, 0, 0, 99},
		},
		{
			Start: []int{2, 3, 0, 3, 99},
			End:   []int{2, 3, 0, 6, 99},
		},
		{
			Start: []int{2, 4, 4, 5, 99, 0},
			End:   []int{2, 4, 4, 5, 99, 9801},
		},
		{
			Start: []int{1, 1, 1, 4, 99, 5, 6, 0, 99},
			End:   []int{30, 1, 1, 4, 2, 5, 6, 0, 99},
		},

		// Using position mode, consider whether the input is equal to 8; output 1
		// (if it is) or 0 (if it is not).
		{
			Start:   []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8},
			Inputs:  []int{7},
			Outputs: []int{0},
		},
		{
			Start:   []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8},
			Inputs:  []int{8},
			Outputs: []int{1},
		},
		{
			Start:   []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8},
			Inputs:  []int{9},
			Outputs: []int{0},
		},

		// Using position mode, consider whether the input is less than 8; output 1
		// (if it is) or 0 (if it is not).
		{
			Start:   []int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8},
			Inputs:  []int{7},
			Outputs: []int{1},
		},
		{
			Start:   []int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8},
			Inputs:  []int{8},
			Outputs: []int{0},
		},
		{
			Start:   []int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8},
			Inputs:  []int{9},
			Outputs: []int{0},
		},

		// Using immediate mode, consider whether the input is equal to 8; output 1
		// (if it is) or 0 (if it is not).
		{
			Start:   []int{3, 3, 1108, -1, 8, 3, 4, 3, 99},
			Inputs:  []int{7},
			Outputs: []int{0},
		},
		{
			Start:   []int{3, 3, 1108, -1, 8, 3, 4, 3, 99},
			Inputs:  []int{8},
			Outputs: []int{1},
		},
		{
			Start:   []int{3, 3, 1108, -1, 8, 3, 4, 3, 99},
			Inputs:  []int{9},
			Outputs: []int{0},
		},

		// Using immediate mode, consider whether the input is less than 8; output
		// 1 (if it is) or 0 (if it is not).
		{
			Start:   []int{3, 3, 1107, -1, 8, 3, 4, 3, 99},
			Inputs:  []int{7},
			Outputs: []int{1},
		},
		{
			Start:   []int{3, 3, 1107, -1, 8, 3, 4, 3, 99},
			Inputs:  []int{8},
			Outputs: []int{0},
		},
		{
			Start:   []int{3, 3, 1107, -1, 8, 3, 4, 3, 99},
			Inputs:  []int{9},
			Outputs: []int{0},
		},

		// Here are some jump tests that take an input, then output 0 if the input
		// was zero or 1 if the input was non-zero:

		// position mode
		{
			Start:   []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9},
			Inputs:  []int{0},
			Outputs: []int{0},
		},
		{
			Start:   []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9},
			Inputs:  []int{1},
			Outputs: []int{1},
		},
		{
			Start:   []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9},
			Inputs:  []int{-1},
			Outputs: []int{1},
		},

		// immediate mode
		{
			Start:   []int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1},
			Inputs:  []int{0},
			Outputs: []int{0},
		},
		{
			Start:   []int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1},
			Inputs:  []int{1},
			Outputs: []int{1},
		},
		{
			Start:   []int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1},
			Inputs:  []int{-1},
			Outputs: []int{1},
		},

		// The example program uses an input instruction to ask for a single
		// number. The program will then output 999 if the input value is below 8,
		// output 1000 if the input value is equal to 8, or output 1001 if the
		// input value is greater than 8.
		{
			Start: []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006,
				20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46,
				104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99},
			Inputs:  []int{7},
			Outputs: []int{999},
		},
		{
			Start: []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006,
				20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46,
				104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99},
			Inputs:  []int{8},
			Outputs: []int{1000},
		},
		{
			Start: []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006,
				20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46,
				104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99},
			Inputs:  []int{9},
			Outputs: []int{1001},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			computer := NewComputer(tt.Start)

			computer.Inputs(tt.Inputs...)

			go computer.Run()

			outputs := computer.Outputs()

			if tt.Outputs != nil {
				assert.Equal(t, tt.Outputs, outputs)
			}

			if tt.End != nil {
				assert.Equal(t, tt.End, computer.Memory)
			}
		})
	}
}
