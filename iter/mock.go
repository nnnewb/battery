package iter

type MockIterator[T any] struct {
	sequence      []T
	idx           int
	nextCallCount int
}

func (i MockIterator[T]) Next() Iterator[T] {
	i.nextCallCount++
	if i.idx < len(i.sequence) {
		i.idx++
	}

	return i
}

func (i MockIterator[T]) Value() T {
	return i.sequence[i.idx]
}

func (i MockIterator[T]) Exhausted() bool {
	return i.idx >= len(i.sequence)
}

func (i MockIterator[T]) NextCallCount() int {
	return i.nextCallCount
}

func Mock[T any](seq ...T) MockIterator[T] {
	return MockIterator[T]{
		sequence:      seq,
		idx:           -1,
		nextCallCount: 0,
	}
}

var _ Iterator[struct{}] = MockIterator[struct{}]{}
