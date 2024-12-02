package pkg

import "io"

func Prepare(r io.Reader) (Pair[[]int], error) {
	p, err := Read(r)
	if err != nil {
		return Pair[[]int]{}, err
	}
	Sort(p)
	return p, nil
}
