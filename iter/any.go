package iter

func Any[T any](it Iterator[T], predicate func(T) bool) bool {
	for it.Next() {
		if predicate(it.Value()) {
			return true
		}
	}
	return false
}
