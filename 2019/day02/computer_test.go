package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComputer(t *testing.T) {
	tests := []struct {
		Start, End []uint64
	}{
		{
			Start: []uint64{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50},
			End:   []uint64{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50},
		},
		{
			Start: []uint64{1, 0, 0, 0, 99},
			End:   []uint64{2, 0, 0, 0, 99},
		},
		{
			Start: []uint64{2, 3, 0, 3, 99},
			End:   []uint64{2, 3, 0, 6, 99},
		},
		{
			Start: []uint64{2, 4, 4, 5, 99, 0},
			End:   []uint64{2, 4, 4, 5, 99, 9801},
		},
		{
			Start: []uint64{1, 1, 1, 4, 99, 5, 6, 0, 99},
			End:   []uint64{30, 1, 1, 4, 2, 5, 6, 0, 99},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			computer := NewComputer(tt.Start)
			if err := computer.Run(); err != nil {
				t.Errorf("computer failed to run: %v", err)
			}

			assert.Equal(t, tt.End, computer.Memory)
		})
	}
}
