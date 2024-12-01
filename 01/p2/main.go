package main

import (
	"flag"
	"fmt"
	"io"
	"math"
	"os"

	"github.com/garrettladley/advent-of-code-24/01/pkg"
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

	r, err := Run(f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error running: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(r)
}

func Run(r io.Reader) (int, error) {
	p, err := pkg.Prepare(r)
	if err != nil {
		return 0, fmt.Errorf("error preparing: %w", err)
	}
	return PairwiseSimilarity(p), nil
}

func Counter[T comparable](s []T) map[T]int {
	m := make(map[T]int)
	for _, v := range s {
		m[v]++
	}
	return m
}

func PairwiseSimilarity(p pkg.Pair[[]int]) int {
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
