//go:build go1.18

package slices

// Count returns the number of elements in the slice that satisfy the given predicate function.
func Count[T any](s []T, predicate func(T) bool) int {
	count := 0
	for _, elem := range s {
		if predicate(elem) {
			count++
		}
	}
	return count
}
