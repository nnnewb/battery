package slices

import (
	"reflect"
	"testing"
)

// Unit test for Chunk function
func TestChunk(t *testing.T) {
	// Test case 1: Chunk size is smaller than the length of the slice
	s := Slice[int]{1, 2, 3, 4, 5, 6, 7, 8, 9}
	n := 3
	expected := Slice[any]{
		Slice[int]{1, 2, 3},
		Slice[int]{4, 5, 6},
		Slice[int]{7, 8, 9},
	}
	result := s.Chunk(n)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Chunk(%v) = %v, expected %v", n, result, expected)
	}

	// Test case 2: Chunk size is greater than the length of the slice
	s = Slice[int]{1, 2, 3}
	n = 5
	expected = Slice[any]{
		Slice[int]{1, 2, 3},
	}
	result = s.Chunk(n)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Chunk(%v) = %v, expected %v", n, result, expected)
	}

	// Test case 3: Chunk size is equal to the length of the slice
	s = Slice[int]{1, 2, 3}
	n = 3
	expected = Slice[any]{
		Slice[int]{1, 2, 3},
	}
	result = s.Chunk(n)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Chunk(%v) = %v, expected %v", n, result, expected)
	}
}
