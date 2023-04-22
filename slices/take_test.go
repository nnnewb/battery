package slices

import (
	"reflect"
	"testing"
)

// TestTake tests the Take function
func TestTake(t *testing.T) {
	s := Slice[int]{1, 2, 3, 4, 5}

	// Test taking 0 elements
	result := s.Take(0)
	if len(result) != 0 {
		t.Errorf("Take(0) = %v; want []", result)
	}

	// Test taking all elements
	result = s.Take(len(s))
	if !reflect.DeepEqual(result, s) {
		t.Errorf("Take(len(s)) = %v; want %v", result, s)
	}

	// Test taking some elements
	result = s.Take(3)
	if !reflect.DeepEqual(result, Slice[int]{1, 2, 3}) {
		t.Errorf("Take(3) = %v; want [1 2 3]", result)
	}

	// Test taking more elements than available
	result = s.Take(10)
	if !reflect.DeepEqual(result, s) {
		t.Errorf("Take(10) = %v; want %v", result, s)
	}
}
