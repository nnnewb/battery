package slices

import "github.com/nnnewb/battery/iter"

func Reduce[T, R any](s Slice[T], initial R, f func(R, T) R) R {
	return iter.Reduce[T, R](iter.Lift(s), initial, f)
}
