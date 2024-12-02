package pkg

import (
	"bufio"
	"io"
)

func Read(r io.Reader) (p Pair[[]int], err error) {
	return readN(r, Rows)
}

func readN(r io.Reader, rows uint) (p Pair[[]int], err error) {
	p.A = make([]int, rows)
	p.B = make([]int, rows)

	var (
		scanner = bufio.NewScanner(r)
		idx     int
	)

	for scanner.Scan() {
		line := scanner.Bytes()

		p.A[idx] = (int(line[0]-'0')*10000 +
			int(line[1]-'0')*1000 +
			int(line[2]-'0')*100 +
			int(line[3]-'0')*10 +
			int(line[4]-'0'))

		p.B[idx] = (int(line[8]-'0')*10000 +
			int(line[9]-'0')*1000 +
			int(line[10]-'0')*100 +
			int(line[11]-'0')*10 +
			int(line[12]-'0'))

		idx++
	}

	if err := scanner.Err(); err != nil {
		return p, err
	}

	return p, nil
}
