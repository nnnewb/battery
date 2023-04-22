package maps

import "testing"

// Unit test for Copy function
func TestCopy(t *testing.T) {
	// Create source map
	src := make(map[string]int)
	src["a"] = 1
	src["b"] = 2
	src["c"] = 3

	// Create destination map
	dst := make(map[string]int)
	dst["b"] = 20
	dst["d"] = 40

	// Call Copy function
	Copy(dst, src)

	// Check if all key/value pairs from src are added to dst
	if len(dst) != 4 {
		t.Errorf("Expected length of dst to be 4, but got %d", len(dst))
	}
	if dst["a"] != 1 {
		t.Errorf("Expected dst[\"a\"] to be 1, but got %d", dst["a"])
	}
	if dst["b"] != 2 {
		t.Errorf("Expected dst[\"b\"] to be 2, but got %d", dst["b"])
	}
	if dst["c"] != 3 {
		t.Errorf("Expected dst[\"c\"] to be 3, but got %d", dst["c"])
	}
	if dst["d"] != 40 {
		t.Errorf("Expected dst[\"d\"] to be 40, but got %d", dst["d"])
	}
}
