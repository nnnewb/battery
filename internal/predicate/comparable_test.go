package predicate

import "testing"

// Unit test for Equal function
func TestEqual(t *testing.T) {
	// Test case 1: Equal values
	if !Equal(5, 5) {
		t.Errorf("Equal(5, 5) returned false, expected true")
	}

	// Test case 2: Unequal values
	if Equal(5, 6) {
		t.Errorf("Equal(5, 6) returned true, expected false")
	}

	// Test case 3: Equal strings
	if !Equal("hello", "hello") {
		t.Errorf(`Equal("hello", "hello") returned false, expected true`)
	}

	// Test case 4: Unequal strings
	if Equal("hello", "world") {
		t.Errorf(`Equal("hello", "world") returned true, expected false`)
	}
}
