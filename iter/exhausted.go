package iter

type exhaustedIter[T any] struct{}

func (i exhaustedIter[T]) Next() Iterator[T] {
	return i
}

func (i exhaustedIter[T]) Value() T {
	var t T
	return t
}

func (i exhaustedIter[T]) Exhausted() bool {
	return true
}

// Exhausted 返回一个已经穷尽的迭代器
func Exhausted[T any]() Iterator[T] {
	return exhaustedIter[T]{}
}

var _ Iterator[struct{}] = exhaustedIter[struct{}]{}
