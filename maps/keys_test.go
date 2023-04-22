package maps

import "testing"

// Unit test for the Keys function
func TestKeys(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	expected := []string{"a", "b", "c"}
	keys := Keys(m)
	if len(keys) != len(expected) {
		t.Errorf("Keys() returned %d keys, expected %d", len(keys), len(expected))
	}
	for _, k := range expected {
		if !contains(keys, k) {
			t.Errorf("Keys() returned keys %v, expected %v", keys, expected)
			break
		}
	}
}

// Helper function to check if a slice contains a given element
func contains(slice []string, elem string) bool {
	for _, e := range slice {
		if e == elem {
			return true
		}
	}
	return false
}
