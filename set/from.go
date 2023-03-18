package set

// NewSetFromSlice 从切片创建一个集合。
//
//	NewSetFromSlice([]int{1,2,3})
func NewSetFromSlice[T comparable](s []T) Set[T] {
	ret := make(Set[T])
	for _, v := range s {
		ret.Add(v)
	}
	return ret
}

// NewSetFromMapKey 从 map 的键创建一个集合。
//
//	NewSetFromMapKey(map[int]int{1:1, 2:2})
func NewSetFromMapKey[K comparable, V any](m map[K]V) Set[K] {
	ret := make(Set[K])
	for k := range m {
		ret.Add(k)
	}
	return ret
}
