package maps

import "testing"

// Unit test for Equal function
func TestEqual(t *testing.T) {
	m1 := map[string]int{"a": 1, "b": 2, "c": 3}
	m2 := map[string]int{"a": 1, "b": 2, "c": 3}
	if !Equal(m1, m2) {
		t.Errorf("Equal(m1, m2) = false, want true")
	}

	m3 := map[string]int{"a": 1, "b": 2}
	if Equal(m1, m3) {
		t.Errorf("Equal(m1, m3) = true, want false")
	}

	m4 := map[string]int{"a": 1, "b": 2, "c": 4}
	if Equal(m1, m4) {
		t.Errorf("Equal(m1, m4) = true, want false")
	}

	m5 := map[string]int{"a": 1, "b": 2, "d": 3}
	if Equal(m1, m5) {
		t.Errorf("Equal(m1, m5) = true, want false")
	}
}
