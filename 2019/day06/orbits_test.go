package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTotalOrbits(t *testing.T) {
	tests := []struct {
		lines []string
		total int
	}{
		{
			lines: []string{
				"COM)B", "B)C", "C)D", "D)E", "E)F", "B)G", "G)H", "D)I", "E)J", "J)K",
				"K)L",
			},
			total: 42,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			u := ParseUniverse(tt.lines)
			actual := TotalOrbits(u)
			assert.Equal(t, tt.total, actual)
		})
	}
}

func TestTotalTransfers(t *testing.T) {
	tests := []struct {
		lines []string
		total int
	}{
		{
			lines: []string{
				"COM)B", "B)C", "C)D", "D)E", "E)F", "B)G", "G)H", "D)I", "E)J", "J)K",
				"K)L", "K)YOU", "I)SAN",
			},
			total: 4,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			u := ParseUniverse(tt.lines)
			source := u.Planets["YOU"]
			target := u.Planets["SAN"]

			actual := TotalTransfers(u, source, target)
			assert.Equal(t, tt.total, actual)
		})
	}
}
