package slices

import "github.com/nnnewb/battery/iter"

func (s Slice[T]) Contains(v T, equal func(T, T) bool) bool {
	return iter.Contains(iter.Lift(s), v, equal)
}
