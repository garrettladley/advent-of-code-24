package pkg

import (
	"slices"

	"golang.org/x/exp/constraints"
)

func Sort[T constraints.Integer | constraints.Float](p Pair[[]T]) {
	slices.SortFunc(p.A, func(a T, b T) int {
		return int(a - b)
	})

	slices.SortFunc(p.B, func(a T, b T) int {
		return int(a - b)
	})
}
