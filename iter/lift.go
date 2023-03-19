package iter

// Lift 兼容 go-functional 的写法，等同于 iter.FromSlice
func Lift[T any](s []T) Iterator[T] {
	return FromSlice(s)
}
