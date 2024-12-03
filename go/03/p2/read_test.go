package main

import (
	"bytes"
	"testing"

	"github.com/garrettladley/advent-of-code-24/go/aoc"
)

func TestRead(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		in          string
		expectedOut []aoc.Pair[uint16]
	}{
		{
			name: "prompt",
			in:   "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))",
			expectedOut: []aoc.Pair[uint16]{
				{A: 2, B: 4},
				{A: 8, B: 5},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualOut, err := Read(bytes.NewBuffer([]byte(tt.in)))
			if err != nil {
				t.Errorf("non-nil error encountered while reading: %v", err)
			}
			if len(actualOut) != len(tt.expectedOut) {
				t.Fatalf("expected len(actual) to be %d, but was %d", len(tt.expectedOut), len(actualOut))
			}

			for idx := range actualOut {
				if actualOut[idx] != tt.expectedOut[idx] {
					t.Errorf("expected actual[%d] to be %v, but was %v", idx, tt.expectedOut[idx], actualOut[idx])
				}
			}
		})
	}
}
