package main

import (
	"bufio"
	"io"
	"regexp"
	"strconv"

	"github.com/garrettladley/advent-of-code-24/go/aoc"
)

var mulPattern = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

func Read(r io.Reader) ([]aoc.Pair[uint16], error) {
	scanner := bufio.NewScanner(r)
	var pairs []aoc.Pair[uint16]
	for scanner.Scan() {
		matches := mulPattern.FindAllStringSubmatch(scanner.Text(), -1)
		for _, match := range matches {
			first, err := strconv.ParseUint(match[1], 10, 16)
			if err != nil {
				return nil, err
			}
			second, err := strconv.ParseUint(match[2], 10, 16)
			if err != nil {
				return nil, err
			}
			pairs = append(pairs, aoc.Pair[uint16]{A: uint16(first), B: uint16(second)})
		}
	}
	return pairs, scanner.Err()
}
