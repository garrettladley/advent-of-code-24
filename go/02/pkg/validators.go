package pkg

import (
	"math"

	"golang.org/x/exp/constraints"
)

func IsValidRow[T constraints.Integer | constraints.Float](row []T) bool {
	if len(row) < 2 {
		return true
	}

	isIncreasing := row[0] < row[1]

	for i := 0; i < len(row)-1; i++ {
		if !IsValidDelta(row[i], row[i+1]) || ((row[i] < row[i+1]) != isIncreasing) {
			return false
		}
	}
	return true
}

func IsValidDelta[T constraints.Integer | constraints.Float](a T, b T) bool {
	delta := math.Abs(float64(a - b))
	return delta >= 1 && delta <= 3
}
