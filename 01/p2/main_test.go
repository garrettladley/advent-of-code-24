package main

import (
	"bytes"
	"testing"
)

func TestRun(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		in          string
		expectedOut int
	}{
		{
			name:        "prompt",
			in:          "3   4\n4   3\n2   5\n1   3\n3   9\n3   3",
			expectedOut: 31,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := Run(bytes.NewBuffer([]byte(tt.in)))
			if err != nil {
				t.Errorf("non-nil error encountered while reading: %v", err)
			}
			if actual != tt.expectedOut {
				t.Errorf("expected output to be %d, but was %d", tt.expectedOut, actual)
			}
		})
	}
}
