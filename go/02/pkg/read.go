package pkg

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func Read(r io.Reader) (g Grid, err error) {
	s := bufio.NewScanner(r)
	for s.Scan() {
		line := s.Text()
		var row []int8
		for _, raw := range strings.Split(line, " ") {
			num, err := strconv.Atoi(raw)
			if err != nil {
				return nil, fmt.Errorf("failed to parse number: %w", err)
			}
			row = append(row, int8(num))
		}
		g = append(g, row)
	}
	return g, nil
}
