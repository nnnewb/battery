package slices

import "github.com/nnnewb/battery/iter"

func (s ComparableElemSlice[T]) Unique() Slice[T] {
	return iter.Collect(iter.Unique(iter.Lift(s)))
}
