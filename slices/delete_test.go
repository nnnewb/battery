package slices

import (
	"reflect"
	"testing"
)

// Unit test for Delete function
func TestDelete(t *testing.T) {
	// Test case 1: delete element at index 0
	s1 := Slice[int]{1, 2, 3}
	s1 = s1.Delete(0)
	if !reflect.DeepEqual(s1, Slice[int]{2, 3}) {
		t.Errorf("Delete failed. Expected [2, 3], but got %v", s1)
	}

	// Test case 2: delete element at index len(s)-1
	s2 := Slice[string]{"apple", "banana", "cherry"}
	s2 = s2.Delete(len(s2) - 1)
	if !reflect.DeepEqual(s2, Slice[string]{"apple", "banana"}) {
		t.Errorf("Delete failed. Expected [apple, banana], but got %v", s2)
	}

	// Test case 3: delete element at index out of bounds
	s3 := Slice[float64]{1.1, 2.2, 3.3}
	s3 = s3.Delete(3)
	if !reflect.DeepEqual(s3, Slice[float64]{1.1, 2.2, 3.3}) {
		t.Errorf("Delete failed. Expected [1.1, 2.2, 3.3], but got %v", s3)
	}
}
