package slices

func (s Slice[T]) Cap() int {
	return cap(s)
}
