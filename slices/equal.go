package slices

// Equal 比较两个切片是否相等。传入比较函数 eq，此函数返回 true 表示元素相同。
// 元素数量、顺序都一致则认为切片相等。注意 nil 切片和空切片会视作相等。
func (s Slice[T]) Equal(other []T, eq func(T, T) bool) bool {
	if len(s) != len(other) {
		return false
	}
	for i, v1 := range s {
		v2 := other[i]
		if !eq(v1, v2) {
			return false
		}
	}
	return true
}
