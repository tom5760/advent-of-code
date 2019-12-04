package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	expectedPart1 = 1716
	expectedPart2 = 1163
)

func TestPart1(t *testing.T) {
	assert.Equal(t, expectedPart1, Part1())
}

func TestPart2(t *testing.T) {
	assert.Equal(t, expectedPart2, Part2())
}
