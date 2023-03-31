package iter

// Contains 是否包含特定元素。元素必须是可比较的。
func Contains[T comparable](it Iterator[T], v T) bool {
	return Any(it, func(t T) bool { return t == v })
}
