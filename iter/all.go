package iter

// AllFunc 遍历迭代器，若所有元素都能使 predicate 返回 true，则函数返回 true。
// 反之则返回 false 。若迭代器已经穷尽，也返回 true。
func AllFunc[T any](it Iterator[T], predicate func(T) bool) bool {
	for it.Next() {
		if !predicate(it.Value()) {
			return false
		}
	}
	return true
}

// All 遍历迭代器，若所有元素都是 true 则返回 true。反之返回 false。
func All(it Iterator[bool]) bool {
	for it.Next() {
		if !it.Value() {
			return false
		}
	}
	return true
}

// AllEqual 遍历迭代器，若所有元素都和 v 相等，则返回 true。
// 反之则返回 false 。若迭代器已经穷尽，也返回 true。
func AllEqual[T comparable](it Iterator[T], v T) bool {
	for it.Next() {
		if v != it.Value() {
			return false
		}
	}
	return true
}
