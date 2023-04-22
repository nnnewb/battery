package slices

import "github.com/nnnewb/battery/iter"

// All 测试是否所有元素都满足条件，若 s 长度为 或存在元素不符合条件
// 则返回 false 。
func (s Slice[T]) All(predicate func(T) bool) bool {
	if len(s) == 0 {
		return false
	}
	return iter.All(iter.Lift(s), predicate)
}

// All 测试是否所有元素都满足条件，若 s 长度为 或存在元素不符合条件
// 则返回 false 。
func (s ComparableElemSlice[T]) All(predicate func(T) bool) bool {
	return Slice[T](s).All(predicate)
}

// All 测试是否所有元素都满足条件，若 s 长度为 或存在元素不符合条件
// 则返回 false 。
func (s OrderedElemSlice[T]) All(predicate func(T) bool) bool {
	return Slice[T](s).All(predicate)
}
