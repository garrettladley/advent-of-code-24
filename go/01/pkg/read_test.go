package pkg

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
		rows        uint
		numberWidth uint
		spaceWidth  uint
		expectedOut aoc.Pair[[]int]
	}{
		{
			name:        "two rows",
			in:          "80784   47731\n81682   36089\n",
			rows:        2,
			numberWidth: NumberWidth,
			spaceWidth:  SpaceWidth,
			expectedOut: aoc.Pair[[]int]{
				A: []int{80784, 81682},
				B: []int{47731, 36089},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualOut, err := ReadN(bytes.NewBuffer([]byte(tt.in)), tt.rows, tt.numberWidth, tt.spaceWidth)
			if err != nil {
				t.Errorf("non-nil error encountered while reading: %v", err)
			}
			if len(actualOut.A) != len(tt.expectedOut.A) {
				t.Fatalf("expected len(actual.a) to be %d, but was %d", len(actualOut.A), len(tt.expectedOut.A))
			}
			if len(actualOut.B) != len(tt.expectedOut.B) {
				t.Fatalf("expected len(actual.b) to be %d, but was %d", len(actualOut.B), len(tt.expectedOut.B))
			}
			if len(actualOut.A) != len(actualOut.B) {
				t.Fatalf("expected len(actual.a) == len(actual.b), but len(actual.a)=%d and len(actual.b)=%d", len(actualOut.A), len(actualOut.B))
			}

			for idx := range actualOut.A {
				if actualOut.A[idx] != tt.expectedOut.A[idx] {
					t.Errorf("expected value in index %d of actual.a to be %d, but was %d", idx, actualOut.A[idx], tt.expectedOut.A[idx])
				}
				if actualOut.B[idx] != tt.expectedOut.B[idx] {
					t.Errorf("expected value in index %d of actual.b to be %d, but was %d", idx, actualOut.B[idx], tt.expectedOut.B[idx])
				}
			}
		})
	}
}
