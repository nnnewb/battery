package assert

import "testing"

func Assert(t *testing.T, expr bool) {
	t.Helper()

	if !expr {
		t.Fatal("expect `false` to be `true`")
	}
}
