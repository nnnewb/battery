//go:build go1.18

package slices

// Any tests whether any element in the slice satisfies the predicate function.
// Returns true if any element satisfies the predicate, false otherwise.
func Any[T any](s []T, predicate func(T) bool) bool {
	for _, elem := range s {
		if predicate(elem) {
			return true
		}
	}
	return false
}
