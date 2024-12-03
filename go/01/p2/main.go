package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"

	"github.com/garrettladley/advent-of-code-24/go/01/pkg"
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
	return run(ctx, r, pkg.Rows, pkg.NumberWidth, pkg.SpaceWidth)
}

func run(_ context.Context, r io.Reader, rows uint, numberWidth uint, numSpaces uint) (string, error) {
	p, err := pkg.ReadN(r, rows, numberWidth, numSpaces)
	if err != nil {
		return "", fmt.Errorf("error reading: %w", err)
	}
	return strconv.Itoa(PairwiseSimilarity(p)), nil
}

func Counter[T comparable](s []T) map[T]int {
	m := make(map[T]int)
	for _, v := range s {
		m[v]++
	}
	return m
}

func PairwiseSimilarity(p aoc.Pair[[]int]) int {
	var (
		a        = p.A
		bCounter = Counter(p.B)
		sum      int
	)
	for idx := range a {
		var score int
		if count, ok := bCounter[a[idx]]; ok {
			score = count
		}
		sum += int(math.Abs(float64(a[idx] * score)))
	}
	return sum
}
