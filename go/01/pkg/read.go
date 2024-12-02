package pkg

import (
	"fmt"
	"io"
	"strconv"
	"text/scanner"
)

func Read(r io.Reader) (p Pair[[]int], err error) {
	var (
		s   scanner.Scanner
		isB bool
		a   []int
		b   []int
	)

	s.Init(r)
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		parsed, err := strconv.Atoi(s.TokenText())
		if err != nil {
			return p, err
		}
		if isB {
			b = append(b, int(parsed))
		} else {
			a = append(a, int(parsed))
		}
		isB = !isB
	}

	if len(a) != len(b) {
		return p, fmt.Errorf("expected a and b to be the same length, but len(a)=%d, len(b)=%d", len(a), len(b))
	}

	p.A = a
	p.B = b
	return p, nil
}
