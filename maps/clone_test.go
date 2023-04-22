package maps

import "testing"

// Unit test for Clone function
func TestClone(t *testing.T) {
	// create a map with some key-value pairs
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	// clone the map
	clonedMap := Clone(m)
	// check that the cloned map is not the same as the original map
	if &clonedMap == &m {
		t.Errorf("Clone did not create a new map")
	}
	// check that the cloned map has the same key-value pairs as the original map
	for k, v := range m {
		if clonedMap[k] != v {
			t.Errorf("Clone did not copy key-value pair correctly")
		}
	}
}
