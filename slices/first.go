package slices

import "github.com/nnnewb/battery/iter"

func (s Slice[T]) First(predicate func(T) bool) (T, bool) {
	return iter.First(iter.Lift(s), predicate)
}
