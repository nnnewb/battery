package maps

import "testing"

// Unit test for EqualFunc
func TestEqualFunc(t *testing.T) {
	m1 := map[string]int{"a": 1, "b": 2, "c": 3}
	m2 := map[string]int{"a": 1, "b": 4, "c": 3}
	eq := func(v1, v2 int) bool {
		return v1 == v2
	}
	if EqualFunc(m1, m2, eq) {
		t.Errorf("EqualFunc(%v, %v) = true, want false", m1, m2)
	}
}
