//go:build go1.18

package slices

import (
	"math/rand"
)

// Shuffle randomly shuffles the elements of the slice and returns the shuffled slice.
// This implementation creates a new slice to avoid modifying the original slice.
func Shuffle[T any](s []T) []T {
	shuffled := make([]T, len(s))
	copy(shuffled, s)
	for i := len(shuffled) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	}
	return shuffled
}
