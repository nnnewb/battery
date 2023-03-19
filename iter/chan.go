package iter

type chanIter[T any] struct {
	c         <-chan T
	cur       T
	exhausted bool
}

func (i chanIter[T]) Exhausted() bool {
	return i.exhausted
}

func (i chanIter[T]) Next() Iterator[T] {
	var ok bool
	i.cur, ok = <-i.c
	i.exhausted = !ok
	return i
}

func (i chanIter[T]) Value() T {
	return i.cur
}

var _ Iterator[struct{}] = chanIter[struct{}]{}

// FromChan 从 go 内置的 chan 创建 Iterator
func FromChan[T any](c <-chan T) Iterator[T] {
	return chanIter[T]{c: c}
}
