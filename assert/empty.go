package assert

import "testing"

func Empty[T any](t *testing.T, items []T) {
	t.Helper()

	if len(items) != 0 {
		t.Errorf("expected `%v` to be empty", items)
	}
}
