package slices

import (
	"reflect"
	"testing"
)

// Unit tests for Unique function
func TestUnique(t *testing.T) {
	// Test with empty slice
	emptySlice := make(Slice[int], 0)
	emptyUniqueSlice := Unique(emptySlice)
	if len(emptyUniqueSlice) != 0 {
		t.Errorf("Unique(%v) = %v; want []", emptySlice, emptyUniqueSlice)
	}

	// Test with slice containing only unique elements
	uniqueSlice := Slice[int]{1, 2, 3, 4, 5}
	uniqueUniqueSlice := Unique(uniqueSlice)
	if !reflect.DeepEqual(uniqueSlice, uniqueUniqueSlice) {
		t.Errorf("Unique(%v) = %v; want %v", uniqueSlice, uniqueUniqueSlice, uniqueSlice)
	}

	// Test with slice containing duplicate elements
	duplicateSlice := Slice[int]{1, 2, 3, 2, 4, 1, 5}
	duplicateUniqueSlice := Unique(duplicateSlice)
	expectedSlice := Slice[int]{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(duplicateUniqueSlice, expectedSlice) {
		t.Errorf("Unique(%v) = %v; want %v", duplicateSlice, duplicateUniqueSlice, expectedSlice)
	}
}

func TestSlice_Unique(t *testing.T) {
	// Test with a slice of integers
	intSlice := Slice[int]{1, 2, 3, 2, 4, 5, 3}
	expectedIntSlice := Slice[int]{1, 2, 3, 4, 5}
	uniqueIntSlice := intSlice.Unique(func(x, y int) bool {
		return x == y
	})
	if !reflect.DeepEqual(uniqueIntSlice, expectedIntSlice) {
		t.Errorf("Unique failed for integers. Expected %v but got %v", expectedIntSlice, uniqueIntSlice)
	}

	// Test with a slice of strings
	strSlice := Slice[string]{"apple", "banana", "orange", "banana", "kiwi", "orange"}
	expectedStrSlice := Slice[string]{"apple", "banana", "orange", "kiwi"}
	uniqueStrSlice := strSlice.Unique(func(x, y string) bool {
		return x == y
	})
	if !reflect.DeepEqual(uniqueStrSlice, expectedStrSlice) {
		t.Errorf("Unique failed for strings. Expected %v but got %v", expectedStrSlice, uniqueStrSlice)
	}

	// Test with a slice of structs
	type person struct {
		name string
		age  int
	}
	personSlice := Slice[person]{
		{"Alice", 25},
		{"Bob", 30},
		{"Charlie", 25},
		{"Bob", 30},
		{"David", 35},
		{"Charlie", 25},
	}
	expectedPersonSlice := Slice[person]{
		{"Alice", 25},
		{"Bob", 30},
		{"Charlie", 25},
		{"David", 35},
	}
	uniquePersonSlice := personSlice.Unique(func(x, y person) bool {
		return x.name == y.name && x.age == y.age
	})
	if !reflect.DeepEqual(uniquePersonSlice, expectedPersonSlice) {
		t.Errorf("Unique failed for structs. Expected %v but got %v", expectedPersonSlice, uniquePersonSlice)
	}
}
