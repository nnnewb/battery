package slices

import "github.com/nnnewb/battery/iter"

// Map 逐个元素执行操作，返回新的切片
func Map[T, R any](s Slice[T], f func(T) R) Slice[R] {
	return iter.Collect(iter.Map[T, R](iter.Lift(s), f))
}
