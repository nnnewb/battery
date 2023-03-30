package assert

import "testing"

// Equal 比较两个可比较类型的值是否一致。若不一致则测试失败。
func Equal[T comparable](t *testing.T, expr1, expr2 T) {
	t.Helper()

	if expr1 != expr2 {
		t.Fatalf("expect %v equals to %v", expr1, expr2)
	}
}

// SliceEqual 比较两个切片是否相同，要求两个切片大小、元素、顺序都一致。
// 元素必须是可比较的。若不一致则测试失败。
func SliceEqual[T comparable](t *testing.T, s1, s2 []T) {
	t.Helper()

	if len(s1) != len(s2) {
		t.Errorf("%v and %v has different length", s1, s2)
		return
	}

	for i, v := range s1 {
		if v != s2[i] {
			t.Errorf("expected `%v` to equal `%v`", s1, s2)
		}
	}
}
