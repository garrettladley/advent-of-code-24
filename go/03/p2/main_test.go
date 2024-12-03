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
			in:          "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))",
			expectedOut: "48",
		},
		{
			name:        "adjacent instructions",
			in:          "do()mul(1,2)don't()mul(3,4)",
			expectedOut: "2",
		},
		{
			name:        "multiple controls same position",
			in:          "do()don't()mul(1,2)",
			expectedOut: "0",
		},
		{
			name:        "nested instructions",
			in:          "do(mul(1,2))mul(3,4)",
			expectedOut: "14",
		},
		{
			name:        "large numbers",
			in:          "mul(999,999)do()mul(1000,1000)",
			expectedOut: "998001",
		},
		{
			name:        "empty input",
			in:          "",
			expectedOut: "0",
		},
		{
			name:        "malformed instructions",
			in:          "mul(1,)mul(,2)mul()mul(1.5,2)mul(-1,2)",
			expectedOut: "0",
		},
		{
			name:        "multiple do/dont toggles",
			in:          "mul(1,1)don't()do()don't()do()mul(2,2)",
			expectedOut: "5",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
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
