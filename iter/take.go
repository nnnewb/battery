package iter

type takeIter[T any] struct {
	cur        int
	limit      int
	underlying Iterator[T]
}

func (i takeIter[T]) Next() Iterator[T] {
	if i.cur < i.limit {
		i.cur++
		i.underlying = i.underlying.Next()
	}
	return i
}

func (i takeIter[T]) Value() T {
	return i.underlying.Value()
}

func (i takeIter[T]) Exhausted() bool {
	return i.cur >= i.limit || i.underlying.Exhausted()
}

// Take 返回一个迭代器，取指定数量的元素
func Take[T any](iterator Iterator[T], n int) Iterator[T] {
	return takeIter[T]{
		cur:        -1,
		limit:      n,
		underlying: iterator,
	}
}
