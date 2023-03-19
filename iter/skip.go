package iter

type skipIter[T any] struct {
	skip       int
	underlying Iterator[T]
}

// Skip 返回一个跳过n个元素的迭代器
func Skip[T any](iterator Iterator[T], n int) Iterator[T] {
	return skipIter[T]{
		skip:       n,
		underlying: iterator,
	}
}

// Drop 是 Skip 的别名。
func Drop[T any](iterator Iterator[T], n int) Iterator[T] {
	return Skip(iterator, n)
}

func (i skipIter[T]) Exhausted() bool {
	return i.underlying.Exhausted()
}

func (i skipIter[T]) Next() Iterator[T] {
	if i.underlying.Exhausted() {
		return i
	}

	for i.skip > 0 {
		i.underlying = i.underlying.Next()
		i.skip--
	}

	i.underlying = i.underlying.Next()
	return i
}

func (i skipIter[T]) Value() T {
	return i.underlying.Value()
}
