package iter

import "github.com/nnnewb/battery/internal/constraints"

// Min 取最小的元素
func Min[T any, K constraints.Ordered](it Iterator[T], keyFunc func(T) K) (T, bool) {
	var zero T
	i := Sort(it, keyFunc)
	for i.Next() {
		return i.Value(), true
	}
	return zero, false
}

// Max 取最大的元素
func Max[T any, K constraints.Ordered](it Iterator[T], keyFunc func(T) K) (T, bool) {
	var last T
	var found bool
	i := Sort(it, keyFunc)
	for i.Next() {
		last = i.Value()
		found = true
	}
	return last, found
}
