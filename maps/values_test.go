package maps

import (
	"testing"

	"github.com/nnnewb/battery/slices"
)

// Unit test for Values function
func TestValues(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	expected := []int{1, 2, 3}
	result := Values(m)
	for _, v := range expected {
		if !slices.Contains(result, v) {
			t.Errorf("Values(%v) = %v; expected %v", m, result, expected)
		}
	}
}
