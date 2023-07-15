//go:build go1.18

package slices

import (
	"reflect"
	"testing"
)

// TestDrop tests the Drop function
func TestDrop(t *testing.T) {
	// Test case 1: Drop 0 elements from a non-empty slice
	s1 := []int{1, 2, 3, 4, 5}
	dropped1 := Drop(s1, 0)
	expected1 := []int{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(dropped1, expected1) {
		t.Errorf("Test case 1 failed: expected %v but got %v", expected1, dropped1)
	}

	// Test case 2: Drop 3 elements from a non-empty slice
	s2 := []int{1, 2, 3, 4, 5}
	dropped2 := Drop(s2, 3)
	expected2 := []int{4, 5}
	if !reflect.DeepEqual(dropped2, expected2) {
		t.Errorf("Test case 2 failed: expected %v but got %v", expected2, dropped2)
	}

	// Test case 3: Drop more elements than the length of the slice
	s3 := []int{1, 2, 3, 4, 5}
	dropped3 := Drop(s3, 10)
	expected3 := []int{}
	if !reflect.DeepEqual(dropped3, expected3) {
		t.Errorf("Test case 3 failed: expected %v but got %v", expected3, dropped3)
	}

	// Test case 4: Drop all elements from a non-empty slice
	s4 := []int{1, 2, 3, 4, 5}
	dropped4 := Drop(s4, len(s4))
	expected4 := []int{}
	if !reflect.DeepEqual(dropped4, expected4) {
		t.Errorf("Test case 4 failed: expected %v but got %v", expected4, dropped4)
	}

	// Test case 5: Drop elements from an empty slice
	s5 := []int{}
	dropped5 := Drop(s5, 10)
	expected5 := []int{}
	if !reflect.DeepEqual(dropped5, expected5) {
		t.Errorf("Test case 5 failed: expected %v but got %v", expected5, dropped5)
	}
}
