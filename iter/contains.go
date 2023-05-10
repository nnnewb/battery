package iter

// ContainsFunc 是否包含特定元素。元素必须是可比较的。
func ContainsFunc[T any](it Iterator[T], v T, equal func(T, T) bool) bool {
	return AnyFunc(it, func(t T) bool { return equal(t, v) })
}

// Contains 遍历迭代器，检查是否存在元素 v 。存在则返回 true，反之返回 false。
func Contains[T comparable](it Iterator[T], v T) bool {
	return AnyEqual(it, v)
}
