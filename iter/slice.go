package iter

type sliceIter[T any] struct {
	idx        int
	underlying []T
}

func (s sliceIter[T]) Exhausted() bool {
	return s.idx >= len(s.underlying)
}

func (s sliceIter[T]) Next() Iterator[T] {
	s.idx++
	return s
}

func (s sliceIter[T]) Value() T {
	return s.underlying[s.idx]
}

var _ Iterator[struct{}] = sliceIter[struct{}]{}
