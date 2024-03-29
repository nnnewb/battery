//go:build go1.18

package slices

// Copy returns a new slice with the same elements as the original slice. The original slice is not modified.
func Copy[T any](s []T) []T {
	// create a new slice with the same capacity as the original slice
	copySlice := make([]T, len(s), cap(s))
	// copy the elements from the original slice to the new slice
	copy(copySlice, s)
	return copySlice
}
