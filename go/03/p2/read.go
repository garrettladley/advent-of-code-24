package main

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"

	"github.com/garrettladley/advent-of-code-24/go/aoc"
)

var mulRegex = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

func Read(r io.Reader) ([]aoc.Pair[uint16], error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanRunes)

	var results []aoc.Pair[uint16]
	var buffer strings.Builder
	enabled := true

	for scanner.Scan() {
		char := scanner.Text()
		buffer.WriteString(char)

		// check for do/don't instructions
		if strings.HasSuffix(buffer.String(), "do()") {
			enabled = true
			buffer.Reset()
			continue
		}
		if strings.HasSuffix(buffer.String(), "don't()") {
			enabled = false
			buffer.Reset()
			continue
		}

		// look for mul instructions
		if matches := mulRegex.FindStringSubmatch(buffer.String()); matches != nil {
			if enabled {
				a, err := strconv.ParseUint(matches[1], 10, 16)
				if err != nil {
					return nil, fmt.Errorf("invalid first number: %v", err)
				}

				b, err := strconv.ParseUint(matches[2], 10, 16)
				if err != nil {
					return nil, fmt.Errorf("invalid second number: %v", err)
				}

				results = append(results, aoc.Pair[uint16]{
					A: uint16(a),
					B: uint16(b),
				})
			}
			buffer.Reset()
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("scanner error: %v", err)
	}

	return results, nil
}
