package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tom5760/advent-of-code/2019/common"
)

func TestBodies_Step(t *testing.T) {
	tests := []struct {
		input string
		steps [][]Body
	}{
		{
			input: `<x=-1, y=0, z=2>
<x=2, y=-10, z=-7>
<x=4, y=-8, z=8>
<x=3, y=5, z=-1>`,
			steps: [][]Body{
				[]Body{
					{
						X: -1, Y: 0, Z: 2,
						DX: 0, DY: 0, DZ: 0,
					},
					{
						X: 2, Y: -10, Z: -7,
						DX: 0, DY: 0, DZ: 0,
					},
					{
						X: 4, Y: -8, Z: 8,
						DX: 0, DY: 0, DZ: 0,
					},
					{
						X: 3, Y: 5, Z: -1,
						DX: 0, DY: 0, DZ: 0,
					},
				},
				[]Body{
					{
						X: 2, Y: -1, Z: 1,
						DX: 3, DY: -1, DZ: -1,
					},
					{
						X: 3, Y: -7, Z: -4,
						DX: 1, DY: 3, DZ: 3,
					},
					{
						X: 1, Y: -7, Z: 5,
						DX: -3, DY: 1, DZ: -3,
					},
					{
						X: 2, Y: 2, Z: 0,
						DX: -1, DY: -3, DZ: 1,
					},
				},
			},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			lines, err := common.ReadStringSlice(strings.NewReader(tt.input), nil)
			if !assert.NoError(t, err) {
				return
			}

			s, err := ParseSystem(lines)
			if !assert.NoError(t, err) {
				return
			}

			for _, step := range tt.steps {
				assert.Equal(t, step, s.Bodies)
				s.Step()
			}
		})
	}
}
