package iter

type filterIter[T any] struct {
	f          func(T) bool
	underlying Iterator[T]
}

func (i filterIter[T]) Next() Iterator[T] {
	it := i.underlying.Next()
	for ; !it.Exhausted(); it = it.Next() {
		if i.f(it.Value()) {
			break
		}
	}

	i.underlying = it
	return i
}

func (i filterIter[T]) Value() T {
	return i.underlying.Value()
}

func (i filterIter[T]) Exhausted() bool {
	return i.underlying.Exhausted()
}

// Filter 创建一个迭代器，过滤出符合条件的元素。
func Filter[T any](iterator Iterator[T], f func(T) bool) Iterator[T] {
	return filterIter[T]{
		f:          f,
		underlying: iterator,
	}
}

var _ Iterator[struct{}] = filterIter[struct{}]{}
