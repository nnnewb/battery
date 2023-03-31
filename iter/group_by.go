package iter

// GroupBy 将元素按指定键分组
func GroupBy[T any, K comparable](it Iterator[T], keyFunc func(T) K) map[K][]T {
	ret := make(map[K][]T)
	for it.Next() {
		k := keyFunc(it.Value())
		if _, ok := ret[k]; !ok {
			ret[k] = make([]T, 0, 1)
		}
		ret[k] = append(ret[k], it.Value())
	}
	return ret
}
