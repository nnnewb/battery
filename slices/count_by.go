package slices

import "github.com/nnnewb/battery/iter"

func CountBy[T any, K comparable](s Slice[T], keyFunc func(T) K) map[K]int {
	return iter.CountBy(iter.Lift(s), keyFunc)
}
