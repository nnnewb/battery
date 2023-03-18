package iter

type chainIter[T any] struct {
	it1 Iterator[T]
	it2 Iterator[T]
}

func (i chainIter[T]) Next() Iterator[T] {
	if !i.it1.Exhausted() {
		i.it1 = i.it1.Next()
		return i
	}

	if !i.it2.Exhausted() {
		i.it2 = i.it2.Next()
		return i
	}

	return i
}

func (i chainIter[T]) Value() T {
	if !i.it1.Exhausted() {
		return i.it1.Value()
	}

	return i.it2.Value()
}

func (i chainIter[T]) Exhausted() bool {
	return i.it1.Exhausted() && i.it2.Exhausted()
}
