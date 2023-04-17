package slices

import "github.com/nnnewb/battery/iter"

// Any 测试切片元素是否有任意一个元素满足条件，有则返回 true 。
// 切片为空或没有满足条件的元素则返回 false 。
func (s Slice[T]) Any(predicate func(T) bool) bool {
	return iter.Any(iter.Lift(s), predicate)
}
