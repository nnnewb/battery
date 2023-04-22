package slices

import "testing"

// Unit test for Count function
func TestCount(t *testing.T) {
	// Test case 1: Counting even numbers in a slice
	s1 := Slice[int]{2, 4, 6, 7, 8, 9}
	evenPredicate := func(n int) bool { return n%2 == 0 }
	expectedCount1 := 4
	actualCount1 := s1.Count(evenPredicate)
	if actualCount1 != expectedCount1 {
		t.Errorf("Test case 1 failed: expected %d but got %d", expectedCount1, actualCount1)
	}

	// Test case 2: Counting strings with length greater than 5
	s2 := Slice[string]{"apple", "banana", "orange", "grapefruit", "kiwi"}
	longStringPredicate := func(s string) bool { return len(s) > 5 }
	expectedCount2 := 3
	actualCount2 := s2.Count(longStringPredicate)
	if actualCount2 != expectedCount2 {
		t.Errorf("Test case 2 failed: expected %d but got %d", expectedCount2, actualCount2)
	}
}
