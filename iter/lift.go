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

// FromSlice 创建一个切片迭代器
func FromSlice[T any](s []T) Iterator[T] {
	return sliceIter[T]{
		underlying: s,
		idx:        -1,
	}
}

// Lift 兼容 go-functional 的写法，等同于 FromSlice
func Lift[T any](s []T) Iterator[T] {
	return FromSlice(s)
}
