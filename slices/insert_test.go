//go:build go1.18

package slices

import (
	"reflect"
	"testing"
)

// Unit test for Insert function
func TestInsert(t *testing.T) {
	// Test inserting an element within range
	s := []int{1, 2, 3}
	expected := []int{1, 2, 4, 3}
	result := Insert(s, 2, 4)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Insert failed: expected %v, but got %v", expected, result)
	}

	// Test inserting an element at the end of the slice
	s = []int{1, 2, 3}
	expected = []int{1, 2, 3, 4}
	result = Insert(s, 3, 4)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Insert failed: expected %v, but got %v", expected, result)
	}

	// Test inserting an element at the beginning of the slice
	s = []int{1, 2, 3}
	expected = []int{4, 1, 2, 3}
	result = Insert(s, 0, 4)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Insert failed: expected %v, but got %v", expected, result)
	}

	// Test inserting an element into an empty slice
	s = []int{}
	expected = []int{1}
	result = Insert(s, 0, 1)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Insert failed: expected %v, but got %v", expected, result)
	}

	// Test inserting an element at an index greater than the length of the slice
	s = []int{1, 2, 3}
	expected = []int{1, 2, 3, 4}
	result = Insert(s, 10, 4)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Insert failed: expected %v, but got %v", expected, result)
	}
}
