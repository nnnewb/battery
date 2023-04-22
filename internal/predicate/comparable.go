package predicate

// Equal returns true if v1 and v2 are equal, false otherwise.
func Equal[T comparable](v1, v2 T) bool {
	return v1 == v2
}
