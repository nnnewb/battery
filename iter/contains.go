package iter

// Contains 是否包含特定元素。元素必须是可比较的。
func Contains[T any](it Iterator[T], v T, equal func(T, T) bool) bool {
	return Any(it, func(t T) bool { return equal(t, v) })
}
