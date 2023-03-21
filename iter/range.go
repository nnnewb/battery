package iter

import "github.com/nnnewb/battery/constraints"

// Range 返回一个范围迭代器，从指定起始值迭代到结束值。
func Range[T constraints.Numeric](from T, to T, step T) Iterator[T] {
	return Generator(func() GenFunc[T] {
		cur := from
		return func() (T, bool) {
			if cur+step > to {
				return 0, false
			}
			cur += step
			return cur, true
		}
	}())
}
