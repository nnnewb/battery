package slices

import "github.com/nnnewb/battery/iter"

func (s Slice[T]) All(predicate func(T) bool) bool {
	return iter.All(iter.Lift(s), predicate)
}
