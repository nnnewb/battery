package slices

import "github.com/nnnewb/battery/iter"

func GroupBy[T any, K comparable](s Slice[T], keyFunc func(T) K) map[K][]T {
	return iter.GroupBy(iter.Lift(s), keyFunc)
}
