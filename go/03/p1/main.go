package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/garrettladley/advent-of-code-24/go/aoc"
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
	p, err := Read(r)
	if err != nil {
		return "", fmt.Errorf("error reading: %w", err)
	}
	return strconv.Itoa(int(apply(p))), nil
}

func apply(p []aoc.Pair[uint16]) (total uint64) {
	for _, pair := range p {
		total += uint64(pair.A) * uint64(pair.B)
	}
	return
}
