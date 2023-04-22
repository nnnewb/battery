package slices

import "github.com/nnnewb/battery/iter"

// Unique returns a new slice containing only the unique elements of the original slice, in the order they first appear.
func (s ComparableElemSlice[T]) Unique() Slice[T] {
	return iter.Collect(iter.Unique(iter.Lift(s)))
}
