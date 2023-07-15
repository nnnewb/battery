//go:build go1.18

package slices

import (
	"reflect"
	"testing"
)

// TestTake tests the Take function
func TestTake(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}

	// Test taking 0 elements
	result := Take(s, 0)
	if len(result) != 0 {
		t.Errorf("Take(0) = %v; want []", result)
	}

	// Test taking all elements
	result = Take(s, len(s))
	if !reflect.DeepEqual(result, s) {
		t.Errorf("Take(len(s)) = %v; want %v", result, s)
	}

	// Test taking some elements
	result = Take(s, 3)
	if !reflect.DeepEqual(result, []int{1, 2, 3}) {
		t.Errorf("Take(3) = %v; want [1 2 3]", result)
	}

	// Test taking more elements than available
	result = Take(s, 10)
	if !reflect.DeepEqual(result, s) {
		t.Errorf("Take(10) = %v; want %v", result, s)
	}
}
