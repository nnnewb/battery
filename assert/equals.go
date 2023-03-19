package assert

import "testing"

func Equal[T comparable](t *testing.T, expr1, expr2 T) {
	t.Helper()

	if expr1 != expr2 {
		t.Fatalf("expect %v equals to %v", expr1, expr2)
	}
}
