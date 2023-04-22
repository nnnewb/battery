package maps

import "github.com/nnnewb/battery/slices"

// Keys returns the keys of the map m. The keys will be in an indeterminate order.
func Keys[M ~map[K]V, K comparable, V any](m M) slices.Slice[K] {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
