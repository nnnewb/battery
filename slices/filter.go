package slices

import "github.com/nnnewb/battery/iter"

// Filter 过滤切片元素，返回新切片
func (s Slice[T]) Filter(predicate func(T) bool) Slice[T] {
	if len(s) == 0 {
		return make(Slice[T], 0)
	}

	return iter.Collect(iter.Filter(iter.Lift(s), predicate))
}
