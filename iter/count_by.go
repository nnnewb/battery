package iter

// CountBy 按指定键分组统计元素数量.
func CountBy[T any, K comparable](it Iterator[T], keyFunc func(T) K) map[K]int {
	ret := make(map[K]int)
	for it.Next() {
		k := keyFunc(it.Value())
		if _, ok := ret[k]; !ok {
			ret[k] = 1
			continue
		}
		ret[k]++
	}
	return ret
}
