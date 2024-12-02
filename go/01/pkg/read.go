package pkg

import (
	"bytes"
	"fmt"
	"io"
)

func Read(r io.Reader) (p Pair[[]int], err error) {
	return readN(r, Rows)
}

const rowLen uint = 14

func readN(r io.Reader, rows uint) (p Pair[[]int], err error) {
	buf := new(bytes.Buffer)
	n, err := buf.ReadFrom(r)
	if err != nil {
		return p, fmt.Errorf("failed to read input: %w", err)
	}
	if n == 0 {
		return p, fmt.Errorf("no data read")
	}

	data := buf.Bytes()
	expectedLen := rowLen * rows // 13 chars + newline per row
	if len(data) != int(expectedLen) {
		return p, fmt.Errorf("expected %d bytes, got %d", expectedLen, len(data))
	}

	p.A = make([]int, rows)
	p.B = make([]int, rows)

	for i, pos := uint(0), uint(0); i < rows; i++ {
		p.A[i] = (int(data[pos]-'0')*10000 +
			int(data[pos+1]-'0')*1000 +
			int(data[pos+2]-'0')*100 +
			int(data[pos+3]-'0')*10 +
			int(data[pos+4]-'0'))

		p.B[i] = (int(data[pos+8]-'0')*10000 +
			int(data[pos+9]-'0')*1000 +
			int(data[pos+10]-'0')*100 +
			int(data[pos+11]-'0')*10 +
			int(data[pos+12]-'0'))

		pos += rowLen
	}

	return p, nil
}
