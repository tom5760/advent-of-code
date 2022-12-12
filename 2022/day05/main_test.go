package day05

import (
	"testing"

	"github.com/tom5760/advent-of-code/2022/testutils"
)

func TestDay05(t *testing.T) {
	t.Parallel()

	testutils.Run(t, Parse, Part1, Part2, []testutils.Test[string, string]{
		{
			Name:  "example",
			Part1: "CMZ",
			Part2: "MCD",
		},
		{
			Name:  "input",
			Part1: "MQTPGLLDN",
			Part2: "LVZPSTTCZ",
		},
	})
}
