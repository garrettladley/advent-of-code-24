package pkg

import "slices"

func Sort(p Pair[[]int]) {
	slices.SortFunc(p.A, func(a int, b int) int {
		return int(a - b)
	})

	slices.SortFunc(p.B, func(a int, b int) int {
		return int(a - b)
	})
}
