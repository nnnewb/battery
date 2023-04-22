package maps

import "testing"

// Unit test for Clear function
func TestClear(t *testing.T) {
	m := make(map[string]int)
	m["a"] = 1
	m["b"] = 2
	m["c"] = 3

	Clear(m)

	if len(m) != 0 {
		t.Errorf("Clear did not remove all entries from map")
	}
}
