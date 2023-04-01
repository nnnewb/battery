package slices

// Equal 比较两个切片是否相等，要求元素可比较。若元素
// 数量、顺序都一致则认为相同。注意 nil 切片和空切片会视作相等。
func Equal[T comparable](s1, s2 []T) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

// EqualFunc 比较两个切片是否相等。传入比较函数 eq，此函数返回 true 表示元素相同。
// 元素数量、顺序都一致则认为切片相等。注意 nil 切片和空切片会视作相等。
func EqualFunc[T1, T2 any](s1 []T1, s2 []T2, eq func(T1, T2) bool) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i, v1 := range s1 {
		v2 := s2[i]
		if !eq(v1, v2) {
			return false
		}
	}
	return true
}
