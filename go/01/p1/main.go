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
	p, err := pkg.Read(r)
	if err != nil {
		return "", fmt.Errorf("error reading: %w", err)
	}
	pkg.Sort(p)
	return strconv.Itoa(PairwiseDelta(p)), nil
}

func PairwiseDelta(p pkg.Pair[[]int]) int {
	var (
		a   = p.A
		b   = p.B
		sum int
	)
	for idx := range a {
		sum += int(math.Abs(float64(a[idx] - b[idx])))
	}
	return sum
}
