package iter

// CountByFunc 按指定键分组统计元素数量.
func CountByFunc[T any, K comparable](it Iterator[T], keyFunc func(T) K) map[K]int {
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

// CountBy 遍历迭代器，分组统计数量。返回一个 map，键为迭代器元素，值为元素出现的次数。
func CountBy[T comparable](it Iterator[T]) map[T]int {
	ret := make(map[T]int)
	for it.Next() {
		v := it.Value()
		if _, ok := ret[v]; !ok {
			ret[v] = 1
		} else {
			ret[v]++
		}
	}
	return nil
}
