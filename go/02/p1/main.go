package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/garrettladley/advent-of-code-24/go/02/pkg"
)

func main() {
	inputFlag := flag.String("input", "../input.txt", "the location of the input file")
	flag.Parse()

	f, err := os.Open(*inputFlag)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error opening file: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	ctx := context.Background()
	r, err := Run(ctx, f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error running: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(r)
}

func Run(ctx context.Context, r io.Reader) (string, error) {
	return run(ctx, r)
}

func run(_ context.Context, r io.Reader) (string, error) {
	g, err := pkg.Read(r)
	if err != nil {
		return "", fmt.Errorf("failed to read grid: %w", err)
	}
	return strconv.Itoa(int(SafeCount(g))), nil
}

func SafeCount(g pkg.Grid) uint {
	var unsafeCount uint
	for _, row := range g {
		if len(row) == 1 {
			continue
		}

		if !pkg.IsValidDelta(row[0], row[1]) {
			unsafeCount++
			continue
		}

		if !pkg.IsValidRow(row) {
			unsafeCount++
		}
	}
	return uint(len(g)) - unsafeCount
}
