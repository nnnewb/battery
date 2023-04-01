package slices

import "github.com/nnnewb/battery/iter"

// Take 取前n个元素
func (s Slice[T]) Take(n int) Slice[T] {
	return iter.Collect(iter.Take(iter.Lift(s), n))
}
