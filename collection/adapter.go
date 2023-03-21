package collection

import (
	"github.com/nnnewb/battery/iter"
)

type Slice[T any] []T

func (s Slice[T]) Iterator() iter.Iterator[T] {
	return iter.Lift(s)
}

type Chan[T any] <-chan T

func (c Chan[T]) Iterator() iter.Iterator[T] {
	return iter.FromChan(c)
}

type Generator[T any] func() (T, bool)

func (g Generator[T]) Iterator() iter.Iterator[T] {
	return iter.Generator(iter.GenFunc[T](g))
}

type it[T any] struct {
	it iter.Iterator[T]
}

func (i it[T]) Iterator() iter.Iterator[T] {
	return i.it
}

func Iter[T any](iterator iter.Iterator[T]) Iterable[T] {
	return it[T]{it: iterator}
}
