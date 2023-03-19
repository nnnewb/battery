package iter

type mapIter[T, R any] struct {
	f          func(T) R
	underlying Iterator[T]
	cur        R
}

func (i mapIter[T, R]) Next() Iterator[R] {
	i.underlying = i.underlying.Next()
	if !i.underlying.Exhausted() {
		i.cur = i.f(i.underlying.Value())
	}
	return i
}

func (i mapIter[T, R]) Value() R {
	return i.cur
}

func (i mapIter[T, R]) Exhausted() bool {
	return i.underlying.Exhausted()
}

func Map[T, R any](iterator Iterator[T], f func(T) R) Iterator[R] {
	it := mapIter[T, R]{
		f:          f,
		underlying: iterator,
	}
	return it
}
