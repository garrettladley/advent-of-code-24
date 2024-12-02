package pkg

import (
	"bytes"
	"fmt"
	"io"
)

func Read(r io.Reader) (p Pair[[]int], err error) {
	buf := new(bytes.Buffer)
	n, err := buf.ReadFrom(r)
	if err != nil {
		return p, fmt.Errorf("failed to read input: %w", err)
	}
	if n == 0 {
		return p, fmt.Errorf("no data read")
	}

	data := buf.Bytes()
	rowLen := NumberWidth + SpaceWidth + NumberWidth + 1 // number A + space + number B + newline
	expectedLen := (rowLen) * Rows
	if len(data) != int(expectedLen) {
		return p, fmt.Errorf("expected %d bytes, got %d", expectedLen, len(data))
	}

	p.A = make([]int, Rows)
	p.B = make([]int, Rows)

	for i, pos := uint(0), uint(0); i < Rows; i++ {
		// unwrap the loop
		p.A[i] = (int(data[pos]-'0')*10000 +
			int(data[pos+1]-'0')*1000 +
			int(data[pos+2]-'0')*100 +
			int(data[pos+3]-'0')*10 +
			int(data[pos+4]-'0'))

		// unwrap the loop
		p.B[i] = (int(data[pos+8]-'0')*10000 +
			int(data[pos+9]-'0')*1000 +
			int(data[pos+10]-'0')*100 +
			int(data[pos+11]-'0')*10 +
			int(data[pos+12]-'0'))

		pos += rowLen
	}

	return p, nil
}

func ReadN(r io.Reader, rows uint, numberWidth uint, spaceWidth uint) (p Pair[[]int], err error) {
	// i don't love but it helps us unwrap the loop in the "production" case
	if rows == Rows && numberWidth == NumberWidth && spaceWidth == SpaceWidth {
		return Read(r)
	}

	buf := new(bytes.Buffer)
	n, err := buf.ReadFrom(r)
	if err != nil {
		return p, fmt.Errorf("failed to read input: %w", err)
	}
	if n == 0 {
		return p, fmt.Errorf("no data read")
	}

	data := buf.Bytes()
	rowLen := numberWidth + spaceWidth + numberWidth + 1 // number A + space + number B + newline
	expectedLen := rowLen * rows
	if len(data) != int(expectedLen) {
		return p, fmt.Errorf("expected %d bytes, got %d", expectedLen, len(data))
	}

	p.A = make([]int, rows)
	p.B = make([]int, rows)

	for idx, pos := uint(0), uint(0); idx < rows; idx++ {
		var a int
		for p := uint(0); p < numberWidth; p++ {
			a = a*10 + int(data[pos+p]-'0')
		}
		p.A[idx] = a

		var b int
		for p := uint(0); p < numberWidth; p++ {
			b = b*10 + int(data[pos+numberWidth+spaceWidth+p]-'0')
		}
		p.B[idx] = b

		pos += rowLen
	}

	return p, nil
}
