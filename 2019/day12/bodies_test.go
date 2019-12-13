package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBodies_Step(t *testing.T) {
	tests := []struct {
		bodies Bodies
		steps  []Bodies
	}{
		{
			bodies: Bodies{
				{X: -1, Y: 0, Z: 2},
				{X: 2, Y: -10, Z: -7},
				{X: 4, Y: -8, Z: 8},
				{X: 3, Y: 5, Z: -1},
			},
			steps: []Bodies{
				Bodies{
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
				Bodies{
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
			for _, step := range tt.steps {
				assert.Equal(t, step, tt.bodies)
				tt.bodies.Step()
			}
		})
	}
}
