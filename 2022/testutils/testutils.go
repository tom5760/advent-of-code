package testutils

import "testing"

type (
	Test[T any] struct {
		Name  string
		Part1 T
		Part2 T
	}

	Tests[T any] []Test[T]

	ParseFunc[T any]   func(name string) (T, error)
	PartFunc[I, O any] func(input I) O
)

var GT *testing.T

func Run[I any, O comparable](
	t *testing.T,
	parse ParseFunc[I],
	part1, part2 PartFunc[I, O],
	tests []Test[O],
) {
	t.Helper()

	for _, test := range tests {
		test := test

		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			t.Helper()
			GT = t

			input, err := parse(test.Name)
			if err != nil {
				t.Fatal("failed to parse input:", err)
				return
			}

			p1 := part1(input)
			if p1 != test.Part1 {
				t.Errorf("part 1: got=%v, want=%v", p1, test.Part1)
			}

			p2 := part2(input)
			if p2 != test.Part2 {
				t.Errorf("part 2: got=%v, want=%v", p2, test.Part2)
			}
		})
	}
}
