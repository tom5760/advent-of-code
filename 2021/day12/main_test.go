package main

import (
	"strings"
	"testing"
)

const (
	exampleInput1 = `start-A
start-b
A-c
A-b
b-d
A-end
b-end
`

	exampleInput2 = `dc-end
HN-start
start-kj
dc-start
dc-HN
LN-dc
HN-end
kj-sa
kj-HN
kj-dc
`

	exampleInput3 = `fs-end
he-DX
fs-he
start-DX
pj-DX
end-zg
zg-sl
zg-pj
pj-he
RW-he
fs-DX
pj-RW
zg-RW
start-pj
he-WI
zg-he
pj-fs
start-RW
`
)

func TestPart1(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		expected uint64
		input    string
	}{
		{
			name:     "example1",
			expected: 10,
			input:    exampleInput1,
		},
		{
			name:     "example2",
			expected: 19,
			input:    exampleInput2,
		},
		{
			name:     "example3",
			expected: 226,
			input:    exampleInput3,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			graph, err := ParseInput(strings.NewReader(test.input))
			if err != nil {
				t.Fatalf("unexpected input parse error: %v", err)
			}

			actual := Part1(graph)

			if actual != test.expected {
				t.Errorf("got %v; want %v", actual, test.expected)
			}
		})
	}
}
