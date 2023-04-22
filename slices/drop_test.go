package slices

import (
	"reflect"
	"testing"
)

// TestDrop tests the Drop function
func TestDrop(t *testing.T) {
	// Test case 1: Drop 0 elements from a non-empty slice
	s1 := Slice[int]{1, 2, 3, 4, 5}
	dropped1 := s1.Drop(0)
	expected1 := Slice[int]{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(dropped1, expected1) {
		t.Errorf("Test case 1 failed: expected %v but got %v", expected1, dropped1)
	}

	// Test case 2: Drop 3 elements from a non-empty slice
	s2 := Slice[int]{1, 2, 3, 4, 5}
	dropped2 := s2.Drop(3)
	expected2 := Slice[int]{4, 5}
	if !reflect.DeepEqual(dropped2, expected2) {
		t.Errorf("Test case 2 failed: expected %v but got %v", expected2, dropped2)
	}

	// Test case 3: Drop more elements than the length of the slice
	s3 := Slice[int]{1, 2, 3, 4, 5}
	dropped3 := s3.Drop(10)
	expected3 := Slice[int]{}
	if !reflect.DeepEqual(dropped3, expected3) {
		t.Errorf("Test case 3 failed: expected %v but got %v", expected3, dropped3)
	}

	// Test case 4: Drop all elements from a non-empty slice
	s4 := Slice[int]{1, 2, 3, 4, 5}
	dropped4 := s4.Drop(len(s4))
	expected4 := Slice[int]{}
	if !reflect.DeepEqual(dropped4, expected4) {
		t.Errorf("Test case 4 failed: expected %v but got %v", expected4, dropped4)
	}

	// Test case 5: Drop elements from an empty slice
	s5 := Slice[int]{}
	dropped5 := s5.Drop(10)
	expected5 := Slice[int]{}
	if !reflect.DeepEqual(dropped5, expected5) {
		t.Errorf("Test case 5 failed: expected %v but got %v", expected5, dropped5)
	}
}
