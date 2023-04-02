package slices

import "github.com/nnnewb/battery/iter"

func (s Slice[T]) Last(predicate func(T) bool) (T, bool) {
	return iter.Last(iter.Lift(s), predicate)
}
