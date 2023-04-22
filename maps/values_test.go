package maps

import (
	"testing"
)

// Unit test for Values function
func TestValues(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	expected := []int{1, 2, 3}
	result := Values(m)
	for _, v := range expected {
		if !result.Contains(v, func(i int, i2 int) bool { return i == i2 }) {
			t.Errorf("Values(%v) = %v; expected %v", m, result, expected)
		}
	}
}
