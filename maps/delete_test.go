package maps

import "testing"

// Unit test for DeleteFunc
func TestDeleteFunc(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	del := func(k string, v int) bool {
		return v == 2
	}
	DeleteFunc(m, del)
	if len(m) != 2 {
		t.Errorf("DeleteFunc did not delete the correct number of key/value pairs")
	}
	if _, ok := m["b"]; ok {
		t.Errorf("DeleteFunc did not delete the correct key/value pairs")
	}
}
