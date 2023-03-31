package iter

// All 检测元素是否都符合条件。
func All[T any](it Iterator[T], predicate func(T) bool) bool {
	for it.Next() {
		if !predicate(it.Value()) {
			return false
		}
	}
	return true
}
