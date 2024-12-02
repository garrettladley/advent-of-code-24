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
			name:        "prompt",
			in:          "7 6 4 2 1\n1 2 7 8 9\n9 7 6 2 1\n1 3 2 4 5\n8 6 4 4 1\n1 3 6 7 9\n",
			expectedOut: "4",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			actual, err := run(ctx, bytes.NewBuffer([]byte(tt.in)))
			if err != nil {
				t.Errorf("non-nil error encountered while reading: %v", err)
			}
			if actual != tt.expectedOut {
				t.Errorf("expected output to be %s, but was %s", tt.expectedOut, actual)
			}
		})
	}
}
