package slices

import (
	"reflect"
	"testing"
)

// Unit test for Insert function
func TestInsert(t *testing.T) {
	// Test inserting an element within range
	s := Slice[int]{1, 2, 3}
	expected := Slice[int]{1, 2, 4, 3}
	result := s.Insert(2, 4)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Insert failed: expected %v, but got %v", expected, result)
	}

	// Test inserting an element at the end of the slice
	s = Slice[int]{1, 2, 3}
	expected = Slice[int]{1, 2, 3, 4}
	result = s.Insert(3, 4)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Insert failed: expected %v, but got %v", expected, result)
	}

	// Test inserting an element at the beginning of the slice
	s = Slice[int]{1, 2, 3}
	expected = Slice[int]{4, 1, 2, 3}
	result = s.Insert(0, 4)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Insert failed: expected %v, but got %v", expected, result)
	}

	// Test inserting an element into an empty slice
	s = Slice[int]{}
	expected = Slice[int]{1}
	result = s.Insert(0, 1)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Insert failed: expected %v, but got %v", expected, result)
	}

	// Test inserting an element at an index greater than the length of the slice
	s = Slice[int]{1, 2, 3}
	expected = Slice[int]{1, 2, 3, 4}
	result = s.Insert(10, 4)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Insert failed: expected %v, but got %v", expected, result)
	}
}
