package iter

// Lift 兼容 go-functional 的写法，等同于 FromSlice
func Lift[T any](s []T) Iterator[T] {
	return FromSlice(s)
}
