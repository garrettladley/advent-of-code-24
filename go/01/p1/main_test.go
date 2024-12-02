package main

import (
	"bytes"
	"context"
	"testing"
)

func TestRun(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		in          string
		expectedOut string
	}{
		{
			name:        "two rows",
			in:          "80784   47731\n81682   36089\n",
			expectedOut: "78646",
		},
		{
			name:        "prompt",
			in:          "3   4\n4   3\n2   5\n1   3\n3   9\n3   3",
			expectedOut: "11",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			actual, err := Run(ctx, bytes.NewBuffer([]byte(tt.in)))
			if err != nil {
				t.Errorf("non-nil error encountered while reading: %v", err)
			}
			if actual != tt.expectedOut {
				t.Errorf("expected output to be %s, but was %s", tt.expectedOut, actual)
			}
		})
	}
}
