package slices

import "github.com/nnnewb/battery/iter"

func (s Slice[T]) Any(predicate func(T) bool) bool {
	return iter.Any(iter.Lift(s), predicate)
}
