package assert

import "testing"

func Equal[T comparable](t *testing.T, expr1, expr2 T) {
	t.Helper()

	if expr1 != expr2 {
		t.Fatalf("expect %v equals to %v", expr1, expr2)
	}
}

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
