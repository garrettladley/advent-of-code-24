package pkg

type Pair[T any] struct {
	A T
	B T
}

func NewPair[T any](a T, b T) Pair[T] {
	return Pair[T]{A: a, B: b}
}
