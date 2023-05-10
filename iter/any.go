package iter

// AnyFunc 遍历迭代器，若存在元素使 predicate 函数返回 true，则返回 true。反之返回 false。
func AnyFunc[T any](it Iterator[T], predicate func(T) bool) bool {
	for it.Next() {
		if predicate(it.Value()) {
			return true
		}
	}
	return false
}

// AnyEqual 遍历迭代器，若存在 v 则返回 true。反之返回 false。
func AnyEqual[T comparable](it Iterator[T], v T) bool {
	for it.Next() {
		if it.Value() == v {
			return true
		}
	}
	return false
}

// Any 遍历迭代器，若存在元素为 true 则返回 true。反之返回 false。
// 若没有元素，也返回 false。
func Any(it Iterator[bool]) bool {
	for it.Next() {
		if it.Value() {
			return true
		}
	}
	return false
}
